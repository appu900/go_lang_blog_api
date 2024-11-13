package services

import (
	"blog-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService struct {
	DB *mongo.Collection
}

func (ps *PostService) createPost() (*models.Post, error) {
	return nil, nil
}
