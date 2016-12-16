package tests

import (
	"SideBySideGorm/new_types"
	"fmt"
	newGorm "github.com/badu/gorm"
	"os"
	"reflect"
	"sort"
	"testing"
)

func SkipSaveAssociation(t *testing.T) {
	type Company struct {
		newGorm.Model
		Name string
	}

	type User struct {
		newGorm.Model
		Name      string
		CompanyID uint
		Company   Company `gorm:"save_associations:false"`
	}
	new_types.TestDB.AutoMigrate(&Company{}, &User{})

	new_types.TestDB.Save(&User{Name: "jinzhu", Company: Company{Name: "skip_save_association"}})

	if !new_types.TestDB.Where("name = ?", "skip_save_association").First(&Company{}).RecordNotFound() {
		t.Errorf("Company skip_save_association should not been saved")
	}
}

func BelongsTo(t *testing.T) {
	post := new_types.Post{
		Title:        "post belongs to",
		Body:         "body belongs to",
		Category:     new_types.Category{Name: "Category 1"},
		MainCategory: new_types.Category{Name: "Main Category 1"},
	}

	if err := new_types.TestDB.Save(&post).Error; err != nil {
		t.Error("Got errors when save post", err)
	}

	if post.Category.ID == 0 || post.MainCategory.ID == 0 {
		t.Errorf("Category's primary key should be updated")
	}

	if post.CategoryId.Int64 == 0 || post.MainCategoryId == 0 {
		t.Errorf("post's foreign key should be updated")
	}

	// Query
	var category1 new_types.Category
	new_types.TestDB.Model(&post).Association("Category").Find(&category1)
	if category1.Name != "Category 1" {
		t.Errorf("Query belongs to relations with Association")
	}

	var mainCategory1 new_types.Category
	new_types.TestDB.Model(&post).Association("MainCategory").Find(&mainCategory1)
	if mainCategory1.Name != "Main Category 1" {
		t.Errorf("Query belongs to relations with Association")
	}

	var category11 new_types.Category
	new_types.TestDB.Model(&post).Related(&category11)
	if category11.Name != "Category 1" {
		t.Errorf("Query belongs to relations with Related")
	}

	if new_types.TestDB.Model(&post).Association("Category").Count() != 1 {
		t.Errorf("Post's category count should be 1")
	}

	if new_types.TestDB.Model(&post).Association("MainCategory").Count() != 1 {
		t.Errorf("Post's main category count should be 1")
	}

	// Append
	var category2 = new_types.Category{
		Name: "Category 2",
	}
	new_types.TestDB.Model(&post).Association("Category").Append(&category2)

	if category2.ID == 0 {
		t.Errorf("Category should has ID when created with Append")
	}

	var category21 new_types.Category
	new_types.TestDB.Model(&post).Related(&category21)

	if category21.Name != "Category 2" {
		t.Errorf("Category should be updated with Append")
	}

	if new_types.TestDB.Model(&post).Association("Category").Count() != 1 {
		t.Errorf("Post's category count should be 1")
	}

	// Replace
	var category3 = new_types.Category{
		Name: "Category 3",
	}
	new_types.TestDB.Model(&post).Association("Category").Replace(&category3)

	if category3.ID == 0 {
		t.Errorf("Category should has ID when created with Replace")
	}

	var category31 new_types.Category
	new_types.TestDB.Model(&post).Related(&category31)
	if category31.Name != "Category 3" {
		t.Errorf("Category should be updated with Replace")
	}

	if new_types.TestDB.Model(&post).Association("Category").Count() != 1 {
		t.Errorf("Post's category count should be 1")
	}

	// Delete
	new_types.TestDB.Model(&post).Association("Category").Delete(&category2)
	if new_types.TestDB.Model(&post).Related(&new_types.Category{}).RecordNotFound() {
		t.Errorf("Should not delete any category when Delete a unrelated Category")
	}

	if post.Category.Name == "" {
		t.Errorf("Post's category should not be reseted when Delete a unrelated Category")
	}

	new_types.TestDB.Model(&post).Association("Category").Delete(&category3)

	if post.Category.Name != "" {
		t.Errorf("Post's category should be reseted after Delete")
	}

	var category41 new_types.Category
	new_types.TestDB.Model(&post).Related(&category41)
	if category41.Name != "" {
		t.Errorf("Category should be deleted with Delete")
	}

	if count := new_types.TestDB.Model(&post).Association("Category").Count(); count != 0 {
		t.Errorf("Post's category count should be 0 after Delete, but got %v", count)
	}

	// Clear
	new_types.TestDB.Model(&post).Association("Category").Append(&new_types.Category{
		Name: "Category 2",
	})

	if new_types.TestDB.Model(&post).Related(&new_types.Category{}).RecordNotFound() {
		t.Errorf("Should find category after append")
	}

	if post.Category.Name == "" {
		t.Errorf("Post's category should has value after Append")
	}

	new_types.TestDB.Model(&post).Association("Category").Clear()

	if post.Category.Name != "" {
		t.Errorf("Post's category should be cleared after Clear")
	}

	if !new_types.TestDB.Model(&post).Related(&new_types.Category{}).RecordNotFound() {
		t.Errorf("Should not find any category after Clear")
	}

	if count := new_types.TestDB.Model(&post).Association("Category").Count(); count != 0 {
		t.Errorf("Post's category count should be 0 after Clear, but got %v", count)
	}

	// Check Association mode with soft delete
	category6 := new_types.Category{
		Name: "Category 6",
	}
	new_types.TestDB.Model(&post).Association("Category").Append(&category6)

	if count := new_types.TestDB.Model(&post).Association("Category").Count(); count != 1 {
		t.Errorf("Post's category count should be 1 after Append, but got %v", count)
	}

	new_types.TestDB.Delete(&category6)

	if count := new_types.TestDB.Model(&post).Association("Category").Count(); count != 0 {
		t.Errorf("Post's category count should be 0 after the category has been deleted, but got %v", count)
	}

	if err := new_types.TestDB.Model(&post).Association("Category").Find(&new_types.Category{}).Error; err == nil {
		t.Errorf("Post's category is not findable after Delete")
	}

	if count := new_types.TestDB.Unscoped().Model(&post).Association("Category").Count(); count != 1 {
		t.Errorf("Post's category count should be 1 when query with Unscoped, but got %v", count)
	}

	if err := new_types.TestDB.Unscoped().Model(&post).Association("Category").Find(&new_types.Category{}).Error; err != nil {
		t.Errorf("Post's category should be findable when query with Unscoped, got %v", err)
	}
}

func BelongsToOverrideForeignKey1(t *testing.T) {
	type Profile struct {
		newGorm.Model
		Name string
	}

	type User struct {
		newGorm.Model
		Profile      Profile `gorm:"ForeignKey:ProfileRefer"`
		ProfileRefer int
	}

	if field, ok := new_types.TestDB.NewScope(&User{}).FieldByName("Profile"); ok {

		ForeignFieldNames := field.GetForeignFieldNames()
		AssociationForeignFieldNames := field.GetAssociationForeignFieldNames()

		if !field.RelationIsBelongsTo() ||
			!reflect.DeepEqual(ForeignFieldNames, newGorm.StrSlice{"ProfileRefer"}) ||
			!reflect.DeepEqual(AssociationForeignFieldNames, newGorm.StrSlice{"ID"}) {
			t.Errorf("Override belongs to foreign key with tag")
		}
	}
}

func BelongsToOverrideForeignKey2(t *testing.T) {
	type Profile struct {
		newGorm.Model
		Refer string
		Name  string
	}

	type User struct {
		newGorm.Model
		Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:Refer"`
		ProfileID int
	}

	if field, ok := new_types.TestDB.NewScope(&User{}).FieldByName("Profile"); ok {

		ForeignFieldNames := field.GetForeignFieldNames()
		AssociationForeignFieldNames := field.GetAssociationForeignFieldNames()

		if !field.RelationIsBelongsTo()  ||
			!reflect.DeepEqual(ForeignFieldNames, newGorm.StrSlice{"ProfileID"}) ||
			!reflect.DeepEqual(AssociationForeignFieldNames, newGorm.StrSlice{"Refer"}) {
			t.Errorf("Override belongs to foreign key with tag")
		}
	}
}

func HasOne(t *testing.T) {
	user := new_types.User{
		Name:       "has one",
		CreditCard: new_types.CreditCard{Number: "411111111111"},
	}

	if err := new_types.TestDB.Save(&user).Error; err != nil {
		t.Error("Got errors when save user", err.Error())
	}

	if user.CreditCard.UserId.Int64 == 0 {
		t.Errorf("CreditCard's foreign key should be updated")
	}

	// Query
	var creditCard1 new_types.CreditCard
	new_types.TestDB.Model(&user).Related(&creditCard1)

	if creditCard1.Number != "411111111111" {
		t.Errorf("Query has one relations with Related")
	}

	var creditCard11 new_types.CreditCard
	new_types.TestDB.Model(&user).Association("CreditCard").Find(&creditCard11)

	if creditCard11.Number != "411111111111" {
		t.Errorf("Query has one relations with Related")
	}

	if new_types.TestDB.Model(&user).Association("CreditCard").Count() != 1 {
		t.Errorf("User's credit card count should be 1")
	}

	// Append
	var creditcard2 = new_types.CreditCard{
		Number: "411111111112",
	}
	new_types.TestDB.Model(&user).Association("CreditCard").Append(&creditcard2)

	if creditcard2.ID == 0 {
		t.Errorf("Creditcard should has ID when created with Append")
	}

	var creditcard21 new_types.CreditCard
	new_types.TestDB.Model(&user).Related(&creditcard21)
	if creditcard21.Number != "411111111112" {
		t.Errorf("CreditCard should be updated with Append")
	}

	if new_types.TestDB.Model(&user).Association("CreditCard").Count() != 1 {
		t.Errorf("User's credit card count should be 1")
	}

	// Replace
	var creditcard3 = new_types.CreditCard{
		Number: "411111111113",
	}
	new_types.TestDB.Model(&user).Association("CreditCard").Replace(&creditcard3)

	if creditcard3.ID == 0 {
		t.Errorf("Creditcard should has ID when created with Replace")
	}

	var creditcard31 new_types.CreditCard
	new_types.TestDB.Model(&user).Related(&creditcard31)
	if creditcard31.Number != "411111111113" {
		t.Errorf("CreditCard should be updated with Replace")
	}

	if new_types.TestDB.Model(&user).Association("CreditCard").Count() != 1 {
		t.Errorf("User's credit card count should be 1")
	}

	// Delete
	new_types.TestDB.Model(&user).Association("CreditCard").Delete(&creditcard2)
	var creditcard4 new_types.CreditCard
	new_types.TestDB.Model(&user).Related(&creditcard4)
	if creditcard4.Number != "411111111113" {
		t.Errorf("Should not delete credit card when Delete a unrelated CreditCard")
	}

	if new_types.TestDB.Model(&user).Association("CreditCard").Count() != 1 {
		t.Errorf("User's credit card count should be 1")
	}

	new_types.TestDB.Model(&user).Association("CreditCard").Delete(&creditcard3)
	if !new_types.TestDB.Model(&user).Related(&new_types.CreditCard{}).RecordNotFound() {
		t.Errorf("Should delete credit card with Delete")
	}

	if new_types.TestDB.Model(&user).Association("CreditCard").Count() != 0 {
		t.Errorf("User's credit card count should be 0 after Delete")
	}

	// Clear
	var creditcard5 = new_types.CreditCard{
		Number: "411111111115",
	}
	new_types.TestDB.Model(&user).Association("CreditCard").Append(&creditcard5)

	if new_types.TestDB.Model(&user).Related(&new_types.CreditCard{}).RecordNotFound() {
		t.Errorf("Should added credit card with Append")
	}

	if new_types.TestDB.Model(&user).Association("CreditCard").Count() != 1 {
		t.Errorf("User's credit card count should be 1")
	}

	new_types.TestDB.Model(&user).Association("CreditCard").Clear()
	if !new_types.TestDB.Model(&user).Related(&new_types.CreditCard{}).RecordNotFound() {
		t.Errorf("Credit card should be deleted with Clear")
	}

	if new_types.TestDB.Model(&user).Association("CreditCard").Count() != 0 {
		t.Errorf("User's credit card count should be 0 after Clear")
	}

	// Check Association mode with soft delete
	var creditcard6 = new_types.CreditCard{
		Number: "411111111116",
	}
	new_types.TestDB.Model(&user).Association("CreditCard").Append(&creditcard6)

	if count := new_types.TestDB.Model(&user).Association("CreditCard").Count(); count != 1 {
		t.Errorf("User's credit card count should be 1 after Append, but got %v", count)
	}

	new_types.TestDB.Delete(&creditcard6)

	if count := new_types.TestDB.Model(&user).Association("CreditCard").Count(); count != 0 {
		t.Errorf("User's credit card count should be 0 after credit card deleted, but got %v", count)
	}

	if err := new_types.TestDB.Model(&user).Association("CreditCard").Find(&new_types.CreditCard{}).Error; err == nil {
		t.Errorf("User's creditcard is not findable after Delete")
	}

	if count := new_types.TestDB.Unscoped().Model(&user).Association("CreditCard").Count(); count != 1 {
		t.Errorf("User's credit card count should be 1 when query with Unscoped, but got %v", count)
	}

	if err := new_types.TestDB.Unscoped().Model(&user).Association("CreditCard").Find(&new_types.CreditCard{}).Error; err != nil {
		t.Errorf("User's creditcard should be findable when query with Unscoped, got %v", err)
	}
}

func HasOneOverrideForeignKey1(t *testing.T) {
	type Profile struct {
		newGorm.Model
		Name      string
		UserRefer uint
	}

	type User struct {
		newGorm.Model
		Profile Profile `gorm:"ForeignKey:UserRefer"`
	}

	if field, ok := new_types.TestDB.NewScope(&User{}).FieldByName("Profile"); ok {


		ForeignFieldNames := field.GetForeignFieldNames()
		AssociationForeignFieldNames := field.GetAssociationForeignFieldNames()
		if !field.RelationIsHasOne() ||
			!reflect.DeepEqual(ForeignFieldNames, newGorm.StrSlice{"UserRefer"}) ||
			!reflect.DeepEqual(AssociationForeignFieldNames, newGorm.StrSlice{"ID"}) {
			t.Errorf("Override belongs to foreign key with tag")
		}
	}
}

func HasOneOverrideForeignKey2(t *testing.T) {
	type Profile struct {
		newGorm.Model
		Name   string
		UserID uint
	}

	type User struct {
		newGorm.Model
		Refer   string
		Profile Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	}

	if field, ok := new_types.TestDB.NewScope(&User{}).FieldByName("Profile"); ok {

		ForeignFieldNames := field.GetForeignFieldNames()
		AssociationForeignFieldNames := field.GetAssociationForeignFieldNames()

		if !field.RelationIsHasOne() ||
			!reflect.DeepEqual(ForeignFieldNames, newGorm.StrSlice{"UserID"}) ||
			!reflect.DeepEqual(AssociationForeignFieldNames, newGorm.StrSlice{"Refer"}) {
			t.Errorf("Override belongs to foreign key with tag")
		}
	}
}

func HasMany(t *testing.T) {
	post := new_types.Post{
		Title:    "post has many",
		Body:     "body has many",
		Comments: []*new_types.Comment{{Content: "Comment 1"}, {Content: "Comment 2"}},
	}

	if err := new_types.TestDB.Save(&post).Error; err != nil {
		t.Error("Got errors when save post")
		t.Errorf("ERROR : %v", err)
	}

	for _, comment := range post.Comments {
		if comment.PostId == 0 {
			t.Errorf("comment's PostID should be updated")
		}
	}

	var compareComments = func(comments []new_types.Comment, contents []string) bool {
		var commentContents []string
		for _, comment := range comments {
			commentContents = append(commentContents, comment.Content)
		}
		sort.Strings(commentContents)
		sort.Strings(contents)
		return reflect.DeepEqual(commentContents, contents)
	}

	// Query
	if new_types.TestDB.First(&new_types.Comment{}, "content = ?", "Comment 1").Error != nil {
		t.Errorf("Comment 1 should be saved")
	}

	var comments1 []new_types.Comment
	new_types.TestDB.Model(&post).Association("Comments").Find(&comments1)
	if !compareComments(comments1, []string{"Comment 1", "Comment 2"}) {
		t.Errorf("Query has many relations with Association")
	}

	var comments11 []new_types.Comment
	new_types.TestDB.Model(&post).Related(&comments11)
	if !compareComments(comments11, []string{"Comment 1", "Comment 2"}) {
		t.Errorf("Query has many relations with Related")
	}

	if new_types.TestDB.Model(&post).Association("Comments").Count() != 2 {
		t.Errorf("Post's comments count should be 2")
	}

	// Append
	new_types.TestDB.Model(&post).Association("Comments").Append(&new_types.Comment{Content: "Comment 3"})

	var comments2 []new_types.Comment
	new_types.TestDB.Model(&post).Related(&comments2)
	if !compareComments(comments2, []string{"Comment 1", "Comment 2", "Comment 3"}) {
		t.Errorf("Append new record to has many relations")
	}

	if new_types.TestDB.Model(&post).Association("Comments").Count() != 3 {
		t.Errorf("Post's comments count should be 3 after Append")
	}

	// Delete
	new_types.TestDB.Model(&post).Association("Comments").Delete(comments11)

	var comments3 []new_types.Comment
	new_types.TestDB.Model(&post).Related(&comments3)
	if !compareComments(comments3, []string{"Comment 3"}) {
		t.Errorf("Delete an existing resource for has many relations")
	}

	if new_types.TestDB.Model(&post).Association("Comments").Count() != 1 {
		t.Errorf("Post's comments count should be 1 after Delete 2")
	}

	// Replace
	new_types.TestDB.Model(&new_types.Post{Id: 999}).Association("Comments").Replace()

	var comments4 []new_types.Comment
	new_types.TestDB.Model(&post).Related(&comments4)
	if len(comments4) == 0 {
		t.Errorf("Replace for other resource should not clear all comments")
	}

	new_types.TestDB.Model(&post).Association("Comments").Replace(&new_types.Comment{Content: "Comment 4"}, &new_types.Comment{Content: "Comment 5"})

	var comments41 []new_types.Comment
	new_types.TestDB.Model(&post).Related(&comments41)
	if !compareComments(comments41, []string{"Comment 4", "Comment 5"}) {
		t.Errorf("Replace has many relations")
	}

	// Clear
	new_types.TestDB.Model(&new_types.Post{Id: 999}).Association("Comments").Clear()

	var comments5 []new_types.Comment
	new_types.TestDB.Model(&post).Related(&comments5)
	if len(comments5) == 0 {
		t.Errorf("Clear should not clear all comments")
	}

	new_types.TestDB.Model(&post).Association("Comments").Clear()

	var comments51 []new_types.Comment
	new_types.TestDB.Model(&post).Related(&comments51)
	if len(comments51) != 0 {
		t.Errorf("Clear has many relations")
	}

	// Check Association mode with soft delete
	var comment6 = new_types.Comment{
		Content: "comment 6",
	}
	new_types.TestDB.Model(&post).Association("Comments").Append(&comment6)

	if count := new_types.TestDB.Model(&post).Association("Comments").Count(); count != 1 {
		t.Errorf("post's comments count should be 1 after Append, but got %v", count)
	}

	new_types.TestDB.Delete(&comment6)

	if count := new_types.TestDB.Model(&post).Association("Comments").Count(); count != 0 {
		t.Errorf("post's comments count should be 0 after comment been deleted, but got %v", count)
	}

	var comments6 []new_types.Comment
	if new_types.TestDB.Model(&post).Association("Comments").Find(&comments6); len(comments6) != 0 {
		t.Errorf("post's comments count should be 0 when find with Find, but got %v", len(comments6))
	}

	if count := new_types.TestDB.Unscoped().Model(&post).Association("Comments").Count(); count != 1 {
		t.Errorf("post's comments count should be 1 when query with Unscoped, but got %v", count)
	}

	var comments61 []new_types.Comment
	if new_types.TestDB.Unscoped().Model(&post).Association("Comments").Find(&comments61); len(comments61) != 1 {
		t.Errorf("post's comments count should be 1 when query with Unscoped, but got %v", len(comments61))
	}
}

func HasManyOverrideForeignKey1(t *testing.T) {
	type Profile struct {
		newGorm.Model
		Name      string
		UserRefer uint
	}

	type User struct {
		newGorm.Model
		Profile []Profile `gorm:"ForeignKey:UserRefer"`
	}

	if field, ok := new_types.TestDB.NewScope(&User{}).FieldByName("Profile"); ok {

		ForeignFieldNames := field.GetForeignFieldNames()
		AssociationForeignFieldNames := field.GetAssociationForeignFieldNames()

		if !field.RelationIsHasMany() ||
			!reflect.DeepEqual(ForeignFieldNames, newGorm.StrSlice{"UserRefer"}) ||
			!reflect.DeepEqual(AssociationForeignFieldNames, newGorm.StrSlice{"ID"}) {
			t.Errorf("Override belongs to foreign key with tag")
		}
	}
}

func HasManyOverrideForeignKey2(t *testing.T) {
	type Profile struct {
		newGorm.Model
		Name   string
		UserID uint
	}

	type User struct {
		newGorm.Model
		Refer   string
		Profile []Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	}

	if field, ok := new_types.TestDB.NewScope(&User{}).FieldByName("Profile"); ok {

		ForeignFieldNames := field.GetForeignFieldNames()
		AssociationForeignFieldNames := field.GetAssociationForeignFieldNames()

		if !field.RelationIsHasMany() ||
			!reflect.DeepEqual(ForeignFieldNames, newGorm.StrSlice{"UserID"}) ||
			!reflect.DeepEqual(AssociationForeignFieldNames, newGorm.StrSlice{"Refer"}) {
			t.Errorf("Override belongs to foreign key with tag")
		}
	}
}

func ManyToMany(t *testing.T) {
	new_types.TestDB.Raw("delete from languages")
	var languages = []new_types.Language{{Name: "ZH"}, {Name: "EN"}}
	user := new_types.User{Name: "Many2Many", Languages: languages}
	new_types.TestDB.Save(&user)

	// Query
	var newLanguages []new_types.Language
	new_types.TestDB.Model(&user).Related(&newLanguages, "Languages")
	if len(newLanguages) != len([]string{"ZH", "EN"}) {
		t.Errorf("Query many to many relations")
	}

	new_types.TestDB.Model(&user).Association("Languages").Find(&newLanguages)
	if len(newLanguages) != len([]string{"ZH", "EN"}) {
		t.Errorf("Should be able to find many to many relations")
	}

	if new_types.TestDB.Model(&user).Association("Languages").Count() != len([]string{"ZH", "EN"}) {
		t.Errorf("Count should return correct result")
	}

	// Append
	new_types.TestDB.Model(&user).Association("Languages").Append(&new_types.Language{Name: "DE"})
	if new_types.TestDB.Where("name = ?", "DE").First(&new_types.Language{}).RecordNotFound() {
		t.Errorf("New record should be saved when append")
	}

	languageA := new_types.Language{Name: "AA"}
	new_types.TestDB.Save(&languageA)
	new_types.TestDB.Model(&new_types.User{Id: user.Id}).Association("Languages").Append(&languageA)

	languageC := new_types.Language{Name: "CC"}
	new_types.TestDB.Save(&languageC)
	new_types.TestDB.Model(&user).Association("Languages").Append(&[]new_types.Language{{Name: "BB"}, languageC})

	new_types.TestDB.Model(&new_types.User{Id: user.Id}).Association("Languages").Append(&[]new_types.Language{{Name: "DD"}, {Name: "EE"}})

	totalLanguages := []string{"ZH", "EN", "DE", "AA", "BB", "CC", "DD", "EE"}

	if new_types.TestDB.Model(&user).Association("Languages").Count() != len(totalLanguages) {
		t.Errorf("All appended languages should be saved")
	}

	// Delete
	user.Languages = []new_types.Language{}
	new_types.TestDB.Model(&user).Association("Languages").Find(&user.Languages)

	var language new_types.Language
	new_types.TestDB.Where("name = ?", "EE").First(&language)
	new_types.TestDB.Model(&user).Association("Languages").Delete(language, &language)

	if new_types.TestDB.Model(&user).Association("Languages").Count() != len(totalLanguages)-1 || len(user.Languages) != len(totalLanguages)-1 {
		t.Errorf("Relations should be deleted with Delete")
	}
	if new_types.TestDB.Where("name = ?", "EE").First(&new_types.Language{}).RecordNotFound() {
		t.Errorf("Language EE should not be deleted")
	}

	new_types.TestDB.Where("name IN (?)", []string{"CC", "DD"}).Find(&languages)

	user2 := new_types.User{Name: "Many2Many_User2", Languages: languages}
	new_types.TestDB.Save(&user2)

	new_types.TestDB.Model(&user).Association("Languages").Delete(languages, &languages)
	if new_types.TestDB.Model(&user).Association("Languages").Count() != len(totalLanguages)-3 || len(user.Languages) != len(totalLanguages)-3 {
		t.Errorf("Relations should be deleted with Delete")
	}

	if new_types.TestDB.Model(&user2).Association("Languages").Count() == 0 {
		t.Errorf("Other user's relations should not be deleted")
	}

	// Replace
	var languageB new_types.Language
	new_types.TestDB.Where("name = ?", "BB").First(&languageB)
	new_types.TestDB.Model(&user).Association("Languages").Replace(languageB)
	if len(user.Languages) != 1 || new_types.TestDB.Model(&user).Association("Languages").Count() != 1 {
		t.Errorf("Relations should be replaced")
	}

	new_types.TestDB.Model(&user).Association("Languages").Replace()
	if len(user.Languages) != 0 || new_types.TestDB.Model(&user).Association("Languages").Count() != 0 {
		t.Errorf("Relations should be replaced with empty")
	}

	new_types.TestDB.Model(&user).Association("Languages").Replace(&[]new_types.Language{{Name: "FF"}, {Name: "JJ"}})
	if len(user.Languages) != 2 || new_types.TestDB.Model(&user).Association("Languages").Count() != len([]string{"FF", "JJ"}) {
		t.Errorf("Relations should be replaced")
	}

	// Clear
	new_types.TestDB.Model(&user).Association("Languages").Clear()
	if len(user.Languages) != 0 || new_types.TestDB.Model(&user).Association("Languages").Count() != 0 {
		t.Errorf("Relations should be cleared")
	}

	// Check Association mode with soft delete
	var language6 = new_types.Language{
		Name: "language 6",
	}
	new_types.TestDB.Model(&user).Association("Languages").Append(&language6)

	if count := new_types.TestDB.Model(&user).Association("Languages").Count(); count != 1 {
		t.Errorf("user's languages count should be 1 after Append, but got %v", count)
	}

	new_types.TestDB.Delete(&language6)

	if count := new_types.TestDB.Model(&user).Association("Languages").Count(); count != 0 {
		t.Errorf("user's languages count should be 0 after language been deleted, but got %v", count)
	}

	var languages6 []new_types.Language
	if new_types.TestDB.Model(&user).Association("Languages").Find(&languages6); len(languages6) != 0 {
		t.Errorf("user's languages count should be 0 when find with Find, but got %v", len(languages6))
	}

	if count := new_types.TestDB.Unscoped().Model(&user).Association("Languages").Count(); count != 1 {
		t.Errorf("user's languages count should be 1 when query with Unscoped, but got %v", count)
	}

	var languages61 []new_types.Language
	if new_types.TestDB.Unscoped().Model(&user).Association("Languages").Find(&languages61); len(languages61) != 1 {
		t.Errorf("user's languages count should be 1 when query with Unscoped, but got %v", len(languages61))
	}
}

func Related(t *testing.T) {
	user := new_types.User{
		Name:            "jinzhu",
		BillingAddress:  new_types.Address{Address1: "Billing Address - Address 1"},
		ShippingAddress: new_types.Address{Address1: "Shipping Address - Address 1"},
		Emails:          []new_types.Email{{Email: "jinzhu@example.com"}, {Email: "jinzhu-2@example@example.com"}},
		CreditCard:      new_types.CreditCard{Number: "1234567890"},
		Company:         new_types.Company{Name: "company1"},
	}

	if err := new_types.TestDB.Save(&user).Error; err != nil {
		t.Errorf("No error should happen when saving user")
		t.Errorf("ERROR : %v", err)
	}

	if user.CreditCard.ID == 0 {
		t.Errorf("After user save, credit card should have id")
	}

	if user.BillingAddress.ID == 0 {
		t.Errorf("After user save, billing address should have id")
	}

	if user.Emails[0].Id == 0 {
		t.Errorf("After user save, billing address should have id")
	}

	var emails []new_types.Email
	new_types.TestDB.Model(&user).Related(&emails)
	if len(emails) != 2 {
		t.Errorf("Should have two emails")
	}

	var emails2 []new_types.Email
	new_types.TestDB.Model(&user).Where("email = ?", "jinzhu@example.com").Related(&emails2)
	if len(emails2) != 1 {
		t.Errorf("Should have two emails")
	}

	var emails3 []*new_types.Email
	new_types.TestDB.Model(&user).Related(&emails3)
	if len(emails3) != 2 {
		t.Errorf("Should have two emails")
	}

	var user1 new_types.User
	new_types.TestDB.Model(&user).Related(&user1.Emails)
	if len(user1.Emails) != 2 {
		t.Errorf("Should have only one email match related condition")
	}

	var address1 new_types.Address
	new_types.TestDB.Model(&user).Related(&address1, "BillingAddressId")
	if address1.Address1 != "Billing Address - Address 1" {
		t.Errorf("Should get billing address from user correctly")
	}

	user1 = new_types.User{}
	new_types.TestDB.Model(&address1).Related(&user1, "BillingAddressId")
	if new_types.TestDB.NewRecord(user1) {
		t.Errorf("Should get user from address correctly")
	}

	var user2 new_types.User
	new_types.TestDB.Model(&emails[0]).Related(&user2)
	if user2.Id != user.Id || user2.Name != user.Name {
		t.Errorf("Should get user from email correctly")
	}

	var creditcard new_types.CreditCard
	var user3 new_types.User
	new_types.TestDB.First(&creditcard, "number = ?", "1234567890")
	new_types.TestDB.Model(&creditcard).Related(&user3)
	if user3.Id != user.Id || user3.Name != user.Name {
		t.Errorf("Should get user from credit card correctly")
	}

	if !new_types.TestDB.Model(&new_types.CreditCard{}).Related(&new_types.User{}).RecordNotFound() {
		t.Errorf("RecordNotFound for Related")
	}

	var company new_types.Company
	if new_types.TestDB.Model(&user).Related(&company, "Company").RecordNotFound() || company.Name != "company1" {
		t.Errorf("RecordNotFound for Related")
	}
}

func ForeignKey(t *testing.T) {
	for _, structField := range new_types.TestDB.NewScope(&new_types.User{}).GetModelStruct().StructFields() {
		for _, foreignKey := range []string{"BillingAddressID", "ShippingAddressId", "CompanyID"} {
			if structField.StructName == foreignKey && !structField.IsForeignKey() {
				t.Errorf(fmt.Sprintf("%v should be foreign key", foreignKey))
			}
		}
	}

	for _, structField := range new_types.TestDB.NewScope(&new_types.Email{}).GetModelStruct().StructFields() {
		for _, foreignKey := range []string{"UserId"} {
			if structField.StructName == foreignKey && !structField.IsForeignKey() {
				t.Errorf(fmt.Sprintf("%v should be foreign key", foreignKey))
			}
		}
	}

	for _, structField := range new_types.TestDB.NewScope(&new_types.Post{}).GetModelStruct().StructFields() {
		for _, foreignKey := range []string{"CategoryId", "MainCategoryId"} {
			if structField.StructName == foreignKey && !structField.IsForeignKey() {
				t.Errorf(fmt.Sprintf("%v should be foreign key", foreignKey))
			}
		}
	}

	for _, structField := range new_types.TestDB.NewScope(&new_types.Comment{}).GetModelStruct().StructFields() {
		for _, foreignKey := range []string{"PostId"} {
			if structField.StructName == foreignKey && !structField.IsForeignKey() {
				t.Errorf(fmt.Sprintf("%v should be foreign key", foreignKey))
			}
		}
	}
}

func newTestForeignKey(t *testing.T, source interface{}, sourceFieldName string, target interface{}, targetFieldName string) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect == "" || dialect == "sqlite" {
		// sqlite does not support ADD CONSTRAINT in ALTER TABLE
		return
	}
	targetScope := new_types.TestDB.NewScope(target)
	targetTableName := targetScope.TableName()
	modelScope := new_types.TestDB.NewScope(source)
	modelField, ok := modelScope.FieldByName(sourceFieldName)
	if !ok {
		t.Fatalf(fmt.Sprintf("Failed to get field by name: %v", sourceFieldName))
	}
	targetField, ok := targetScope.FieldByName(targetFieldName)
	if !ok {
		t.Fatalf(fmt.Sprintf("Failed to get field by name: %v", targetFieldName))
	}
	dest := fmt.Sprintf("%v(%v)", targetTableName, targetField.DBName)
	err := new_types.TestDB.Model(source).AddForeignKey(modelField.DBName, dest, "CASCADE", "CASCADE").Error
	if err != nil {
		t.Fatalf(fmt.Sprintf("Failed to create foreign key: %v", err))
	}
}

func LongForeignKey(t *testing.T) {
	newTestForeignKey(t, &new_types.NotSoLongTableName{}, "ReallyLongThingID", &new_types.ReallyLongTableNameToTestMySQLNameLengthLimit{}, "ID")
}

func LongForeignKeyWithShortDest(t *testing.T) {
	newTestForeignKey(t, &new_types.ReallyLongThingThatReferencesShort{}, "ShortID", &new_types.Short{}, "ID")
}

func HasManyChildrenWithOneStruct(t *testing.T) {
	category := new_types.Category{
		Name: "main",
		Categories: []new_types.Category{
			{Name: "sub1"},
			{Name: "sub2"},
		},
	}

	new_types.TestDB.Save(&category)
}
