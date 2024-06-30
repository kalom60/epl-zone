package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kalom60/epl-zone/internal/models"
	"github.com/kalom60/epl-zone/internal/scrape"
)

type PgTeamRepository struct {
	db *sqlx.DB
}

func NewPgTeamRepository(db *sqlx.DB) TeamRepository {
	return &PgTeamRepository{db: db}
}

func (r *PgTeamRepository) GetTeams() (*[]models.Team, error) {
	var teams []models.Team

	err := r.db.Select(&teams, "SELECT * FROM teams")
	if err != nil {
		return nil, err
	}

	return &teams, nil
}

func (r *PgTeamRepository) AddTeams(teams []scrape.Team) error {
	for _, team := range teams {
		_, err := r.db.Exec("INSERT INTO teams (team, logo) VALUES ($1, $2)", team.Name, team.Logo)
		if err != nil {
			return err
		}
	}

	return nil
}
