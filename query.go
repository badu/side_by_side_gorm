package tests

import (
	"fmt"
	"reflect"

	"SideBySideGorm/new_types"
	newGorm "github.com/badu/gorm"
	"testing"
	"time"
)

func FirstAndLast(t *testing.T) {
	new_types.TestDB.Save(&new_types.User{Name: "user1", Emails: []new_types.Email{{Email: "user1@example.com"}}})
	new_types.TestDB.Save(&new_types.User{Name: "user2", Emails: []new_types.Email{{Email: "user2@example.com"}}})

	var user1, user2, user3, user4 new_types.User
	new_types.TestDB.First(&user1)
	new_types.TestDB.Order("id").Limit(1).Find(&user2)

	new_types.TestDB.Last(&user3)
	new_types.TestDB.Order("id desc").Limit(1).Find(&user4)
	//TODO : @Badu - simplify
	if user1.Id != user2.Id || user3.Id != user4.Id {
		t.Errorf("First and Last should by order by primary key")
	}

	var users []new_types.User
	new_types.TestDB.First(&users)
	if len(users) != 1 {
		t.Errorf("Find first record as slice")
	}

	var user new_types.User
	if new_types.TestDB.Joins("left join emails on emails.user_id = users.id").First(&user).Error != nil {
		t.Errorf("Should not raise any error when order with Join table")
	}

	if user.Email != "" {
		t.Errorf("User's Email should be blank as no one set it")
	}
}

func FirstAndLastWithNoStdPrimaryKey(t *testing.T) {
	new_types.TestDB.Save(&new_types.Animal{Name: "animal1"})
	new_types.TestDB.Save(&new_types.Animal{Name: "animal2"})

	var animal1, animal2, animal3, animal4 new_types.Animal
	new_types.TestDB.First(&animal1)
	new_types.TestDB.Order("counter").Limit(1).Find(&animal2)

	new_types.TestDB.Last(&animal3)
	new_types.TestDB.Order("counter desc").Limit(1).Find(&animal4)
	//TODO : @Badu - simplify
	if animal1.Counter != animal2.Counter || animal3.Counter != animal4.Counter {
		t.Errorf("First and Last should work correctly")
	}
}

func FirstAndLastWithRaw(t *testing.T) {
	user1 := new_types.User{Name: "user", Emails: []new_types.Email{{Email: "user1@example.com"}}}
	user2 := new_types.User{Name: "user", Emails: []new_types.Email{{Email: "user2@example.com"}}}
	new_types.TestDB.Save(&user1)
	new_types.TestDB.Save(&user2)

	var user3, user4 new_types.User
	new_types.TestDB.Raw("select * from users WHERE name = ?", "user").First(&user3)
	if user3.Id != user1.Id {
		t.Errorf("Find first record with raw")
	}

	new_types.TestDB.Raw("select * from users WHERE name = ?", "user").Last(&user4)
	if user4.Id != user2.Id {
		t.Errorf("Find last record with raw")
	}
}

func UIntPrimaryKey(t *testing.T) {
	var animal new_types.Animal
	new_types.TestDB.First(&animal, uint64(1))
	if animal.Counter != 1 {
		t.Errorf("Fetch a record from with a non-int primary key should work, but failed")
	}

	new_types.TestDB.Model(new_types.Animal{}).Where(new_types.Animal{Counter: uint64(2)}).Scan(&animal)
	if animal.Counter != 2 {
		t.Errorf("Fetch a record from with a non-int primary key should work, but failed")
	}
}

func StringPrimaryKeyForNumericValueStartingWithZero(t *testing.T) {
	type AddressByZipCode struct {
		ZipCode string `gorm:"primary_key"`
		Address string
	}

	new_types.TestDB.AutoMigrate(&AddressByZipCode{})
	new_types.TestDB.Create(&AddressByZipCode{ZipCode: "00501", Address: "Holtsville"})

	var address AddressByZipCode
	new_types.TestDB.First(&address, "00501")
	if address.ZipCode != "00501" {
		t.Errorf("Fetch a record from with a string primary key for a numeric value starting with zero should work, but failed")
	}
}

func FindAsSliceOfPointers(t *testing.T) {
	new_types.TestDB.Save(&new_types.User{Name: "user"})

	var users []new_types.User
	new_types.TestDB.Find(&users)

	var userPointers []*new_types.User
	new_types.TestDB.Find(&userPointers)

	if len(users) == 0 || len(users) != len(userPointers) {
		t.Errorf("Find slice of pointers")
	}
}

func SearchWithPlainSQL(t *testing.T) {
	user1 := new_types.User{Name: "PlainSqlUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "PlainSqlUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "PlainSqlUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)
	scopedb := new_types.TestDB.Where("name LIKE ?", "%PlainSqlUser%")

	if new_types.TestDB.Where("name = ?", user1.Name).First(&new_types.User{}).RecordNotFound() {
		t.Errorf("Search with plain SQL")
	}

	if new_types.TestDB.Where("name LIKE ?", "%"+user1.Name+"%").First(&new_types.User{}).RecordNotFound() {
		t.Errorf("Search with plan SQL (regexp)")
	}

	var users []new_types.User
	new_types.TestDB.Find(&users, "name LIKE ? and age > ?", "%PlainSqlUser%", 1)
	if len(users) != 2 {
		t.Errorf("Should found 2 users that age > 1, but got %v", len(users))
	}

	new_types.TestDB.Where("name LIKE ?", "%PlainSqlUser%").Where("age >= ?", 1).Find(&users)
	if len(users) != 3 {
		t.Errorf("Should found 3 users that age >= 1, but got %v", len(users))
	}

	scopedb.Where("age <> ?", 20).Find(&users)
	if len(users) != 2 {
		t.Errorf("Should found 2 users age != 20, but got %v", len(users))
	}

	scopedb.Where("birthday > ?", newparseTime("2000-1-1")).Find(&users)
	if len(users) != 2 {
		t.Errorf("Should found 2 users's birthday > 2000-1-1, but got %v", len(users))
	}

	scopedb.Where("birthday > ?", "2002-10-10").Find(&users)
	if len(users) != 2 {
		t.Errorf("Should found 2 users's birthday >= 2002-10-10, but got %v", len(users))
	}

	scopedb.Where("birthday >= ?", "2010-1-1").Where("birthday < ?", "2020-1-1").Find(&users)
	if len(users) != 1 {
		t.Errorf("Should found 1 users's birthday < 2020-1-1 and >= 2010-1-1, but got %v", len(users))
	}

	new_types.TestDB.Where("name in (?)", []string{user1.Name, user2.Name}).Find(&users)
	if len(users) != 2 {
		t.Errorf("Should found 2 users, but got %v", len(users))
	}

	new_types.TestDB.Where("id in (?)", []int64{user1.Id, user2.Id, user3.Id}).Find(&users)
	if len(users) != 3 {
		t.Errorf("Should found 3 users, but got %v", len(users))
	}

	new_types.TestDB.Where("id in (?)", user1.Id).Find(&users)
	if len(users) != 1 {
		t.Errorf("Should found 1 users, but got %v", len(users))
	}

	if err := new_types.TestDB.Where("id IN (?)", []string{}).Find(&users).Error; err != nil {
		t.Error("no error should happen when query with empty slice, but got: ", err)
	}

	if err := new_types.TestDB.Not("id IN (?)", []string{}).Find(&users).Error; err != nil {
		t.Error("no error should happen when query with empty slice, but got: ", err)
	}

	if new_types.TestDB.Where("name = ?", "none existing").Find(&[]new_types.User{}).RecordNotFound() {
		t.Errorf("Should not get RecordNotFound error when looking for none existing records")
	}
}

func SearchWithStruct(t *testing.T) {
	user1 := new_types.User{Name: "StructSearchUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "StructSearchUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "StructSearchUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	if new_types.TestDB.Where(user1.Id).First(&new_types.User{}).RecordNotFound() {
		t.Errorf("Search with primary key")
	}

	if new_types.TestDB.First(&new_types.User{}, user1.Id).RecordNotFound() {
		t.Errorf("Search with primary key as inline condition")
	}

	if new_types.TestDB.First(&new_types.User{}, fmt.Sprintf("%v", user1.Id)).RecordNotFound() {
		t.Errorf("Search with primary key as inline condition")
	}

	var users []new_types.User
	new_types.TestDB.Where([]int64{user1.Id, user2.Id, user3.Id}).Find(&users)
	if len(users) != 3 {
		t.Errorf("Should found 3 users when search with primary keys, but got %v", len(users))
	}

	var user new_types.User
	new_types.TestDB.First(&user, &new_types.User{Name: user1.Name})
	if user.Id == 0 || user.Name != user1.Name {
		t.Errorf("Search first record with inline pointer of struct")
	}

	new_types.TestDB.First(&user, new_types.User{Name: user1.Name})
	if user.Id == 0 || user.Name != user.Name {
		t.Errorf("Search first record with inline struct")
	}

	new_types.TestDB.Where(&new_types.User{Name: user1.Name}).First(&user)
	if user.Id == 0 || user.Name != user1.Name {
		t.Errorf("Search first record with where struct")
	}

	new_types.TestDB.Find(&users, &new_types.User{Name: user2.Name})
	if len(users) != 1 {
		t.Errorf("Search all records with inline struct")
	}
}

func SearchWithMap(t *testing.T) {
	companyID := 1
	user1 := new_types.User{Name: "MapSearchUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "MapSearchUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "MapSearchUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	user4 := new_types.User{Name: "MapSearchUser4", Age: 30, Birthday: newparseTime("2020-1-1"), CompanyID: &companyID}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3).Save(&user4)

	var user new_types.User
	new_types.TestDB.First(&user, map[string]interface{}{"name": user1.Name})
	if user.Id == 0 || user.Name != user1.Name {
		t.Errorf("Search first record with inline map")
	}

	user = new_types.User{}
	new_types.TestDB.Where(map[string]interface{}{"name": user2.Name}).First(&user)
	if user.Id == 0 || user.Name != user2.Name {
		t.Errorf("Search first record with where map")
	}

	var users []new_types.User
	new_types.TestDB.Where(map[string]interface{}{"name": user3.Name}).Find(&users)
	if len(users) != 1 {
		t.Errorf("Search all records with inline map")
	}

	new_types.TestDB.Find(&users, map[string]interface{}{"name": user3.Name})
	if len(users) != 1 {
		t.Errorf("Search all records with inline map")
	}

	new_types.TestDB.Find(&users, map[string]interface{}{"name": user4.Name, "company_id": nil})
	if len(users) != 0 {
		t.Errorf("Search all records with inline map containing null value finding 0 records")
	}

	new_types.TestDB.Find(&users, map[string]interface{}{"name": user1.Name, "company_id": nil})
	if len(users) != 1 {
		t.Errorf("Search all records with inline map containing null value finding 1 record")
	}

	new_types.TestDB.Find(&users, map[string]interface{}{"name": user4.Name, "company_id": companyID})
	if len(users) != 1 {
		t.Errorf("Search all records with inline multiple value map")
	}
}

func SearchWithEmptyChain(t *testing.T) {
	user1 := new_types.User{Name: "ChainSearchUser1", Age: 1, Birthday: newparseTime("2000-1-1")}
	user2 := new_types.User{Name: "ChainearchUser2", Age: 10, Birthday: newparseTime("2010-1-1")}
	user3 := new_types.User{Name: "ChainearchUser3", Age: 20, Birthday: newparseTime("2020-1-1")}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	if new_types.TestDB.Where("").Where("").First(&new_types.User{}).Error != nil {
		t.Errorf("Should not raise any error if searching with empty strings")
	}

	if new_types.TestDB.Where(&new_types.User{}).Where("name = ?", user1.Name).First(&new_types.User{}).Error != nil {
		t.Errorf("Should not raise any error if searching with empty struct")
	}

	if new_types.TestDB.Where(map[string]interface{}{}).Where("name = ?", user1.Name).First(&new_types.User{}).Error != nil {
		t.Errorf("Should not raise any error if searching with empty map")
	}
}

func Select(t *testing.T) {
	user1 := new_types.User{Name: "SelectUser1"}
	new_types.TestDB.Save(&user1)

	var user new_types.User
	new_types.TestDB.Where("name = ?", user1.Name).Select("name").Find(&user)
	if user.Id != 0 {
		t.Errorf("Should not have ID because only selected name, %+v", user.Id)
	}

	if user.Name != user1.Name {
		t.Errorf("Should have user Name when selected it")
	}
}

func OrderAndPluck(t *testing.T) {
	user1 := new_types.User{Name: "OrderPluckUser1", Age: 1}
	user2 := new_types.User{Name: "OrderPluckUser2", Age: 10}
	user3 := new_types.User{Name: "OrderPluckUser3", Age: 20}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)
	scopedb := new_types.TestDB.Model(&new_types.User{}).Where("name like ?", "%OrderPluckUser%")

	var user new_types.User
	scopedb.Order(newGorm.SqlExpr("name = ? DESC", "OrderPluckUser2")).First(&user)
	if user.Name != "OrderPluckUser2" {
		t.Errorf("Order with sql expression")
	}

	var ages []int64
	scopedb.Order("age desc").Pluck("age", &ages)
	if ages[0] != 20 {
		t.Errorf("The first age should be 20 when order with age desc")
	}

	var ages1, ages2 []int64
	scopedb.Order("age desc").Pluck("age", &ages1).Pluck("age", &ages2)
	if !reflect.DeepEqual(ages1, ages2) {
		t.Errorf("The first order is the primary order")
	}

	var ages3, ages4 []int64
	scopedb.Model(&new_types.User{}).Order("age desc").Pluck("age", &ages3).Order("age", true).Pluck("age", &ages4)
	if reflect.DeepEqual(ages3, ages4) {
		t.Errorf("Reorder should work")
	}

	var names []string
	var ages5 []int64
	scopedb.Model(new_types.User{}).Order("name").Order("age desc").Pluck("age", &ages5).Pluck("name", &names)
	if names != nil && ages5 != nil {
		if !(names[0] == user1.Name && names[1] == user2.Name && names[2] == user3.Name && ages5[2] == 20) {
			t.Errorf("Order with multiple orders")
		}
	} else {
		t.Errorf("Order with multiple orders")
	}

	new_types.TestDB.Model(new_types.User{}).Select("name, age").Find(&[]new_types.User{})
}

func Limit(t *testing.T) {
	user1 := new_types.User{Name: "LimitUser1", Age: 1}
	user2 := new_types.User{Name: "LimitUser2", Age: 10}
	user3 := new_types.User{Name: "LimitUser3", Age: 20}
	user4 := new_types.User{Name: "LimitUser4", Age: 10}
	user5 := new_types.User{Name: "LimitUser5", Age: 20}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3).Save(&user4).Save(&user5)

	var users1, users2, users3 []new_types.User
	new_types.TestDB.Order("age desc").Limit(3).Find(&users1).Limit(5).Find(&users2).Limit(-1).Find(&users3)

	if len(users1) != 3 || len(users2) != 5 || len(users3) <= 5 {
		t.Errorf("Limit should works")
	}
}

func Offset(t *testing.T) {
	for i := 0; i < 20; i++ {
		new_types.TestDB.Save(&new_types.User{Name: fmt.Sprintf("OffsetUser%v", i)})
	}
	var users1, users2, users3, users4 []new_types.User
	new_types.TestDB.Limit(100).Order("age desc").Find(&users1).Offset(3).Find(&users2).Offset(5).Find(&users3).Offset(-1).Find(&users4)

	if (len(users1) != len(users4)) || (len(users1)-len(users2) != 3) || (len(users1)-len(users3) != 5) {
		t.Errorf("Offset should work")
	}
}

func Or(t *testing.T) {
	user1 := new_types.User{Name: "OrUser1", Age: 1}
	user2 := new_types.User{Name: "OrUser2", Age: 10}
	user3 := new_types.User{Name: "OrUser3", Age: 20}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	var users []new_types.User
	new_types.TestDB.Where("name = ?", user1.Name).Or("name = ?", user2.Name).Find(&users)
	if len(users) != 2 {
		t.Errorf("Find users with or")
	}
}

func Count(t *testing.T) {
	user1 := new_types.User{Name: "CountUser1", Age: 1}
	user2 := new_types.User{Name: "CountUser2", Age: 10}
	user3 := new_types.User{Name: "CountUser3", Age: 20}

	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)
	var count, count1, count2 int64
	var users []new_types.User

	if err := new_types.TestDB.Where("name = ?", user1.Name).Or("name = ?", user3.Name).Find(&users).Count(&count).Error; err != nil {
		t.Errorf(fmt.Sprintf("Count should work, but got err %v", err))
	}

	if count != int64(len(users)) {
		t.Errorf("Count() method should get correct value")
	}

	new_types.TestDB.Model(&new_types.User{}).Where("name = ?", user1.Name).Count(&count1).Or("name in (?)", []string{user2.Name, user3.Name}).Count(&count2)
	if count1 != 1 || count2 != 3 {
		t.Errorf("Multiple count in chain")
	}
}

func Not(t *testing.T) {
	new_types.TestDB.Unscoped().Delete(new_types.User{})
	new_types.TestDB.Create(newGetPreparedUser("user1", "not"))
	new_types.TestDB.Create(newGetPreparedUser("user2", "not"))
	new_types.TestDB.Create(newGetPreparedUser("user3", "not"))

	user4 := newGetPreparedUser("user4", "not")
	user4.Company = new_types.Company{}
	new_types.TestDB.Create(user4)

	var users1, users2, users3, users4, users5, users6, users7, users8, users9 []new_types.User
	rowsAffected := new_types.TestDB.Where("role = ?", "not").Find(&users1).RowsAffected
	if rowsAffected != 4 {
		t.Errorf("should find 4 `not` users")
	} else {
	}
	new_types.TestDB.Not(users1[0].Id).Find(&users2)

	if len(users1)-len(users2) != 1 {
		t.Errorf("Should ignore the first users with Not")
	} else {
	}

	new_types.TestDB.Not([]int{}).Find(&users3)
	if len(users1)-len(users3) != 0 {
		t.Errorf("Should find all users with a blank condition")
	} else {
	}

	var name3Count int64
	new_types.TestDB.Table("users").Where("name = ?", "user3").Count(&name3Count)
	new_types.TestDB.Not("name", "user3").Find(&users4)
	if len(users1)-len(users4) != int(name3Count) {
		t.Errorf("Should find all users's name not equal 3")
	} else {
	}

	new_types.TestDB.Not("name = ?", "user3").Find(&users4)
	if len(users1)-len(users4) != int(name3Count) {
		t.Errorf("Should find all users's name not equal 3")
	} else {
	}

	new_types.TestDB.Not("name <> ?", "user3").Find(&users4)
	if len(users4) != int(name3Count) {
		t.Errorf("Should find all users's name not equal 3")
	} else {
	}

	new_types.TestDB.Not(new_types.User{Name: "user3"}).Find(&users5)

	if len(users1)-len(users5) != int(name3Count) {
		t.Errorf("Should find all users's name not equal 3")
	} else {
	}

	new_types.TestDB.Not(map[string]interface{}{"name": "user3"}).Find(&users6)
	if len(users1)-len(users6) != int(name3Count) {
		t.Errorf("Should find all users's name not equal 3")
	} else {
	}

	new_types.TestDB.Not(map[string]interface{}{"name": "user3", "company_id": nil}).Find(&users7)
	if len(users1)-len(users7) != 2 { // not user3 or user4
		t.Errorf("Should find all user's name not equal to 3 who do not have company id")
	} else {
	}

	new_types.TestDB.Not("name", []string{"user3"}).Find(&users8)
	if len(users1)-len(users8) != int(name3Count) {
		t.Errorf("Should find all users's name not equal 3")
	} else {
	}

	var name2Count int64
	new_types.TestDB.Table("users").Where("name = ?", "user2").Count(&name2Count)
	new_types.TestDB.Not("name", []string{"user3", "user2"}).Find(&users9)
	if len(users1)-len(users9) != (int(name3Count) + int(name2Count)) {
		t.Errorf("Should find all users's name not equal 3")
	} else {
	}
}

func FillSmallerStruct(t *testing.T) {
	user1 := new_types.User{Name: "SmallerUser", Age: 100}
	new_types.TestDB.Save(&user1)
	type SimpleUser struct {
		Name      string
		Id        int64
		UpdatedAt time.Time
		CreatedAt time.Time
	}

	var simpleUser SimpleUser
	new_types.TestDB.Table("users").Where("name = ?", user1.Name).First(&simpleUser)

	if simpleUser.Id == 0 || simpleUser.Name == "" {
		t.Errorf("Should fill data correctly into smaller struct")
	}
}

func FindOrInitialize(t *testing.T) {
	var user1, user2, user3, user4, user5, user6 new_types.User
	new_types.TestDB.Where(&new_types.User{Name: "find or init", Age: 33}).FirstOrInit(&user1)
	if user1.Name != "find or init" || user1.Id != 0 || user1.Age != 33 {
		t.Errorf("user should be initialized with search value")
	}

	new_types.TestDB.Where(new_types.User{Name: "find or init", Age: 33}).FirstOrInit(&user2)
	if user2.Name != "find or init" || user2.Id != 0 || user2.Age != 33 {
		t.Errorf("user should be initialized with search value")
	}

	new_types.TestDB.FirstOrInit(&user3, map[string]interface{}{"name": "find or init 2"})
	if user3.Name != "find or init 2" || user3.Id != 0 {
		t.Errorf("user should be initialized with inline search value")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or init"}).Attrs(new_types.User{Age: 44}).FirstOrInit(&user4)
	if user4.Name != "find or init" || user4.Id != 0 || user4.Age != 44 {
		t.Errorf("user should be initialized with search value and attrs")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or init"}).Assign("age", 44).FirstOrInit(&user4)
	if user4.Name != "find or init" || user4.Id != 0 || user4.Age != 44 {
		t.Errorf("user should be initialized with search value and assign attrs")
	}

	new_types.TestDB.Save(&new_types.User{Name: "find or init", Age: 33})
	new_types.TestDB.Where(&new_types.User{Name: "find or init"}).Attrs("age", 44).FirstOrInit(&user5)
	if user5.Name != "find or init" || user5.Id == 0 || user5.Age != 33 {
		t.Errorf("user should be found and not initialized by Attrs")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or init", Age: 33}).FirstOrInit(&user6)
	if user6.Name != "find or init" || user6.Id == 0 || user6.Age != 33 {
		t.Errorf("user should be found with FirstOrInit")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or init"}).Assign(new_types.User{Age: 44}).FirstOrInit(&user6)
	if user6.Name != "find or init" || user6.Id == 0 || user6.Age != 44 {
		t.Errorf("user should be found and updated with assigned attrs")
	}
}

func FindOrCreate(t *testing.T) {
	var user1, user2, user3, user4, user5, user6, user7, user8 new_types.User
	new_types.TestDB.Where(&new_types.User{Name: "find or create", Age: 33}).FirstOrCreate(&user1)
	if user1.Name != "find or create" || user1.Id == 0 || user1.Age != 33 {
		t.Errorf("user should be created with search value")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or create", Age: 33}).FirstOrCreate(&user2)
	if user1.Id != user2.Id || user2.Name != "find or create" || user2.Id == 0 || user2.Age != 33 {
		t.Errorf("user should be created with search value")
	}

	new_types.TestDB.FirstOrCreate(&user3, map[string]interface{}{"name": "find or create 2"})
	if user3.Name != "find or create 2" || user3.Id == 0 {
		t.Errorf("user should be created with inline search value")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or create 3"}).Attrs("age", 44).FirstOrCreate(&user4)
	if user4.Name != "find or create 3" || user4.Id == 0 || user4.Age != 44 {
		t.Errorf("user should be created with search value and attrs")
	}

	updatedAt1 := user4.UpdatedAt
	new_types.TestDB.Where(&new_types.User{Name: "find or create 3"}).Assign("age", 55).FirstOrCreate(&user4)
	if updatedAt1.Format(time.RFC3339Nano) == user4.UpdatedAt.Format(time.RFC3339Nano) {
		t.Errorf("UpdateAt should be changed when update values with assign")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or create 4"}).Assign(new_types.User{Age: 44}).FirstOrCreate(&user4)
	if user4.Name != "find or create 4" || user4.Id == 0 || user4.Age != 44 {
		t.Errorf("user should be created with search value and assigned attrs")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or create"}).Attrs("age", 44).FirstOrInit(&user5)
	if user5.Name != "find or create" || user5.Id == 0 || user5.Age != 33 {
		t.Errorf("user should be found and not initialized by Attrs")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or create"}).Assign(new_types.User{Age: 44}).FirstOrCreate(&user6)
	if user6.Name != "find or create" || user6.Id == 0 || user6.Age != 44 {
		t.Errorf("user should be found and updated with assigned attrs")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or create"}).Find(&user7)
	if user7.Name != "find or create" || user7.Id == 0 || user7.Age != 44 {
		t.Errorf("user should be found and updated with assigned attrs")
	}

	new_types.TestDB.Where(&new_types.User{Name: "find or create embedded struct"}).Assign(new_types.User{Age: 44, CreditCard: new_types.CreditCard{Number: "1231231231"}, Emails: []new_types.Email{{Email: "jinzhu@assign_embedded_struct.com"}, {Email: "jinzhu-2@assign_embedded_struct.com"}}}).FirstOrCreate(&user8)
	if new_types.TestDB.Where("email = ?", "jinzhu-2@assign_embedded_struct.com").First(&new_types.Email{}).RecordNotFound() {
		t.Errorf("embedded struct email should be saved")
	}

	if new_types.TestDB.Where("email = ?", "1231231231").First(&new_types.CreditCard{}).RecordNotFound() {
		t.Errorf("embedded struct credit card should be saved")
	}
}

func SelectWithEscapedFieldName(t *testing.T) {
	user1 := new_types.User{Name: "EscapedFieldNameUser", Age: 1}
	user2 := new_types.User{Name: "EscapedFieldNameUser", Age: 10}
	user3 := new_types.User{Name: "EscapedFieldNameUser", Age: 20}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	var names []string
	new_types.TestDB.Model(new_types.User{}).Where(&new_types.User{Name: "EscapedFieldNameUser"}).Pluck("\"name\"", &names)

	if len(names) != 3 {
		t.Errorf("Expected 3 name, but got: %d", len(names))
	}
}

func SelectWithVariables(t *testing.T) {
	new_types.TestDB.Save(&new_types.User{Name: "jinzhu"})

	rows, _ := new_types.TestDB.Table("users").Select("? as fake", newGorm.SqlExpr("name")).Rows()

	if !rows.Next() {
		t.Errorf("Should have returned at least one row")
	} else {
		columns, _ := rows.Columns()
		if !reflect.DeepEqual(columns, []string{"fake"}) {
			t.Errorf("Should only contains one column")
		}
	}

	rows.Close()
}
