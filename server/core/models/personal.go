package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Personal struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name,omitempty"`
	ShortDescription string             `json:"shortDescription,omitempty"`
	AboutMe          string             `json:"aboutMe,omitempty"`
	BOD              time.Time          `json:"bOD,omitempty"`
	Location         string             `json:"location,omitempty"`
	Links            Links              `json:"links,omitempty"`
	ImageUrl         string             `json:"imageUrl,omitempty"`
}

type Links struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	LinkedIn string             `json:"linkedIn,omitempty"`
	GitHub   string             `json:"github,omitempty"`
	Medium   string             `json:"medium,omitempty"`
	Resume   string             `json:"resume,omitempty"`
	Email    string             `json:"email,omitempty"`
}
