package tests

import (
	"SideBySideGorm/new_types"
	newGorm "github.com/badu/gorm"
	"testing"
	"time"
)

func DoCustomizeColumn(t *testing.T) {
	col := "mapped_name"
	new_types.TestDB.DropTable(&new_types.CustomizeColumn{})
	new_types.TestDB.AutoMigrate(&new_types.CustomizeColumn{})

	scope := new_types.TestDB.NewScope(&new_types.CustomizeColumn{})
	if !new_types.TestDB.Dialect().HasColumn(scope.TableName(), col) {
		t.Errorf("CustomizeColumn should have column %s", col)
	}

	col = "mapped_id"
	if scope.PKName() != col {
		t.Errorf("CustomizeColumn should have primary key %s, but got %q", col, scope.PKName())
	}

	expected := "foo"
	now := time.Now()
	cc := new_types.CustomizeColumn{ID: 666, Name: expected, Date: &now}

	if count := new_types.TestDB.Create(&cc).RowsAffected; count != 1 {
		t.Error("There should be one record be affected when create record")
	}

	var cc1 new_types.CustomizeColumn
	new_types.TestDB.First(&cc1, 666)

	if cc1.Name != expected {
		t.Errorf("Failed to query CustomizeColumn")
	}

	cc.Name = "bar"
	new_types.TestDB.Save(&cc)

	var cc2 new_types.CustomizeColumn
	new_types.TestDB.First(&cc2, 666)
	if cc2.Name != "bar" {
		t.Errorf("Failed to query CustomizeColumn")
	}
}

func DoCustomColumnAndIgnoredFieldClash(t *testing.T) {
	new_types.TestDB.DropTable(&new_types.CustomColumnAndIgnoredFieldClash{})
	if err := new_types.TestDB.AutoMigrate(&new_types.CustomColumnAndIgnoredFieldClash{}).Error; err != nil {
		t.Errorf("Should not raise error: %s", err)
	}
}

func ManyToManyWithCustomizedColumn(t *testing.T) {
	new_types.TestDB.DropTable(&new_types.CustomizePerson{}, &new_types.CustomizeAccount{}, "PersonAccount")
	new_types.TestDB.AutoMigrate(&new_types.CustomizePerson{}, &new_types.CustomizeAccount{})

	account := new_types.CustomizeAccount{IdAccount: "account", Name: "id1"}
	person := new_types.CustomizePerson{
		IdPerson: "person",
		Accounts: []new_types.CustomizeAccount{account},
	}

	if err := new_types.TestDB.Create(&account).Error; err != nil {
		t.Errorf("no error should happen, but got %v", err)
	}

	if err := new_types.TestDB.Create(&person).Error; err != nil {
		t.Errorf("no error should happen, but got %v", err)
	}

	var person1 new_types.CustomizePerson
	new_types.TestDB.NewScope(nil)
	if err := new_types.TestDB.Preload("Accounts").First(&person1, newGorm.Quote("idPerson", new_types.TestDB.Dialect())+" = ?", person.IdPerson).Error; err != nil {
		t.Errorf("no error should happen when preloading customized column many2many relations, but got %v", err)
	}

	if len(person1.Accounts) != 1 || person1.Accounts[0].IdAccount != "account" {
		t.Errorf("should preload correct accounts")
	}
}

func OneToOneWithCustomizedColumn(t *testing.T) {
	new_types.TestDB.DropTable(&new_types.CustomizeUser{}, &new_types.CustomizeInvitation{})
	new_types.TestDB.AutoMigrate(&new_types.CustomizeUser{}, &new_types.CustomizeInvitation{})

	user := new_types.CustomizeUser{
		Email: "hello@example.com",
	}
	invitation := new_types.CustomizeInvitation{
		Address: "hello@example.com",
	}

	err := new_types.TestDB.Create(&user).Error
	if err != nil {
		t.Errorf("no error should happen on create user, but got %v", err)
	}
	err = new_types.TestDB.Create(&invitation).Error
	if err != nil {
		t.Errorf("no error should happen on create invitation, but got %v", err)
	}
	var invitation2 new_types.CustomizeInvitation
	err = new_types.TestDB.Preload("Person").Find(&invitation2, invitation.ID).Error
	if err != nil {
		t.Errorf("no error should happen, but got %v", err)
	}
	if invitation2.Person.Email != user.Email {
		t.Errorf("Should preload one to one relation with customize foreign keys")
	}
}

func OneToManyWithCustomizedColumn(t *testing.T) {
	new_types.TestDB.DropTable(&new_types.PromotionDiscount{}, &new_types.PromotionCoupon{})
	new_types.TestDB.AutoMigrate(&new_types.PromotionDiscount{}, &new_types.PromotionCoupon{})

	discount := new_types.PromotionDiscount{
		Name: "Happy New Year",
		Coupons: []*new_types.PromotionCoupon{
			{Code: "newyear1"},
			{Code: "newyear2"},
		},
	}

	if err := new_types.TestDB.Create(&discount).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	var discount1 new_types.PromotionDiscount
	if err := new_types.TestDB.Preload("Coupons").First(&discount1, "id = ?", discount.ID).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if len(discount.Coupons) != 2 {
		t.Errorf("should find two coupons")
	}

	var coupon new_types.PromotionCoupon
	if err := new_types.TestDB.Preload("Discount").First(&coupon, "code = ?", "newyear1").Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if coupon.Discount.Name != "Happy New Year" {
		t.Errorf("should preload discount from coupon")
	}
}

func HasOneWithPartialCustomizedColumn(t *testing.T) {
	new_types.TestDB.DropTable(&new_types.PromotionDiscount{}, &new_types.PromotionRule{})
	new_types.TestDB.AutoMigrate(&new_types.PromotionDiscount{}, &new_types.PromotionRule{})

	var begin = time.Now()
	var end = time.Now().Add(24 * time.Hour)
	discount := new_types.PromotionDiscount{
		Name: "Happy New Year 2",
		Rule: &new_types.PromotionRule{
			Name:  "time_limited",
			Begin: &begin,
			End:   &end,
		},
	}

	if err := new_types.TestDB.Create(&discount).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	var discount1 new_types.PromotionDiscount
	if err := new_types.TestDB.Preload("Rule").First(&discount1, "id = ?", discount.ID).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if discount.Rule.Begin.Format(time.RFC3339Nano) != begin.Format(time.RFC3339Nano) {
		t.Errorf("Should be able to preload Rule")
	}

	var rule new_types.PromotionRule
	if err := new_types.TestDB.Preload("Discount").First(&rule, "name = ?", "time_limited").Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if rule.Discount.Name != "Happy New Year 2" {
		t.Errorf("should preload discount from rule")
	}
}

func BelongsToWithPartialCustomizedColumn(t *testing.T) {
	new_types.TestDB.DropTable(&new_types.PromotionDiscount{}, &new_types.PromotionBenefit{})
	new_types.TestDB.AutoMigrate(&new_types.PromotionDiscount{}, &new_types.PromotionBenefit{})

	discount := new_types.PromotionDiscount{
		Name: "Happy New Year 3",
		Benefits: []new_types.PromotionBenefit{
			{Name: "free cod"},
			{Name: "free shipping"},
		},
	}

	if err := new_types.TestDB.Create(&discount).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	var discount1 new_types.PromotionDiscount
	if err := new_types.TestDB.Preload("Benefits").First(&discount1, "id = ?", discount.ID).Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if len(discount.Benefits) != 2 {
		t.Errorf("should find two benefits")
	}

	var benefit new_types.PromotionBenefit
	if err := new_types.TestDB.Preload("Discount").First(&benefit, "name = ?", "free cod").Error; err != nil {
		t.Errorf("no error should happen but got %v", err)
	}

	if benefit.Discount.Name != "Happy New Year 3" {
		t.Errorf("should preload discount from coupon")
	}
}
