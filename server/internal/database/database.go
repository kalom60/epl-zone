package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	DB() *sqlx.DB
	Health() map[string]string
	Close() error
	ConvertCSVToDB() error
	FlushPlayerTable() error
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

func New() (Service, error) {
	if dbInstace != nil {
		return dbInstace, nil
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, err
	}

	dbInstace = &service{
		db: db,
	}

	if err := dbInstace.CreatePlayerTable(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create database players table: %w", err)
	}

	if err := dbInstace.CreateTeamTable(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create database teams table: %w", err)
	}

	return dbInstace, nil
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

func (s *service) ConvertCSVToDB() error {
	relativeFilePath := "data/stats.csv"
	absoluteFilePath, err := filepath.Abs(relativeFilePath)
	if err != nil {
		return fmt.Errorf("could not determine absolute file path: %w", err)
	}

	_, err = os.Stat(absoluteFilePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("CSV file %s does not exist", absoluteFilePath)
	}

	// Escape single quotes in the file path
	escapedFilePath := strings.ReplaceAll(absoluteFilePath, "'", "''")

	query := fmt.Sprintf("COPY players (Player, Nation, Pos, Age, MP, Starts, Min, Gls, Ast, PK, CrdY, CrdR, xG, xAG, Team) FROM '%s' DELIMITER ',' CSV HEADER", escapedFilePath)

	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreatePlayerTable() error {
	createTableQuery := `
	    CREATE TABLE IF NOT EXISTS players (
        	id SERIAL PRIMARY KEY,
        	player VARCHAR(255),
        	nation VARCHAR(255),
        	pos VARCHAR(255),
        	age FLOAT,
        	mp INT,
        	starts INT,
        	min FLOAT,
        	gls FLOAT,
        	ast FLOAT,
        	pk FLOAT,
        	crdy FLOAT,
        	crdr FLOAT,
        	xg FLOAT,
        	xag FLOAT,
        	team VARCHAR(255)
    	);
    `
	_, err := s.db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create players table: %w", err)
	}
	return nil
}

func (s *service) CreateTeamTable() error {
	createTableQuery := `
	    CREATE TABLE IF NOT EXISTS teams(
        	id SERIAL PRIMARY KEY,
        	team VARCHAR(255),
        	logo VARCHAR(255)
    	);
    `
	_, err := s.db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create teams table: %w", err)
	}
	return nil
}

func (s *service) FlushPlayerTable() error {
	flushTableQuery := `
        DELETE FROM players;
        DELETE FROM teams;
    `
	_, err := s.db.Exec(flushTableQuery)
	if err != nil {
		return err
	}
	return nil
}
