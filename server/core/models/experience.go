package models

import (
	"time"
)

type Experience struct {
	Organization    string
	OrganizationUrl string
	Role            string
	Description     string
	Type            string // Remote/Freelance, FullTime, PartTime, School
	StartAt         time.Time
	EndAt           time.Time
}
