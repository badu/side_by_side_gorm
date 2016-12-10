package tests

import (
	"SideBySideGorm/new_types"
	"fmt"
	"testing"
	"time"
)

func Indexes(t *testing.T) {
	if err := new_types.TestDB.Model(&new_types.Email{}).AddIndex("idx_email_email", "email").Error; err != nil {
		t.Errorf("Got error when tried to create index: %+v", err)
	}

	scope := new_types.TestDB.NewScope(&new_types.Email{})
	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_email_email") {
		t.Errorf("Email should have index idx_email_email")
	}

	if err := new_types.TestDB.Model(&new_types.Email{}).RemoveIndex("idx_email_email").Error; err != nil {
		t.Errorf("Got error when tried to remove index: %+v", err)
	}

	if new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_email_email") {
		t.Errorf("Email's index idx_email_email should be deleted")
	}

	if err := new_types.TestDB.Model(&new_types.Email{}).AddIndex("idx_email_email_and_user_id", "user_id", "email").Error; err != nil {
		t.Errorf("Got error when tried to create index: %+v", err)
	}

	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_email_email_and_user_id") {
		t.Errorf("Email should have index idx_email_email_and_user_id")
	}

	if err := new_types.TestDB.Model(&new_types.Email{}).RemoveIndex("idx_email_email_and_user_id").Error; err != nil {
		t.Errorf("Got error when tried to remove index: %+v", err)
	}

	if new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_email_email_and_user_id") {
		t.Errorf("Email's index idx_email_email_and_user_id should be deleted")
	}

	if err := new_types.TestDB.Model(&new_types.Email{}).AddUniqueIndex("idx_email_email_and_user_id", "user_id", "email").Error; err != nil {
		t.Errorf("Got error when tried to create index: %+v", err)
	}

	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_email_email_and_user_id") {
		t.Errorf("Email should have index idx_email_email_and_user_id")
	}

	if new_types.TestDB.Save(&new_types.User{Name: "unique_indexes", Emails: []new_types.Email{{Email: "user1@example.comiii"}, {Email: "user1@example.com"}, {Email: "user1@example.com"}}}).Error == nil {
		t.Errorf("Should get to create duplicate record when having unique index")
	}

	var user = new_types.User{Name: "sample_user"}
	new_types.TestDB.Save(&user)
	if new_types.TestDB.Model(&user).Association("Emails").Append(new_types.Email{Email: "not-1duplicated@gmail.com"}, new_types.Email{Email: "not-duplicated2@gmail.com"}).Error != nil {
		t.Errorf("Should get no error when append two emails for user")
	}

	if new_types.TestDB.Model(&user).Association("Emails").Append(new_types.Email{Email: "duplicated@gmail.com"}, new_types.Email{Email: "duplicated@gmail.com"}).Error == nil {
		t.Errorf("Should get no duplicated email error when insert duplicated emails for a user")
	}

	if err := new_types.TestDB.Model(&new_types.Email{}).RemoveIndex("idx_email_email_and_user_id").Error; err != nil {
		t.Errorf("Got error when tried to remove index: %+v", err)
	}

	if new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_email_email_and_user_id") {
		t.Errorf("Email's index idx_email_email_and_user_id should be deleted")
	}

	if new_types.TestDB.Save(&new_types.User{Name: "unique_indexes", Emails: []new_types.Email{{Email: "user1@example.com"}, {Email: "user1@example.com"}}}).Error != nil {
		t.Errorf("Should be able to create duplicated emails after remove unique index")
	}
}

func AutoMigration(t *testing.T) {
	new_types.TestDB.AutoMigrate(&new_types.Address{})
	if err := new_types.TestDB.Table("emails").AutoMigrate(&new_types.BigEmail{}).Error; err != nil {
		t.Errorf("Auto Migrate should not raise any error")
	}

	now := time.Now()
	new_types.TestDB.Save(&new_types.BigEmail{Email: "jinzhu@example.org", UserAgent: "pc", RegisteredAt: &now})

	scope := new_types.TestDB.NewScope(&new_types.BigEmail{})
	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_email_agent") {
		t.Errorf("Failed to create index")
	}

	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "uix_emails_registered_at") {
		t.Errorf("Failed to create index")
	}

	var bigemail new_types.BigEmail
	new_types.TestDB.First(&bigemail, "user_agent = ?", "pc")
	if bigemail.Email != "jinzhu@example.org" || bigemail.UserAgent != "pc" || bigemail.RegisteredAt.IsZero() {
		t.Error("Big Emails should be saved and fetched correctly")
	}
}

func DoMultipleIndexes(t *testing.T) {
	if err := new_types.TestDB.DropTableIfExists(&new_types.MultipleIndexes{}).Error; err != nil {
		fmt.Printf("Got error when try to delete table multiple_indexes, %+v\n", err)
	}

	new_types.TestDB.AutoMigrate(&new_types.MultipleIndexes{})
	if err := new_types.TestDB.AutoMigrate(&new_types.BigEmail{}).Error; err != nil {
		t.Errorf("Auto Migrate should not raise any error")
	}

	new_types.TestDB.Save(&new_types.MultipleIndexes{UserID: 1, Name: "jinzhu", Email: "jinzhu@example.org", Other: "foo"})

	scope := new_types.TestDB.NewScope(&new_types.MultipleIndexes{})
	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "uix_multipleindexes_user_name") {
		t.Errorf("Failed to create index")
	}

	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "uix_multipleindexes_user_email") {
		t.Errorf("Failed to create index")
	}

	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "uix_multiple_indexes_email") {
		t.Errorf("Failed to create index")
	}

	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_multipleindexes_user_other") {
		t.Errorf("Failed to create index")
	}

	if !new_types.TestDB.Dialect().HasIndex(scope.TableName(), "idx_multiple_indexes_other") {
		t.Errorf("Failed to create index")
	}

	var mutipleIndexes new_types.MultipleIndexes
	new_types.TestDB.First(&mutipleIndexes, "name = ?", "jinzhu")
	if mutipleIndexes.Email != "jinzhu@example.org" || mutipleIndexes.Name != "jinzhu" {
		t.Error("MutipleIndexes should be saved and fetched correctly")
	}

	// Check unique constraints
	if err := new_types.TestDB.Save(&new_types.MultipleIndexes{UserID: 1, Name: "name1", Email: "jinzhu@example.org", Other: "foo"}).Error; err == nil {
		t.Error("MultipleIndexes unique index failed")
	}

	if err := new_types.TestDB.Save(&new_types.MultipleIndexes{UserID: 1, Name: "name1", Email: "foo@example.org", Other: "foo"}).Error; err != nil {
		t.Error("MultipleIndexes unique index failed")
	}

	if err := new_types.TestDB.Save(&new_types.MultipleIndexes{UserID: 2, Name: "name1", Email: "jinzhu@example.org", Other: "foo"}).Error; err == nil {
		t.Error("MultipleIndexes unique index failed")
	}

	if err := new_types.TestDB.Save(&new_types.MultipleIndexes{UserID: 2, Name: "name1", Email: "foo2@example.org", Other: "foo"}).Error; err != nil {
		t.Error("MultipleIndexes unique index failed")
	}
}
