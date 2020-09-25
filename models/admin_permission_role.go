package models

type AdminPermissionRole struct {
	PermissionId int `json:"permission_id" xorm:"not null INT(11)"`
	RoleId       int `json:"role_id" xorm:"not null INT(11)"`
}
