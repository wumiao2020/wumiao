package models

import (
	"html/template"
	"wumiao/extend"
)

type Category struct {
	Id              int           `json:"id" xorm:"not null pk autoincr comment('Entity ID') INT(10)"`
	ParentId        int           `json:"parent_id" xorm:"default 0 comment('Parent Category ID') index INT(10)"`
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
	AuthorId        int           `json:"author_id" xorm:"comment('Page Author Id') index INT(10)"`
	CreatedAt       extend.Time   `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('Page Creation Time') TIMESTAMP"`
	UpdatedAt       extend.Time   `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('Page Updated Time') TIMESTAMP"`
}

func (c Category) GetAll() []Category {
	panic("implement me")
}

func (c Category) GetList(limit int, start int) []Category {
	panic("implement me")
}

func (c Category) Get(string string) *Category {
	panic("implement me")
}

func (c Category) GetByIdentifier(string string) *Category {
	panic("implement me")
}

func (c Category) GetByUuid(string string) *Category {
	panic("implement me")
}

func (c Category) DeleteByID(id int64) error {
	panic("implement me")
}

func (c Category) Update(data *Category, columns []string) error {
	panic("implement me")
}

func (c Category) Create(data *Category) error {
	panic("implement me")
}