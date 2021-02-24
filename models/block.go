package models

import (
	"html/template"
	"wumiao/extend"
)

type Block struct {
	Id         int64         `json:"id" xorm:"pk autoincr comment('Entity ID') BIGINT(20)"`
	Title      string        `json:"title" xorm:"comment('Page Title') index VARCHAR(255)"`
	Identifier string        `json:"identifier" xorm:"comment('Page String Identifier') unique VARCHAR(255)"`
	Content    template.HTML `json:"content" xorm:"comment('Page Content') MEDIUMTEXT"`
	Status     int           `json:"status" xorm:"not null default 1 comment('Is Page Active') index TINYINT(1)"`
	Position   int           `json:"position" xorm:"not null default 0 comment('Page Sort Order') SMALLINT(6)"`
	Author     string        `json:"author" xorm:"comment('Author') VARCHAR(32)"`
	AuthorId   int64         `json:"author_id" xorm:"comment('Page Author Id') index BIGINT(20)"`
	CreatedAt  extend.Time   `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('Page Creation Time') TIMESTAMP"`
	UpdatedAt  extend.Time   `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('Page Updated Time') TIMESTAMP"`
}
