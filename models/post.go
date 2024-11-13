package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Author    primitive.ObjectID `json:"author,omitempty" bson:"author,omitempty"`
	LikeCount int                `json:"like_count,omitempty" bson:"like_count,omitempty"`
}
