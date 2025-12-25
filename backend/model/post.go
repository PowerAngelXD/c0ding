package model

import "time"

// Post下的评论
type Comment struct {
	Index uint `json:"index" gorm:"primaryKey"`
	// 评论者昵称
	UserName string `json:"userName"`
	// 评论者ID
	UserID uint `json:"userId"`
	// 评论内容
	Content string `json:"content"`

	// 赞同该评论的数量
	Likes uint `json:"likes"`
}

type Reply struct {
	Index uint `json:"index" gorm:"primaryKey"`
	// 对哪个Comment做出的回复
	CommentID uint `json:"commentId"`
	// 回复人昵称
	UserName string `json:"userName"`
	// 回复人ID
	UserID uint `json:"userId"`
	// 回复内容
	Content string `json:"content"`

	// 赞同该回复的数量
	Likes uint `json:"likes"`
}

// 帖子
type Post struct {
	// 帖子ID
	PostID     uint   `json:"postId" gorm:"primaryKey"`
	AuthorName string `json:"authorName"`
	AuthorID   uint   `json:"authorId"`
	// 发布时间
	PostTime time.Time `json:"postTime"`
	// 最后更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 帖子内容
	Content string `json:"content"`

	// 帖子的喜欢数量
	Likes uint `json:"likes"`
	// 帖子的浏览量
	Views uint `json:"views"`
	// 帖子的tags
	Tags []string `json:"tags"`
}
