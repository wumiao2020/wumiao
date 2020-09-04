package models

import (
	"html/template"
	"time"
)

type Page struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ParentId    int       `json:"parent_id" xorm:"not null default 0 comment('父级ID') TINYINT(4)"`
	Type        string    `json:"type" xorm:"not null default 'note' comment('文档类型') VARCHAR(255)"`
	Status      int       `json:"status" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Title       string    `json:"title" xorm:"comment('主题') index VARCHAR(255)"`
	Keywords    string    `json:"keywords" xorm:"comment('关键字') index VARCHAR(255)"`
	Description string    `json:"description" xorm:"comment('描述') VARCHAR(255)"`
	CreatedAt   time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	Identifier  string    `json:"identifier" xorm:"unique VARCHAR(36)"`
	Uuid        string    `json:"uuid" xorm:"comment('UUID') unique VARCHAR(36)"`
	Content     template.HTML
}
