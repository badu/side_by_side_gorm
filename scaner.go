package tests

import (
	"SideBySideGorm/new_types"
	"testing"
)

func ScannableSlices(t *testing.T) {
	if err := new_types.TestDB.AutoMigrate(&new_types.RecordWithSlice{}).Error; err != nil {
		t.Errorf("Should create table with slice values correctly: %s", err)
	}

	r1 := new_types.RecordWithSlice{
		Strings: new_types.ExampleStringSlice{"a", "b", "c"},
		Structs: new_types.ExampleStructSlice{
			{"name1", "value1"},
			{"name2", "value2"},
		},
	}

	if err := new_types.TestDB.Save(&r1).Error; err != nil {
		t.Errorf("Should save record with slice values")
	}

	var r2 new_types.RecordWithSlice

	if err := new_types.TestDB.Find(&r2).Error; err != nil {
		t.Errorf("Should fetch record with slice values")
	}

	if len(r2.Strings) != 3 || r2.Strings[0] != "a" || r2.Strings[1] != "b" || r2.Strings[2] != "c" {
		t.Errorf("Should have serialised and deserialised a string array")
	}

	if len(r2.Structs) != 2 || r2.Structs[0].Name != "name1" || r2.Structs[0].Value != "value1" || r2.Structs[1].Name != "name2" || r2.Structs[1].Value != "value2" {
		t.Errorf("Should have serialised and deserialised a struct array")
	}
}
