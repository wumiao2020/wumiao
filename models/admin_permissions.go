package models

import (
	"time"
)

type AdminPermissions struct {
	Id         int64     `json:"id" xorm:"pk autoincr comment('菜单ID') BIGINT(20)"`
	Name       string    `json:"name" xorm:"not null comment('菜单名称') VARCHAR(50)"`
	ParentId   int64     `json:"parent_id" xorm:"default 0 comment('父菜单ID') BIGINT(20)"`
	OrderNum   int       `json:"order_num" xorm:"default 0 comment('显示顺序') INT(4)"`
	Url        string    `json:"url" xorm:"default '#' comment('请求地址') VARCHAR(200)"`
	Target     string    `json:"target" xorm:"default '' comment('打开方式（menuItem页签 menuBlank新窗口）') VARCHAR(20)"`
	Type       string    `json:"type" xorm:"default '' comment('菜单类型（M目录 C菜单 F按钮）') CHAR(1)"`
	Visible    string    `json:"visible" xorm:"default '0' comment('菜单状态（0显示 1隐藏）') CHAR(1)"`
	Perms      string    `json:"perms" xorm:"comment('权限标识') VARCHAR(100)"`
	Icon       string    `json:"icon" xorm:"default '#' comment('菜单图标') VARCHAR(100)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(500)"`
}
