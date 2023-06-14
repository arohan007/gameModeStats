package Service

import "github.com/arohan007/gameModeStats/Models"

type UserService interface {
	CreateUser(name, email, areaCode string) (*Models.User, error)
	GetUser(email string) (*Models.User, error)
	UpdateAreaCode(areaCode string, user *Models.User) bool
}
