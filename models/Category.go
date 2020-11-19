package models

import (
	"html/template"
	"time"
)

type Category struct {
	Id              int64         `json:"id" xorm:"pk autoincr comment('Entity ID') BIGINT(20)"`
	ParentId        int64         `json:"parent_id" xorm:"default 0 comment('Parent Category ID') index BIGINT(20)"`
	Title           string        `json:"title" xorm:"comment('Page Title') index VARCHAR(255)"`
	PageLayout      string        `json:"page_layout" xorm:"comment('Page Layout') VARCHAR(255)"`
	Thumb           string        `json:"thumb" xorm:"VARCHAR(255)"`
	MetaTitle       string        `json:"meta_title" xorm:"comment('Page Meta Title') index VARCHAR(255)"`
	MetaKeywords    string        `json:"meta_keywords" xorm:"comment('Page Meta Keywords') TEXT"`
	MetaDescription string        `json:"meta_description" xorm:"comment('Page Meta Description') TEXT"`
	Identifier      string        `json:"identifier" xorm:"comment('Page String Identifier') unique VARCHAR(64)"`
	Uuid            string        `json:"uuid" xorm:"comment('Page Uuid') unique VARCHAR(36)"`
	ContentHeading  string        `json:"content_heading" xorm:"comment('Page Content Heading') index VARCHAR(255)"`
	Content         template.HTML `json:"content" xorm:"comment('Page Content') MEDIUMTEXT"`
	IsActive        int           `json:"is_active" xorm:"not null default 1 comment('Is Page Active') index TINYINT(1)"`
	Position        int           `json:"position" xorm:"not null default 0 comment('Page Sort Order') SMALLINT(6)"`
	Path            string        `json:"path" xorm:"default '1' comment('Tree Path') VARCHAR(64)"`
	Author          string        `json:"author" xorm:"comment('Author') VARCHAR(32)"`
	AuthorId        int64         `json:"author_id" xorm:"comment('Page Author Id') index BIGINT(20)"`
	CreatedAt       time.Time     `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('Page Creation Time') TIMESTAMP"`
	UpdatedAt       time.Time     `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('Page Updated Time') TIMESTAMP"`
}
