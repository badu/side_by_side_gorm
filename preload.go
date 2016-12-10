package tests

import (
	"SideBySideGorm/new_types"
	"database/sql"
	newGorm "github.com/badu/gorm"
	"os"
	"reflect"
	"testing"
)

func Preload(t *testing.T) {
	user1 := newGetPreloadUser("user1")
	new_types.TestDB.Save(user1)

	preloadDB := new_types.TestDB.Where("role = ?", "Preload").Preload("BillingAddress").Preload("ShippingAddress").
		Preload("CreditCard").Preload("Emails").Preload("Company")
	var user new_types.User
	preloadDB.Find(&user)
	newCheckUserHasPreloadData(user, t)

	user2 := newGetPreloadUser("user2")
	new_types.TestDB.Save(user2)

	user3 := newGetPreloadUser("user3")
	new_types.TestDB.Save(user3)

	var users []new_types.User
	preloadDB.Find(&users)

	for _, user := range users {
		newCheckUserHasPreloadData(user, t)
	}

	var users2 []*new_types.User
	preloadDB.Find(&users2)

	for _, user := range users2 {
		newCheckUserHasPreloadData(*user, t)
	}

	var users3 []*new_types.User
	preloadDB.Preload("Emails", "email = ?", user3.Emails[0].Email).Find(&users3)

	for _, user := range users3 {
		if user.Name == user3.Name {
			if len(user.Emails) != 1 {
				t.Errorf("should only preload one emails for user3 when with condition")
			}
		} else if len(user.Emails) != 0 {
			t.Errorf("should not preload any emails for other users when with condition (%d loaded)", len(user.Emails))
		} else if user.Emails == nil {
			t.Errorf("should return an empty slice to indicate zero results")
		}
	}
}

func NestedPreload1(t *testing.T) {
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
	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{Level2: Level2{Level1: Level1{Value: "value"}}}
	if err := new_types.TestDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := new_types.TestDB.Preload("Level2").Preload("Level2.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}

	if err := new_types.TestDB.Preload("Level2").Preload("Level2.Level1").Find(&got, "name = ?", "not_found").Error; err != newGorm.ErrRecordNotFound {
		t.Error(err)
	}
}

func NestedPreload2(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := new_types.TestDB.Preload("Level2s.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload3(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level3{
		Level2s: []Level2{
			{Level1: Level1{Value: "value1"}},
			{Level1: Level1{Value: "value2"}},
		},
	}
	if err := new_types.TestDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := new_types.TestDB.Preload("Level2s.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload4(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := new_types.TestDB.Preload("Level2.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

// Slice: []Level3
func NestedPreload5(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := make([]Level3, 2)
	want[0] = Level3{Level2: Level2{Level1: Level1{Value: "value"}}}
	if err := new_types.TestDB.Create(&want[0]).Error; err != nil {
		t.Error(err)
	}
	want[1] = Level3{Level2: Level2{Level1: Level1{Value: "value2"}}}
	if err := new_types.TestDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := new_types.TestDB.Preload("Level2").Preload("Level2.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload6(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want[0]).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := new_types.TestDB.Preload("Level2s.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload7(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := make([]Level3, 2)
	want[0] = Level3{
		Level2s: []Level2{
			{Level1: Level1{Value: "value1"}},
			{Level1: Level1{Value: "value2"}},
		},
	}
	if err := new_types.TestDB.Create(&want[0]).Error; err != nil {
		t.Error(err)
	}

	want[1] = Level3{
		Level2s: []Level2{
			{Level1: Level1{Value: "value3"}},
			{Level1: Level1{Value: "value4"}},
		},
	}
	if err := new_types.TestDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := new_types.TestDB.Preload("Level2s.Level1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload8(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want[0]).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := new_types.TestDB.Preload("Level2.Level1s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload9(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level2_1{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists(&Level0{})
	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}, &Level2_1{}, &Level0{}).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want[0]).Error; err != nil {
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
	if err := new_types.TestDB.Create(&want[1]).Error; err != nil {
		t.Error(err)
	}

	var got []Level3
	if err := new_types.TestDB.Preload("Level2").Preload("Level2.Level1s").Preload("Level2_1").Preload("Level2_1.Level1s").Preload("Level2_1.Level1s.Level0s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload10(t *testing.T) {
	new_types.TestDB.DropTableIfExists(&new_types.LevelA3{})
	new_types.TestDB.DropTableIfExists(&new_types.LevelA2{})
	new_types.TestDB.DropTableIfExists(&new_types.LevelA1{})

	if err := new_types.TestDB.AutoMigrate(&new_types.LevelA1{}, &new_types.LevelA2{}, &new_types.LevelA3{}).Error; err != nil {
		t.Error(err)
	}

	levelA1 := &new_types.LevelA1{Value: "foo"}
	if err := new_types.TestDB.Save(levelA1).Error; err != nil {
		t.Error(err)
	}

	want := []*new_types.LevelA2{
		{
			Value: "bar",
			LevelA3s: []*new_types.LevelA3{
				{
					Value:   "qux",
					LevelA1: levelA1,
				},
			},
		},
		{
			Value:    "bar 2",
			LevelA3s: []*new_types.LevelA3{},
		},
	}
	for _, levelA2 := range want {
		if err := new_types.TestDB.Save(levelA2).Error; err != nil {
			t.Error(err)
		}
	}

	var got []*new_types.LevelA2
	if err := new_types.TestDB.Preload("LevelA3s.LevelA1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload11(t *testing.T) {
	new_types.TestDB.DropTableIfExists(&new_types.LevelB2{})
	new_types.TestDB.DropTableIfExists(&new_types.LevelB3{})
	new_types.TestDB.DropTableIfExists(&new_types.LevelB1{})
	if err := new_types.TestDB.AutoMigrate(&new_types.LevelB1{}, &new_types.LevelB2{}, &new_types.LevelB3{}).Error; err != nil {
		t.Error(err)
	}

	levelB1 := &new_types.LevelB1{Value: "foo"}
	if err := new_types.TestDB.Create(levelB1).Error; err != nil {
		t.Error(err)
	}

	levelB3 := &new_types.LevelB3{
		Value:     "bar",
		LevelB1ID: sql.NullInt64{Valid: true, Int64: int64(levelB1.ID)},
	}
	if err := new_types.TestDB.Create(levelB3).Error; err != nil {
		t.Error(err)
	}
	levelB1.LevelB3s = []*new_types.LevelB3{levelB3}

	want := []*new_types.LevelB1{levelB1}
	var got []*new_types.LevelB1
	if err := new_types.TestDB.Preload("LevelB3s.LevelB2s").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func NestedPreload12(t *testing.T) {
	new_types.TestDB.DropTableIfExists(&new_types.LevelC2{})
	new_types.TestDB.DropTableIfExists(&new_types.LevelC3{})
	new_types.TestDB.DropTableIfExists(&new_types.LevelC1{})
	if err := new_types.TestDB.AutoMigrate(&new_types.LevelC1{}, &new_types.LevelC2{}, &new_types.LevelC3{}).Error; err != nil {
		t.Error(err)
	}

	level2 := new_types.LevelC2{
		Value: "c2",
		LevelC1: new_types.LevelC1{
			Value: "c1",
		},
	}
	new_types.TestDB.Create(&level2)

	want := []new_types.LevelC3{
		{
			Value:   "c3-1",
			LevelC2: level2,
		}, {
			Value:   "c3-2",
			LevelC2: level2,
		},
	}

	for i := range want {
		if err := new_types.TestDB.Create(&want[i]).Error; err != nil {
			t.Error(err)
		}
	}

	var got []new_types.LevelC3
	if err := new_types.TestDB.Preload("LevelC2").Preload("LevelC2.LevelC1").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func ManyToManyPreloadWithMultiPrimaryKeys(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists("levels")

	if err := new_types.TestDB.AutoMigrate(&Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level2{Value: "Bob", LanguageCode: "ru", Level1s: []Level1{
		{Value: "ru", LanguageCode: "ru"},
		{Value: "en", LanguageCode: "en"},
	}}
	if err := new_types.TestDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	want2 := Level2{Value: "Tom", LanguageCode: "zh", Level1s: []Level1{
		{Value: "zh", LanguageCode: "zh"},
		{Value: "de", LanguageCode: "de"},
	}}
	if err := new_types.TestDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got Level2
	if err := new_types.TestDB.Preload("Level1s").Find(&got, "value = ?", "Bob").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}

	var got2 Level2
	if err := new_types.TestDB.Preload("Level1s").Find(&got2, "value = ?", "Tom").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %s; want %s", newtoJSONString(got2), newtoJSONString(want2))
	}

	var got3 []Level2
	if err := new_types.TestDB.Preload("Level1s").Find(&got3, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got3, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", newtoJSONString(got3), newtoJSONString([]Level2{got, got2}))
	}

	var got4 []Level2
	if err := new_types.TestDB.Preload("Level1s", "value IN (?)", []string{"zh", "ru"}).Find(&got4, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	var ruLevel1 Level1
	var zhLevel1 Level1
	new_types.TestDB.First(&ruLevel1, "value = ?", "ru")
	new_types.TestDB.First(&zhLevel1, "value = ?", "zh")

	got.Level1s = []Level1{ruLevel1}
	got2.Level1s = []Level1{zhLevel1}
	if !reflect.DeepEqual(got4, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", newtoJSONString(got4), newtoJSONString([]Level2{got, got2}))
	}

	if err := new_types.TestDB.Preload("Level1s").Find(&got4, "value IN (?)", []string{"non-existing"}).Error; err != nil {
		t.Error(err)
	}
}

func ManyToManyPreloadForNestedPointer(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists("levels")

	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
	if err := new_types.TestDB.Save(&want).Error; err != nil {
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
	if err := new_types.TestDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := new_types.TestDB.Preload("Level2.Level1s").Find(&got, "value = ?", "Bob").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}

	var got2 Level3
	if err := new_types.TestDB.Preload("Level2.Level1s").Find(&got2, "value = ?", "Tom").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %s; want %s", newtoJSONString(got2), newtoJSONString(want2))
	}

	var got3 []Level3
	if err := new_types.TestDB.Preload("Level2.Level1s").Find(&got3, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got3, []Level3{got, got2}) {
		t.Errorf("got %s; want %s", newtoJSONString(got3), newtoJSONString([]Level3{got, got2}))
	}

	var got4 []Level3
	if err := new_types.TestDB.Preload("Level2.Level1s", "value IN (?)", []string{"zh", "ru"}).Find(&got4, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	var got5 Level3
	new_types.TestDB.Preload("Level2.Level1s").Find(&got5, "value = ?", "bogus")

	var ruLevel1 Level1
	var zhLevel1 Level1
	new_types.TestDB.First(&ruLevel1, "value = ?", "ru")
	new_types.TestDB.First(&zhLevel1, "value = ?", "zh")

	got.Level2.Level1s = []*Level1{&ruLevel1}
	got2.Level2.Level1s = []*Level1{&zhLevel1}
	if !reflect.DeepEqual(got4, []Level3{got, got2}) {
		t.Errorf("got %s; want %s", newtoJSONString(got4), newtoJSONString([]Level3{got, got2}))
	}
}

func NestedManyToManyPreload(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists("level1_level2")
	new_types.TestDB.DropTableIfExists("level2_level3")

	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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

	if err := new_types.TestDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := new_types.TestDB.Preload("Level2s").Preload("Level2s.Level1s").Find(&got, "value = ?", "Level3").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}

	if err := new_types.TestDB.Preload("Level2s.Level1s").Find(&got, "value = ?", "not_found").Error; err != newGorm.ErrRecordNotFound {
		t.Error(err)
	}
}

func NestedManyToManyPreload2(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists("level1_level2")

	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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

	if err := new_types.TestDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	var got Level3
	if err := new_types.TestDB.Preload("Level2.Level1s").Find(&got, "value = ?", "Level3").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}

	if err := new_types.TestDB.Preload("Level2.Level1s").Find(&got, "value = ?", "not_found").Error; err != newGorm.ErrRecordNotFound {
		t.Error(err)
	}
}

func NestedManyToManyPreload3(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists("level1_level2")

	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
		if err := new_types.TestDB.Save(&want).Error; err != nil {
			t.Error(err)
		}
	}

	var gots []*Level3
	if err := new_types.TestDB.Preload("Level2.Level1s", func(db *newGorm.DBCon) *newGorm.DBCon {
		return db.Order("level1.id ASC")
	}).Find(&gots).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(gots, wants) {
		t.Errorf("got %s; want %s", newtoJSONString(gots), newtoJSONString(wants))
	}
}

func NestedManyToManyPreload3ForStruct(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists("level1_level2")

	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
		if err := new_types.TestDB.Save(&want).Error; err != nil {
			t.Error(err)
		}
	}

	var gots []*Level3
	if err := new_types.TestDB.Preload("Level2.Level1s", func(db *newGorm.DBCon) *newGorm.DBCon {
		return db.Order("level1.id ASC")
	}).Find(&gots).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(gots, wants) {
		t.Errorf("got %s; want %s", newtoJSONString(gots), newtoJSONString(wants))
	}
}

func NestedManyToManyPreload4(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level4{})
	new_types.TestDB.DropTableIfExists("level1_level2")
	new_types.TestDB.DropTableIfExists("level2_level3")

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

	if err := new_types.TestDB.AutoMigrate(&Level4{}, &Level3{}, &Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	if err := new_types.TestDB.Save(&dummy).Error; err != nil {
		t.Error(err)
	}

	var level1 Level1
	if err := new_types.TestDB.Preload("Level2s").Preload("Level2s.Level3s").Preload("Level2s.Level3s.Level4s").First(&level1).Error; err != nil {
		t.Error(err)
	}
}

func ManyToManyPreloadForPointer(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})
	new_types.TestDB.DropTableIfExists("levels")

	if err := new_types.TestDB.AutoMigrate(&Level2{}, &Level1{}).Error; err != nil {
		t.Error(err)
	}

	want := Level2{Value: "Bob", Level1s: []*Level1{
		{Value: "ru"},
		{Value: "en"},
	}}
	if err := new_types.TestDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	want2 := Level2{Value: "Tom", Level1s: []*Level1{
		{Value: "zh"},
		{Value: "de"},
	}}
	if err := new_types.TestDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got Level2
	if err := new_types.TestDB.Preload("Level1s").Find(&got, "value = ?", "Bob").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}

	var got2 Level2
	if err := new_types.TestDB.Preload("Level1s").Find(&got2, "value = ?", "Tom").Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %s; want %s", newtoJSONString(got2), newtoJSONString(want2))
	}

	var got3 []Level2
	if err := new_types.TestDB.Preload("Level1s").Find(&got3, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got3, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", newtoJSONString(got3), newtoJSONString([]Level2{got, got2}))
	}

	var got4 []Level2
	if err := new_types.TestDB.Preload("Level1s", "value IN (?)", []string{"zh", "ru"}).Find(&got4, "value IN (?)", []string{"Bob", "Tom"}).Error; err != nil {
		t.Error(err)
	}

	var got5 Level2
	new_types.TestDB.Preload("Level1s").First(&got5, "value = ?", "bogus")

	var ruLevel1 Level1
	var zhLevel1 Level1
	new_types.TestDB.First(&ruLevel1, "value = ?", "ru")
	new_types.TestDB.First(&zhLevel1, "value = ?", "zh")

	got.Level1s = []*Level1{&ruLevel1}
	got2.Level1s = []*Level1{&zhLevel1}
	if !reflect.DeepEqual(got4, []Level2{got, got2}) {
		t.Errorf("got %s; want %s", newtoJSONString(got4), newtoJSONString([]Level2{got, got2}))
	}
}

func NilPointerSlice(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(&Level3{})
	new_types.TestDB.DropTableIfExists(&Level2{})
	new_types.TestDB.DropTableIfExists(&Level1{})

	if err := new_types.TestDB.AutoMigrate(&Level3{}, &Level2{}, &Level1{}).Error; err != nil {
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
	if err := new_types.TestDB.Save(&want).Error; err != nil {
		t.Error(err)
	}

	want2 := Level1{
		Value:  "Tom",
		Level2: nil,
	}
	if err := new_types.TestDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	var got []Level1
	if err := new_types.TestDB.Preload("Level2").Preload("Level2.Level3").Find(&got).Error; err != nil {
		t.Error(err)
	}

	if len(got) != 2 {
		t.Errorf("got %v items, expected 2", len(got))
	}

	if !reflect.DeepEqual(got[0], want) && !reflect.DeepEqual(got[1], want) {
		t.Errorf("got %s; want array containing %s", newtoJSONString(got), newtoJSONString(want))
	}

	if !reflect.DeepEqual(got[0], want2) && !reflect.DeepEqual(got[1], want2) {
		t.Errorf("got %s; want array containing %s", newtoJSONString(got), newtoJSONString(want2))
	}
}

func NilPointerSlice2(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(new(Level4))
	new_types.TestDB.DropTableIfExists(new(Level3))
	new_types.TestDB.DropTableIfExists(new(Level2))
	new_types.TestDB.DropTableIfExists(new(Level1))

	if err := new_types.TestDB.AutoMigrate(new(Level4), new(Level3), new(Level2), new(Level1)).Error; err != nil {
		t.Error(err)
	}

	want := new(Level1)
	if err := new_types.TestDB.Save(want).Error; err != nil {
		t.Error(err)
	}

	got := new(Level1)
	err := new_types.TestDB.Preload("Level2.Level3s.Level4").Last(&got).Error
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}

func PrefixedPreloadDuplication(t *testing.T) {
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

	new_types.TestDB.DropTableIfExists(new(Level3))
	new_types.TestDB.DropTableIfExists(new(Level4))
	new_types.TestDB.DropTableIfExists(new(Level2))
	new_types.TestDB.DropTableIfExists(new(Level1))

	if err := new_types.TestDB.AutoMigrate(new(Level3), new(Level4), new(Level2), new(Level1)).Error; err != nil {
		t.Error(err)
	}

	lvl := &Level3{}
	if err := new_types.TestDB.Save(lvl).Error; err != nil {
		t.Error(err)
	}

	sublvl1 := &Level4{Level3ID: lvl.ID}
	if err := new_types.TestDB.Save(sublvl1).Error; err != nil {
		t.Error(err)
	}
	sublvl2 := &Level4{Level3ID: lvl.ID}
	if err := new_types.TestDB.Save(sublvl2).Error; err != nil {
		t.Error(err)
	}

	lvl.Level4s = []*Level4{sublvl1, sublvl2}

	want1 := Level1{
		Level2: &Level2{
			Level3: lvl,
		},
	}
	if err := new_types.TestDB.Save(&want1).Error; err != nil {
		t.Error(err)
	}

	want2 := Level1{
		Level2: &Level2{
			Level3: lvl,
		},
	}
	if err := new_types.TestDB.Save(&want2).Error; err != nil {
		t.Error(err)
	}

	want := []Level1{want1, want2}

	var got []Level1
	err := new_types.TestDB.Preload("Level2.Level3.Level4s").Find(&got).Error
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s; want %s", newtoJSONString(got), newtoJSONString(want))
	}
}
