package dto

type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func NewUser(name string, email string, password string, role string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

type UsersList struct {
	UsersArray []User
	UsersCount int
}

func NewUsersList(usersArray []User, usersCount int) *UsersList {
	return &UsersList{
		UsersArray: usersArray,
		UsersCount: usersCount,
	}
}
