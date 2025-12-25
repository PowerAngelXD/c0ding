package model

import "time"

type Comment struct {
	Index    uint   `json:"index" gorm:"primaryKey"`
	UserName string `json:"userName"`
	UserID   uint   `json:"userId"`

	Content string `json:"content"`

	Likes uint `json:"likes"`
	Views uint `json:"views"`
}

type Reply struct {
	Index     uint   `json:"index" gorm:"primaryKey"`
	CommentID uint   `json:"commentId"`
	UserName  string `json:"userName"`
	UserID    uint   `json:"userId"`

	Content string `json:"content"`

	Likes uint `json:"likes"`
	Views uint `json:"views"`
}

type Post struct {
	PostID     uint      `json:"postId" gorm:"primaryKey"`
	AuthorName string    `json:"authorName"`
	AuthorID   uint      `json:"authorId"`
	PostTime   time.Time `json:"postTime"`
	UpdateTime time.Time `json:"updateTime"`

	Content string `json:"content"`

	Likes uint `json:"likes"`
	Views uint `json:"views"`
}
