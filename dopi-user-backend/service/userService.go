package service

import (
	"dopi-user/model"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *Session) *UserService {
	collection := session.session.DB("dopi-user").C("users")
	userService := UserService{collection: collection}
	userService.CreateAdminUser()
	return &userService
}

func (u *UserService) CreateAdminUser() {
	admin, err := u.GetUserByUsername("admin")
	if err != nil {
		log.Println("Creating admin user")
		password := os.Getenv("ADMIN_PASSWORD")
		passwordCrypted, err := Generate(password)
		if err != nil {
			panic("Could not crypt password")
		}
		admin = &model.User{
			Username: "admin",
			Password: passwordCrypted,
			Roles:    []string{"admin"},
		}
		u.CreateUser(admin)
	}
	log.Println("admin user: ", admin.Username)
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	err := u.collection.Insert(&user)
	if err != nil {
		log.Printf("Error creating user: %s", err.Error())
		return user, err
	}
	return u.GetUserByUsername(user.Username)
}

func (u *UserService) UpdateUser(user *model.User) (*model.User, error) {
	log.Println("Save user", user.Roles)
	_, err := u.collection.UpsertId(user.Username, user)
	//err := u.collection.Update(bson.M{"_id": user.Username}, bson.M{"roles": user.Roles})
	if err != nil {
		log.Printf("Error updating user: %s", err.Error())
		return user, err
	}
	return u.GetUserByUsername(user.Username)
}

func (u *UserService) GetUserByUsername(username string) (*model.User, error) {
	user := model.User{}
	err := u.collection.FindId(username).One(&user)
	return &user, err
}

func (u *UserService) GetUsers() ([]model.User, error) {
	users := []model.User{}
	err := u.collection.Find(bson.M{}).All(&users)
	return users, err
}

func (u *UserService) Login(username string, password string) (*model.User, error) {
	user, err := u.GetUserByUsername(username)
	if err != nil {
		return &model.User{}, err
	}
	err = Compare(user.Password, password)
	if err != nil {
		return &model.User{}, err
	}
	return user, err
}
