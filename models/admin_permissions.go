package models

import (
	"wumiao/extend"
)

type AdminPermissions struct {
	Id          int         `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name        string      `json:"name" xorm:"not null comment('权限名') VARCHAR(255)"`
	Label       string      `json:"label" xorm:"comment('权限解释名称') VARCHAR(255)"`
	Description string      `json:"description" xorm:"comment('描述与备注') VARCHAR(255)"`
	Type        string      `json:"type" xorm:"comment('归类') VARCHAR(255)"`
	Cid         int         `json:"cid" xorm:"not null default 0 comment('级别') TINYINT(4)"`
	IsMenu      int         `json:"is_menu" xorm:"not null default 0 TINYINT(1)"`
	Ordering    int         `json:"ordering" xorm:"not null default 0 INT(10)"`
	Icon        string      `json:"icon" xorm:"comment('图标') VARCHAR(255)"`
	CreatedAt   extend.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt   extend.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
}
