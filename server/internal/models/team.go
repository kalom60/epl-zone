package models

type Team struct {
	ID   int    `json:"id" db:"id"`
	Team string `json:"team" db:"team"`
	Logo string `json:"logo" db:"logo"`
}
