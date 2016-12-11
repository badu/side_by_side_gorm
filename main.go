package tests

import (
	"SideBySideGorm/new_types"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	newGorm "github.com/badu/gorm"
	"github.com/erikstmartin/go-testdb"
	"github.com/jinzhu/now"
	pgdialect "gorm/dialects/postgres"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
)

var (
	newCompareToys = func(toys []new_types.Toy, contents []string) bool {
		var toyContents []string
		for _, toy := range toys {
			toyContents = append(toyContents, toy.Name)
		}
		sort.Strings(toyContents)
		sort.Strings(contents)
		return reflect.DeepEqual(toyContents, contents)
	}
)

func newGetPreloadUser(name string) *new_types.User {
	return newGetPreparedUser(name, "Preload")
}

func newCheckUserHasPreloadData(user new_types.User, t *testing.T) {
	u := newGetPreloadUser(user.Name)
	if user.BillingAddress.Address1 != u.BillingAddress.Address1 {
		t.Error("Failed to preload user's BillingAddress")
	}

	if user.ShippingAddress.Address1 != u.ShippingAddress.Address1 {
		t.Error("Failed to preload user's ShippingAddress")
	}

	if user.CreditCard.Number != u.CreditCard.Number {
		t.Error("Failed to preload user's CreditCard")
	}

	if user.Company.Name != u.Company.Name {
		t.Error("Failed to preload user's Company")
	}

	if len(user.Emails) != len(u.Emails) {
		t.Error("Failed to preload user's Emails")
	} else {
		var found int
		for _, e1 := range u.Emails {
			for _, e2 := range user.Emails {
				if e1.Email == e2.Email {
					found++
					break
				}
			}
		}
		if found != len(u.Emails) {
			t.Error("Failed to preload user's email details")
		}
	}
}

func newCompareTags(tags []new_types.Tag, contents []string) bool {
	var tagContents []string
	for _, tag := range tags {
		tagContents = append(tagContents, tag.Value)
	}
	sort.Strings(tagContents)
	sort.Strings(contents)
	return reflect.DeepEqual(tagContents, contents)
}

func newGetPreparedUser(name string, role string) *new_types.User {
	var company new_types.Company
	new_types.TestDB.Where(new_types.Company{Name: role}).FirstOrCreate(&company)

	return &new_types.User{
		Name:            name,
		Age:             20,
		Role:            new_types.Role{role},
		BillingAddress:  new_types.Address{Address1: fmt.Sprintf("Billing Address %v", name)},
		ShippingAddress: new_types.Address{Address1: fmt.Sprintf("Shipping Address %v", name)},
		CreditCard:      new_types.CreditCard{Number: fmt.Sprintf("123456%v", name)},
		Emails: []new_types.Email{
			{Email: fmt.Sprintf("user_%v@example1.com", name)}, {Email: fmt.Sprintf("user_%v@example2.com", name)},
		},
		Company: company,
		Languages: []new_types.Language{
			{Name: fmt.Sprintf("lang_1_%v", name)},
			{Name: fmt.Sprintf("lang_2_%v", name)},
		},
	}
}

func NewNameIn1And2(d *newGorm.DBCon) *newGorm.DBCon {
	return d.Where("name in (?)", []string{"ScopeUser1", "ScopeUser2"})
}

func NewNameIn2And3(d *newGorm.DBCon) *newGorm.DBCon {
	return d.Where("name in (?)", []string{"ScopeUser2", "ScopeUser3"})
}

func NewNameIn(names []string) func(d *newGorm.DBCon) *newGorm.DBCon {
	return func(d *newGorm.DBCon) *newGorm.DBCon {
		return d.Where("name in (?)", names)
	}
}

func newtoJSONString(v interface{}) []byte {
	r, _ := json.MarshalIndent(v, "", "  ")
	return r
}

func newparseTime(str string) *time.Time {
	t := now.MustParse(str)
	return &t
}

func NewDialectHasTzSupport() bool {
	// NB: FoundationDB do not support time zones.
	if dialect := os.Getenv("GORM_DIALECT"); dialect == "foundation" {
		return false
	}
	return true
}

func OpenNewTestConnection(t *testing.T) {
	osDialect := os.Getenv("GORM_DIALECT")
	osDBAddress := os.Getenv("GORM_DBADDRESS")

	switch osDialect {
	case "mysql":
		//osDBAddress = "127.0.0.1:3306"

		// CREATE USER 'gorm'@'localhost' IDENTIFIED BY 'gorm';
		// CREATE DATABASE gorm;
		// GRANT ALL ON * TO 'gorm'@'localhost';
		if osDBAddress != "" {
			osDBAddress = fmt.Sprintf("tcp(%v)", osDBAddress)
		}
		new_types.TestDB, new_types.TestDBErr = newGorm.Open("mysql", fmt.Sprintf("root:@%v/gorm?charset=utf8&parseTime=True", osDBAddress))
		if new_types.TestDBErr != nil {
			t.Fatalf("ERROR : %v", new_types.TestDBErr)
		}
	case "postgres":
		if osDBAddress != "" {
			osDBAddress = fmt.Sprintf("host=%v ", osDBAddress)
		}
		new_types.TestDB, new_types.TestDBErr = newGorm.Open("postgres", fmt.Sprintf("%vuser=gorm password=gorm DB.name=gorm sslmode=disable", osDBAddress))
		if new_types.TestDBErr != nil{
			t.Fatalf("ERROR : %v", new_types.TestDBErr)
		}
	case "foundation":
		new_types.TestDB, new_types.TestDBErr = newGorm.Open("foundation", "dbname=gorm port=15432 sslmode=disable")
		if new_types.TestDBErr != nil{
			t.Fatalf("ERROR : %v", new_types.TestDBErr)
		}
	default:
		new_types.TestDB, new_types.TestDBErr = newGorm.Open("sqlite3", "test.db?cache=shared&mode=memory")
		if new_types.TestDBErr != nil {

			t.Fatalf("ERROR : %v", new_types.TestDBErr)
		}
	}
	//TODO : @Badu - uncomment below if you want full traces
	//TestDB.SetLogMode(newGorm.LOG_DEBUG)

	new_types.TestDB.DB().SetMaxIdleConns(10)
}

func RunNewMigration(t *testing.T) {
	//t.Log("Running migration...")
	if err := new_types.TestDB.DropTableIfExists(&new_types.User{}).Error; err != nil {
		fmt.Printf("Got error when try to delete table users, %+v\n", err)
	}

	for _, table := range []string{"animals", "user_languages"} {
		new_types.TestDB.Exec(fmt.Sprintf("drop table %v;", table))
	}

	values := []interface{}{
		&new_types.Short{},
		&new_types.ReallyLongThingThatReferencesShort{},
		&new_types.ReallyLongTableNameToTestMySQLNameLengthLimit{},
		&new_types.NotSoLongTableName{},
		&new_types.Product{},
		&new_types.Email{},
		&new_types.Address{},
		&new_types.CreditCard{},
		&new_types.Company{},
		&new_types.Role{},
		&new_types.Language{},
		&new_types.HNPost{},
		&new_types.EngadgetPost{},
		&new_types.Animal{},
		&new_types.User{},
		&new_types.JoinTable{},
		&new_types.Post{},
		&new_types.Category{},
		&new_types.Comment{},
		&new_types.Cat{},
		&new_types.Dog{},
		&new_types.Hamster{},
		&new_types.Toy{},
		&new_types.ElementWithIgnoredField{},
	}
	for _, value := range values {
		new_types.TestDB.DropTable(value)
	}
	if err := new_types.TestDB.AutoMigrate(values...).Error; err != nil {
		panic(fmt.Sprintf("No error should happen when create table, but got %+v", err))
	}
	//t.Log("Migration done.")
}

func StringPrimaryKey(t *testing.T) {
	type UUIDStruct struct {
		ID   string `gorm:"primary_key"`
		Name string
	}
	new_types.TestDB.DropTable(&UUIDStruct{})
	new_types.TestDB.AutoMigrate(&UUIDStruct{})

	data := UUIDStruct{ID: "uuid", Name: "hello"}
	if err := new_types.TestDB.Save(&data).Error; err != nil || data.ID != "uuid" || data.Name != "hello" {
		t.Errorf("string primary key should not be populated")
	}

	data = UUIDStruct{ID: "uuid", Name: "hello world"}
	if err := new_types.TestDB.Save(&data).Error; err != nil || data.ID != "uuid" || data.Name != "hello world" {
		t.Errorf("string primary key should not be populated")
	}
}

func ExceptionsWithInvalidSql(t *testing.T) {
	var columns []string
	if new_types.TestDB.Where("sdsd.zaaa = ?", "sd;;;aa").Pluck("aaa", &columns).Error == nil {
		t.Errorf("Should got error with invalid SQL")
	}

	if new_types.TestDB.Model(&new_types.User{}).Where("sdsd.zaaa = ?", "sd;;;aa").Pluck("aaa", &columns).Error == nil {
		t.Errorf("Should got error with invalid SQL")
	}

	if new_types.TestDB.Where("sdsd.zaaa = ?", "sd;;;aa").Find(&new_types.User{}).Error == nil {
		t.Errorf("Should got error with invalid SQL")
	}

	var count1, count2 int64
	new_types.TestDB.Model(&new_types.User{}).Count(&count1)
	if count1 <= 0 {
		t.Errorf("Should find some users")
	}

	if new_types.TestDB.Where("name = ?", "jinzhu; delete * from users").First(&new_types.User{}).Error == nil {
		t.Errorf("Should got error with invalid SQL")
	}

	new_types.TestDB.Model(&new_types.User{}).Count(&count2)
	if count1 != count2 {
		t.Errorf("No user should not be deleted by invalid SQL")
	}
}

func SetTable(t *testing.T) {
	new_types.TestDB.Create(newGetPreparedUser("pluck_user1", "pluck_user"))
	new_types.TestDB.Create(newGetPreparedUser("pluck_user2", "pluck_user"))
	new_types.TestDB.Create(newGetPreparedUser("pluck_user3", "pluck_user"))

	if err := new_types.TestDB.Table("users").Where("role = ?", "pluck_user").Pluck("age", &[]int{}).Error; err != nil {
		t.Error("No errors should happen if set table for pluck", err)
	}

	var users []new_types.User
	if new_types.TestDB.Table("users").Find(&[]new_types.User{}).Error != nil {
		t.Errorf("No errors should happen if set table for find")
	}

	if new_types.TestDB.Table("invalid_table").Find(&users).Error == nil {
		t.Errorf("Should got error when table is set to an invalid table")
	}

	new_types.TestDB.Exec("drop table deleted_users;")
	if new_types.TestDB.Table("deleted_users").CreateTable(&new_types.User{}).Error != nil {
		t.Errorf("Create table with specified table")
	}

	new_types.TestDB.Table("deleted_users").Save(&new_types.User{Name: "DeletedUser"})

	var deletedUsers []new_types.User
	new_types.TestDB.Table("deleted_users").Find(&deletedUsers)
	if len(deletedUsers) != 1 {
		t.Errorf("Query from specified table")
	}

	new_types.TestDB.Save(newGetPreparedUser("normal_user", "reset_table"))
	new_types.TestDB.Table("deleted_users").Save(newGetPreparedUser("deleted_user", "reset_table"))
	var user1, user2, user3 new_types.User
	new_types.TestDB.Where("role = ?", "reset_table").First(&user1).Table("deleted_users").First(&user2).Table("").First(&user3)
	//TODO : @Badu - simplify
	if (user1.Name != "normal_user") || (user2.Name != "deleted_user") || (user3.Name != "normal_user") {
		t.Errorf("unset specified table with blank string")
	}
}

func HasTable(t *testing.T) {
	type Foo struct {
		Id    int
		Stuff string
	}
	new_types.TestDB.DropTable(&Foo{})

	// Table should not exist at this point, HasTable should return false
	if ok := new_types.TestDB.HasTable("foos"); ok {
		t.Errorf("Table should not exist, but does")
	}
	if ok := new_types.TestDB.HasTable(&Foo{}); ok {
		t.Errorf("Table should not exist, but does")
	}

	// We create the table
	if err := new_types.TestDB.CreateTable(&Foo{}).Error; err != nil {
		t.Errorf("Table should be created")
	}

	// And now it should exits, and HasTable should return true
	if ok := new_types.TestDB.HasTable("foos"); !ok {
		t.Errorf("Table should exist, but HasTable informs it does not")
	}
	if ok := new_types.TestDB.HasTable(&Foo{}); !ok {
		t.Errorf("Table should exist, but HasTable informs it does not")
	}
}

func TableName(t *testing.T) {
	DB := new_types.TestDB.Model("")
	if DB.NewScope(new_types.Order{}).TableName() != "orders" {
		t.Errorf("Order's table name should be orders")
	}

	if DB.NewScope(&new_types.Order{}).TableName() != "orders" {
		t.Errorf("&Order's table name should be orders")
	}

	if DB.NewScope([]new_types.Order{}).TableName() != "orders" {
		t.Errorf("[]Order's table name should be orders")
	}

	if DB.NewScope(&[]new_types.Order{}).TableName() != "orders" {
		t.Errorf("&[]Order's table name should be orders")
	}

	DB.SingularTable(true)
	if DB.NewScope(new_types.Order{}).TableName() != "order" {
		t.Errorf("Order's singular table name should be order")
	}

	if DB.NewScope(&new_types.Order{}).TableName() != "order" {
		t.Errorf("&Order's singular table name should be order")
	}

	if DB.NewScope([]new_types.Order{}).TableName() != "order" {
		t.Errorf("[]Order's singular table name should be order")
	}

	if DB.NewScope(&[]new_types.Order{}).TableName() != "order" {
		t.Errorf("&[]Order's singular table name should be order")
	}

	if DB.NewScope(&new_types.Cart{}).TableName() != "shopping_cart" {
		t.Errorf("&Cart's singular table name should be shopping_cart")
	}

	if DB.NewScope(new_types.Cart{}).TableName() != "shopping_cart" {
		t.Errorf("Cart's singular table name should be shopping_cart")
	}

	if DB.NewScope(&[]new_types.Cart{}).TableName() != "shopping_cart" {
		t.Errorf("&[]Cart's singular table name should be shopping_cart")
	}

	if DB.NewScope([]new_types.Cart{}).TableName() != "shopping_cart" {
		t.Errorf("[]Cart's singular table name should be shopping_cart")
	}
	DB.SingularTable(false)
}

func NullValues(t *testing.T) {
	new_types.TestDB.DropTable(&new_types.NullValue{})
	new_types.TestDB.AutoMigrate(&new_types.NullValue{})

	if err := new_types.TestDB.Save(&new_types.NullValue{
		Name:    sql.NullString{String: "hello", Valid: true},
		Gender:  &sql.NullString{String: "M", Valid: true},
		Age:     sql.NullInt64{Int64: 18, Valid: true},
		Male:    sql.NullBool{Bool: true, Valid: true},
		Height:  sql.NullFloat64{Float64: 100.11, Valid: true},
		AddedAt: new_types.NullTime{Time: time.Now(), Valid: true},
	}).Error; err != nil {
		t.Errorf("Not error should raise when test null value")
	}

	var nv new_types.NullValue
	new_types.TestDB.First(&nv, "name = ?", "hello")

	if nv.Name.String != "hello" || nv.Gender.String != "M" || nv.Age.Int64 != 18 || nv.Male.Bool != true || nv.Height.Float64 != 100.11 || nv.AddedAt.Valid != true {
		t.Errorf("Should be able to fetch null value")
	}

	if err := new_types.TestDB.Save(&new_types.NullValue{
		Name:    sql.NullString{String: "hello-2", Valid: true},
		Gender:  &sql.NullString{String: "F", Valid: true},
		Age:     sql.NullInt64{Int64: 18, Valid: false},
		Male:    sql.NullBool{Bool: true, Valid: true},
		Height:  sql.NullFloat64{Float64: 100.11, Valid: true},
		AddedAt: new_types.NullTime{Time: time.Now(), Valid: false},
	}).Error; err != nil {
		t.Errorf("Not error should raise when test null value")
	}

	var nv2 new_types.NullValue
	new_types.TestDB.First(&nv2, "name = ?", "hello-2")
	if nv2.Name.String != "hello-2" || nv2.Gender.String != "F" || nv2.Age.Int64 != 0 || nv2.Male.Bool != true || nv2.Height.Float64 != 100.11 || nv2.AddedAt.Valid != false {
		t.Errorf("Should be able to fetch null value")
	}

	if err := new_types.TestDB.Save(&new_types.NullValue{
		Name:    sql.NullString{String: "hello-3", Valid: false},
		Gender:  &sql.NullString{String: "M", Valid: true},
		Age:     sql.NullInt64{Int64: 18, Valid: false},
		Male:    sql.NullBool{Bool: true, Valid: true},
		Height:  sql.NullFloat64{Float64: 100.11, Valid: true},
		AddedAt: new_types.NullTime{Time: time.Now(), Valid: false},
	}).Error; err == nil {
		t.Errorf("Can't save because of name can't be null")
	}
}

func NullValuesWithFirstOrCreate(t *testing.T) {
	var nv1 = new_types.NullValue{
		Name:   sql.NullString{String: "first_or_create", Valid: true},
		Gender: &sql.NullString{String: "M", Valid: true},
	}

	var nv2 new_types.NullValue
	result := new_types.TestDB.Where(nv1).FirstOrCreate(&nv2)

	if result.RowsAffected != 1 {
		t.Errorf("RowsAffected should be 1 after create some record")
	}

	if result.Error != nil {
		t.Errorf("Should not raise any error, but got %v", result.Error)
	}

	if nv2.Name.String != "first_or_create" || nv2.Gender.String != "M" {
		t.Errorf("first or create with nullvalues")
	}

	if err := new_types.TestDB.Where(nv1).Assign(new_types.NullValue{Age: sql.NullInt64{Int64: 18, Valid: true}}).FirstOrCreate(&nv2).Error; err != nil {
		t.Errorf("Should not raise any error, but got %v", err)
	}

	if nv2.Age.Int64 != 18 {
		t.Errorf("should update age to 18")
	}
}

func Transaction(t *testing.T) {
	tx := new_types.TestDB.Begin()
	u := new_types.User{Name: "transcation"}
	if err := tx.Save(&u).Error; err != nil {
		t.Errorf("No error should raise")
	}

	if err := tx.First(&new_types.User{}, "name = ?", "transcation").Error; err != nil {
		t.Errorf("Should find saved record")
	}

	if sqlTx, ok := tx.AsSQLDB().(*sql.Tx); !ok || sqlTx == nil {
		t.Errorf("Should return the underlying sql.Tx")
	}

	tx.Rollback()

	if err := tx.First(&new_types.User{}, "name = ?", "transcation").Error; err == nil {
		t.Errorf("Should not find record after rollback")
	}

	tx2 := new_types.TestDB.Begin()
	u2 := new_types.User{Name: "transcation-2"}
	if err := tx2.Save(&u2).Error; err != nil {
		t.Errorf("No error should raise")
	}

	if err := tx2.First(&new_types.User{}, "name = ?", "transcation-2").Error; err != nil {
		t.Errorf("Should find saved record")
	}

	tx2.Commit()

	if err := new_types.TestDB.First(&new_types.User{}, "name = ?", "transcation-2").Error; err != nil {
		t.Errorf("Should be able to find committed record")
	}
}

func Row(t *testing.T) {
	user1 := new_types.User{Name: "RowUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "RowUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "RowUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	row := new_types.TestDB.Table("users").Where("name = ?", user2.Name).Select("age").Row()
	var age int64
	row.Scan(&age)
	if age != 10 {
		t.Errorf("Scan with Row")
	}
}

func Rows(t *testing.T) {
	user1 := new_types.User{Name: "RowsUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "RowsUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "RowsUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	rows, err := new_types.TestDB.Table("users").Where("name = ? or name = ?", user2.Name, user3.Name).Select("name, age").Rows()
	if err != nil {
		t.Errorf("Not error should happen, got %v", err)
	}

	count := 0
	for rows.Next() {
		var name string
		var age int64
		rows.Scan(&name, &age)
		count++
	}

	if count != 2 {
		t.Errorf("Should found two records")
	}
}

func ScanRows(t *testing.T) {
	user1 := new_types.User{Name: "ScanRowsUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "ScanRowsUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "ScanRowsUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	rows, err := new_types.TestDB.Table("users").Where("name = ? or name = ?", user2.Name, user3.Name).Select("name, age").Rows()
	if err != nil {
		t.Errorf("Not error should happen, got %v", err)
	}

	type Result struct {
		Name string
		Age  int
	}

	var results []Result
	for rows.Next() {
		var result Result
		if err := new_types.TestDB.ScanRows(rows, &result); err != nil {
			t.Errorf("should get no error, but got %v", err)
		}
		results = append(results, result)
	}

	if !reflect.DeepEqual(results, []Result{{Name: "ScanRowsUser2", Age: 10}, {Name: "ScanRowsUser3", Age: 20}}) {
		t.Errorf("Should find expected results")
	}
}

func Scan(t *testing.T) {
	user1 := new_types.User{Name: "ScanUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "ScanUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "ScanUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	type result struct {
		Name string
		Age  int
	}

	var res result
	new_types.TestDB.Table("users").Select("name, age").Where("name = ?", user3.Name).Scan(&res)
	if res.Name != user3.Name {
		t.Errorf("Scan into struct should work")
	}

	var doubleAgeRes result
	new_types.TestDB.Table("users").Select("age + age as age").Where("name = ?", user3.Name).Scan(&doubleAgeRes)
	if doubleAgeRes.Age != res.Age*2 {
		t.Errorf("Scan double age as age")
	}

	var ress []result
	new_types.TestDB.Table("users").Select("name, age").Where("name in (?)", []string{user2.Name, user3.Name}).Scan(&ress)
	if len(ress) != 2 || ress[0].Name != user2.Name || ress[1].Name != user3.Name {
		t.Errorf("Scan into struct map")
	}
}

func Raw(t *testing.T) {
	user1 := new_types.User{Name: "ExecRawSqlUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "ExecRawSqlUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "ExecRawSqlUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	type result struct {
		Name  string
		Email string
	}

	var ress []result
	new_types.TestDB.Raw("SELECT name, age FROM users WHERE name = ? or name = ?", user2.Name, user3.Name).Scan(&ress)
	if len(ress) != 2 || ress[0].Name != user2.Name || ress[1].Name != user3.Name {
		t.Errorf("Raw with scan")
	}

	rows, _ := new_types.TestDB.Raw("select name, age from users where name = ?", user3.Name).Rows()
	count := 0
	for rows.Next() {
		count++
	}
	if count != 1 {
		t.Errorf("Raw with Rows should find one record with name 3")
	}

	new_types.TestDB.Exec("update users set name=? where name in (?)", "jinzhu", []string{user1.Name, user2.Name, user3.Name})
	if new_types.TestDB.Where("name in (?)", []string{user1.Name, user2.Name, user3.Name}).First(&new_types.User{}).Error != newGorm.ErrRecordNotFound {
		t.Error("Raw sql to update records")
	}
}

func Group(t *testing.T) {
	rows, err := new_types.TestDB.Select("name").Table("users").Group("name").Rows()

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var name string
			rows.Scan(&name)
		}
	} else {
		t.Errorf("Should not raise any error")
	}
}

func Joins(t *testing.T) {
	var user = new_types.User{
		Name:       "joins",
		CreditCard: new_types.CreditCard{Number: "411111111111"},
		Emails:     []new_types.Email{{Email: "join1@example.com"}, {Email: "join2@example.com"}},
	}
	new_types.TestDB.Save(&user)

	var users1 []new_types.User
	new_types.TestDB.Joins("left join emails on emails.user_id = users.id").Where("name = ?", "joins").Find(&users1)
	if len(users1) != 2 {
		t.Errorf("should find two users using left join")
	}

	var users2 []new_types.User
	new_types.TestDB.Joins("left join emails on emails.user_id = users.id AND emails.email = ?", "join1@example.com").Where("users.name = ?", "joins").First(&users2)
	if len(users2) != 1 {
		t.Errorf("should find one users using left join with conditions")
	}

	var users3 []new_types.User
	new_types.TestDB.Joins("join emails on emails.user_id = users.id AND emails.email = ?", "join1@example.com").Joins("join credit_cards on credit_cards.user_id = users.id AND credit_cards.number = ?", "411111111111").Where("name = ?", "joins").First(&users3)
	if len(users3) != 1 {
		t.Errorf("should find one users using multiple left join conditions : %v")
	}

	var users4 []new_types.User
	new_types.TestDB.Joins("join emails on emails.user_id = users.id AND emails.email = ?", "join1@example.com").Joins("join credit_cards on credit_cards.user_id = users.id AND credit_cards.number = ?", "422222222222").Where("name = ?", "joins").First(&users4)
	if len(users4) != 0 {
		t.Errorf("should find no user when searching with unexisting credit card")
	}

	var users5 []new_types.User
	db5 := new_types.TestDB.Joins("join emails on emails.user_id = users.id AND emails.email = ?", "join1@example.com").Joins("join credit_cards on credit_cards.user_id = users.id AND credit_cards.number = ?", "411111111111").Where(new_types.User{Id: 1}).Where(new_types.Email{Id: 1}).Not(new_types.Email{Id: 10}).First(&users5)
	if db5.Error != nil {
		t.Errorf("Should not raise error for join where identical fields in different tables. Error: %s", db5.Error.Error())
	}
}

func JoinsWithSelect(t *testing.T) {
	type result struct {
		Name  string
		Email string
	}

	user := new_types.User{
		Name:   "joins_with_select",
		Emails: []new_types.Email{{Email: "join1@example.com"}, {Email: "join2@example.com"}},
	}
	new_types.TestDB.Save(&user)

	var results []result
	new_types.TestDB.Table("users").Select("name, emails.email").Joins("left join emails on emails.user_id = users.id").Where("name = ?", "joins_with_select").Scan(&results)
	if len(results) != 2 || results[0].Email != "join1@example.com" || results[1].Email != "join2@example.com" {
		t.Errorf("Should find all two emails with Join select")
	}
}

func Having(t *testing.T) {
	rows, err := new_types.TestDB.Select("name, count(*) as total").Table("users").Group("name").Having("name IN (?)", []string{"2", "3"}).Rows()

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var name string
			var total int64
			rows.Scan(&name, &total)

			if name == "2" && total != 1 {
				t.Errorf("Should have one user having name 2")
			}
			if name == "3" && total != 2 {
				t.Errorf("Should have two users having name 3")
			}
		}
	} else {
		t.Errorf("Should not raise any error")
	}
}

func TimeWithZone(t *testing.T) {
	var format = "2006-01-02 15:04:05 -0700"
	var times []time.Time
	GMT8, _ := time.LoadLocation("Asia/Shanghai")
	times = append(times, time.Date(2013, 02, 19, 1, 51, 49, 123456789, GMT8))
	times = append(times, time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.UTC))

	for index, vtime := range times {
		name := "time_with_zone_" + strconv.Itoa(index)
		user := new_types.User{Name: name, Birthday: &vtime}

		if !NewDialectHasTzSupport() {
			// If our driver dialect doesn't support TZ's, just use UTC for everything here.
			utcBirthday := user.Birthday.UTC()
			user.Birthday = &utcBirthday
		}

		new_types.TestDB.Save(&user)
		expectedBirthday := "2013-02-18 17:51:49 +0000"
		foundBirthday := user.Birthday.UTC().Format(format)
		if foundBirthday != expectedBirthday {
			t.Errorf("User's birthday should not be changed after save for name=%s, expected bday=%+v but actual value=%+v", name, expectedBirthday, foundBirthday)
		}

		var findUser, findUser2, findUser3 new_types.User
		new_types.TestDB.First(&findUser, "name = ?", name)
		foundBirthday = findUser.Birthday.UTC().Format(format)
		if foundBirthday != expectedBirthday {
			t.Errorf("User's birthday should not be changed after find for name=%s, expected bday=%+v but actual value=%+v", name, expectedBirthday, foundBirthday)
		}

		if new_types.TestDB.Where("id = ? AND birthday >= ?", findUser.Id, user.Birthday.Add(-time.Minute)).First(&findUser2).RecordNotFound() {
			t.Errorf("User should be found")
		}

		if !new_types.TestDB.Where("id = ? AND birthday >= ?", findUser.Id, user.Birthday.Add(time.Minute)).First(&findUser3).RecordNotFound() {
			t.Errorf("User should not be found")
		}
	}
}

func Hstore(t *testing.T) {
	type Details struct {
		Id   int64
		Bulk pgdialect.Hstore
	}

	if dialect := os.Getenv("GORM_DIALECT"); dialect != "postgres" {
		t.Skip()
	}

	if err := new_types.TestDB.Exec("CREATE EXTENSION IF NOT EXISTS hstore").Error; err != nil {
		//fmt.Println("\033[31mHINT: Must be superuser to create hstore extension (ALTER USER gorm WITH SUPERUSER;)\033[0m")
		panic(fmt.Sprintf("No error should happen when create hstore extension, but got %+v", err))
	}

	new_types.TestDB.Exec("drop table details")

	if err := new_types.TestDB.CreateTable(&Details{}).Error; err != nil {
		panic(fmt.Sprintf("No error should happen when create table, but got %+v", err))
	}

	bankAccountId, phoneNumber, opinion := "123456", "14151321232", "sharkbait"
	bulk := map[string]*string{
		"bankAccountId": &bankAccountId,
		"phoneNumber":   &phoneNumber,
		"opinion":       &opinion,
	}
	d := Details{Bulk: bulk}
	new_types.TestDB.Save(&d)

	var d2 Details
	if err := new_types.TestDB.First(&d2).Error; err != nil {
		t.Errorf("Got error when tried to fetch details: %+v", err)
	}

	for k := range bulk {
		if r, ok := d2.Bulk[k]; ok {
			if res, _ := bulk[k]; *res != *r {
				t.Errorf("Details should be equal")
			}
		} else {
			t.Errorf("Details should be existed")
		}
	}
}

func SetAndGet(t *testing.T) {
	if value, ok := new_types.TestDB.Set("gorm:save_associations", true).Get("gorm:save_associations"); !ok {
		t.Errorf("Should be able to get setting 'gorm:save_associations' after set")
	} else {
		if !value.(bool) {
			t.Errorf("Setted value should be TRUE")
		}
	}

	if _, ok := new_types.TestDB.Get("non_existing"); ok {
		t.Errorf("Get non existing key should return error")
	}
}

func CompatibilityMode(t *testing.T) {
	DB, _ := newGorm.Open("testdb", "")
	testdb.SetQueryFunc(func(query string) (driver.Rows, error) {
		columns := []string{"id", "name", "age"}
		result := `
		1,Tim,20
		2,Joe,25
		3,Bob,30
		`
		return testdb.RowsFromCSVString(columns, result), nil
	})

	var users []new_types.User
	DB.Find(&users)
	if (users[0].Name != "Tim") || len(users) != 3 {
		t.Errorf("Unexcepted result returned")
	}
}

func OpenExistingDB(t *testing.T) {
	new_types.TestDB.Save(&new_types.User{Name: "jnfeinstein"})
	dialect := os.Getenv("GORM_DIALECT")

	db, err := newGorm.Open(dialect, new_types.TestDB.DB())
	if err != nil {
		t.Errorf("Should have wrapped the existing DB connection")
	}

	var user new_types.User
	if db.Where("name = ?", "jnfeinstein").First(&user).Error == newGorm.ErrRecordNotFound {
		t.Errorf("Should have found existing record")
	}
}

func DdlErrors(t *testing.T) {
	var err error

	if err = new_types.TestDB.Close(); err != nil {
		t.Errorf("Closing DDL test db connection err=%s", err)
	}
	defer func() {
		// Reopen DB connection.
		OpenNewTestConnection(t)
		if new_types.TestDBErr != nil {
			t.Fatalf("Failed re-opening db connection: %s", new_types.TestDBErr)
		}
	}()

	if err := new_types.TestDB.Find(&new_types.User{}).Error; err == nil {
		t.Errorf("Expected operation on closed db to produce an error, but err was nil")
	}
}

func OpenWithOneParameter(t *testing.T) {
	db, err := newGorm.Open("dialect")
	if db != nil {
		t.Error("Open with one parameter returned non nil for db")
	}
	if err == nil {
		t.Error("Open with one parameter returned err as nil")
	}
}
