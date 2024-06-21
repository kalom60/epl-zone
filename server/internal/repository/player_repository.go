package repository

import "github.com/kalom60/epl-zone/internal/models"

type PlayerRepository interface {
	GetAllPlayers() (*[]models.Player, error)
}
