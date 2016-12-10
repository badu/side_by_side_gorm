package tests

import (
	"testing"
	"time"
)

func OldDoDelete(t *testing.T) {
	user1, user2 := User{Name: "delete1"}, User{Name: "delete2"}
	OLDDB.Save(&user1)
	OLDDB.Save(&user2)

	if err := OLDDB.Delete(&user1).Error; err != nil {
		t.Errorf("No error should happen when delete a record, err=%s", err)
	}

	if !OLDDB.Where("name = ?", user1.Name).First(&User{}).RecordNotFound() {
		t.Errorf("User can't be found after delete")
	}

	if OLDDB.Where("name = ?", user2.Name).First(&User{}).RecordNotFound() {
		t.Errorf("Other users that not deleted should be found-able")
	}
}

func OldInlineDelete(t *testing.T) {
	user1, user2 := User{Name: "inline_delete1"}, User{Name: "inline_delete2"}
	OLDDB.Save(&user1)
	OLDDB.Save(&user2)

	if OLDDB.Delete(&User{}, user1.Id).Error != nil {
		t.Errorf("No error should happen when delete a record")
	} else if !OLDDB.Where("name = ?", user1.Name).First(&User{}).RecordNotFound() {
		t.Errorf("User can't be found after delete")
	}

	if err := OLDDB.Delete(&User{}, "name = ?", user2.Name).Error; err != nil {
		t.Errorf("No error should happen when delete a record, err=%s", err)
	} else if !OLDDB.Where("name = ?", user2.Name).First(&User{}).RecordNotFound() {
		t.Errorf("User can't be found after delete")
	}
}

func OldSoftDelete(t *testing.T) {
	type User struct {
		Id        int64
		Name      string
		DeletedAt *time.Time
	}
	OLDDB.AutoMigrate(&User{})

	user := User{Name: "soft_delete"}
	OLDDB.Save(&user)
	OLDDB.Delete(&user)

	if OLDDB.First(&User{}, "name = ?", user.Name).Error == nil {
		t.Errorf("Can't find a soft deleted record")
	}

	if err := OLDDB.Unscoped().First(&User{}, "name = ?", user.Name).Error; err != nil {
		t.Errorf("Should be able to find soft deleted record with Unscoped, but err=%s", err)
	}

	OLDDB.Unscoped().Delete(&user)
	if !OLDDB.Unscoped().First(&User{}, "name = ?", user.Name).RecordNotFound() {
		t.Errorf("Can't find permanently deleted record")
	}
}
