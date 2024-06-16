package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	DB() *sqlx.DB
	Health() map[string]string
	Close() error
}

type service struct {
	db *sqlx.DB
}

var (
	database  = os.Getenv("DB_DATABASE")
	password  = os.Getenv("DB_PASSWORD")
	username  = os.Getenv("DB_USERNAME")
	port      = os.Getenv("DB_PORT")
	host      = os.Getenv("DB_HOST")
	dbInstace *service
)

func New() Service {
	if dbInstace != nil {
		return dbInstace
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	dbInstace = &service{
		db: db,
	}

	return dbInstace
}

func (s *service) DB() *sqlx.DB {
	return s.db
}

func (s *service) Health() map[string]string {
	stats := s.db.DB.Stats() // Use the underlying *sql.DB

	healthStats := make(map[string]string)
	healthStats["total_conns"] = fmt.Sprintf("%d", stats.OpenConnections)
	healthStats["idle_conns"] = fmt.Sprintf("%d", stats.Idle)
	healthStats["max_conns"] = fmt.Sprintf("%d", stats.MaxOpenConnections)

	return healthStats
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
