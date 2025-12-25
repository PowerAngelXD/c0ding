package model

import "time"

type Statistics struct {
	Likes uint
	Views uint
}

type Comment struct {
	Index    uint
	UserName string
	UserID   uint

	Content string

	Stats Statistics
}

type Reply struct {
	Index     uint
	CommentID uint
	UserName  string
	UserID    uint

	Content string

	Stats Statistics
}

type Post struct {
	PostID     uint
	AuthorName string
	AuthorID   uint
	PostTime   time.Time
	UpdateTime time.Time

	Content string

	Stats Statistics
}
