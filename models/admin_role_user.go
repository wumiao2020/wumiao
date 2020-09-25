package models

type AdminRoleUser struct {
	RoleId int `json:"role_id" xorm:"not null INT(11)"`
	UserId int `json:"user_id" xorm:"not null INT(11)"`
}
