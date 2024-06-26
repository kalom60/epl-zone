package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", s.healthHandler)

	e.GET("/allplayers", s.playerHandler.GetAllPlayers)
	e.GET("/player/team/:team", s.playerHandler.GetPlayersByTeam)
	e.GET("/player/position/:position", s.playerHandler.GetPlayersByPostiton)

	e.GET("/teams", s.teamHandler.GetTeams)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
