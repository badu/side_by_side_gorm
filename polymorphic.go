package tests

import (
	"SideBySideGorm/new_types"
	"testing"
)

func Polymorphic(t *testing.T) {
	cat := new_types.Cat{Name: "Mr. Bigglesworth", Toy: new_types.Toy{Name: "cat toy"}}
	dog := new_types.Dog{Name: "Pluto", Toys: []new_types.Toy{{Name: "dog toy 1"}, {Name: "dog toy 2"}}}
	new_types.TestDB.Save(&cat).Save(&dog)

	if new_types.TestDB.Model(&cat).Association("Toy").Count() != 1 {
		t.Errorf("Cat's toys count should be 1")
	}

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 2 {
		t.Errorf("Dog's toys count should be 2")
	}

	// Query
	var catToys []new_types.Toy
	if new_types.TestDB.Model(&cat).Related(&catToys, "Toy").RecordNotFound() {
		t.Errorf("Did not find any has one polymorphic association")
	} else if len(catToys) != 1 {
		t.Errorf("Should have found only one polymorphic has one association")
	} else if catToys[0].Name != cat.Toy.Name {
		t.Errorf("Should have found the proper has one polymorphic association")
	}

	var dogToys []new_types.Toy
	if new_types.TestDB.Model(&dog).Related(&dogToys, "Toys").RecordNotFound() {
		t.Errorf("Did not find any polymorphic has many associations")
	} else if len(dogToys) != len(dog.Toys) {
		t.Errorf("Should have found all polymorphic has many associations")
	}

	var catToy new_types.Toy
	new_types.TestDB.Model(&cat).Association("Toy").Find(&catToy)
	if catToy.Name != cat.Toy.Name {
		t.Errorf("Should find has one polymorphic association")
	}

	var dogToys1 []new_types.Toy
	new_types.TestDB.Model(&dog).Association("Toys").Find(&dogToys1)
	if !newCompareToys(dogToys1, []string{"dog toy 1", "dog toy 2"}) {
		t.Errorf("Should find has many polymorphic association")
	}

	// Append
	new_types.TestDB.Model(&cat).Association("Toy").Append(&new_types.Toy{
		Name: "cat toy 2",
	})

	var catToy2 new_types.Toy
	new_types.TestDB.Model(&cat).Association("Toy").Find(&catToy2)
	if catToy2.Name != "cat toy 2" {
		t.Errorf("Should update has one polymorphic association with Append")
	}

	if new_types.TestDB.Model(&cat).Association("Toy").Count() != 1 {
		t.Errorf("Cat's toys count should be 1 after Append")
	}

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 2 {
		t.Errorf("Should return two polymorphic has many associations")
	}

	new_types.TestDB.Model(&dog).Association("Toys").Append(&new_types.Toy{
		Name: "dog toy 3",
	})

	var dogToys2 []new_types.Toy
	new_types.TestDB.Model(&dog).Association("Toys").Find(&dogToys2)
	if !newCompareToys(dogToys2, []string{"dog toy 1", "dog toy 2", "dog toy 3"}) {
		t.Errorf("Dog's toys should be updated with Append")
	}

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 3 {
		t.Errorf("Should return three polymorphic has many associations")
	}

	// Replace
	new_types.TestDB.Model(&cat).Association("Toy").Replace(&new_types.Toy{
		Name: "cat toy 3",
	})

	var catToy3 new_types.Toy
	new_types.TestDB.Model(&cat).Association("Toy").Find(&catToy3)
	if catToy3.Name != "cat toy 3" {
		t.Errorf("Should update has one polymorphic association with Replace")
	}

	if new_types.TestDB.Model(&cat).Association("Toy").Count() != 1 {
		t.Errorf("Cat's toys count should be 1 after Replace")
	}

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 3 {
		t.Errorf("Should return three polymorphic has many associations")
	}

	new_types.TestDB.Model(&dog).Association("Toys").Replace(&new_types.Toy{
		Name: "dog toy 4",
	}, []new_types.Toy{
		{Name: "dog toy 5"}, {Name: "dog toy 6"}, {Name: "dog toy 7"},
	})

	var dogToys3 []new_types.Toy
	new_types.TestDB.Model(&dog).Association("Toys").Find(&dogToys3)
	if !newCompareToys(dogToys3, []string{"dog toy 4", "dog toy 5", "dog toy 6", "dog toy 7"}) {
		t.Errorf("Dog's toys should be updated with Replace")
	}

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 4 {
		t.Errorf("Should return three polymorphic has many associations")
	}

	// Delete
	new_types.TestDB.Model(&cat).Association("Toy").Delete(&catToy2)

	var catToy4 new_types.Toy
	new_types.TestDB.Model(&cat).Association("Toy").Find(&catToy4)
	if catToy4.Name != "cat toy 3" {
		t.Errorf("Should not update has one polymorphic association when Delete a unrelated Toy")
	}

	if new_types.TestDB.Model(&cat).Association("Toy").Count() != 1 {
		t.Errorf("Cat's toys count should be 1")
	}

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 4 {
		t.Errorf("Dog's toys count should be 4")
	}

	new_types.TestDB.Model(&cat).Association("Toy").Delete(&catToy3)

	if !new_types.TestDB.Model(&cat).Related(&new_types.Toy{}, "Toy").RecordNotFound() {
		t.Errorf("Toy should be deleted with Delete")
	}

	if new_types.TestDB.Model(&cat).Association("Toy").Count() != 0 {
		t.Errorf("Cat's toys count should be 0 after Delete")
	}

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 4 {
		t.Errorf("Dog's toys count should not be changed when delete cat's toy")
	}

	new_types.TestDB.Model(&dog).Association("Toys").Delete(&dogToys2)

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 4 {
		t.Errorf("Dog's toys count should not be changed when delete unrelated toys")
	}

	new_types.TestDB.Model(&dog).Association("Toys").Delete(&dogToys3)

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 0 {
		t.Errorf("Dog's toys count should be deleted with Delete")
	}

	// Clear
	new_types.TestDB.Model(&cat).Association("Toy").Append(&new_types.Toy{
		Name: "cat toy 2",
	})

	if new_types.TestDB.Model(&cat).Association("Toy").Count() != 1 {
		t.Errorf("Cat's toys should be added with Append")
	}

	new_types.TestDB.Model(&cat).Association("Toy").Clear()

	if new_types.TestDB.Model(&cat).Association("Toy").Count() != 0 {
		t.Errorf("Cat's toys should be cleared with Clear")
	}

	new_types.TestDB.Model(&dog).Association("Toys").Append(&new_types.Toy{
		Name: "dog toy 8",
	})

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 1 {
		t.Errorf("Dog's toys should be added with Append")
	}

	new_types.TestDB.Model(&dog).Association("Toys").Clear()

	if new_types.TestDB.Model(&dog).Association("Toys").Count() != 0 {
		t.Errorf("Dog's toys should be cleared with Clear")
	}
}

func NamedPolymorphic(t *testing.T) {
	hamster := new_types.Hamster{Name: "Mr. Hammond", PreferredToy: new_types.Toy{Name: "bike"}, OtherToy: new_types.Toy{Name: "treadmill"}}
	new_types.TestDB.Save(&hamster)

	hamster2 := new_types.Hamster{}
	new_types.TestDB.Preload("PreferredToy").Preload("OtherToy").Find(&hamster2, hamster.Id)
	if hamster2.PreferredToy.Id != hamster.PreferredToy.Id || hamster2.PreferredToy.Name != hamster.PreferredToy.Name {
		t.Errorf("Hamster's preferred toy couldn't be preloaded")
	}
	if hamster2.OtherToy.Id != hamster.OtherToy.Id || hamster2.OtherToy.Name != hamster.OtherToy.Name {
		t.Errorf("Hamster's other toy couldn't be preloaded")
	}

	// clear to omit Toy.Id in count
	hamster2.PreferredToy = new_types.Toy{}
	hamster2.OtherToy = new_types.Toy{}

	if new_types.TestDB.Model(&hamster2).Association("PreferredToy").Count() != 1 {
		t.Errorf("Hamster's preferred toy count should be 1")
	}

	if new_types.TestDB.Model(&hamster2).Association("OtherToy").Count() != 1 {
		t.Errorf("Hamster's other toy count should be 1")
	}

	// Query
	var hamsterToys []new_types.Toy
	if new_types.TestDB.Model(&hamster).Related(&hamsterToys, "PreferredToy").RecordNotFound() {
		t.Errorf("Did not find any has one polymorphic association")
	} else if len(hamsterToys) != 1 {
		t.Errorf("Should have found only one polymorphic has one association")
	} else if hamsterToys[0].Name != hamster.PreferredToy.Name {
		t.Errorf("Should have found the proper has one polymorphic association")
	}

	if new_types.TestDB.Model(&hamster).Related(&hamsterToys, "OtherToy").RecordNotFound() {
		t.Errorf("Did not find any has one polymorphic association")
	} else if len(hamsterToys) != 1 {
		t.Errorf("Should have found only one polymorphic has one association")
	} else if hamsterToys[0].Name != hamster.OtherToy.Name {
		t.Errorf("Should have found the proper has one polymorphic association")
	}

	hamsterToy := new_types.Toy{}
	new_types.TestDB.Model(&hamster).Association("PreferredToy").Find(&hamsterToy)
	if hamsterToy.Name != hamster.PreferredToy.Name {
		t.Errorf("Should find has one polymorphic association")
	}
	hamsterToy = new_types.Toy{}
	new_types.TestDB.Model(&hamster).Association("OtherToy").Find(&hamsterToy)
	if hamsterToy.Name != hamster.OtherToy.Name {
		t.Errorf("Should find has one polymorphic association")
	}

	// Append
	new_types.TestDB.Model(&hamster).Association("PreferredToy").Append(&new_types.Toy{
		Name: "bike 2",
	})
	new_types.TestDB.Model(&hamster).Association("OtherToy").Append(&new_types.Toy{
		Name: "treadmill 2",
	})

	hamsterToy = new_types.Toy{}
	new_types.TestDB.Model(&hamster).Association("PreferredToy").Find(&hamsterToy)
	if hamsterToy.Name != "bike 2" {
		t.Errorf("Should update has one polymorphic association with Append")
	}

	hamsterToy = new_types.Toy{}
	new_types.TestDB.Model(&hamster).Association("OtherToy").Find(&hamsterToy)
	if hamsterToy.Name != "treadmill 2" {
		t.Errorf("Should update has one polymorphic association with Append")
	}

	if new_types.TestDB.Model(&hamster2).Association("PreferredToy").Count() != 1 {
		t.Errorf("Hamster's toys count should be 1 after Append")
	}

	if new_types.TestDB.Model(&hamster2).Association("OtherToy").Count() != 1 {
		t.Errorf("Hamster's toys count should be 1 after Append")
	}

	// Replace
	new_types.TestDB.Model(&hamster).Association("PreferredToy").Replace(&new_types.Toy{
		Name: "bike 3",
	})
	new_types.TestDB.Model(&hamster).Association("OtherToy").Replace(&new_types.Toy{
		Name: "treadmill 3",
	})

	hamsterToy = new_types.Toy{}
	new_types.TestDB.Model(&hamster).Association("PreferredToy").Find(&hamsterToy)
	if hamsterToy.Name != "bike 3" {
		t.Errorf("Should update has one polymorphic association with Replace")
	}

	hamsterToy = new_types.Toy{}
	new_types.TestDB.Model(&hamster).Association("OtherToy").Find(&hamsterToy)
	if hamsterToy.Name != "treadmill 3" {
		t.Errorf("Should update has one polymorphic association with Replace")
	}

	if new_types.TestDB.Model(&hamster2).Association("PreferredToy").Count() != 1 {
		t.Errorf("hamster's toys count should be 1 after Replace")
	}

	if new_types.TestDB.Model(&hamster2).Association("OtherToy").Count() != 1 {
		t.Errorf("hamster's toys count should be 1 after Replace")
	}

	// Clear
	new_types.TestDB.Model(&hamster).Association("PreferredToy").Append(&new_types.Toy{
		Name: "bike 2",
	})
	new_types.TestDB.Model(&hamster).Association("OtherToy").Append(&new_types.Toy{
		Name: "treadmill 2",
	})

	if new_types.TestDB.Model(&hamster).Association("PreferredToy").Count() != 1 {
		t.Errorf("Hamster's toys should be added with Append")
	}
	if new_types.TestDB.Model(&hamster).Association("OtherToy").Count() != 1 {
		t.Errorf("Hamster's toys should be added with Append")
	}

	new_types.TestDB.Model(&hamster).Association("PreferredToy").Clear()

	if new_types.TestDB.Model(&hamster2).Association("PreferredToy").Count() != 0 {
		t.Errorf("Hamster's preferred toy should be cleared with Clear")
	}
	if new_types.TestDB.Model(&hamster2).Association("OtherToy").Count() != 1 {
		t.Errorf("Hamster's other toy should be still available")
	}

	new_types.TestDB.Model(&hamster).Association("OtherToy").Clear()
	if new_types.TestDB.Model(&hamster).Association("OtherToy").Count() != 0 {
		t.Errorf("Hamster's other toy should be cleared with Clear")
	}
}
