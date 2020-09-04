package models

import (
	"time"
)

type File struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('主ID') BIGINT(20)"`
	PgeId     int       `json:"pge_id" xorm:"not null default 0 comment('上级ID') index TINYINT(4)"`
	Name      string    `json:"name" xorm:"comment('名字') index VARCHAR(255)"`
	Path      string    `json:"path" xorm:"comment('文件路径') VARCHAR(255)"`
	Size      int       `json:"size" xorm:"comment('文件大小') INT(11)"`
	Ext       int       `json:"ext" xorm:"comment('文件后缀') INT(11)"`
	Status    int       `json:"status" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Type      string    `json:"type" xorm:"comment('文档类型') index VARCHAR(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
}
