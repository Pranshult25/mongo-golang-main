package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Vote struct {
	Author    string              `json:"author" bson:"author"`
	CommentId primitive.ObjectID  `json:"commentId" bson:"commentId"`
	Direction int                 `json:"direction" bson:"direction"`
}