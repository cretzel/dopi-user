package model

type User struct {
	Username string `bson:"_id"`
	Roles    []string
	Password string
}

type UserService interface {
	CreateUser(user *User) (*User, error)
	GetUsers() ([]User, error)
	GetUserByUsername(username string) (*User, error)
	Login(username string, password string) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(username string) error
}
