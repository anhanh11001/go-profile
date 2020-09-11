package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"_id,omitempty`
	CreatedAt        time.Time          `json:"createdAt,omitempty"`
	ShortDescription string             `json:"shortDescription,omitempty"`
	ImageUrl         string             `json:"imageUrl,omitempty"`
	Url              string             `json:"url,omitempty"`
}
