package models

import (
	"wumiao/extend"
)

type Menu struct {
	Id        int64       `json:"id" xorm:"pk autoincr comment('Entity ID') BIGINT(20)"`
	ParentId  int64       `json:"parent_id" xorm:"default 0 comment('Menu Parent Id') BIGINT(20)"`
	Title     string      `json:"title" xorm:"comment('Menu Title') index VARCHAR(255)"`
	Thumb     string      `json:"thumb" xorm:"comment('Menu Content Heading') VARCHAR(255)"`
	Uri       string      `json:"uri" xorm:"comment('Menu String Identifier') VARCHAR(64)"`
	IsActive  int         `json:"is_active" xorm:"not null default 1 comment('Is Menu Active') index TINYINT(1)"`
	Position  int         `json:"position" xorm:"not null default 0 comment('Menu Sort Order') SMALLINT(6)"`
	Author    string      `json:"author" xorm:"comment('Menu Author') VARCHAR(32)"`
	AuthorId  int64       `json:"author_id" xorm:"comment('Menu Author Id') index BIGINT(20)"`
	CreatedAt extend.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('Menu Creation Time') TIMESTAMP"`
	UpdatedAt extend.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('Menu Updated Time') TIMESTAMP"`
}
