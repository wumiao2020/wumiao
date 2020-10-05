package models

import (
	"wumiao/extend"
)

type Menu struct {
	Id        int64       `json:"id" xorm:"not null pk autoincr comment('Entity ID') INT(10)"`
	ParentId  int         `json:"parent_id" xorm:"default 0 comment('Parent Category ID') index INT(10)"`
	Uri       string      `json:"uri" xorm:"comment('Page Title') index VARCHAR(255)"`
	Title     string      `json:"title" xorm:"comment('Page Title') index VARCHAR(255)"`
	Thumb     string      `json:"thumb" xorm:"comment('Page Meta Title') index VARCHAR(255)"`
	IsActive  int         `json:"is_active" xorm:"not null default 1 comment('Is Page Active') index TINYINT(1)"`
	Position  int         `json:"position" xorm:"not null default 0 comment('Page Sort Order') SMALLINT(6)"`
	Author    string      `json:"author" xorm:"comment('Author') VARCHAR(32)"`
	AuthorId  int         `json:"author_id" xorm:"comment('Page Author Id') index INT(10)"`
	CreatedAt extend.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('Page Creation Time') TIMESTAMP"`
	UpdatedAt extend.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('Page Updated Time') TIMESTAMP"`
}
