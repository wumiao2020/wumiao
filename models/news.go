package models

import (
	"wumiao/extend"
)

type News struct {
	Id              int64       `json:"id" xorm:"not null pk autoincr comment('Entity ID') INT(10)"`
	Title           string      `json:"title" xorm:"comment('Page Title') index VARCHAR(255)"`
	PageLayout      string      `json:"page_layout" xorm:"comment('Page Layout') VARCHAR(255)"`
	MetaTitle       string      `json:"meta_title" xorm:"comment('Page Meta Title') VARCHAR(255)"`
	MetaKeywords    string      `json:"meta_keywords" xorm:"comment('Page Meta Keywords') TEXT"`
	MetaDescription string      `json:"meta_description" xorm:"comment('Page Meta Description') TEXT"`
	Identifier      string      `json:"identifier" xorm:"comment('Page String Identifier') unique VARCHAR(64)"`
	Uuid            string      `json:"uuid" xorm:"comment('Page Uuid') unique VARCHAR(36)"`
	ContentHeading  string      `json:"content_heading" xorm:"comment('Page Content Heading') index VARCHAR(255)"`
	IsActive        int         `json:"is_active" xorm:"not null default 1 comment('Is Page Active') index TINYINT(1)"`
	Position        int         `json:"position" xorm:"not null default 0 comment('Page Sort Order') SMALLINT(6)"`
	Author          string      `json:"author" xorm:"comment('Author') VARCHAR(32)"`
	AuthorId        int         `json:"author_id" xorm:"comment('Page Author Id') INT(10)"`
	CreatedAt       extend.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('Page Creation Time') TIMESTAMP"`
	UpdatedAt       extend.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('Page Updated Time') TIMESTAMP"`
}
