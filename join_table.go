package tests

import (
	"SideBySideGorm/new_types"
	"testing"
)

func DoJoinTable(t *testing.T) {
	new_types.TestDB.Exec("drop table person_addresses;")
	new_types.TestDB.AutoMigrate(&new_types.Person{})
	new_types.TestDB.SetJoinTableHandler(&new_types.Person{}, "Addresses", &new_types.PersonAddress{})

	address1 := &new_types.Address{Address1: "address 1 of person 1"}
	address2 := &new_types.Address{Address1: "address 2 of person 1"}
	person := &new_types.Person{Name: "person", Addresses: []*new_types.Address{address1, address2}}
	new_types.TestDB.Save(person)
	person2 := &new_types.Person{}
	new_types.TestDB.Model(person).Where(new_types.Person{Id: person.Id}).Related(&person2.Addresses, "Addresses").Find(&person2)
	//TODO : @Badu - seems to me it fails retrieving with relations as I expect it
	new_types.TestDB.Model(person).Association("Addresses").Delete(address1)

	if new_types.TestDB.Find(&[]new_types.PersonAddress{}, "person_id = ?", person.Id).RowsAffected != 1 {
		t.Errorf("Should found one address")
	}

	if new_types.TestDB.Model(person).Association("Addresses").Count() != 1 {
		t.Errorf("Should found one address")
	}

	if new_types.TestDB.Unscoped().Find(&[]new_types.PersonAddress{}, "person_id = ?", person.Id).RowsAffected != 2 {
		t.Errorf("Found two addresses with Unscoped")
	}

	if new_types.TestDB.Model(person).Association("Addresses").Clear(); new_types.TestDB.Model(person).Association("Addresses").Count() != 0 {
		t.Errorf("Should deleted all addresses")
	}
}
