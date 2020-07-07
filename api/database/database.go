package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Open - Function to open DB connection
func Open() *sql.DB {
	host := "127.0.0.1"
	user := "mattl"
	pw := "matt"
	database := "met_burn"
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pw, database)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open: %v", err))
	}
	err = db.Ping()
	if err == nil {
		fmt.Println("Ping Successful")
		return db
	}
	return db
}
