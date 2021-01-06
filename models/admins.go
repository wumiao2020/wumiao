package models

import (
	"wumiao/extend"
)

type Admins struct {
	Id              int64       `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Name            string      `json:"name" xorm:"not null VARCHAR(255)"`
	Email           string      `json:"email" xorm:"not null unique VARCHAR(255)"`
	EmailVerifiedAt extend.Time `json:"email_verified_at" xorm:"TIMESTAMP"`
	Password        string      `json:"password" xorm:"not null VARCHAR(255)"`
	PasswordConfirm string      `json:"password_confirm" xorm:"-"`
	Status          int         `json:"status" xorm:"not null default 0 INT(1)"`
	RememberToken   string      `json:"remember_token" xorm:"VARCHAR(100)"`
	LockExpires     extend.Time `json:"lock_expires" xorm:"lock_expires TIMESTAMP"`
	CreatedAt       extend.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt       extend.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
}
