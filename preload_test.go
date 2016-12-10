package tests

import (
	"database/sql"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func getPreloadUser(name string) *User {
	return getPreparedUser(name, "Preload")
}

func checkUserHasPreloadData(user User, t *testing.T) {
	u := getPreloadUser(user.Name)
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

func OldPreload(t *testing.T) {
	user1 := getPreloadUser("user1")
	OLDDB.Save(user1)

	preloadDB := OLDDB.Where("role = ?", "Preload").Preload("BillingAddress").Preload("ShippingAddress").
		Preload("CreditCard").Preload("Emails").Preload("Company")
	var user User
	preloadDB.Find(&user)
	checkUserHasPreloadData(user, t)

	user2 := getPreloadUser("user2")
	OLDDB.Save(user2)

	user3 := getPreloadUser("user3")
	OLDDB.Save(user3)

	var users []User
	preloadDB.Find(&users)

	for _, user := range users {
		checkUserHasPreloadData(user, t)
	}

	var users2 []*User
	preloadDB.Find(&users2)

	for _, user := range users2 {
		checkUserHasPreloadData(*user, t)
	}

	var users3 []*User
	preloadDB.Preload("Emails", "email = ?", user3.Emails[0].Email).Find(&users3)

	for _, user := range users3 {
		if user.Name == user3.Name {
			if len(user.Emails) != 1 {
				t.Errorf("should only preload one emails for user3 when with condition")
			}
		} else if len(user.Emails) != 0 {
			t.Errorf("should not preload any emails for other users when with condition")
		} else if user.Emails == nil {
			t.Errorf("should return an empty slice to indicate zero results")
		}
	}
}

func OldNestedPreload1(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1   Level1
			Level3ID uint
		}
		Level3 struct {
			ID     uint
			Name   string
			Level2 Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{Level2: Level2{Level1: Level1{Value: "value"}}}
	if err := OLDDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := OLDDB.Preload("Level2").Preload("Level2.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}

	if err := OLDDB.Preload("Level2").Preload("Level2.Level1").Find(&got, "name = ?", "not_found").Error; err != gorm.ErrRecordNotFound {
		t.Error(err)
	}
}

func OldNestedPreload2(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1s  []*Level1
			Level3ID uint
		}
		Level3 struct {
			ID      uint
			Name    string
			Level2s []Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{
		Level2s: []Level2{
			{
				Level1s: []*Level1{
					{Value: "value1"},
					{Value: "value2"},
				},
			},
			{
				Level1s: []*Level1{
					{Value: "value3"},
				},
			},
		},
	}
	if err := OLDDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := OLDDB.Preload("Level2s.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldNestedPreload3(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1   Level1
			Level3ID uint
		}
		Level3 struct {
			Name    string
			ID      uint
			Level2s []Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{
		Level2s: []Level2{
			{Level1: Level1{Value: "value1"}},
			{Level1: Level1{Value: "value2"}},
		},
	}
	if err := OLDDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := OLDDB.Preload("Level2s.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldNestedPreload4(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1s  []Level1
			Level3ID uint
		}
		Level3 struct {
			ID     uint
			Name   string
			Level2 Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{
		Level2: Level2{
			Level1s: []Level1{
				{Value: "value1"},
				{Value: "value2"},
			},
		},
	}
	if err := OLDDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := OLDDB.Preload("Level2.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

// Slice: []Level3
func OldNestedPreload5(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1   Level1
			Level3ID uint
		}
		Level3 struct {
			ID     uint
			Name   string
			Level2 Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := make([]Level3, 2)
	want[0] = Level3{Level2: Level2{Level1: Level1{Value: "value"}}}
	if err := OLDDB.Create(&want[0]).Error; err != nil {
		t.Error(err)
	}
	want[1] = Level3{Level2: Level2{Level1: Level1{Value: "value2"}}}
	if err := OLDDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := OLDDB.Preload("Level2").Preload("Level2.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldNestedPreload6(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1s  []Level1
			Level3ID uint
		}
		Level3 struct {
			ID      uint
			Name    string
			Level2s []Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := make([]Level3, 2)
	want[0] = Level3{
		Level2s: []Level2{
			{
				Level1s: []Level1{
					{Value: "value1"},
					{Value: "value2"},
				},
			},
			{
				Level1s: []Level1{
					{Value: "value3"},
				},
			},
		},
	}
	if err := OLDDB.Create(&want[0]).Error; err != nil {
		t.Error(err)
	}

	want[1] = Level3{
		Level2s: []Level2{
			{
				Level1s: []Level1{
					{Value: "value3"},
					{Value: "value4"},
				},
			},
			{
				Level1s: []Level1{
					{Value: "value5"},
				},
			},
		},
	}
	if err := OLDDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := OLDDB.Preload("Level2s.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldNestedPreload7(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1   Level1
			Level3ID uint
		}
		Level3 struct {
			ID      uint
			Name    string
			Level2s []Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := make([]Level3, 2)
	want[0] = Level3{
		Level2s: []Level2{
			{Level1: Level1{Value: "value1"}},
			{Level1: Level1{Value: "value2"}},
		},
	}
	if err := OLDDB.Create(&want[0]).Error; err != nil {
		t.Error(err)
	}

	want[1] = Level3{
		Level2s: []Level2{
			{Level1: Level1{Value: "value3"}},
			{Level1: Level1{Value: "value4"}},
		},
	}
	if err := OLDDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := OLDDB.Preload("Level2s.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldNestedPreload8(t *testing.T) {
	type (
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
		}
		Level2 struct {
			ID       uint
			Level1s  []Level1
			Level3ID uint
		}
		Level3 struct {
			ID     uint
			Name   string
			Level2 Level2
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := make([]Level3, 2)
	want[0] = Level3{
		Level2: Level2{
			Level1s: []Level1{
				{Value: "value1"},
				{Value: "value2"},
			},
		},
	}
	if err := OLDDB.Create(&want[0]).Error; err != nil {
		t.Error(err)
	}
	want[1] = Level3{
		Level2: Level2{
			Level1s: []Level1{
				{Value: "value3"},
				{Value: "value4"},
			},
		},
	}
	if err := OLDDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := OLDDB.Preload("Level2.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldNestedPreload9(t *testing.T) {
	type (
		Level0 struct {
			ID       uint
			Value    string
			Level1ID uint
		}
		Level1 struct {
			ID         uint
			Value      string
			Level2ID   uint
			Level2_1ID uint
			Level0s    []Level0
		}
		Level2 struct {
			ID       uint
			Level1s  []Level1
			Level3ID uint
		}
		Level2_1 struct {
			ID       uint
			Level1s  []Level1
			Level3ID uint
		}
		Level3 struct {
			ID       uint
			Name     string
			Level2   Level2
			Level2_1 Level2_1
		}
	)
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level2_1{})
	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists(&Level0{})
	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}, &Level2_1{}, &Level0{}).Error; err != nil {
		t.Error(err)
	}

	want := make([]Level3, 2)
	want[0] = Level3{
		Level2: Level2{
			Level1s: []Level1{
				{Value: "value1"},
				{Value: "value2"},
			},
		},
		Level2_1: Level2_1{
			Level1s: []Level1{
				{
					Value:   "value1-1",
					Level0s: []Level0{{Value: "Level0-1"}},
				},
				{
					Value:   "value2-2",
					Level0s: []Level0{{Value: "Level0-2"}},
				},
			},
		},
	}
	if err := OLDDB.Create(&want[0]).Error; err != nil {
		t.Error(err)
	}
	want[1] = Level3{
		Level2: Level2{
			Level1s: []Level1{
				{Value: "value3"},
				{Value: "value4"},
			},
		},
		Level2_1: Level2_1{
			Level1s: []Level1{
				{
					Value:   "value3-3",
					Level0s: []Level0{},
				},
				{
					Value:   "value4-4",
					Level0s: []Level0{},
				},
			},
		},
	}
	if err := OLDDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := OLDDB.Preload("Level2").Preload("Level2.Level1s").Preload("Level2_1").Preload("Level2_1.Level1s").Preload("Level2_1.Level1s.Level0s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

type LevelA1 struct {
	ID    uint
	Value string
}

type LevelA2 struct {
	ID       uint
	Value    string
	LevelA3s []*LevelA3
}

type LevelA3 struct {
	ID        uint
	Value     string
	LevelA1ID sql.NullInt64
	LevelA1   *LevelA1
	LevelA2ID sql.NullInt64
	LevelA2   *LevelA2
}

func OldNestedPreload10(t *testing.T) {
	OLDDB.DropTableIfExists(&LevelA3{})
	OLDDB.DropTableIfExists(&LevelA2{})
	OLDDB.DropTableIfExists(&LevelA1{})

	if err := OLDDB.AutoMigrate(&LevelA1{}, &LevelA2{}, &LevelA3{}).Error; err != nil {
		t.Error(err)
	}

	levelA1 := &LevelA1{Value: "foo"}
	if err := OLDDB.Save(levelA1).Error; err != nil {
		t.Error(err)
	}

	want := []*LevelA2{
		{
			Value: "bar",
			LevelA3s: []*LevelA3{
				{
					Value:   "qux",
					LevelA1: levelA1,
				},
			},
		},
		{
			Value:    "bar 2",
			LevelA3s: []*LevelA3{},
		},
	}
	for _, levelA2 := range want {
		if err := OLDDB.Save(levelA2).Error; err != nil {
			t.Error(err)
		}
	}

	var got []*LevelA2
	if err := OLDDB.Preload("LevelA3s.LevelA1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

type LevelB1 struct {
	ID       uint
	Value    string
	LevelB3s []*LevelB3
}

type LevelB2 struct {
	ID    uint
	Value string
}

type LevelB3 struct {
	ID        uint
	Value     string
	LevelB1ID sql.NullInt64
	LevelB1   *LevelB1
	LevelB2s  []*LevelB2 `gorm:"many2many:levelb1_levelb3_levelb2s"`
}

func OldNestedPreload11(t *testing.T) {
	OLDDB.DropTableIfExists(&LevelB2{})
	OLDDB.DropTableIfExists(&LevelB3{})
	OLDDB.DropTableIfExists(&LevelB1{})
	if err := OLDDB.AutoMigrate(&LevelB1{}, &LevelB2{}, &LevelB3{}).Error; err != nil {
		t.Error(err)
	}

	levelB1 := &LevelB1{Value: "foo"}
	if err := OLDDB.Create(levelB1).Error; err != nil {
		t.Error(err)
	}

	levelB3 := &LevelB3{
		Value:     "bar",
		LevelB1ID: sql.NullInt64{Valid: true, Int64: int64(levelB1.ID)},
	}
	if err := OLDDB.Create(levelB3).Error; err != nil {
		t.Error(err)
	}
	levelB1.LevelB3s = []*LevelB3{levelB3}

	want := []*LevelB1{levelB1}
	var got []*LevelB1
	if err := OLDDB.Preload("LevelB3s.LevelB2s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

type LevelC1 struct {
	ID        uint
	Value     string
	LevelC2ID uint
}

type LevelC2 struct {
	ID      uint
	Value   string
	LevelC1 LevelC1
}

type LevelC3 struct {
	ID        uint
	Value     string
	LevelC2ID uint
	LevelC2   LevelC2
}

func OldNestedPreload12(t *testing.T) {
	OLDDB.DropTableIfExists(&LevelC2{})
	OLDDB.DropTableIfExists(&LevelC3{})
	OLDDB.DropTableIfExists(&LevelC1{})
	if err := OLDDB.AutoMigrate(&LevelC1{}, &LevelC2{}, &LevelC3{}).Error; err != nil {
		t.Error(err)
	}

	level2 := LevelC2{
		Value: "c2",
		LevelC1: LevelC1{
			Value: "c1",
		},
	}
	OLDDB.Create(&level2)

	want := []LevelC3{
		{
			Value:   "c3-1",
			LevelC2: level2,
		}, {
			Value:   "c3-2",
			LevelC2: level2,
		},
	}

	for i := range want {
		if err := OLDDB.Create(&want[i]).Error; err != nil {
			t.Error(err)
		}
	}

	var got []LevelC3
	if err := OLDDB.Preload("LevelC2").Preload("LevelC2.LevelC1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldManyToManyPreloadWithMultiPrimaryKeys(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect == "" || dialect == "sqlite" {
		return
	}

	type (
		Level1 struct {
			ID           uint   `gorm:"primary_key;"`
			LanguageCode string `gorm:"primary_key"`
			Value        string
		}
		Level2 struct {
			ID           uint   `gorm:"primary_key;"`
			LanguageCode string `gorm:"primary_key"`
			Value        string
			Level1s      []Level1 `gorm:"many2many:levels;"`
		}
	)

	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists("levels")

	if err := OLDDB.AutoMigrate(&Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level2{Value: "Bob", LanguageCode: "ru", Level1s: []Level1{
		{Value: "ru", LanguageCode: "ru"},
		{Value: "en", LanguageCode: "en"},
	}}
	if err := OLDDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	want2 := Level2{Value: "Tom", LanguageCode: "zh", Level1s: []Level1{
		{Value: "zh", LanguageCode: "zh"},
		{Value: "de", LanguageCode: "de"},
	}}
	if err := OLDDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got Level2
	if err := OLDDB.Preload("Level1s").Find(&got, "value = ?", "Bob").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}

	var got2 Level2
	if err := OLDDB.Preload("Level1s").Find(&got2, "value = ?", "Tom").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %s; want %s", toJSONString(got2), toJSONString(want2))
	}

	var got3 []Level2
	if err := OLDDB.Preload("Level1s").Find(&got3, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got3, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", toJSONString(got3), toJSONString([]Level2{got, got2}))
	}

	var got4 []Level2
	if err := OLDDB.Preload("Level1s", "value IN (?)", []string{"zh", "ru"}).Find(&got4, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	var ruLevel1 Level1
	var zhLevel1 Level1
	OLDDB.First(&ruLevel1, "value = ?", "ru")
	OLDDB.First(&zhLevel1, "value = ?", "zh")

	got.Level1s = []Level1{ruLevel1}
	got2.Level1s = []Level1{zhLevel1}
	if !reflect.DeepEqual(got4, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", toJSONString(got4), toJSONString([]Level2{got, got2}))
	}

	if err := OLDDB.Preload("Level1s").Find(&got4, "value IN (?)", []string{"non-existing"}).Error; err != nil {
		t.Error(err)
	}
}

func OldManyToManyPreloadForNestedPointer(t *testing.T) {
	type (
		Level1 struct {
			ID    uint
			Value string
		}
		Level2 struct {
			ID      uint
			Value   string
			Level1s []*Level1 `gorm:"many2many:levels;"`
		}
		Level3 struct {
			ID       uint
			Value    string
			Level2ID sql.NullInt64
			Level2   *Level2
		}
	)

	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists("levels")

	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{
		Value: "Bob",
		Level2: &Level2{
			Value: "Foo",
			Level1s: []*Level1{
				{Value: "ru"},
				{Value: "en"},
			},
		},
	}
	if err := OLDDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	want2 := Level3{
		Value: "Tom",
		Level2: &Level2{
			Value: "Bar",
			Level1s: []*Level1{
				{Value: "zh"},
				{Value: "de"},
			},
		},
	}
	if err := OLDDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := OLDDB.Preload("Level2.Level1s").Find(&got, "value = ?", "Bob").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}

	var got2 Level3
	if err := OLDDB.Preload("Level2.Level1s").Find(&got2, "value = ?", "Tom").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %s; want %s", toJSONString(got2), toJSONString(want2))
	}

	var got3 []Level3
	if err := OLDDB.Preload("Level2.Level1s").Find(&got3, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got3, []Level3{got, got2}) {
		t.Errorf("got %s; want %s", toJSONString(got3), toJSONString([]Level3{got, got2}))
	}

	var got4 []Level3
	if err := OLDDB.Preload("Level2.Level1s", "value IN (?)", []string{"zh", "ru"}).Find(&got4, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	var got5 Level3
	OLDDB.Preload("Level2.Level1s").Find(&got5, "value = ?", "bogus")

	var ruLevel1 Level1
	var zhLevel1 Level1
	OLDDB.First(&ruLevel1, "value = ?", "ru")
	OLDDB.First(&zhLevel1, "value = ?", "zh")

	got.Level2.Level1s = []*Level1{&ruLevel1}
	got2.Level2.Level1s = []*Level1{&zhLevel1}
	if !reflect.DeepEqual(got4, []Level3{got, got2}) {
		t.Errorf("got %s; want %s", toJSONString(got4), toJSONString([]Level3{got, got2}))
	}
}

func OldNestedManyToManyPreload(t *testing.T) {
	type (
		Level1 struct {
			ID    uint
			Value string
		}
		Level2 struct {
			ID      uint
			Value   string
			Level1s []*Level1 `gorm:"many2many:level1_level2;"`
		}
		Level3 struct {
			ID      uint
			Value   string
			Level2s []Level2 `gorm:"many2many:level2_level3;"`
		}
	)

	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists("level1_level2")
	OLDDB.DropTableIfExists("level2_level3")

	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{
		Value: "Level3",
		Level2s: []Level2{
			{
				Value: "Bob",
				Level1s: []*Level1{
					{Value: "ru"},
					{Value: "en"},
				},
			}, {
				Value: "Tom",
				Level1s: []*Level1{
					{Value: "zh"},
					{Value: "de"},
				},
			},
		},
	}

	if err := OLDDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := OLDDB.Preload("Level2s").Preload("Level2s.Level1s").Find(&got, "value = ?", "Level3").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}

	if err := OLDDB.Preload("Level2s.Level1s").Find(&got, "value = ?", "not_found").Error; err != gorm.ErrRecordNotFound {
		t.Error(err)
	}
}

func OldNestedManyToManyPreload2(t *testing.T) {
	type (
		Level1 struct {
			ID    uint
			Value string
		}
		Level2 struct {
			ID      uint
			Value   string
			Level1s []*Level1 `gorm:"many2many:level1_level2;"`
		}
		Level3 struct {
			ID       uint
			Value    string
			Level2ID sql.NullInt64
			Level2   *Level2
		}
	)

	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists("level1_level2")

	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{
		Value: "Level3",
		Level2: &Level2{
			Value: "Bob",
			Level1s: []*Level1{
				{Value: "ru"},
				{Value: "en"},
			},
		},
	}

	if err := OLDDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := OLDDB.Preload("Level2.Level1s").Find(&got, "value = ?", "Level3").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}

	if err := OLDDB.Preload("Level2.Level1s").Find(&got, "value = ?", "not_found").Error; err != gorm.ErrRecordNotFound {
		t.Error(err)
	}
}

func OldNestedManyToManyPreload3(t *testing.T) {
	type (
		Level1 struct {
			ID    uint
			Value string
		}
		Level2 struct {
			ID      uint
			Value   string
			Level1s []*Level1 `gorm:"many2many:level1_level2;"`
		}
		Level3 struct {
			ID       uint
			Value    string
			Level2ID sql.NullInt64
			Level2   *Level2
		}
	)

	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists("level1_level2")

	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	level1Zh := &Level1{Value: "zh"}
	level1Ru := &Level1{Value: "ru"}
	level1En := &Level1{Value: "en"}

	level21 := &Level2{
		Value:   "Level2-1",
		Level1s: []*Level1{level1Zh, level1Ru},
	}

	level22 := &Level2{
		Value:   "Level2-2",
		Level1s: []*Level1{level1Zh, level1En},
	}

	wants := []*Level3{
		{
			Value:  "Level3-1",
			Level2: level21,
		},
		{
			Value:  "Level3-2",
			Level2: level22,
		},
		{
			Value:  "Level3-3",
			Level2: level21,
		},
	}

	for _, want := range wants {
		if err := OLDDB.Save(&want).Error; err != nil {
			t.Error(err)
		}
	}

	var gots []*Level3
	if err := OLDDB.Preload("Level2.Level1s", func(db *gorm.DB) *gorm.DB {
		return db.Order("level1.id ASC")
	}).Find(&gots).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(gots, wants) {
		t.Errorf("got %s; want %s", toJSONString(gots), toJSONString(wants))
	}
}

func OldNestedManyToManyPreload3ForStruct(t *testing.T) {
	type (
		Level1 struct {
			ID    uint
			Value string
		}
		Level2 struct {
			ID      uint
			Value   string
			Level1s []Level1 `gorm:"many2many:level1_level2;"`
		}
		Level3 struct {
			ID       uint
			Value    string
			Level2ID sql.NullInt64
			Level2   Level2
		}
	)

	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists("level1_level2")

	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	level1Zh := Level1{Value: "zh"}
	level1Ru := Level1{Value: "ru"}
	level1En := Level1{Value: "en"}

	level21 := Level2{
		Value:   "Level2-1",
		Level1s: []Level1{level1Zh, level1Ru},
	}

	level22 := Level2{
		Value:   "Level2-2",
		Level1s: []Level1{level1Zh, level1En},
	}

	wants := []*Level3{
		{
			Value:  "Level3-1",
			Level2: level21,
		},
		{
			Value:  "Level3-2",
			Level2: level22,
		},
		{
			Value:  "Level3-3",
			Level2: level21,
		},
	}

	for _, want := range wants {
		if err := OLDDB.Save(&want).Error; err != nil {
			t.Error(err)
		}
	}

	var gots []*Level3
	if err := OLDDB.Preload("Level2.Level1s", func(db *gorm.DB) *gorm.DB {
		return db.Order("level1.id ASC")
	}).Find(&gots).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(gots, wants) {
		t.Errorf("got %s; want %s", toJSONString(gots), toJSONString(wants))
	}
}

func OldNestedManyToManyPreload4(t *testing.T) {
	type (
		Level4 struct {
			ID       uint
			Value    string
			Level3ID uint
		}
		Level3 struct {
			ID      uint
			Value   string
			Level4s []*Level4
		}
		Level2 struct {
			ID      uint
			Value   string
			Level3s []*Level3 `gorm:"many2many:level2_level3;"`
		}
		Level1 struct {
			ID      uint
			Value   string
			Level2s []*Level2 `gorm:"many2many:level1_level2;"`
		}
	)

	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level4{})
	OLDDB.DropTableIfExists("level1_level2")
	OLDDB.DropTableIfExists("level2_level3")

	dummy := Level1{
		Value: "Level1",
		Level2s: []*Level2{{
			Value: "Level2",
			Level3s: []*Level3{{
				Value: "Level3",
				Level4s: []*Level4{{
					Value: "Level4",
				}},
			}},
		}},
	}

	if err := OLDDB.AutoMigrate(&Level4{}, &Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	if err := OLDDB.Save(&dummy).Error; err != nil {
		t.Error(err)
	}

	var level1 Level1
	if err := OLDDB.Preload("Level2s").Preload("Level2s.Level3s").Preload("Level2s.Level3s.Level4s").First(&level1).Error; err != nil {
		t.Error(err)
	}
}

func OldManyToManyPreloadForPointer(t *testing.T) {
	type (
		Level1 struct {
			ID    uint
			Value string
		}
		Level2 struct {
			ID      uint
			Value   string
			Level1s []*Level1 `gorm:"many2many:levels;"`
		}
	)

	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})
	OLDDB.DropTableIfExists("levels")

	if err := OLDDB.AutoMigrate(&Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level2{Value: "Bob", Level1s: []*Level1{
		{Value: "ru"},
		{Value: "en"},
	}}
	if err := OLDDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	want2 := Level2{Value: "Tom", Level1s: []*Level1{
		{Value: "zh"},
		{Value: "de"},
	}}
	if err := OLDDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got Level2
	if err := OLDDB.Preload("Level1s").Find(&got, "value = ?", "Bob").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}

	var got2 Level2
	if err := OLDDB.Preload("Level1s").Find(&got2, "value = ?", "Tom").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %s; want %s", toJSONString(got2), toJSONString(want2))
	}

	var got3 []Level2
	if err := OLDDB.Preload("Level1s").Find(&got3, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got3, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", toJSONString(got3), toJSONString([]Level2{got, got2}))
	}

	var got4 []Level2
	if err := OLDDB.Preload("Level1s", "value IN (?)", []string{"zh", "ru"}).Find(&got4, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	var got5 Level2
	OLDDB.Preload("Level1s").First(&got5, "value = ?", "bogus")

	var ruLevel1 Level1
	var zhLevel1 Level1
	OLDDB.First(&ruLevel1, "value = ?", "ru")
	OLDDB.First(&zhLevel1, "value = ?", "zh")

	got.Level1s = []*Level1{&ruLevel1}
	got2.Level1s = []*Level1{&zhLevel1}
	if !reflect.DeepEqual(got4, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", toJSONString(got4), toJSONString([]Level2{got, got2}))
	}
}

func OldNilPointerSlice(t *testing.T) {
	type (
		Level3 struct {
			ID    uint
			Value string
		}
		Level2 struct {
			ID       uint
			Value    string
			Level3ID uint
			Level3   *Level3
		}
		Level1 struct {
			ID       uint
			Value    string
			Level2ID uint
			Level2   *Level2
		}
	)

	OLDDB.DropTableIfExists(&Level3{})
	OLDDB.DropTableIfExists(&Level2{})
	OLDDB.DropTableIfExists(&Level1{})

	if err := OLDDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level1{
		Value: "Bob",
		Level2: &Level2{
			Value: "en",
			Level3: &Level3{
				Value: "native",
			},
		},
	}
	if err := OLDDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	want2 := Level1{
		Value:  "Tom",
		Level2: nil,
	}
	if err := OLDDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got []Level1
	if err := OLDDB.Preload("Level2").Preload("Level2.Level3").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if len(got) != 2 {
		t.Errorf("got %v items, expected 2", len(got))
	}

	if !reflect.DeepEqual(got[0], want) && !reflect.DeepEqual(got[1], want) {
		t.Errorf("got %s; want array containing %s", toJSONString(got), toJSONString(want))
	}

	if !reflect.DeepEqual(got[0], want2) && !reflect.DeepEqual(got[1], want2) {
		t.Errorf("got %s; want array containing %s", toJSONString(got), toJSONString(want2))
	}
}

func OldNilPointerSlice2(t *testing.T) {
	type (
		Level4 struct {
			ID uint
		}
		Level3 struct {
			ID       uint
			Level4ID sql.NullInt64 `sql:"index"`
			Level4   *Level4
		}
		Level2 struct {
			ID      uint
			Level3s []*Level3 `gorm:"many2many:level2_level3s"`
		}
		Level1 struct {
			ID       uint
			Level2ID sql.NullInt64 `sql:"index"`
			Level2   *Level2
		}
	)

	OLDDB.DropTableIfExists(new(Level4))
	OLDDB.DropTableIfExists(new(Level3))
	OLDDB.DropTableIfExists(new(Level2))
	OLDDB.DropTableIfExists(new(Level1))

	if err := OLDDB.AutoMigrate(new(Level4), new(Level3), new(Level2), new(Level1)).Error; err != nil {
		t.Error(err)
	}

	want := new(Level1)
	if err := OLDDB.Save(want).Error; err != nil {
		t.Error(err)
	}

	got := new(Level1)
	err := OLDDB.Preload("Level2.Level3s.Level4").Last(&got).Error
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func OldPrefixedPreloadDuplication(t *testing.T) {
	type (
		Level4 struct {
			ID       uint
			Name     string
			Level3ID uint
		}
		Level3 struct {
			ID      uint
			Name    string
			Level4s []*Level4
		}
		Level2 struct {
			ID       uint
			Name     string
			Level3ID sql.NullInt64 `sql:"index"`
			Level3   *Level3
		}
		Level1 struct {
			ID       uint
			Name     string
			Level2ID sql.NullInt64 `sql:"index"`
			Level2   *Level2
		}
	)

	OLDDB.DropTableIfExists(new(Level3))
	OLDDB.DropTableIfExists(new(Level4))
	OLDDB.DropTableIfExists(new(Level2))
	OLDDB.DropTableIfExists(new(Level1))

	if err := OLDDB.AutoMigrate(new(Level3), new(Level4), new(Level2), new(Level1)).Error; err != nil {
		t.Error(err)
	}

	lvl := &Level3{}
	if err := OLDDB.Save(lvl).Error; err != nil {
		t.Error(err)
	}

	sublvl1 := &Level4{Level3ID: lvl.ID}
	if err := OLDDB.Save(sublvl1).Error; err != nil {
		t.Error(err)
	}
	sublvl2 := &Level4{Level3ID: lvl.ID}
	if err := OLDDB.Save(sublvl2).Error; err != nil {
		t.Error(err)
	}

	lvl.Level4s = []*Level4{sublvl1, sublvl2}

	want1 := Level1{
		Level2: &Level2{
			Level3: lvl,
		},
	}
	if err := OLDDB.Save(&want1).Error; err != nil {
		t.Error(err)
	}

	want2 := Level1{
		Level2: &Level2{
			Level3: lvl,
		},
	}
	if err := OLDDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	want := []Level1{want1, want2}

	var got []Level1
	err := OLDDB.Preload("Level2.Level3.Level4s").Find(&got).Error
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", toJSONString(got), toJSONString(want))
	}
}

func toJSONString(v interface{}) []byte {
	r, _ := json.MarshalIndent(v, "", "  ")
	return r
}
