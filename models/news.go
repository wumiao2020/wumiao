package models

import (
	"time"
)

type News struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Catid       int       `json:"catid" xorm:"not null comment('栏目id') index(category) index SMALLINT(5)"`
	Title       string    `json:"title" xorm:"comment('主题') VARCHAR(255)"`
	Thumb       string    `json:"thumb" xorm:"comment('缩略图') VARCHAR(255)"`
	Keywords    string    `json:"keywords" xorm:"comment('关键字') VARCHAR(255)"`
	Description string    `json:"description" xorm:"comment('描述') TEXT"`
	Hits        int       `json:"hits" xorm:"comment('浏览数') index INT(10)"`
	Uid         int       `json:"uid" xorm:"not null comment('作者id') index INT(10)"`
	Author      string    `json:"author" xorm:"not null comment('作者名称') VARCHAR(50)"`
	Status      int       `json:"status" xorm:"not null comment('状态') index(category) index TINYINT(2)"`
	Identifier  string    `json:"identifier" xorm:"comment('地址') VARCHAR(255)"`
	Inputip     string    `json:"inputip" xorm:"comment('录入者ip') VARCHAR(15)"`
	SortOrder   int       `json:"sort_order" xorm:"default 0 comment('排序值') index INT(10)"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('录入时间') TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') index TIMESTAMP"`
}
