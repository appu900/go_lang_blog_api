package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PostId  primitive.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	UserId  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
}