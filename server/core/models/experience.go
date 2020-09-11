package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Organization    string             `json:"organization,omitempty"`
	OrganizationUrl string             `json:"organizationUrl,omitempty"`
	Role            string             `json:"role,omitempty"`
	Description     string             `json:"description,omitempty"`
	Type            string             `json:"type,omitempty"` // Remote/Freelance, FullTime, PartTime, School
	StartAt         time.Time          `json:"startAt,omitempty"`
	EndAt           time.Time          `json:"endAt,omitempty"`
}

type ExperienceType string

const (
	FreelanceExp ExperienceType = "freelance"
	FullTimeExp  ExperienceType = "fulltime"
	PartTimeExp  ExperienceType = "parttime"
	SchoolExp    ExperienceType = "school"
)
