package model

type Post struct {
	PostId    int    `json:"postId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    int    `json:"userId"`
	CateId    int    `json:"cateId"`
	ViewCount int    `json:"viewCount"`
	CreateAt  string `json:"createAt"`
	UpdateAt  string `json:"updateAt"`
}
