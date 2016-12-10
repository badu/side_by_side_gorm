package tests

import (
	"SideBySideGorm/new_types"
	"os"
	"testing"
)

func ManyToManyWithMultiPrimaryKeys(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect != "" && dialect != "sqlite" {
		new_types.TestDB.DropTable(&new_types.Blog{}, &new_types.Tag{})
		new_types.TestDB.DropTable("blog_tags")
		new_types.TestDB.CreateTable(&new_types.Blog{}, &new_types.Tag{})
		blog := new_types.Blog{
			Locale:  "ZH",
			Subject: "subject",
			Body:    "body",
			Tags: []new_types.Tag{
				{Locale: "ZH", Value: "tag1"},
				{Locale: "ZH", Value: "tag2"},
			},
		}

		new_types.TestDB.Save(&blog)
		if !newCompareTags(blog.Tags, []string{"tag1", "tag2"}) {
			t.Errorf("Blog should has two tags")
		}

		// Append
		var tag3 = &new_types.Tag{Locale: "ZH", Value: "tag3"}
		new_types.TestDB.Model(&blog).Association("Tags").Append([]*new_types.Tag{tag3})
		if !newCompareTags(blog.Tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Blog should has three tags after Append")
		}

		if new_types.TestDB.Model(&blog).Association("Tags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		var tags []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags, "Tags")
		if !newCompareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		var blog1 new_types.Blog
		new_types.TestDB.Preload("Tags").Find(&blog1)
		if !newCompareTags(blog1.Tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Preload many2many relations")
		}

		// Replace
		var tag5 = &new_types.Tag{Locale: "ZH", Value: "tag5"}
		var tag6 = &new_types.Tag{Locale: "ZH", Value: "tag6"}
		new_types.TestDB.Model(&blog).Association("Tags").Replace(tag5, tag6)
		var tags2 []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags2, "Tags")
		if !newCompareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		if new_types.TestDB.Model(&blog).Association("Tags").Count() != 2 {
			t.Errorf("Blog should has three tags after Replace")
		}

		// Delete
		new_types.TestDB.Model(&blog).Association("Tags").Delete(tag5)
		var tags3 []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags3, "Tags")
		if !newCompareTags(tags3, []string{"tag6"}) {
			t.Errorf("Should find 1 tags after Delete")
		}

		if new_types.TestDB.Model(&blog).Association("Tags").Count() != 1 {
			t.Errorf("Blog should has three tags after Delete")
		}

		new_types.TestDB.Model(&blog).Association("Tags").Delete(tag3)
		var tags4 []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags4, "Tags")
		if !newCompareTags(tags4, []string{"tag6"}) {
			t.Errorf("Tag should not be deleted when Delete with a unrelated tag")
		}

		// Clear
		new_types.TestDB.Model(&blog).Association("Tags").Clear()
		if new_types.TestDB.Model(&blog).Association("Tags").Count() != 0 {
			t.Errorf("All tags should be cleared")
		}
	}
}

func ManyToManyWithCustomizedForeignKeys(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect != "" && dialect != "sqlite" {
		new_types.TestDB.DropTable(&new_types.Blog{}, &new_types.Tag{})
		new_types.TestDB.DropTable("shared_blog_tags")
		new_types.TestDB.CreateTable(&new_types.Blog{}, &new_types.Tag{})
		blog := new_types.Blog{
			Locale:  "ZH",
			Subject: "subject",
			Body:    "body",
			SharedTags: []new_types.Tag{
				{Locale: "ZH", Value: "tag1"},
				{Locale: "ZH", Value: "tag2"},
			},
		}
		new_types.TestDB.Save(&blog)

		blog2 := new_types.Blog{
			ID:     blog.ID,
			Locale: "EN",
		}
		new_types.TestDB.Create(&blog2)

		if !newCompareTags(blog.SharedTags, []string{"tag1", "tag2"}) {
			t.Errorf("Blog should has two tags")
		}

		// Append
		var tag3 = &new_types.Tag{Locale: "ZH", Value: "tag3"}
		new_types.TestDB.Model(&blog).Association("SharedTags").Append([]*new_types.Tag{tag3})
		if !newCompareTags(blog.SharedTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Blog should has three tags after Append")
		}

		if new_types.TestDB.Model(&blog).Association("SharedTags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		if new_types.TestDB.Model(&blog2).Association("SharedTags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		var tags []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags, "SharedTags")
		if !newCompareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		new_types.TestDB.Model(&blog2).Related(&tags, "SharedTags")
		if !newCompareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		var blog1 new_types.Blog
		new_types.TestDB.Preload("SharedTags").Find(&blog1)
		if !newCompareTags(blog1.SharedTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Preload many2many relations")
		}

		var tag4 = &new_types.Tag{Locale: "ZH", Value: "tag4"}
		new_types.TestDB.Model(&blog2).Association("SharedTags").Append(tag4)

		new_types.TestDB.Model(&blog).Related(&tags, "SharedTags")
		if !newCompareTags(tags, []string{"tag1", "tag2", "tag3", "tag4"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		new_types.TestDB.Model(&blog2).Related(&tags, "SharedTags")
		if !newCompareTags(tags, []string{"tag1", "tag2", "tag3", "tag4"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		// Replace
		var tag5 = &new_types.Tag{Locale: "ZH", Value: "tag5"}
		var tag6 = &new_types.Tag{Locale: "ZH", Value: "tag6"}
		new_types.TestDB.Model(&blog2).Association("SharedTags").Replace(tag5, tag6)
		var tags2 []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags2, "SharedTags")
		if !newCompareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		new_types.TestDB.Model(&blog2).Related(&tags2, "SharedTags")
		if !newCompareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		if new_types.TestDB.Model(&blog).Association("SharedTags").Count() != 2 {
			t.Errorf("Blog should has three tags after Replace")
		}

		// Delete
		new_types.TestDB.Model(&blog).Association("SharedTags").Delete(tag5)
		var tags3 []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags3, "SharedTags")
		if !newCompareTags(tags3, []string{"tag6"}) {
			t.Errorf("Should find 1 tags after Delete")
		}

		if new_types.TestDB.Model(&blog).Association("SharedTags").Count() != 1 {
			t.Errorf("Blog should has three tags after Delete")
		}

		new_types.TestDB.Model(&blog2).Association("SharedTags").Delete(tag3)
		var tags4 []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags4, "SharedTags")
		if !newCompareTags(tags4, []string{"tag6"}) {
			t.Errorf("Tag should not be deleted when Delete with a unrelated tag")
		}

		// Clear
		new_types.TestDB.Model(&blog2).Association("SharedTags").Clear()
		if new_types.TestDB.Model(&blog).Association("SharedTags").Count() != 0 {
			t.Errorf("All tags should be cleared")
		}
	}
}

func ManyToManyWithCustomizedForeignKeys2(t *testing.T) {
	if dialect := os.Getenv("GORM_DIALECT"); dialect != "" && dialect != "sqlite" {
		new_types.TestDB.DropTable(&new_types.Blog{}, &new_types.Tag{})
		new_types.TestDB.DropTable("locale_blog_tags")
		new_types.TestDB.CreateTable(&new_types.Blog{}, &new_types.Tag{})
		blog := new_types.Blog{
			Locale:  "ZH",
			Subject: "subject",
			Body:    "body",
			LocaleTags: []new_types.Tag{
				{Locale: "ZH", Value: "tag1"},
				{Locale: "ZH", Value: "tag2"},
			},
		}
		new_types.TestDB.Save(&blog)

		blog2 := new_types.Blog{
			ID:     blog.ID,
			Locale: "EN",
		}
		new_types.TestDB.Create(&blog2)

		// Append
		var tag3 = &new_types.Tag{Locale: "ZH", Value: "tag3"}
		new_types.TestDB.Model(&blog).Association("LocaleTags").Append([]*new_types.Tag{tag3})
		if !newCompareTags(blog.LocaleTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Blog should has three tags after Append")
		}

		if new_types.TestDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("Blog should has three tags after Append")
		}

		if new_types.TestDB.Model(&blog2).Association("LocaleTags").Count() != 0 {
			t.Errorf("EN Blog should has 0 tags after ZH Blog Append")
		}

		var tags []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags, "LocaleTags")
		if !newCompareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related")
		}

		new_types.TestDB.Model(&blog2).Related(&tags, "LocaleTags")
		if len(tags) != 0 {
			t.Errorf("Should find 0 tags with Related for EN Blog")
		}

		var blog1 new_types.Blog
		new_types.TestDB.Preload("LocaleTags").Find(&blog1, "locale = ? AND id = ?", "ZH", blog.ID)
		if !newCompareTags(blog1.LocaleTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Preload many2many relations")
		}

		var tag4 = &new_types.Tag{Locale: "ZH", Value: "tag4"}
		new_types.TestDB.Model(&blog2).Association("LocaleTags").Append(tag4)

		new_types.TestDB.Model(&blog).Related(&tags, "LocaleTags")
		if !newCompareTags(tags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Should find 3 tags with Related for EN Blog")
		}

		new_types.TestDB.Model(&blog2).Related(&tags, "LocaleTags")
		if !newCompareTags(tags, []string{"tag4"}) {
			t.Errorf("Should find 1 tags with Related for EN Blog")
		}

		// Replace
		var tag5 = &new_types.Tag{Locale: "ZH", Value: "tag5"}
		var tag6 = &new_types.Tag{Locale: "ZH", Value: "tag6"}
		new_types.TestDB.Model(&blog2).Association("LocaleTags").Replace(tag5, tag6)

		var tags2 []new_types.Tag
		new_types.TestDB.Model(&blog).Related(&tags2, "LocaleTags")
		if !newCompareTags(tags2, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("CN Blog's tags should not be changed after EN Blog Replace")
		}

		var blog11 new_types.Blog
		new_types.TestDB.Preload("LocaleTags").First(&blog11, "id = ? AND locale = ?", blog.ID, blog.Locale)
		if !newCompareTags(blog11.LocaleTags, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("CN Blog's tags should not be changed after EN Blog Replace")
		}

		new_types.TestDB.Model(&blog2).Related(&tags2, "LocaleTags")
		if !newCompareTags(tags2, []string{"tag5", "tag6"}) {
			t.Errorf("Should find 2 tags after Replace")
		}

		var blog21 new_types.Blog
		new_types.TestDB.Preload("LocaleTags").First(&blog21, "id = ? AND locale = ?", blog2.ID, blog2.Locale)
		if !newCompareTags(blog21.LocaleTags, []string{"tag5", "tag6"}) {
			t.Errorf("EN Blog's tags should be changed after Replace")
		}

		if new_types.TestDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog should has three tags after Replace")
		}

		if new_types.TestDB.Model(&blog2).Association("LocaleTags").Count() != 2 {
			t.Errorf("EN Blog should has two tags after Replace")
		}

		// Delete
		new_types.TestDB.Model(&blog).Association("LocaleTags").Delete(tag5)

		if new_types.TestDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog should has three tags after Delete with EN's tag")
		}

		if new_types.TestDB.Model(&blog2).Association("LocaleTags").Count() != 2 {
			t.Errorf("EN Blog should has two tags after ZH Blog Delete with EN's tag")
		}

		new_types.TestDB.Model(&blog2).Association("LocaleTags").Delete(tag5)

		if new_types.TestDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog should has three tags after EN Blog Delete with EN's tag")
		}

		if new_types.TestDB.Model(&blog2).Association("LocaleTags").Count() != 1 {
			t.Errorf("EN Blog should has 1 tags after EN Blog Delete with EN's tag")
		}

		// Clear
		new_types.TestDB.Model(&blog2).Association("LocaleTags").Clear()
		if new_types.TestDB.Model(&blog).Association("LocaleTags").Count() != 3 {
			t.Errorf("ZH Blog's tags should not be cleared when clear EN Blog's tags")
		}

		if new_types.TestDB.Model(&blog2).Association("LocaleTags").Count() != 0 {
			t.Errorf("EN Blog's tags should be cleared when clear EN Blog's tags")
		}

		new_types.TestDB.Model(&blog).Association("LocaleTags").Clear()
		if new_types.TestDB.Model(&blog).Association("LocaleTags").Count() != 0 {
			t.Errorf("ZH Blog's tags should be cleared when clear ZH Blog's tags")
		}

		if new_types.TestDB.Model(&blog2).Association("LocaleTags").Count() != 0 {
			t.Errorf("EN Blog's tags should be cleared")
		}
	}
}
