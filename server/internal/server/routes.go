package server

import (
	"net/http"

	"github.com/kalom60/epl-zone/internal/scrape"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", s.scrapeHandler)
    e.GET("/health", s.healthHandler)

    e.GET("/allplayers", s.playerHandler.GetAllPlayers)

    return e
}

func (s *Server) scrapeHandler(c echo.Context) error {
    scrape.Scrapper()
    return c.JSON(http.StatusOK, "Done")
}

func (s *Server) healthHandler(c echo.Context) error {
    return c.JSON(http.StatusOK, s.db.Health())
}
