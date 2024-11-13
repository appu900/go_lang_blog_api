
package services

import (
	"blog-api/models"
	"blog-api/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DB *mongo.Collection
}

type UserResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

// ** signup a user
func (us *UserService) RegisterUser(username, password, email string) (*UserResponse, error) {
	user := models.User{Username: username, Password: password, Email: email}
	user.ID = primitive.NewObjectID()

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	_, err := us.DB.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	// jenrate jwt token
	token, err := utils.GenerateToken(user.ID.Hex(), user.Username)
	if err != nil {
		return nil, err
	}

	response := &UserResponse{
		User:  user,
		Token: token,
	}

	return response, nil
}

func (us *UserService) LoginUser(username string, password string) (*UserResponse, error) {
	var user models.User
	err := us.DB.FindOne(context.TODO(), models.User{Username: username}).Decode(&user)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, errors.New("invalid password")
	}

	// ** create a jwt token

	token, err := utils.GenerateToken(user.ID.Hex(), user.Username)
	if err != nil {
		return nil, err
	}

	response := &UserResponse{
		User:  user,
		Token: token,
	}

	return response, nil
}



