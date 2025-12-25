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

type LikeChain struct {
	Index  uint
	UserID uint
	PostID uint
}
