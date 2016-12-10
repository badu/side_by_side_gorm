package tests

import (
	"SideBySideGorm/new_types"
	"testing"
)

func PrefixColumnNameForEmbeddedStruct(t *testing.T) {
	new_types.TestDB.NewScope(&new_types.EngadgetPost{})
	dialect := new_types.TestDB.Dialect()
	if !dialect.HasColumn(new_types.TestDB.NewScope(&new_types.EngadgetPost{}).TableName(), "author_name") || !dialect.HasColumn(new_types.TestDB.NewScope(&new_types.EngadgetPost{}).TableName(), "author_email") {
		t.Errorf("should has prefix for embedded columns")
	}

	if !dialect.HasColumn(new_types.TestDB.NewScope(&new_types.HNPost{}).TableName(), "user_name") || !dialect.HasColumn(new_types.TestDB.NewScope(&new_types.HNPost{}).TableName(), "user_email") {
		t.Errorf("should has prefix for embedded columns")
	}
}

func SaveAndQueryEmbeddedStruct(t *testing.T) {
	new_types.TestDB.Save(&new_types.HNPost{BasePost: new_types.BasePost{Title: "news"}})
	new_types.TestDB.Save(&new_types.HNPost{BasePost: new_types.BasePost{Title: "hn_news"}})
	var news new_types.HNPost
	if err := new_types.TestDB.First(&news, "title = ?", "hn_news").Error; err != nil {
		t.Errorf("no error should happen when query with embedded struct, but got %v", err)
	} else if news.Title != "hn_news" {
		t.Errorf("embedded struct's value should be scanned correctly")
	}

	new_types.TestDB.Save(&new_types.EngadgetPost{BasePost: new_types.BasePost{Title: "engadget_news"}})
	var egNews new_types.EngadgetPost
	if err := new_types.TestDB.First(&egNews, "title = ?", "engadget_news").Error; err != nil {
		t.Errorf("no error should happen when query with embedded struct, but got %v", err)
	} else if egNews.BasePost.Title != "engadget_news" {
		t.Errorf("embedded struct's value should be scanned correctly")
	}

	if new_types.TestDB.NewScope(&new_types.HNPost{}).PK() == nil {
		t.Errorf("primary key with embedded struct should works")
	}

	for _, field := range new_types.TestDB.NewScope(&new_types.HNPost{}).Fields() {
		if field.StructName == "BasePost" {
			t.Errorf("scope Fields should not contain embedded struct")
		}
	}
}
