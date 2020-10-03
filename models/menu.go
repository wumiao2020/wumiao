package models

import (
	"wumiao/extend"
)

type Menu struct {
	Id         int64       `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Status     int         `json:"status" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Title      string      `json:"title" xorm:"comment('主题') index VARCHAR(255)"`
	Identifier string      `json:"identifier" xorm:"unique VARCHAR(36)"`
	CreatedAt  extend.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt  extend.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
}
