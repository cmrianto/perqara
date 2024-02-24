package model

import "time"

type ContentType struct {
	tableName struct{} `gorm:"content_types"`
	Id        int64    `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
