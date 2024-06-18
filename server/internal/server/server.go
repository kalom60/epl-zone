package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kalom60/epl-zone/internal/cron"
	"github.com/kalom60/epl-zone/internal/database"
	"github.com/kalom60/epl-zone/internal/handlers"
	"github.com/kalom60/epl-zone/internal/repository"
	"github.com/kalom60/epl-zone/internal/scrape"
)

type Server struct {
	port          int
	db            database.Service
	playerRepo    repository.PlayerRepository
	playerHandler handlers.PlayerHandlers
	cronJob       cron.Jobber
}

func NewServer() (*Server, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}

	dbService, err := database.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	playerRepo := repository.NewPgPlayerRepository(dbService.DB())
	playerHandler := handlers.NewPlayerHandler(playerRepo)

	cronJob := cron.NewCronJob(dbService)
	cronJob.Start()

	err = scrape.Scrapper()
	if err != nil {
		return nil, fmt.Errorf("failed to scrape data and save to CSV: %w", err)
	}

    err = dbService.FlushPlayerTable()
    if err != nil {
        return nil, fmt.Errorf("failed to delete all players record from players table: %w", err)
    }

	err = dbService.ConvertCSVToDB()
	if err != nil {
		return nil, fmt.Errorf("failed to convert CSV to DB: %w", err)
	}

	s := &Server{
		port:          port,
		db:            dbService,
		playerRepo:    playerRepo,
		playerHandler: playerHandler,
		cronJob:       cronJob,
	}

	return s, nil
}

func (s *Server) NewHTTPServer() *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
