package models

import (
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "gopkg.in/mgo.v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Author   string              `json:"author" bson:"author"`
	Title    string              `json:"title" bson:"title,omitempty"`
	Body     string              `json:"body" bson:"body"`
	Category string              `json:"category" bson:"category"`
	PostedAt time.Time           `json:"postedAt" bson:"postedAt"`
	ParentId bson.ObjectId       `json:"parentId,omitempty" bson:"parentId,omitempty"`
	RootId   bson.ObjectId       `json:"rootId,omitempty" bson:"rootId,omitempty"`
}