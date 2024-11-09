package user

import (
	"golang.org/x/crypto/bcrypt"
	"tm/src/user/dto"
	"tm/src/user/persistence"
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

func (userService *UserService) GetUserById(userId int) *dto.User {
	user := userService.userDao.GetUserById(userId)
	return user
}

func (userService *UserService) CreateUser(name string, email string, password string, role string) *dto.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := dto.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}
	userService.userDao.CreateUser(&newUser)
	return &newUser

}

func (userService *UserService) UpdateUser(name string, email string, password string, role string, userId int) *dto.User {
	userService.userDao.EditUser(name, email, password, role, userId)
	return userService.userDao.GetUserById(userId)
}

func (userService *UserService) DeleteUser(userId int) {
	userService.userDao.DeleteUser(userId)
}

func (userService *UserService) GetUsersList(sort string, page int, pageSize int) *dto.UsersList {
	usersList := dto.NewUsersList(
		userService.userDao.GetUsersList(sort, page, pageSize),
		userService.userDao.GetUsersCount(),
	)
	return usersList
}
