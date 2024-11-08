package user

import (
	"awesomeProject/src/user/domain"
	"awesomeProject/src/user/persistence"
	"golang.org/x/crypto/bcrypt"
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

func (userService *UserService) CreateUser(name string, email string, password string, role string) *domain.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}
	userService.userDao.CreateUser(&newUser)
	return &newUser

}

func (userService *UserService) UpdateUser(name string, email string, password string, role string, userId int) *domain.User {
	userService.userDao.EditUser(name, email, password, role, userId)
	return userService.userDao.GetUserById(userId)
}

func (userService *UserService) DeleteUser(userId int) {
	userService.userDao.DeleteUser(userId)
}

func (userService *UserService) GetUsersList(sort string, page int, pageSize int) *domain.UsersList {
	usersList := domain.NewUsersList(
		userService.userDao.GetUsersList(sort, page, pageSize),
		userService.userDao.GetUsersCount(),
	)
	return usersList
}
