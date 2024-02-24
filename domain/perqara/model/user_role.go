package model

import (
	"time"
)

type UserRole struct {
	tableName struct{} `gorm:"user_roles"`
	Id        int64    `gorm:"primaryKey"`
	UserId    int64
	RoleId    int64
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (rp *UserRole) MakeMapArrayByUserId(datas []*UserRole) map[int64][]*UserRole {
	userRolenMap := make(map[int64][]*UserRole, 0)
	for i := range datas {
		data := datas[i]
		userRolenMap[data.UserId] = append(userRolenMap[data.UserId], data)
	}
	return userRolenMap
}
