package model

import "time"

type Role struct {
	tableName struct{} `gorm:"roles"`
	Id        int64    `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
