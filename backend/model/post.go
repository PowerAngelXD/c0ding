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

	Content string

	Likes uint
	Views uint
}

type Post struct {
	PostID     uint
	AuthorName string
	AuthorID   uint
	PostTime   time.Time
	UpdateTime time.Time

	Content string

	Likes uint
	Views uint
}
