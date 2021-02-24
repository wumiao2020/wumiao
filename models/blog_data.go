package models

import (
	"html/template"
	"wumiao/extend"
)

type BlogData struct {
	Id        int64         `json:"id" xorm:"pk autoincr comment('Entity ID') BIGINT(20)"`
	PageId    int64         `json:"page_id" xorm:"default 0 comment('Parent Page ID') index BIGINT(20)"`
	Content   template.HTML `json:"content" xorm:"comment('Page Content') index MEDIUMTEXT"`
	CreatedAt extend.Time   `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created TIMESTAMP"`
	UpdatedAt extend.Time   `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated TIMESTAMP"`
}
