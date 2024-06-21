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
