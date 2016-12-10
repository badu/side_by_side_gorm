package tests

import (
	"SideBySideGorm/new_types"
	"testing"
)

func Scopes(t *testing.T) {
	user1 := new_types.User{Name: "ScopeUser1", Age: 1}
	user2 := new_types.User{Name: "ScopeUser2", Age: 1}
	user3 := new_types.User{Name: "ScopeUser3", Age: 2}
	new_types.TestDB.Save(&user1).Save(&user2).Save(&user3)

	var users1, users2, users3 []new_types.User
	new_types.TestDB.Scopes(NewNameIn1And2).Find(&users1)
	if len(users1) != 2 {
		t.Errorf("Should found two users's name in 1, 2")
	}

	new_types.TestDB.Scopes(NewNameIn1And2, NewNameIn2And3).Find(&users2)
	if len(users2) != 1 {
		t.Errorf("Should found one user's name is 2")
	}

	new_types.TestDB.Scopes(NewNameIn([]string{user1.Name, user3.Name})).Find(&users3)
	if len(users3) != 2 {
		t.Errorf("Should found two users's name in 1, 3")
	}
}
