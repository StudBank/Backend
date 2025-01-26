package models

import (
	"time"

	"gorm.io/gen"
)

type User struct {
	ID   uint
	Name string

	Email string
	Phone string

	Password  string
	LastLogin time.Time

	ReqCount int

	TimeRegistered time.Time
}

type UserQuerier interface {
	// SELECT * FROM @@table WHERE id=@id LIMIT 1
	GetByID(id int) (gen.T, error)
}
