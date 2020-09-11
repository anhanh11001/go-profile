package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Personal struct {
	ID               primitive.ObjectID
	Name             string
	ShortDescription string
	AboutMe          string
	BOD              time.Time
	Location         string
	Links            Links
	ImageUrl         string
}

type Links struct {
	LinkedIn string
	GitHub   string
	Medium   string
	Resume   string
	Email    string
}
