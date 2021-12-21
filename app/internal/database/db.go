package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Store struct {
	DB *gorm.DB
}

// Ping - pings the database to check if there are any issues
func (s *Store) Ping() error {
	if err := s.DB.DB().Ping(); err != nil {
		return err
	}
	return nil
}

// SetupDB returns a pointer to a database connection
// calling functions need to ensure they defer the closing
// of this connection
func SetupDB() (Store, error) {
	var err error
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbName, dbPassword, dbSSL)

	database, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Print("Failed to open connection")
		return Store{}, err
	}

	return Store{
		DB: database,
	}, nil
}

func (s *Store) Close() error {
	err := s.DB.Close()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
