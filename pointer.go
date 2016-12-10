package tests

import (
	"SideBySideGorm/new_types"
	"testing"
)

func PointerFields(t *testing.T) {

	new_types.TestDB.DropTable(&new_types.PointerStruct{})
	new_types.TestDB.AutoMigrate(&new_types.PointerStruct{})
	var name = "pointer struct 1"
	var num = 100
	pointerStruct := new_types.PointerStruct{Name: &name, Num: &num}
	if new_types.TestDB.Create(&pointerStruct).Error != nil {
		t.Errorf("Failed to save pointer struct")
	}

	var pointerStructResult new_types.PointerStruct
	if err := new_types.TestDB.First(&pointerStructResult, "id = ?", pointerStruct.ID).Error; err != nil || *pointerStructResult.Name != name || *pointerStructResult.Num != num {
		t.Errorf("Failed to query saved pointer struct")
	}

	var tableName = new_types.TestDB.NewScope(&new_types.PointerStruct{}).TableName()

	var normalStruct new_types.NormalStruct
	new_types.TestDB.Table(tableName).First(&normalStruct)
	if normalStruct.Name != name || normalStruct.Num != num {
		t.Errorf("Failed to query saved Normal struct")
	}

	var nilPointerStruct = new_types.PointerStruct{}
	if err := new_types.TestDB.Create(&nilPointerStruct).Error; err != nil {
		t.Error("Failed to save nil pointer struct", err)
	}

	var pointerStruct2 new_types.PointerStruct
	if err := new_types.TestDB.First(&pointerStruct2, "id = ?", nilPointerStruct.ID).Error; err != nil {
		t.Error("Failed to query saved nil pointer struct", err)
	}

	var normalStruct2 new_types.NormalStruct
	if err := new_types.TestDB.Table(tableName).First(&normalStruct2, "id = ?", nilPointerStruct.ID).Error; err != nil {
		t.Error("Failed to query saved nil pointer struct", err)
	}

	var partialNilPointerStruct1 = new_types.PointerStruct{Num: &num}
	if err := new_types.TestDB.Create(&partialNilPointerStruct1).Error; err != nil {
		t.Error("Failed to save partial nil pointer struct", err)
	}

	var pointerStruct3 new_types.PointerStruct
	if err := new_types.TestDB.First(&pointerStruct3, "id = ?", partialNilPointerStruct1.ID).Error; err != nil || *pointerStruct3.Num != num {
		t.Error("Failed to query saved partial nil pointer struct", err)
	}

	var normalStruct3 new_types.NormalStruct
	if err := new_types.TestDB.Table(tableName).First(&normalStruct3, "id = ?", partialNilPointerStruct1.ID).Error; err != nil || normalStruct3.Num != num {
		t.Error("Failed to query saved partial pointer struct", err)
	}

	var partialNilPointerStruct2 = new_types.PointerStruct{Name: &name}
	if err := new_types.TestDB.Create(&partialNilPointerStruct2).Error; err != nil {
		t.Error("Failed to save partial nil pointer struct", err)
	}

	var pointerStruct4 new_types.PointerStruct
	if err := new_types.TestDB.First(&pointerStruct4, "id = ?", partialNilPointerStruct2.ID).Error; err != nil || *pointerStruct4.Name != name {
		t.Error("Failed to query saved partial nil pointer struct", err)
	}

	var normalStruct4 new_types.NormalStruct
	if err := new_types.TestDB.Table(tableName).First(&normalStruct4, "id = ?", partialNilPointerStruct2.ID).Error; err != nil || normalStruct4.Name != name {
		t.Error("Failed to query saved partial pointer struct", err)
	}
}
