package models

type Stadium struct {
	Id        int
	Name      string
	Location  string
	Capacity  int
	Altitude  int
	YearBuilt int `db:"year_built"`
}
