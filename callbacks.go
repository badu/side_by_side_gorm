package tests

import (
	"SideBySideGorm/new_types"
	"reflect"
	"testing"
)

func RunCallbacks(t *testing.T) {
	p := new_types.Product{Code: "unique_code", Price: 100}
	new_types.TestDB.Save(&p)

	if !reflect.DeepEqual(p.GetCallTimes(), []int64{1, 1, 0, 1, 1, 0, 0, 0, 0}) {
		t.Errorf("Callbacks should be invoked successfully, %v", p.GetCallTimes())
	}

	new_types.TestDB.Where("Code = ?", "unique_code").First(&p)
	if !reflect.DeepEqual(p.GetCallTimes(), []int64{1, 1, 0, 1, 0, 0, 0, 0, 1}) {
		t.Errorf("After callbacks values are not saved, %v", p.GetCallTimes())
	}

	p.Price = 200
	new_types.TestDB.Save(&p)
	if !reflect.DeepEqual(p.GetCallTimes(), []int64{1, 2, 1, 1, 1, 1, 0, 0, 1}) {
		t.Errorf("After update callbacks should be invoked successfully, %v", p.GetCallTimes())
	}

	var products []new_types.Product
	new_types.TestDB.Find(&products, "code = ?", "unique_code")
	if products[0].AfterFindCallTimes != 2 {
		t.Errorf("AfterFind callbacks should work with slice")
	}

	new_types.TestDB.Where("Code = ?", "unique_code").First(&p)
	if !reflect.DeepEqual(p.GetCallTimes(), []int64{1, 2, 1, 1, 0, 0, 0, 0, 2}) {
		t.Errorf("After update callbacks values are not saved, %v", p.GetCallTimes())
	}

	new_types.TestDB.Delete(&p)
	if !reflect.DeepEqual(p.GetCallTimes(), []int64{1, 2, 1, 1, 0, 0, 1, 1, 2}) {
		t.Errorf("After delete callbacks should be invoked successfully, %v", p.GetCallTimes())
	}

	if new_types.TestDB.Where("Code = ?", "unique_code").First(&p).Error == nil {
		t.Errorf("Can't find a deleted record")
	}
}

func CallbacksWithErrors(t *testing.T) {
	p := new_types.Product{Code: "Invalid", Price: 100}
	if new_types.TestDB.Save(&p).Error == nil {
		t.Errorf("An error from before create callbacks happened when create with invalid value")
	}

	if new_types.TestDB.Where("code = ?", "Invalid").First(&new_types.Product{}).Error == nil {
		t.Errorf("Should not save record that have errors")
	}

	if new_types.TestDB.Save(&new_types.Product{Code: "dont_save", Price: 100}).Error == nil {
		t.Errorf("An error from after create callbacks happened when create with invalid value")
	}

	p2 := new_types.Product{Code: "update_callback", Price: 100}
	new_types.TestDB.Save(&p2)

	p2.Code = "dont_update"
	if new_types.TestDB.Save(&p2).Error == nil {
		t.Errorf("An error from before update callbacks happened when update with invalid value")
	}

	if new_types.TestDB.Where("code = ?", "update_callback").First(&new_types.Product{}).Error != nil {
		t.Errorf("Record Should not be updated due to errors happened in before update callback")
	}

	if new_types.TestDB.Where("code = ?", "dont_update").First(&new_types.Product{}).Error == nil {
		t.Errorf("Record Should not be updated due to errors happened in before update callback")
	}

	p2.Code = "dont_save"
	if new_types.TestDB.Save(&p2).Error == nil {
		t.Errorf("An error from before save callbacks happened when update with invalid value")
	}

	p3 := new_types.Product{Code: "dont_delete", Price: 100}
	new_types.TestDB.Save(&p3)
	if new_types.TestDB.Delete(&p3).Error == nil {
		t.Errorf("An error from before delete callbacks happened when delete")
	}

	if new_types.TestDB.Where("Code = ?", "dont_delete").First(&p3).Error != nil {
		t.Errorf("An error from before delete callbacks happened")
	}

	p4 := new_types.Product{Code: "after_save_error", Price: 100}
	new_types.TestDB.Save(&p4)
	if err := new_types.TestDB.First(&new_types.Product{}, "code = ?", "after_save_error").Error; err == nil {
		t.Errorf("Record should be reverted if get an error in after save callback")
	}

	p5 := new_types.Product{Code: "after_delete_error", Price: 100}
	new_types.TestDB.Save(&p5)
	if err := new_types.TestDB.First(&new_types.Product{}, "code = ?", "after_delete_error").Error; err != nil {
		t.Errorf("Record should be found")
	}

	new_types.TestDB.Delete(&p5)
	if err := new_types.TestDB.First(&new_types.Product{}, "code = ?", "after_delete_error").Error; err != nil {
		t.Errorf("Record shouldn't be deleted because of an error happened in after delete callback")
	}
}
