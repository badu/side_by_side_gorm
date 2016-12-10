package tests

import (
	"SideBySideGorm/new_types"
	"fmt"
	newGorm "github.com/badu/gorm"
	"testing"
)

func DoCalculateField(t *testing.T) {
	var field new_types.CalculateField
	var scope = new_types.TestDB.NewScope(&field)
	if field, ok := scope.FieldByName("Children"); !ok || !field.HasRelations() {
		t.Errorf("Should calculate fields correctly for the first time")
	}

	if field, ok := scope.FieldByName("Category"); !ok || !field.HasRelations() {
		t.Errorf("Should calculate fields correctly for the first time")
	}

	if field, ok := scope.FieldByName("EmbeddedName"); !ok {
		t.Errorf("should find embedded field")
	} else if !field.HasSetting(newGorm.NOT_NULL) {
		t.Errorf(fmt.Sprintf("Should find embedded field's tag settings\n%s", field))
	}
}
