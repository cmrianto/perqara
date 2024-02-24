package model

import "time"

type Permission struct {
	tableName     struct{} `gorm:"permissions"`
	Id            int64    `gorm:"primaryKey"`
	Name          string
	ContentTypeId int64
	CreatedAt     time.Time `gorm:"<-create"`
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
}
