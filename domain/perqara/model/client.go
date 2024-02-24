package model

import (
	"time"
)

type Client struct {
	tableName struct{} `gorm:"clients"`
	Id        int64    `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
