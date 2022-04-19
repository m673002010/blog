package model

import "time"

type Article struct {
	Id        int
	Title     string
	Content   string
	UserId    int
	ViewCount int
	CreateAt  time.Time
	UpdateAt  time.Time
}

func (Article) TableName() string {
	return "article"
}
