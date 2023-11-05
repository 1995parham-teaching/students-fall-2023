package model

type Student struct {
	ID           uint64   `json:"id,omitempty"`
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	EntranceYear int      `json:"entrance_year,omitempty"`
	Courses      []Course `json:"courses,omitempty"`
}
