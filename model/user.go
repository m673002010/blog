package model

import "time"

type User struct {
	Id       int
	Username string
	Account  string
	Password string
	CreateAt time.Time
	UpdateAt time.Time
}

func (User) TableName() string {
	return "user"
}
