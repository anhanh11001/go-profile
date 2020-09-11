package models

import (
	"time"
)

type Blog struct {
	Name             string
	CreatedAt        time.Time
	ShortDescription string
	ImageUrl         string
	Url              string
}
