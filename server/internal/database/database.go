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
		return nil, fmt.Errorf("failed to create database tables: %w", err)
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
	filePath := "data/stats.csv"

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("CSV file %s does not exist", filePath)
	}

	_, err = s.db.Exec(fmt.Sprintf("COPY players FROM %s DELIMITER ',' CSV HEADER", filePath))
	if err != nil {
        fmt.Println("8", err)
		return err
	}

	return nil
}

func (s *service) CreatePlayerTable() error {
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS players (
            player_name VARCHAR(255) NOT NULL,
            nation VARCHAR(255),
            position VARCHAR(255),
            age FLOAT,
            matches_played INT,
            starts INT,
            minutes_played FLOAT,
            goals FLOAT,
            assists FLOAT,
            penalities_scored FLOAT,
            yellow_cards FLOAT,
            red_cards FLOAT,
            expected_goals FLOAT,
            expected_assists FLOAT,
            team_name VARCHAR(255),
            PRIMARY KEY (player_name)
        );
    `
	_, err := s.db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create players table: %w", err)
	}
	return nil
}

func (s *service) FlushPlayerTable() error {
	flushTableQuery := `
        DELETE FROM players;
    `
	_, err := s.db.Exec(flushTableQuery)
	if err != nil {
		return err
	}
	return nil
}
