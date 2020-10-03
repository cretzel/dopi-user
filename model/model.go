package model

type User struct {
	Username string
	Password string
}

type UserService interface {
	CreateUser(user *User) (*User, error)
	//GetUsers() ([]*User, error)
	GetUserByUsername(username string) (*User, error)
	Login(username string, password string) (*User, error)
}
