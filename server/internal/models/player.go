package models

import (
	"github.com/kalom60/epl-zone/pkg/nulls"
)

type Player struct {
	ID              int               `json:"id" db:"id"`
	Player          string            `json:"player" db:"player"`
	Nation          nulls.NullString  `json:"nation" db:"nation"`
	Position        nulls.NullString  `json:"position" db:"pos"`
	Age             nulls.NullFloat64 `json:"age" db:"age"`
	MatchesPlayed   nulls.NullInt64   `json:"matchesPlayed" db:"mp"`
	Starts          nulls.NullInt64   `json:"starts" db:"starts"`
	MinutesPlayed   nulls.NullFloat64 `json:"minutesPlayed" db:"min"`
	Goals           nulls.NullFloat64 `json:"goals" db:"gls"`
	Assists         nulls.NullFloat64 `json:"assists" db:"ast"`
	PenaltiesScored nulls.NullFloat64 `json:"penalitiesScored" db:"pk"`
	YellowCards     nulls.NullFloat64 `json:"yellowCards" db:"crdy"`
	RedCards        nulls.NullFloat64 `json:"redCards" db:"crdr"`
	ExpectedGoals   nulls.NullFloat64 `json:"expectedGoals" db:"xg"`
	ExpectedAssists nulls.NullFloat64 `json:"expectedAssists" db:"xag"`
	TeamName        string            `json:"teamName" db:"team"`
}
