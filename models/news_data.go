package models

import (
	"time"
)

type NewsData struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('Entity ID') BIGINT(20)"`
	PageId    int64     `json:"page_id" xorm:"default 0 comment('Parent Page ID') index BIGINT(20)"`
	Content   string    `json:"content" xorm:"comment('Page Content') index MEDIUMTEXT"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated TIMESTAMP"`
}
