package handlers

import (
	"net/http"

	"github.com/kalom60/epl-zone/internal/repository"
	"github.com/labstack/echo/v4"
)

type PlayerHandlers interface {
	GetAllPlayers(c echo.Context) error
	GetPlayersByTeam(c echo.Context) error
	GetPlayersByPostiton(c echo.Context) error
}

type PlayerHandler struct {
	playerRepo repository.PlayerRepository
}

func NewPlayerHandler(playerRepo repository.PlayerRepository) PlayerHandlers {
	return &PlayerHandler{playerRepo: playerRepo}
}

func (h *PlayerHandler) GetAllPlayers(c echo.Context) error {
	players, err := h.playerRepo.GetAllPlayers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	return c.JSON(http.StatusOK, players)
}

func (h *PlayerHandler) GetPlayersByTeam(c echo.Context) error {
	team := c.Param("team")

	players, err := h.playerRepo.GetPlayersByTeam(team)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	return c.JSON(http.StatusOK, players)
}

func (h *PlayerHandler) GetPlayersByPostiton(c echo.Context) error {
	position := c.Param("position")

	players, err := h.playerRepo.GetPlayersByPostiton(position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	return c.JSON(http.StatusOK, players)
}
