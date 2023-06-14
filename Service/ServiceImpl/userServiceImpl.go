package ServiceImpl

import (
	"LilaGames/Models"
	"errors"
)

type UserServiceImpl struct {
}

// Users : In Memory User DataBase
var UserData map[string]*Models.User //Mapping of user to there emailId.

func NewInitializeUser() {
	UserData = make(map[string]*Models.User)
	return
}

func (u *UserServiceImpl) CreateUser(name, email, areaCode string) (*Models.User, error) {
	validUser := validateUser(email) // todo : Validate in controller
	if !validUser {
		err := errors.New("user already exist")
		return nil, err
	}
	user := &Models.User{
		Name:        name,
		EmailID:     email,
		AreaCode:    areaCode,
		CurrentMode: "",
	}
	UserData[email] = user

	return user, nil
}

func validateUser(email string) bool {
	if _, ok := UserData[email]; ok {
		return false
	}
	return true
}

func (u *UserServiceImpl) GetUser(email string) (*Models.User, error) {
	if _, ok := UserData[email]; ok {
		return UserData[email], nil
	}
	err := errors.New("user not found")
	return nil, err
}

//todo : if player is in gmaeplay how to change his areaCode
func (u *UserServiceImpl) UpdateAreaCode(areaCode string, user *Models.User) bool {
	user.AreaCode = areaCode
	return true
}
