package models

type AdminPermissionRole struct {
	PermissionId int64 `json:"permission_id" xorm:"not null BIGINT(20)"`
	RoleId       int64 `json:"role_id" xorm:"not null BIGINT(20)"`
}
