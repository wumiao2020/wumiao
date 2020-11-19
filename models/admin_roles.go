package models

import (
	"time"
)

type AdminRoles struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Name        string    `json:"name" xorm:"not null comment('角色名称') VARCHAR(255)"`
	Description string    `json:"description" xorm:"comment('备注') VARCHAR(255)"`
	CreatedAt   time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
}
