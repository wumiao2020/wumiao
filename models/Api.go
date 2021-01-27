package models

import (
	"wumiao/extend"
)

type Api struct {
	Id          int64       `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ProductName string      `json:"product_name" xorm:"not null comment('商品名称') VARCHAR(191)"`
	Spu         string      `json:"spu" xorm:"not null comment('货号') VARCHAR(191)"`
	Price       string      `json:"price" xorm:"not null default 0.00 comment('折扣价') DECIMAL(8,2)"`
	Number      int         `json:"number" xorm:"not null default 0 comment('数量') INT(11)"`
	CreatedAt   extend.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt   extend.Time `json:"updated_at" xorm:"not null updated comment('更新时间') DATETIME"`
}
