package model

type User struct {
	ID       uint
	UserName string
	Password string

	// 用户等级
	Level uint
	// 用户经验
	Exp uint
	// 用户获得的赞
	RecvLikes uint
	// 关注该用户的用户数量
	RecvSubscription uint
}

type Like struct {
	// 主键索引
	Index uint
	// 用户ID
	UserID uint
	// 被点赞的帖子ID
	PostID uint
}

type View struct {
	// 主键索引
	Index uint
	// 用户ID
	UserID uint
	// 被浏览的帖子ID
	PostID uint
}
