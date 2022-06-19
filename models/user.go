package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Age     int                `json:"age" bson:"age"`
	Address Address            `json:"address" bson:"address"`
}

type Shoe struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	User_id string             `json:"user_id" bson:"user_id"`
	Brand   string             `json:"brand" bson:"brand"`
	Size    int                `json:"size" bson:"size"`
}
