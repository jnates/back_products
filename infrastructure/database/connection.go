package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	// registering database driver
	_ "github.com/lib/pq"
)

var (
	data *Data
	once sync.Once
)

//Data aa
type Data struct {
	DB *sql.DB
}

// New returns a new instance of Data with the database connection ready.
func New() *Data {
	once.Do(initDB)
	return data
}
func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Println("Cannot connect to database test")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Println("We are connected to the database test")
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	data = &Data{
		DB: db,
	}
}

func getConnection() (*sql.DB, error) {
	DbHost := os.Getenv("DB_HOST") //"127.0.0.1"
	DbDriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	return sql.Open(DbDriver, uri)
}
