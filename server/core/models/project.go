package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name,omitempty"`
	ShortDescription string             `json:"shortDescription,omitempty"`
	Tags             string             `json:"tags,omitempty"` // TODO: Make this a list of string or define some kind of formatting
	Link             string             `json:"link,omitempty"`
	ImageUrl         string             `json:"imageUrl,omitempty"`
	Type             ProjectType        `json:"type,omitempty"` // Work , Personal, Open Source, School
}

type ProjectType string

const (
	WorkProj       ProjectType = "work"
	PersonalProj               = "personal"
	OpenSourceProj             = "open_source"
	SchoolProj                 = "school"
)
