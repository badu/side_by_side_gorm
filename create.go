package tests

import (
	"SideBySideGorm/new_types"
	"os"
	"reflect"
	"testing"
	"time"
)

func Create(t *testing.T) {
	float := 35.03554004971999
	now := time.Now()
	user := new_types.User{Name: "CreateUser", Age: 18, Birthday: &now, UserNum: new_types.Num(111), PasswordHash: []byte{'f', 'a', 'k', '4'}, Latitude: float}

	if !new_types.TestDB.NewRecord(user) || !new_types.TestDB.NewRecord(&user) {
		t.Error("User should be new record before create")
	}

	if count := new_types.TestDB.Save(&user).RowsAffected; count != 1 {
		t.Error("There should be one record be affected when create record")
	}

	if new_types.TestDB.NewRecord(user) || new_types.TestDB.NewRecord(&user) {
		t.Error("User should not new record after save")
	}

	var newUser new_types.User
	new_types.TestDB.First(&newUser, user.Id)

	if !reflect.DeepEqual(newUser.PasswordHash, []byte{'f', 'a', 'k', '4'}) {
		t.Errorf("User's PasswordHash should be saved ([]byte)")
	}

	if newUser.Age != 18 {
		t.Errorf("User's Age should be saved (int)")
	}

	if newUser.UserNum != new_types.Num(111) {
		t.Errorf("User's UserNum should be saved (custom type)")
	}

	if newUser.Latitude != float {
		t.Errorf("Float64 should not be changed after save")
	}

	if user.CreatedAt.IsZero() {
		t.Errorf("Should have created_at after create")
	}

	if newUser.CreatedAt.IsZero() {
		t.Errorf("Should have created_at after create")
	}

	new_types.TestDB.Model(user).Update("name", "create_user_new_name")
	new_types.TestDB.First(&user, user.Id)
	if user.CreatedAt.Format(time.RFC3339Nano) != newUser.CreatedAt.Format(time.RFC3339Nano) {
		t.Errorf("CreatedAt should not be changed after update")
	}
}

func CreateWithAutoIncrement(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect != "postgres" {
		t.Skip("Skipping this because only postgres properly support auto_increment on a non-primary_key column")
	}
	user1 := new_types.User{}
	user2 := new_types.User{}

	new_types.TestDB.Create(&user1)
	new_types.TestDB.Create(&user2)

	if user2.Sequence-user1.Sequence != 1 {
		t.Errorf("Auto increment should apply on Sequence")
	}
}

func CreateWithNoGORMPrimayKey(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect == "mssql" {
		t.Skip("Skipping this because MSSQL will return identity only if the table has an Id column")
	}

	jt := new_types.JoinTable{From: 1, To: 2}
	err := new_types.TestDB.Create(&jt).Error
	if err != nil {
		t.Errorf("No error should happen when create a record without a GORM primary key. But in the database this primary key exists and is the union of 2 or more fields\n But got: %s", err)
	}
}

func CreateWithNoStdPrimaryKeyAndDefaultValues(t *testing.T) {
	animal := new_types.Animal{Name: "Ferdinand"}
	if new_types.TestDB.Save(&animal).Error != nil {
		t.Errorf("No error should happen when create a record without std primary key")
	}

	if animal.Counter == 0 {
		t.Errorf("No std primary key should be filled value after create")
	}

	if animal.Name != "Ferdinand" {
		t.Errorf("Default value should be overrided")
	}

	// Test create with default value not overrided
	an := new_types.Animal{From: "nerdz"}

	if new_types.TestDB.Save(&an).Error != nil {
		t.Errorf("No error should happen when create an record without std primary key")
	}

	// We must fetch the value again, to have the default fields updated
	// (We can't do this in the update statements, since sql default can be expressions
	// And be different from the fields' type (eg. a time.Time fields has a default value of "now()"
	new_types.TestDB.Model(new_types.Animal{}).Where(&new_types.Animal{Counter: an.Counter}).First(&an)

	if an.Name != "galeone" {
		t.Errorf("Default value should fill the field. But got %v", an.Name)
	}
}

func AnonymousScanner(t *testing.T) {
	user := new_types.User{Name: "anonymous_scanner", Role: new_types.Role{Name: "admin"}}
	new_types.TestDB.Save(&user)

	var user2 new_types.User
	new_types.TestDB.First(&user2, "name = ?", "anonymous_scanner")
	if user2.Role.Name != "admin" {
		t.Errorf("Should be able to get anonymous scanner")
	}

	if !user2.IsAdmin() {
		t.Errorf("Should be able to get anonymous scanner")
	}
}

func AnonymousField(t *testing.T) {
	user := new_types.User{Name: "anonymous_field", Company: new_types.Company{Name: "company"}}
	new_types.TestDB.Save(&user)

	var user2 new_types.User
	new_types.TestDB.First(&user2, "name = ?", "anonymous_field")
	new_types.TestDB.Model(&user2).Related(&user2.Company)
	if user2.Company.Name != "company" {
		t.Errorf("Should be able to get anonymous field")
	}
}

func SelectWithCreate(t *testing.T) {
	user := newGetPreparedUser("select_user", "select_with_create")
	new_types.TestDB.Select("Name", "BillingAddress", "CreditCard", "Company", "Emails").Create(user)

	var queryuser new_types.User
	new_types.TestDB.Preload("BillingAddress").Preload("ShippingAddress").
		Preload("CreditCard").Preload("Emails").Preload("Company").First(&queryuser, user.Id)

	if queryuser.Name != user.Name || queryuser.Age == user.Age {
		t.Errorf("Should only create users with name column")
	}

	if queryuser.BillingAddressID.Int64 == 0 || queryuser.ShippingAddressId != 0 ||
		queryuser.CreditCard.ID == 0 || len(queryuser.Emails) == 0 {
		t.Errorf("Should only create selected relationships")
	}
}

func OmitWithCreate(t *testing.T) {
	user := newGetPreparedUser("omit_user", "omit_with_create")
	new_types.TestDB.Omit("Name", "BillingAddress", "CreditCard", "Company", "Emails").Create(user)

	var queryuser new_types.User
	new_types.TestDB.Preload("BillingAddress").Preload("ShippingAddress").
		Preload("CreditCard").Preload("Emails").Preload("Company").First(&queryuser, user.Id)

	if queryuser.Name == user.Name || queryuser.Age != user.Age {
		t.Errorf("Should only create users with age column")
	}

	if queryuser.BillingAddressID.Int64 != 0 || queryuser.ShippingAddressId == 0 ||
		queryuser.CreditCard.ID != 0 || len(queryuser.Emails) != 0 {
		t.Errorf("Should not create omited relationships")
	}
}
