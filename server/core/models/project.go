package models

type Project struct {
	Name             string
	ShortDescription string
	Tags             string // TODO: Make this a list of string or define some kind of formatting
	Link             string
	ImageUrl         string
	Type             string // Work , Personal, Open Source
}
