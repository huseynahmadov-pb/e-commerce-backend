package db

import (
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type PostgresDB struct {
	db *sqlx.DB	
}
func GetDB(secrets map[string]string) *PostgresDB {
	db := createDB(secrets)
	return &PostgresDB{
		db: db,
	}
}

func createDB(secrets map[string]string) *sqlx.DB {
	if err:= godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s",
		 dbHost, dbPort, secrets["username"], secrets["password"])

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}	

	return db
}