package models

import (
	"time"
)

type Admins struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name            string    `json:"name" xorm:"not null VARCHAR(255)"`
	Email           string    `json:"email" xorm:"not null unique VARCHAR(255)"`
	EmailVerifiedAt time.Time `json:"email_verified_at" xorm:"TIMESTAMP"`
	Password        string    `json:"password" xorm:"not null VARCHAR(255)"`
	Status          int       `json:"status" xorm:"not null default 0 TINYINT(1)"`
	RememberToken   string    `json:"remember_token" xorm:"VARCHAR(100)"`
	CreatedAt       time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
}
