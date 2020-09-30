package models

import (
	"html/template"
	"wumiao/extend"
)

type Page struct {
	Id              int64         `json:"id" xorm:"not null pk autoincr comment('Entity ID') INT(10)"`
	ParentId        int           `json:"parent_id" xorm:"default 0 comment('Parent Category ID') index INT(10)"`
	Type            string        `json:"type" xorm:"not null default 'cms' comment('Page Type') index ENUM('category','cms')"`
	Title           string        `json:"title" xorm:"comment('Page Title') index(CMS_PAGE_TITLE_META_KEYWORDS_META_DESCRIPTION_IDENTIFIER_CONTENT) VARCHAR(255)"`
	PageLayout      string        `json:"page_layout" xorm:"comment('Page Layout') VARCHAR(255)"`
	MetaTitle       string        `json:"meta_title" xorm:"comment('Page Meta Title') VARCHAR(255)"`
	MetaKeywords    string        `json:"meta_keywords" xorm:"comment('Page Meta Keywords') index(CMS_PAGE_TITLE_META_KEYWORDS_META_DESCRIPTION_IDENTIFIER_CONTENT) TEXT"`
	MetaDescription string        `json:"meta_description" xorm:"comment('Page Meta Description') index(CMS_PAGE_TITLE_META_KEYWORDS_META_DESCRIPTION_IDENTIFIER_CONTENT) TEXT"`
	Identifier      string        `json:"identifier" xorm:"comment('Page String Identifier') index index(CMS_PAGE_TITLE_META_KEYWORDS_META_DESCRIPTION_IDENTIFIER_CONTENT) VARCHAR(64)"`
	Uuid            string        `json:"uuid" xorm:"comment('Page String Uuid') index index(CMS_PAGE_TITLE_META_KEYWORDS_META_DESCRIPTION_IDENTIFIER_CONTENT) VARCHAR(32)"`
	ContentHeading  string        `json:"content_heading" xorm:"comment('Page Content Heading') VARCHAR(255)"`
	Content         template.HTML `json:"content" xorm:"comment('Page Content') index(CMS_PAGE_TITLE_META_KEYWORDS_META_DESCRIPTION_IDENTIFIER_CONTENT) MEDIUMTEXT"`
	IsActive        int           `json:"is_active" xorm:"not null default 1 comment('Is Page Active') TINYINT(1)"`
	SortOrder       int           `json:"sort_order" xorm:"not null default 0 comment('Page Sort Order') SMALLINT(6)"`
	Path            string        `json:"path" xorm:"default '1' comment('Tree Path') index VARCHAR(64)"`
	CreatedAt       extend.Time   `json:"created_at" xorm:"created"`
	UpdatedAt       extend.Time   `json:"updated_at" xorm:"created"`
}
