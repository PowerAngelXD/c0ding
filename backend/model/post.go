package model

import "time"

type Statistics struct {
	Likes uint
	Views uint
}

type Comment struct {
	UserName string
	UserID   uint

	Content string

	Stats Statistics

	Replies []Comment
}

type Post struct {
	PostID     uint
	AuthorName string
	AuthorID   uint
	PostTime   time.Time
	UpdateTime time.Time

	Content      string
	ResourceURLs []string

	Stats Statistics

	Comments []Comment
}
