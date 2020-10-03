package models

import (
	"time"
)

type AdminRoles struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name        string    `json:"name" xorm:"not null comment('角色名称') VARCHAR(255)"`
	Description string    `json:"description" xorm:"comment('备注') VARCHAR(255)"`
	CreatedAt   time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
}
