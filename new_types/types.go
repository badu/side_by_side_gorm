package new_types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	newGorm "github.com/badu/gorm"
	_ "github.com/badu/gorm/dialects/sqlite"
	"reflect"
	"time"
)

var (
	TestDB      *newGorm.DBCon
	TestDBErr   error

)

type (
	ElementWithIgnoredField struct {
		Id           int64
		Value        string
		IgnoredField int64 `sql:"-"`
	}

	RecordWithSlice struct {
		ID      uint64
		Strings ExampleStringSlice `sql:"type:text"`
		Structs ExampleStructSlice `sql:"type:text"`
	}

	ExampleStringSlice []string

	ExampleStruct struct {
		Name  string
		Value string
	}

	ExampleStructSlice []ExampleStruct

	BasePost struct {
		Id    int64
		Title string
		URL   string
	}

	Author struct {
		Name  string
		Email string
	}

	HNPost struct {
		BasePost
		Author  `gorm:"embedded_prefix:user_"` // Embedded struct
		Upvotes int32
	}

	EngadgetPost struct {
		BasePost BasePost `gorm:"embedded"`
		Author   Author   `gorm:"embedded;embedded_prefix:author_"` // Embedded struct
		ImageUrl string
	}

	LevelA1 struct {
		ID    uint
		Value string
	}

	LevelA2 struct {
		ID       uint
		Value    string
		LevelA3s []*LevelA3
	}

	LevelA3 struct {
		ID        uint
		Value     string
		LevelA1ID sql.NullInt64
		LevelA1   *LevelA1
		LevelA2ID sql.NullInt64
		LevelA2   *LevelA2
	}

	LevelB1 struct {
		ID       uint
		Value    string
		LevelB3s []*LevelB3
	}

	LevelB2 struct {
		ID    uint
		Value string
	}

	LevelB3 struct {
		ID        uint
		Value     string
		LevelB1ID sql.NullInt64
		LevelB1   *LevelB1
		LevelB2s  []*LevelB2 `gorm:"many2many:levelb1_levelb3_levelb2s"`
	}

	LevelC1 struct {
		ID        uint
		Value     string
		LevelC2ID uint
	}

	LevelC2 struct {
		ID      uint
		Value   string
		LevelC1 LevelC1
	}

	LevelC3 struct {
		ID        uint
		Value     string
		LevelC2ID uint
		LevelC2   LevelC2
	}

	Cat struct {
		Id   int
		Name string
		Toy  Toy `gorm:"polymorphic:Owner;"`
	}

	Dog struct {
		Id   int
		Name string
		Toys []Toy `gorm:"polymorphic:Owner;"`
	}

	Hamster struct {
		Id           int
		Name         string
		PreferredToy Toy `gorm:"polymorphic:Owner;polymorphic_value:hamster_preferred"`
		OtherToy     Toy `gorm:"polymorphic:Owner;polymorphic_value:hamster_other"`
	}

	Toy struct {
		Id        int
		Name      string
		OwnerId   int
		OwnerType string
	}

	PointerStruct struct {
		ID   int64
		Name *string
		Num  *int
	}

	NormalStruct struct {
		ID   int64
		Name string
		Num  int
	}

	NotSoLongTableName struct {
		Id                int64
		ReallyLongThingID int64
		ReallyLongThing   ReallyLongTableNameToTestMySQLNameLengthLimit
	}

	ReallyLongTableNameToTestMySQLNameLengthLimit struct {
		Id int64
	}

	ReallyLongThingThatReferencesShort struct {
		Id      int64
		ShortID int64
		Short   Short
	}

	Short struct {
		Id int64
	}

	Num int64

	User struct {
		Id                int64
		Age               int64
		UserNum           Num
		Name              string `sql:"size:255"`
		Email             string
		Birthday          *time.Time    // Time
		CreatedAt         time.Time     // CreatedAt: Time of record is created, will be insert automatically
		UpdatedAt         time.Time     // UpdatedAt: Time of record is updated, will be updated automatically
		Emails            []Email       // Embedded structs
		BillingAddress    Address       // Embedded struct
		BillingAddressID  sql.NullInt64 // Embedded struct's foreign key
		ShippingAddress   Address       // Embedded struct
		ShippingAddressId int64         // Embedded struct's foreign key
		CreditCard        CreditCard
		Latitude          float64
		Languages         []Language `gorm:"many2many:user_languages;"`
		CompanyID         *int
		Company           Company
		Role
		PasswordHash      []byte
		Sequence          uint                  `gorm:"AUTO_INCREMENT"`
		IgnoreMe          int64                 `sql:"-"`
		IgnoreStringSlice []string              `sql:"-"`
		Ignored           struct{ Name string } `sql:"-"`
		IgnoredPointer    *User                 `sql:"-"`
	}

	CreditCard struct {
		ID        int8
		Number    string
		UserId    sql.NullInt64
		CreatedAt time.Time `sql:"not null"`
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	Blog struct {
		ID         uint   `gorm:"primary_key"`
		Locale     string `gorm:"primary_key"`
		Subject    string
		Body       string
		Tags       []Tag `gorm:"many2many:blog_tags;"`
		SharedTags []Tag `gorm:"many2many:shared_blog_tags;ForeignKey:id;AssociationForeignKey:id"`
		LocaleTags []Tag `gorm:"many2many:locale_blog_tags;ForeignKey:id,locale;AssociationForeignKey:id"`
	}

	Tag struct {
		ID     uint   `gorm:"primary_key"`
		Locale string `gorm:"primary_key"`
		Value  string
		Blogs  []*Blog `gorm:"many2many:blogs_tags"`
	}

	Email struct {
		Id        int16
		UserId    int
		Email     string `sql:"type:varchar(100);"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Address struct {
		ID        int
		Address1  string
		Address2  string
		Post      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	Language struct {
		newGorm.Model
		Name  string
		Users []User `gorm:"many2many:user_languages;"`
	}

	Product struct {
		Id                    int64
		Code                  string
		Price                 int64
		CreatedAt             time.Time
		UpdatedAt             time.Time
		AfterFindCallTimes    int64
		BeforeCreateCallTimes int64
		AfterCreateCallTimes  int64
		BeforeUpdateCallTimes int64
		AfterUpdateCallTimes  int64
		BeforeSaveCallTimes   int64
		AfterSaveCallTimes    int64
		BeforeDeleteCallTimes int64
		AfterDeleteCallTimes  int64
	}

	Company struct {
		Id    int64
		Name  string
		Owner *User `sql:"-"`
	}

	Role struct {
		Name string `gorm:"size:256"`
	}

	Animal struct {
		Counter    uint64    `gorm:"primary_key:yes"`
		Name       string    `sql:"DEFAULT:'galeone'"`
		From       string    //test reserved sql keyword as field name
		Age        time.Time `sql:"DEFAULT:current_timestamp"`
		unexported string    // unexported value
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	JoinTable struct {
		From uint64
		To   uint64
		Time time.Time `sql:"default: null"`
	}

	Post struct {
		Id             int64
		CategoryId     sql.NullInt64
		MainCategoryId int64
		Title          string
		Body           string
		Comments       []*Comment
		Category       Category
		MainCategory   Category
	}

	Category struct {
		newGorm.Model
		Name string

		Categories []Category
		CategoryID *uint
	}

	Comment struct {
		newGorm.Model
		PostId  int64
		Content string
		Post    Post
	}

	// Scanner
	NullValue struct {
		Id      int64
		Name    sql.NullString  `sql:"not null"`
		Gender  *sql.NullString `sql:"not null"`
		Age     sql.NullInt64
		Male    sql.NullBool
		Height  sql.NullFloat64
		AddedAt NullTime
	}

	NullTime struct {
		Time  time.Time
		Valid bool
	}

	BigEmail struct {
		Id           int64
		UserId       int64
		Email        string     `sql:"index:idx_email_agent"`
		UserAgent    string     `sql:"index:idx_email_agent"`
		RegisteredAt *time.Time `sql:"unique_index"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	MultipleIndexes struct {
		ID     int64
		UserID int64  `sql:"unique_index:uix_multipleindexes_user_name,uix_multipleindexes_user_email;index:idx_multipleindexes_user_other"`
		Name   string `sql:"unique_index:uix_multipleindexes_user_name"`
		Email  string `sql:"unique_index:,uix_multipleindexes_user_email"`
		Other  string `sql:"index:,idx_multipleindexes_user_other"`
	}

	Person struct {
		Id        int
		Name      string
		Addresses []*Address `gorm:"many2many:person_addresses;"`
	}

	PersonAddress struct {
		newGorm.JoinTableHandler
		PersonID  int
		AddressID int
		DeletedAt *time.Time
		CreatedAt time.Time
	}

	CalculateField struct {
		newGorm.Model
		Name     string
		Children []CalculateFieldChild
		Category CalculateFieldCategory
		EmbeddedField
	}

	EmbeddedField struct {
		EmbeddedName string `sql:"NOT NULL;DEFAULT:'hello'"`
	}

	CalculateFieldChild struct {
		newGorm.Model
		CalculateFieldID uint
		Name             string
	}

	CalculateFieldCategory struct {
		newGorm.Model
		CalculateFieldID uint
		Name             string
	}

	CustomizeColumn struct {
		ID   int64      `gorm:"column:mapped_id; primary_key:yes"`
		Name string     `gorm:"column:mapped_name"`
		Date *time.Time `gorm:"column:mapped_time"`
	}

	// Make sure an ignored field does not interfere with another field's custom
	// column name that matches the ignored field.
	CustomColumnAndIgnoredFieldClash struct {
		Body    string `sql:"-"`
		RawBody string `gorm:"column:body"`
	}

	CustomizePerson struct {
		IdPerson string             `gorm:"column:idPerson;primary_key:true"`
		Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;associationforeignkey:idAccount;foreignkey:idPerson"`
	}

	CustomizeAccount struct {
		IdAccount string `gorm:"column:idAccount;primary_key:true"`
		Name      string
	}

	CustomizeUser struct {
		newGorm.Model
		Email string `sql:"column:email_address"`
	}

	CustomizeInvitation struct {
		newGorm.Model
		Address string         `sql:"column:invitation"`
		Person  *CustomizeUser `gorm:"foreignkey:Email;associationforeignkey:invitation"`
	}

	PromotionDiscount struct {
		newGorm.Model
		Name     string
		Coupons  []*PromotionCoupon `gorm:"ForeignKey:discount_id"`
		Rule     *PromotionRule     `gorm:"ForeignKey:discount_id"`
		Benefits []PromotionBenefit `gorm:"ForeignKey:promotion_id"`
	}

	PromotionBenefit struct {
		newGorm.Model
		Name        string
		PromotionID uint
		Discount    PromotionDiscount `gorm:"ForeignKey:promotion_id"`
	}

	PromotionCoupon struct {
		newGorm.Model
		Code       string
		DiscountID uint
		Discount   PromotionDiscount
	}

	PromotionRule struct {
		newGorm.Model
		Name       string
		Begin      *time.Time
		End        *time.Time
		DiscountID uint
		Discount   *PromotionDiscount
	}

	Order struct {
	}

	Cart struct {
	}
)

func (e ElementWithIgnoredField) TableName() string {
	return "element_with_ignored_field"
}

func (s *Product) BeforeCreate() (err error) {
	if s.Code == "Invalid" {
		err = errors.New("BeforeCreate invalid product")
	}
	s.BeforeCreateCallTimes = s.BeforeCreateCallTimes + 1
	return
}

func (s *Product) BeforeUpdate() (err error) {
	if s.Code == "dont_update" {
		err = errors.New("BeforeUpdate can't update")
	}
	s.BeforeUpdateCallTimes = s.BeforeUpdateCallTimes + 1
	return
}

func (s *Product) BeforeSave() (err error) {
	if s.Code == "dont_save" {
		err = errors.New("BeforeSave can't save")
	}
	s.BeforeSaveCallTimes = s.BeforeSaveCallTimes + 1
	return
}

func (s *Product) AfterFind() {
	s.AfterFindCallTimes = s.AfterFindCallTimes + 1
}

func (s *Product) AfterCreate(tx *newGorm.DBCon) {
	tx.Model(s).UpdateColumn(Product{AfterCreateCallTimes: s.AfterCreateCallTimes + 1})
}

func (s *Product) AfterUpdate() {
	s.AfterUpdateCallTimes = s.AfterUpdateCallTimes + 1
}

func (s *Product) AfterSave() (err error) {
	if s.Code == "after_save_error" {
		err = errors.New("AfterSave can't save")
	}
	s.AfterSaveCallTimes = s.AfterSaveCallTimes + 1
	return
}

func (s *Product) BeforeDelete() (err error) {
	if s.Code == "dont_delete" {
		err = errors.New("BeforeDelete can't delete")
	}
	s.BeforeDeleteCallTimes = s.BeforeDeleteCallTimes + 1
	return
}

func (s *Product) AfterDelete() (err error) {
	if s.Code == "after_delete_error" {
		err = errors.New("AfterDelete can't delete")
	}
	s.AfterDeleteCallTimes = s.AfterDeleteCallTimes + 1
	return
}

func (s *Product) GetCallTimes() []int64 {
	return []int64{s.BeforeCreateCallTimes, s.BeforeSaveCallTimes, s.BeforeUpdateCallTimes, s.AfterCreateCallTimes, s.AfterSaveCallTimes, s.AfterUpdateCallTimes, s.BeforeDeleteCallTimes, s.AfterDeleteCallTimes, s.AfterFindCallTimes}
}

func (l ExampleStringSlice) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *ExampleStringSlice) Scan(input interface{}) error {
	switch value := input.(type) {
	case string:
		return json.Unmarshal([]byte(value), l)
	case []byte:
		return json.Unmarshal(value, l)
	default:
		return errors.New("not supported")
	}
}

func (l ExampleStructSlice) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *ExampleStructSlice) Scan(input interface{}) error {
	switch value := input.(type) {
	case string:
		return json.Unmarshal([]byte(value), l)
	case []byte:
		return json.Unmarshal(value, l)
	default:
		return errors.New("not supported")
	}
}

func (b BigEmail) TableName() string {
	return "emails"
}

func (c Cart) TableName() string {
	return "shopping_cart"
}

func (p Person) String() string {
	optionals := fmt.Sprintf("%q:%d,%q:%q",
		"id", p.Id,
		"type", p.Name)
	if len(p.Addresses) > 0 {
		optionals += fmt.Sprintf(",%q:%d", "addresses", len(p.Addresses))
	}
	return fmt.Sprint(optionals)
}

func (*PersonAddress) Add(handler newGorm.JoinTableHandlerInterface, db *newGorm.DBCon, foreignValue interface{}, associationValue interface{}) error {
	return db.Where(map[string]interface{}{
		"person_id":  db.NewScope(foreignValue).PrimaryKeyValue(),
		"address_id": db.NewScope(associationValue).PrimaryKeyValue(),
	}).Assign(map[string]interface{}{
		"person_id":  foreignValue,
		"address_id": associationValue,
		"deleted_at": newGorm.SqlExpr("NULL"),
	}).FirstOrCreate(&PersonAddress{}).Error
}

func (*PersonAddress) Delete(handler newGorm.JoinTableHandlerInterface, db *newGorm.DBCon) error {
	return db.Delete(&PersonAddress{}).Error
}

func (pa *PersonAddress) JoinWith(handler newGorm.JoinTableHandlerInterface, db *newGorm.DBCon, source interface{}) *newGorm.DBCon {
	table := pa.Table(db)
	return db.Joins("INNER JOIN person_addresses ON person_addresses.address_id = addresses.id").Where(fmt.Sprintf("%v.deleted_at IS NULL OR %v.deleted_at <= '0001-01-02'", table, table))
}

func (role *Role) Scan(value interface{}) error {
	if b, ok := value.([]uint8); ok {
		role.Name = string(b)
	} else {
		role.Name = value.(string)
	}
	return nil
}

func (role Role) Value() (driver.Value, error) {
	return role.Name, nil
}

func (role Role) IsAdmin() bool {
	return role.Name == "admin"
}

func (i *Num) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
	case int64:
		//TODO : @Badu - assignment to method receiver propagates only to callees but not to callers
		*i = Num(s)
	default:
		return errors.New("Cannot scan NamedInt from " + reflect.ValueOf(src).String())
	}
	return nil
}

func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}
	nt.Time, nt.Valid = value.(time.Time), true
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
