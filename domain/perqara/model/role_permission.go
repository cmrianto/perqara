package model

import "time"

type RolePermission struct {
	tableName    struct{} `gorm:"role_permissions"`
	Id           int64    `gorm:"primaryKey"`
	RoleId       int64
	PermissionId int64
	CreatedAt    time.Time `gorm:"<-create"`
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

func (rp *RolePermission) MakeMapArrayByRoleId(datas []*RolePermission) map[int64][]*RolePermission {
	rolePermissionMap := make(map[int64][]*RolePermission, 0)
	for i := range datas {
		data := datas[i]
		rolePermissionMap[data.RoleId] = append(rolePermissionMap[data.RoleId], data)
	}
	return rolePermissionMap
}
