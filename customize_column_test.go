package tests

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

type CustomizeColumn struct {
	ID   int64      `gorm:"column:mapped_id; primary_key:yes"`
	Name string     `gorm:"column:mapped_name"`
	Date *time.Time `gorm:"column:mapped_time"`
}

// Make sure an ignored field does not interfere with another field's custom
// column name that matches the ignored field.
type CustomColumnAndIgnoredFieldClash struct {
	Body    string `sql:"-"`
	RawBody string `gorm:"column:body"`
}

func OldDoCustomizeColumn(t *testing.T) {
	col := "mapped_name"
	OLDDB.DropTable(&CustomizeColumn{})
	OLDDB.AutoMigrate(&CustomizeColumn{})

	scope := OLDDB.NewScope(&CustomizeColumn{})
	if !scope.Dialect().HasColumn(scope.TableName(), col) {
		t.Errorf("CustomizeColumn should have column %s", col)
	}

	col = "mapped_id"
	if scope.PrimaryKey() != col {
		t.Errorf("CustomizeColumn should have primary key %s, but got %q", col, scope.PrimaryKey())
	}

	expected := "foo"
	now := time.Now()
	cc := CustomizeColumn{ID: 666, Name: expected, Date: &now}

	if count := OLDDB.Create(&cc).RowsAffected; count != 1 {
		t.Error("There should be one record be affected when create record")
	}

	var cc1 CustomizeColumn
	OLDDB.First(&cc1, 666)

	if cc1.Name != expected {
		t.Errorf("Failed to query CustomizeColumn")
	}

	cc.Name = "bar"
	OLDDB.Save(&cc)

	var cc2 CustomizeColumn
	OLDDB.First(&cc2, 666)
	if cc2.Name != "bar" {
		t.Errorf("Failed to query CustomizeColumn")
	}
}

func OldDoCustomColumnAndIgnoredFieldClash(t *testing.T) {
	OLDDB.DropTable(&CustomColumnAndIgnoredFieldClash{})
	if err := OLDDB.AutoMigrate(&CustomColumnAndIgnoredFieldClash{}).Error; err != nil {
		t.Errorf("Should not raise error: %s", err)
	}
}

type CustomizePerson struct {
	IdPerson string             `gorm:"column:idPerson;primary_key:true"`
	Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;associationforeignkey:idAccount;foreignkey:idPerson"`
}

type CustomizeAccount struct {
	IdAccount string `gorm:"column:idAccount;primary_key:true"`
	Name      string
}

func OldManyToManyWithCustomizedColumn(t *testing.T) {
	OLDDB.DropTable(&CustomizePerson{}, &CustomizeAccount{}, "PersonAccount")
	OLDDB.AutoMigrate(&CustomizePerson{}, &CustomizeAccount{})

	account := CustomizeAccount{IdAccount: "account", Name: "id1"}
	person := CustomizePerson{
		IdPerson: "person",
		Accounts: []CustomizeAccount{account},
	}

	if err := OLDDB.Create(&account).Error; err != nil {
		t.Errorf("no error should happen, but got %v", err)
	}

	if err := OLDDB.Create(&person).Error; err != nil {
		t.Errorf("no error should happen, but got %v", err)
	}

	var person1 CustomizePerson
	scope := OLDDB.NewScope(nil)
	if err := OLDDB.Preload("Accounts").First(&person1, scope.Quote("idPerson")+" = ?", person.IdPerson).Error; err != nil {
		t.Errorf("no error should happen when preloading customized column many2many relations, but got %v", err)
	}

	if len(person1.Accounts) != 1 || person1.Accounts[0].IdAccount != "account" {
		t.Errorf("should preload correct accounts")
	}
}

type CustomizeUser struct {
	gorm.Model
	Email string `sql:"column:email_address"`
}

type CustomizeInvitation struct {
	gorm.Model
	Address string         `sql:"column:invitation"`
	Person  *CustomizeUser `gorm:"foreignkey:Email;associationforeignkey:invitation"`
}

func OldOneToOneWithCustomizedColumn(t *testing.T) {
	OLDDB.DropTable(&CustomizeUser{}, &CustomizeInvitation{})
	OLDDB.AutoMigrate(&CustomizeUser{}, &CustomizeInvitation{})

	user := CustomizeUser{
		Email: "hello@example.com",
	}
	invitation := CustomizeInvitation{
		Address: "hello@example.com",
	}

	OLDDB.Create(&user)
	OLDDB.Create(&invitation)

	var invitation2 CustomizeInvitation
	if err := OLDDB.Preload("Person").Find(&invitation2, invitation.ID).Error; err != nil {
		t.Errorf("no error should happen, but got %v", err)
	}

	if invitation2.Person.Email != user.Email {
		t.Errorf("Should preload one to one relation with customize foreign keys")
	}
}

type PromotionDiscount struct {
	gorm.Model
	Name     string
	Coupons  []*PromotionCoupon `gorm:"ForeignKey:discount_id"`
	Rule     *PromotionRule     `gorm:"ForeignKey:discount_id"`
	Benefits []PromotionBenefit `gorm:"ForeignKey:promotion_id"`
}

type PromotionBenefit struct {
	gorm.Model
	Name        string
	PromotionID uint
	Discount    PromotionDiscount `gorm:"ForeignKey:promotion_id"`
}

type PromotionCoupon struct {
	gorm.Model
	Code       string
	DiscountID uint
	Discount   PromotionDiscount
}

type PromotionRule struct {
	gorm.Model
	Name       string
	Begin      *time.Time
	End        *time.Time
	DiscountID uint
	Discount   *PromotionDiscount
}

func OldOneToManyWithCustomizedColumn(t *testing.T) {
	OLDDB.DropTable(&PromotionDiscount{}, &PromotionCoupon{})
	OLDDB.AutoMigrate(&PromotionDiscount{}, &PromotionCoupon{})

	discount := PromotionDiscount{
		Name: "Happy New Year",
		Coupons: []*PromotionCoupon{
			{Code: "newyear1"},
			{Code: "newyear2"},
		},
	}

	if err := OLDDB.Create(&discount).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	var discount1 PromotionDiscount
	if err := OLDDB.Preload("Coupons").First(&discount1, "id = ?", discount.ID).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if len(discount.Coupons) != 2 {
		t.Errorf("should find two coupons")
	}

	var coupon PromotionCoupon
	if err := OLDDB.Preload("Discount").First(&coupon, "code = ?", "newyear1").Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if coupon.Discount.Name != "Happy New Year" {
		t.Errorf("should preload discount from coupon")
	}
}

func OldHasOneWithPartialCustomizedColumn(t *testing.T) {
	OLDDB.DropTable(&PromotionDiscount{}, &PromotionRule{})
	OLDDB.AutoMigrate(&PromotionDiscount{}, &PromotionRule{})

	var begin = time.Now()
	var end = time.Now().Add(24 * time.Hour)
	discount := PromotionDiscount{
		Name: "Happy New Year 2",
		Rule: &PromotionRule{
			Name:  "time_limited",
			Begin: &begin,
			End:   &end,
		},
	}

	if err := OLDDB.Create(&discount).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	var discount1 PromotionDiscount
	if err := OLDDB.Preload("Rule").First(&discount1, "id = ?", discount.ID).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if discount.Rule.Begin.Format(time.RFC3339Nano) != begin.Format(time.RFC3339Nano) {
		t.Errorf("Should be able to preload Rule")
	}

	var rule PromotionRule
	if err := OLDDB.Preload("Discount").First(&rule, "name = ?", "time_limited").Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if rule.Discount.Name != "Happy New Year 2" {
		t.Errorf("should preload discount from rule")
	}
}

func OldBelongsToWithPartialCustomizedColumn(t *testing.T) {
	OLDDB.DropTable(&PromotionDiscount{}, &PromotionBenefit{})
	OLDDB.AutoMigrate(&PromotionDiscount{}, &PromotionBenefit{})

	discount := PromotionDiscount{
		Name: "Happy New Year 3",
		Benefits: []PromotionBenefit{
			{Name: "free cod"},
			{Name: "free shipping"},
		},
	}

	if err := OLDDB.Create(&discount).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	var discount1 PromotionDiscount
	if err := OLDDB.Preload("Benefits").First(&discount1, "id = ?", discount.ID).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if len(discount.Benefits) != 2 {
		t.Errorf("should find two benefits")
	}

	var benefit PromotionBenefit
	if err := OLDDB.Preload("Discount").First(&benefit, "name = ?", "free cod").Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if benefit.Discount.Name != "Happy New Year 3" {
		t.Errorf("should preload discount from coupon")
	}
}
