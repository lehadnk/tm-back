package dto

type User struct {
	Id       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
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
	Users []User `json:"users"`
	Count int    `json:"count"`
}

func NewUsersList(usersArray []User, usersCount int) *UsersList {
	return &UsersList{
		Users: usersArray,
		Count: usersCount,
	}
}
