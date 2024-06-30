package repository

import (
	"github.com/kalom60/epl-zone/internal/models"
	"github.com/kalom60/epl-zone/internal/scrape"
)

type TeamRepository interface {
	GetTeams() (*[]models.Team, error)
	AddTeams([]scrape.Team) error
}
