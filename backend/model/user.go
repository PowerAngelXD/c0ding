package model

type User struct {
	ID       uint
	UserName string
	Password string

	RecvLikes        uint
	RecvSubscription uint

	LikePosts      []uint
	HistoryWatched []uint
}
