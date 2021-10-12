package service

import (
	"context"
	"dopi-user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type UserService struct {
	client     *Client
	collection *mongo.Collection
}

func NewUserService(client *Client) *UserService {
	collection := client.GetCollection("dopi-user", "users")
	userService := UserService{client: client, collection: collection}
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
		_, err = u.CreateUser(admin)
		if err != nil {
			return
		}
	}
	log.Println("admin user:", admin.Username)
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	_, err := u.collection.InsertOne(context.TODO(), &user)
	if err != nil {
		log.Printf("Error creating user: %s", err.Error())
		return user, err
	}
	return u.GetUserByUsername(user.Username)
}

func (u *UserService) UpdateUser(user *model.User) (*model.User, error) {
	_, err := u.collection.ReplaceOne(context.TODO(), bson.M{"_id": user.Username}, &user)
	if err != nil {
		log.Printf("Error updating user: %s", err.Error())
		return user, err
	}
	return u.GetUserByUsername(user.Username)
}

func (u *UserService) GetUserByUsername(username string) (*model.User, error) {
	user := model.User{}
	err := u.collection.FindOne(context.TODO(), bson.M{"_id": username}).Decode(&user)
	return &user, err
}

func (u *UserService) GetUsers() ([]model.User, error) {
	var users []model.User
	cursor, err := u.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error getting users: %s", err.Error())
		return users, err
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Printf("Error getting users: %s", err.Error())
		return users, err
	}

	return users, nil
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
