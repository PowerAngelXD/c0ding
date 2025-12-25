package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName" gorm:"unique"`
	Password string `json:"-"`

	// 用户等级
	Level uint `json:"level"`
	// 用户经验
	Exp uint `json:"exp"`
	// 用户获得的赞
	RecvLikes uint `json:"recvLikes"`
	// 关注该用户的用户数量
	RecvSubscription uint `json:"recvSubscription"`
}

type Like struct {
	// 主键索引
	Index uint `json:"index" gorm:"primaryKey"`
	// 用户ID
	UserID uint `json:"userId"`
	// 被点赞的帖子ID
	PostID uint `json:"postId"`
}

type View struct {
	// 主键索引
	Index uint `json:"index" gorm:"primaryKey"`
	// 用户ID
	UserID uint `json:"userId"`
	// 被浏览的帖子ID
	PostID uint `json:"postId"`
}
