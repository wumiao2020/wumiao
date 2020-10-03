package models

import (
	"html/template"
	"time"
)

type NewsData struct {
	Id        int           `json:"id" xorm:"not null pk autoincr comment('Entity ID') INT(10)"`
	PageId    int           `json:"page_id" xorm:"default 0 comment('Parent Page ID') index INT(10)"`
	Content   template.HTML `json:"content" xorm:"comment('Page Content') index MEDIUMTEXT"`
	CreatedAt time.Time     `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created TIMESTAMP"`
	UpdatedAt time.Time     `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated TIMESTAMP"`
}
