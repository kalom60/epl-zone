package models

import "database/sql"

type Player struct {
	ID              int             `json:"id" db:"id"`
	Player          string          `json:"player" db:"player"`
	Nation          sql.NullString  `json:"nation" db:"nation"`
	Position        sql.NullString  `json:"position" db:"pos"`
	Age             sql.NullFloat64 `json:"age" db:"age"`
	MatchesPlayed   sql.NullInt64   `json:"matchesPlayed" db:"mp"`
	Starts          sql.NullInt64   `json:"starts" db:"starts"`
	MinutesPlayed   sql.NullFloat64 `json:"minutesPlayed" db:"min"`
	Goals           sql.NullFloat64 `json:"goals" db:"gls"`
	Assists         sql.NullFloat64 `json:"assists" db:"ast"`
	PenaltiesScored sql.NullFloat64 `json:"penalitiesScored" db:"pk"`
	YellowCards     sql.NullFloat64 `json:"yellowCards" db:"crdy"`
	RedCards        sql.NullFloat64 `json:"redCards" db:"crdr"`
	ExpectedGoals   sql.NullFloat64 `json:"expectedGoals" db:"xg"`
	ExpectedAssists sql.NullFloat64 `json:"expectedAssists" db:"xag"`
	TeamName        string          `json:"teamName" db:"team"`
}
