package model

import (
	"database/sql"
	"time"
)

type User struct {
	tableName struct{} `gorm:"users"`
	Id        int64    `gorm:"primaryKey"`
	Name      string
	UserName  string
	Type      int64
	Password  string
	ClientId  sql.NullInt64 `gorm:"default:null"`
	CreatedAt time.Time     `gorm:"<-create"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
