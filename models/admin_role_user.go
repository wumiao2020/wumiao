package models

type AdminRoleUser struct {
	RoleId int64 `json:"role_id" xorm:"not null BIGINT(20)"`
	UserId int64 `json:"user_id" xorm:"not null BIGINT(20)"`
}
