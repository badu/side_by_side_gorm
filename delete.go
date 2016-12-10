package tests

import (
	"SideBySideGorm/new_types"
	"testing"
	"time"
)

func DoDelete(t *testing.T) {
	user1, user2 := new_types.User{Name: "delete1"}, new_types.User{Name: "delete2"}
	new_types.TestDB.Save(&user1)
	new_types.TestDB.Save(&user2)

	if err := new_types.TestDB.Delete(&user1).Error; err != nil {
		t.Errorf("No error should happen when delete a record, err=%s", err)
	}

	if !new_types.TestDB.Where("name = ?", user1.Name).First(&new_types.User{}).RecordNotFound() {
		t.Errorf("User can't be found after delete")
	}

	if new_types.TestDB.Where("name = ?", user2.Name).First(&new_types.User{}).RecordNotFound() {
		t.Errorf("Other users that not deleted should be found-able")
	}
}

func InlineDelete(t *testing.T) {
	user1, user2 := new_types.User{Name: "inline_delete1"}, new_types.User{Name: "inline_delete2"}
	new_types.TestDB.Save(&user1)
	new_types.TestDB.Save(&user2)

	if new_types.TestDB.Delete(&new_types.User{}, user1.Id).Error != nil {
		t.Errorf("No error should happen when delete a record")
	} else if !new_types.TestDB.Where("name = ?", user1.Name).First(&new_types.User{}).RecordNotFound() {
		t.Errorf("User can't be found after delete")
	}

	if err := new_types.TestDB.Delete(&new_types.User{}, "name = ?", user2.Name).Error; err != nil {
		t.Errorf("No error should happen when delete a record, err=%s", err)
	} else if !new_types.TestDB.Where("name = ?", user2.Name).First(&new_types.User{}).RecordNotFound() {
		t.Errorf("User can't be found after delete")
	}
}

func SoftDelete(t *testing.T) {
	type User struct {
		Id        int64
		Name      string
		DeletedAt *time.Time
	}
	new_types.TestDB.AutoMigrate(&User{})

	user := User{Name: "soft_delete"}
	new_types.TestDB.Save(&user)
	new_types.TestDB.Delete(&user)

	if new_types.TestDB.First(&User{}, "name = ?", user.Name).Error == nil {
		t.Errorf("Can't find a soft deleted record")
	}

	if err := new_types.TestDB.Unscoped().First(&User{}, "name = ?", user.Name).Error; err != nil {
		t.Errorf("Should be able to find soft deleted record with Unscoped, but err=%s", err)
	}

	new_types.TestDB.Unscoped().Delete(&user)
	if !new_types.TestDB.Unscoped().First(&User{}, "name = ?", user.Name).RecordNotFound() {
		t.Errorf("Can't find permanently deleted record")
	}
}
