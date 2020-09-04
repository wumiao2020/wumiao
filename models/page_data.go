package models

import (
	"time"
)

type PageData struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('主建') BIGINT(20)"`
	PageId    int       `json:"page_id" xorm:"not null comment('父ID') index TINYINT(4)"`
	Content   string    `json:"content" xorm:"comment('描述') TEXT"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
}
