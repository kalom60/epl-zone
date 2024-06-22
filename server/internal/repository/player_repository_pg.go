package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kalom60/epl-zone/internal/models"
)

type PgPlayerRepository struct {
	db *sqlx.DB
}

func NewPgPlayerRepository(db *sqlx.DB) PlayerRepository {
	return &PgPlayerRepository{db: db}
}

func (r *PgPlayerRepository) GetAllPlayers() (*[]models.Player, error) {
	var players []models.Player

	err := r.db.Select(&players, "SELECT * FROM players")
	if err != nil {
		return nil, err
	}

	return &players, nil
}

func (r *PgPlayerRepository) GetPlayersByTeam(team string) (*[]models.Player, error) {
	var players []models.Player

	err := r.db.Select(&players, "SELECT * FROM players WHERE team=$1", team)
	if err != nil {
		return nil, err
	}

	return &players, nil
}

func (r *PgPlayerRepository) GetPlayersByPostiton(pos string) (*[]models.Player, error) {
	players := []models.Player{}

	err := r.db.Select(&players, "SELECT * FROM players WHERE pos=$1", pos)
	if err != nil {
		return nil, err
	}

	return &players, nil
}
