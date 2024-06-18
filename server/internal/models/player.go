package models

type Player struct {
	PlayerName      string  `json:"playerName" db:"player_name"`
	Nation          string  `json:"nation" db:"nation"`
	Position        string  `json:"position" db:"position"`
	Age             float64 `json:"age" db:"age"`
	MatchesPlayed   int     `json:"matchesPlayed" db:"matches_played"`
	Starts          int     `json:"starts" db:"starts"`
	MinutesPlayed   float64 `json:"minutesPlayed" db:"minutes_played"`
	Goals           float64 `json:"goals" db:"goals"`
	Assists         float64 `json:"assists" db:"assists"`
	PenaltiesScored float64 `json:"penalitiesScored" db:"penalities_scored"`
	YellowCards     float64 `json:"yellowCards" db:"yellow_cards"`
	RedCards        float64 `json:"redCards" db:"red_cards"`
	ExpectedGoals   float64 `json:"expectedGoals" db:"expected_goals"`
	ExpectedAssists float64 `json:"expectedAssists" db:"expected_assists"`
	TeamName        string  `json:"teamName" db:"team_name"`
}
