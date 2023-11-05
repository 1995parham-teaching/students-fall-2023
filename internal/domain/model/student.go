package model

type Student struct {
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	EnteranceYear int    `json:"enterance_year,omitempty"`
	Courses       []Course
}
