package tests

import (
	"os"
	"reflect"
	"sort"
	"testing"
)

type Blog struct {
	ID         uint   `gorm:"primary_key"`
	Locale     string `gorm:"primary_key"`
	Subject    string
	Body       string
	Tags       []Tag `gorm:"many2many:blog_tags;"`
	SharedTags []Tag `gorm:"many2many:shared_blog_tags;ForeignKey:id;AssociationForeignKey:id"`
	LocaleTags []Tag `gorm:"many2many:locale_blog_tags;ForeignKey:id,locale;AssociationForeignKey:id"`
}

type Tag struct {
	ID     uint   `gorm:"primary_key"`
	Locale string `gorm:"primary_key"`
	Value  string
	Blogs  []*Blog `gorm:"many2many:blogs_tags"`
}

func compareTags(tags []Tag, contents []string) bool {
	var tagContents []string
	for _, tag := range tags {
		tagContents = append(tagContents, tag.Value)
	}
	sort.Strings(tagContents)
	sort.Strings(contents)
	return reflect.DeepEqual(tagContents, contents)
}

func OldManyToManyWithMultiPrimaryKeys(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect != "" && dialect != "sqlite" {
		OLDDB.DropTable(&Blog{}, &Tag{})
		OLDDB.DropTable("blog_tags")
		OLDDB.CreateTable(&Blog{}, &Tag{})
		blog := Blog{
			Locale:  "ZH",
			Subject: "subject",
			Body:    "body",
			Tags: []Tag{
				{Locale: "ZH", Value: "tag1"},
				{Locale: "ZH", Value: "tag2"},
			},
		}

		OLDDB.Save(&blog)
		if !compareTags(blog.Tags, []string{"tag1", "tag2"}) {
			t.Errorf("Blog should has two tags")
		}

		// Append
		var tag3 = &Tag{Locale: "ZH", Value: "tag3"}
		OLDDB.Model(&blog).Association("Tags").Append([]*Tag{tag3})
		if !compareTags(blog.Tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Blog should has three tags after Append")
		}

		if OLDDB.Model(&blog).Association("Tags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		var tags []Tag
		OLDDB.Model(&blog).Related(&tags, "Tags")
		if !compareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		var blog1 Blog
		OLDDB.Preload("Tags").Find(&blog1)
		if !compareTags(blog1.Tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Preload many2many relations")
		}

		// Replace
		var tag5 = &Tag{Locale: "ZH", Value: "tag5"}
		var tag6 = &Tag{Locale: "ZH", Value: "tag6"}
		OLDDB.Model(&blog).Association("Tags").Replace(tag5, tag6)
		var tags2 []Tag
		OLDDB.Model(&blog).Related(&tags2, "Tags")
		if !compareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		if OLDDB.Model(&blog).Association("Tags").Count() != 2 {
			t.Errorf("Blog should has three tags after Replace")
		}

		// Delete
		OLDDB.Model(&blog).Association("Tags").Delete(tag5)
		var tags3 []Tag
		OLDDB.Model(&blog).Related(&tags3, "Tags")
		if !compareTags(tags3, []string{"tag6"}) {
			t.Errorf("Should find 1 tags after Delete")
		}

		if OLDDB.Model(&blog).Association("Tags").Count() != 1 {
			t.Errorf("Blog should has three tags after Delete")
		}

		OLDDB.Model(&blog).Association("Tags").Delete(tag3)
		var tags4 []Tag
		OLDDB.Model(&blog).Related(&tags4, "Tags")
		if !compareTags(tags4, []string{"tag6"}) {
			t.Errorf("Tag should not be deleted when Delete with a unrelated tag")
		}

		// Clear
		OLDDB.Model(&blog).Association("Tags").Clear()
		if OLDDB.Model(&blog).Association("Tags").Count() != 0 {
			t.Errorf("All tags should be cleared")
		}
	}
}

func OldManyToManyWithCustomizedForeignKeys(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect != "" && dialect != "sqlite" {
		OLDDB.DropTable(&Blog{}, &Tag{})
		OLDDB.DropTable("shared_blog_tags")
		OLDDB.CreateTable(&Blog{}, &Tag{})
		blog := Blog{
			Locale:  "ZH",
			Subject: "subject",
			Body:    "body",
			SharedTags: []Tag{
				{Locale: "ZH", Value: "tag1"},
				{Locale: "ZH", Value: "tag2"},
			},
		}
		OLDDB.Save(&blog)

		blog2 := Blog{
			ID:     blog.ID,
			Locale: "EN",
		}
		OLDDB.Create(&blog2)

		if !compareTags(blog.SharedTags, []string{"tag1", "tag2"}) {
			t.Errorf("Blog should has two tags")
		}

		// Append
		var tag3 = &Tag{Locale: "ZH", Value: "tag3"}
		OLDDB.Model(&blog).Association("SharedTags").Append([]*Tag{tag3})
		if !compareTags(blog.SharedTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Blog should has three tags after Append")
		}

		if OLDDB.Model(&blog).Association("SharedTags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		if OLDDB.Model(&blog2).Association("SharedTags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		var tags []Tag
		OLDDB.Model(&blog).Related(&tags, "SharedTags")
		if !compareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		OLDDB.Model(&blog2).Related(&tags, "SharedTags")
		if !compareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		var blog1 Blog
		OLDDB.Preload("SharedTags").Find(&blog1)
		if !compareTags(blog1.SharedTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Preload many2many relations")
		}

		var tag4 = &Tag{Locale: "ZH", Value: "tag4"}
		OLDDB.Model(&blog2).Association("SharedTags").Append(tag4)

		OLDDB.Model(&blog).Related(&tags, "SharedTags")
		if !compareTags(tags, []string{"tag1", "tag2", "tag3", "tag4"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		OLDDB.Model(&blog2).Related(&tags, "SharedTags")
		if !compareTags(tags, []string{"tag1", "tag2", "tag3", "tag4"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		// Replace
		var tag5 = &Tag{Locale: "ZH", Value: "tag5"}
		var tag6 = &Tag{Locale: "ZH", Value: "tag6"}
		OLDDB.Model(&blog2).Association("SharedTags").Replace(tag5, tag6)
		var tags2 []Tag
		OLDDB.Model(&blog).Related(&tags2, "SharedTags")
		if !compareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		OLDDB.Model(&blog2).Related(&tags2, "SharedTags")
		if !compareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		if OLDDB.Model(&blog).Association("SharedTags").Count() != 2 {
			t.Errorf("Blog should has three tags after Replace")
		}

		// Delete
		OLDDB.Model(&blog).Association("SharedTags").Delete(tag5)
		var tags3 []Tag
		OLDDB.Model(&blog).Related(&tags3, "SharedTags")
		if !compareTags(tags3, []string{"tag6"}) {
			t.Errorf("Should find 1 tags after Delete")
		}

		if OLDDB.Model(&blog).Association("SharedTags").Count() != 1 {
			t.Errorf("Blog should has three tags after Delete")
		}

		OLDDB.Model(&blog2).Association("SharedTags").Delete(tag3)
		var tags4 []Tag
		OLDDB.Model(&blog).Related(&tags4, "SharedTags")
		if !compareTags(tags4, []string{"tag6"}) {
			t.Errorf("Tag should not be deleted when Delete with a unrelated tag")
		}

		// Clear
		OLDDB.Model(&blog2).Association("SharedTags").Clear()
		if OLDDB.Model(&blog).Association("SharedTags").Count() != 0 {
			t.Errorf("All tags should be cleared")
		}
	}
}

func OldManyToManyWithCustomizedForeignKeys2(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect != "" && dialect != "sqlite" {
		OLDDB.DropTable(&Blog{}, &Tag{})
		OLDDB.DropTable("locale_blog_tags")
		OLDDB.CreateTable(&Blog{}, &Tag{})
		blog := Blog{
			Locale:  "ZH",
			Subject: "subject",
			Body:    "body",
			LocaleTags: []Tag{
				{Locale: "ZH", Value: "tag1"},
				{Locale: "ZH", Value: "tag2"},
			},
		}
		OLDDB.Save(&blog)

		blog2 := Blog{
			ID:     blog.ID,
			Locale: "EN",
		}
		OLDDB.Create(&blog2)

		// Append
		var tag3 = &Tag{Locale: "ZH", Value: "tag3"}
		OLDDB.Model(&blog).Association("LocaleTags").Append([]*Tag{tag3})
		if !compareTags(blog.LocaleTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Blog should has three tags after Append")
		}

		if OLDDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		if OLDDB.Model(&blog2).Association("LocaleTags").Count() != 0 {
			t.Errorf("EN Blog should has 0 tags after ZH Blog Append")
		}

		var tags []Tag
		OLDDB.Model(&blog).Related(&tags, "LocaleTags")
		if !compareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		OLDDB.Model(&blog2).Related(&tags, "LocaleTags")
		if len(tags) != 0 {
			t.Errorf("Should find 0 tags with Related for EN Blog")
		}

		var blog1 Blog
		OLDDB.Preload("LocaleTags").Find(&blog1, "locale = ? AND id = ?", "ZH", blog.ID)
		if !compareTags(blog1.LocaleTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Preload many2many relations")
		}

		var tag4 = &Tag{Locale: "ZH", Value: "tag4"}
		OLDDB.Model(&blog2).Association("LocaleTags").Append(tag4)

		OLDDB.Model(&blog).Related(&tags, "LocaleTags")
		if !compareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related for EN Blog")
		}

		OLDDB.Model(&blog2).Related(&tags, "LocaleTags")
		if !compareTags(tags, []string{"tag4"}) {
			t.Errorf("Should find 1 tags with Related for EN Blog")
		}

		// Replace
		var tag5 = &Tag{Locale: "ZH", Value: "tag5"}
		var tag6 = &Tag{Locale: "ZH", Value: "tag6"}
		OLDDB.Model(&blog2).Association("LocaleTags").Replace(tag5, tag6)

		var tags2 []Tag
		OLDDB.Model(&blog).Related(&tags2, "LocaleTags")
		if !compareTags(tags2, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("CN Blog's tags should not be changed after EN Blog Replace")
		}

		var blog11 Blog
		OLDDB.Preload("LocaleTags").First(&blog11, "id = ? AND locale = ?", blog.ID, blog.Locale)
		if !compareTags(blog11.LocaleTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("CN Blog's tags should not be changed after EN Blog Replace")
		}

		OLDDB.Model(&blog2).Related(&tags2, "LocaleTags")
		if !compareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		var blog21 Blog
		OLDDB.Preload("LocaleTags").First(&blog21, "id = ? AND locale = ?", blog2.ID, blog2.Locale)
		if !compareTags(blog21.LocaleTags, []string{"tag5", "tag6"}) {
			t.Errorf("EN Blog's tags should be changed after Replace")
		}

		if OLDDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog should has three tags after Replace")
		}

		if OLDDB.Model(&blog2).Association("LocaleTags").Count() != 2 {
			t.Errorf("EN Blog should has two tags after Replace")
		}

		// Delete
		OLDDB.Model(&blog).Association("LocaleTags").Delete(tag5)

		if OLDDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog should has three tags after Delete with EN's tag")
		}

		if OLDDB.Model(&blog2).Association("LocaleTags").Count() != 2 {
			t.Errorf("EN Blog should has two tags after ZH Blog Delete with EN's tag")
		}

		OLDDB.Model(&blog2).Association("LocaleTags").Delete(tag5)

		if OLDDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog should has three tags after EN Blog Delete with EN's tag")
		}

		if OLDDB.Model(&blog2).Association("LocaleTags").Count() != 1 {
			t.Errorf("EN Blog should has 1 tags after EN Blog Delete with EN's tag")
		}

		// Clear
		OLDDB.Model(&blog2).Association("LocaleTags").Clear()
		if OLDDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog's tags should not be cleared when clear EN Blog's tags")
		}

		if OLDDB.Model(&blog2).Association("LocaleTags").Count() != 0 {
			t.Errorf("EN Blog's tags should be cleared when clear EN Blog's tags")
		}

		OLDDB.Model(&blog).Association("LocaleTags").Clear()
		if OLDDB.Model(&blog).Association("LocaleTags").Count() != 0 {
			t.Errorf("ZH Blog's tags should be cleared when clear ZH Blog's tags")
		}

		if OLDDB.Model(&blog2).Association("LocaleTags").Count() != 0 {
			t.Errorf("EN Blog's tags should be cleared")
		}
	}
}
