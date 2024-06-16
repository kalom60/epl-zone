package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kalom60/epl-zone/internal/database"
)

type Server struct {
	port int

    db database.Service
}

func NewServer() *http.Server {
    port, _ := strconv.Atoi(os.Getenv("PORT"))

    dbService := database.New()

    s := &Server{
        port: port,
        db: dbService,
    }

    server := &http.Server{
        Addr: fmt.Sprintf(":%d", s.port),
        Handler: s.RegisterRoutes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    return server
}
