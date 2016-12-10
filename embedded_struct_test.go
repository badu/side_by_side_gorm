package tests

import "testing"

type BasePost struct {
	Id    int64
	Title string
	URL   string
}

type Author struct {
	Name  string
	Email string
}

type HNPost struct {
	BasePost
	Author  `gorm:"embedded_prefix:user_"` // Embedded struct
	Upvotes int32
}

type EngadgetPost struct {
	BasePost BasePost `gorm:"embedded"`
	Author   Author   `gorm:"embedded;embedded_prefix:author_"` // Embedded struct
	ImageUrl string
}

func OldPrefixColumnNameForEmbeddedStruct(t *testing.T) {
	dialect := OLDDB.NewScope(&EngadgetPost{}).Dialect()
	if !dialect.HasColumn(OLDDB.NewScope(&EngadgetPost{}).TableName(), "author_name") || !dialect.HasColumn(OLDDB.NewScope(&EngadgetPost{}).TableName(), "author_email") {
		t.Errorf("should has prefix for embedded columns")
	}

	if !dialect.HasColumn(OLDDB.NewScope(&HNPost{}).TableName(), "user_name") || !dialect.HasColumn(OLDDB.NewScope(&HNPost{}).TableName(), "user_email") {
		t.Errorf("should has prefix for embedded columns")
	}
}

func OldSaveAndQueryEmbeddedStruct(t *testing.T) {
	OLDDB.Save(&HNPost{BasePost: BasePost{Title: "news"}})
	OLDDB.Save(&HNPost{BasePost: BasePost{Title: "hn_news"}})
	var news HNPost
	if err := OLDDB.First(&news, "title = ?", "hn_news").Error; err != nil {
		t.Errorf("no error should happen when query with embedded struct, but got %v", err)
	} else if news.Title != "hn_news" {
		t.Errorf("embedded struct's value should be scanned correctly")
	}

	OLDDB.Save(&EngadgetPost{BasePost: BasePost{Title: "engadget_news"}})
	var egNews EngadgetPost
	if err := OLDDB.First(&egNews, "title = ?", "engadget_news").Error; err != nil {
		t.Errorf("no error should happen when query with embedded struct, but got %v", err)
	} else if egNews.BasePost.Title != "engadget_news" {
		t.Errorf("embedded struct's value should be scanned correctly")
	}

	if OLDDB.NewScope(&HNPost{}).PrimaryField() == nil {
		t.Errorf("primary key with embedded struct should works")
	}

	for _, field := range OLDDB.NewScope(&HNPost{}).Fields() {
		if field.Name == "BasePost" {
			t.Errorf("scope Fields should not contain embedded struct")
		}
	}
}
