package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Gender      string             `json:"gender"`
	FirstName   string             `json:"firstname,omitempty"`
	LastName    string             `json:"lastname,omitempty"`
	PhoneNumber int                `json:"number"`
	Salary      float64            `json:"salary"`
	Active      bool               `json:"active"`
}
