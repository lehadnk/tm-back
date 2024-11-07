package user

import (
	"awesomeProject/src/user/domain"
	"awesomeProject/src/user/persistence"
)

type UserService struct {
	userDao *persistence.UserDao
}

func NewUserService(userDao *persistence.UserDao) *UserService {
	var newUserService = UserService{
		userDao,
	}
	return &newUserService
}

func (userService *UserService) GetUserById(userId int) *domain.User {
	user := userService.userDao.GetUserById(userId)
	return user
}
