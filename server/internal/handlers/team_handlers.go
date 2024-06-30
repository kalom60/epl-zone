package handlers

import (
	"net/http"

	"github.com/kalom60/epl-zone/internal/repository"
	"github.com/labstack/echo/v4"
)

type TeamHandlers interface {
	GetTeams(c echo.Context) error
}

type TeamHandler struct {
	teamRepo repository.TeamRepository
}

func NewTeamHandler(teamRepo repository.TeamRepository) TeamHandlers {
	return &TeamHandler{teamRepo: teamRepo}
}

func (h *TeamHandler) GetTeams(c echo.Context) error {
	teams, err := h.teamRepo.GetTeams()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	return c.JSON(http.StatusOK, teams)
}
