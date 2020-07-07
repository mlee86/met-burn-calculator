package mets

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/mlee86/met-burn-calculator/database"
)

var db *sql.DB

type Exercise struct {
	MetCode     int    `json:"met_code" db:"met_code"`
	Catagory    string `json:"category" db:"category"`
	Description string `json:"description" db:"description"`
}

// Init - initalize database
func Init() {
	db = database.Open()
}

// GetAllCategories - method to return all catagories in the table
func GetAllCategories() ([]string, error) {
	rows, err := db.Query(`SELECT DISTINCT category FROM mets;`)
	var catagories []string
	for rows.Next() {
		var catagory string

		if err := rows.Scan(&catagory); err != nil {
			fmt.Println(err)
			fmt.Println("Error Retrieving Catagories")
		}
		catagories = append(catagories, catagory)
	}
	return catagories, err
}

func GetExercisesByCatagory(category string) ([]Exercise, error) {
	query := fmt.Sprintf("SELECT met_code, category, description FROM mets where category = '%v';", category)

	rows, err := db.Query(query)
	var exercises []Exercise
	for rows.Next() {
		var (
			code string
			cat  string
			desc string
		)

		if err := rows.Scan(&code, &cat, &desc); err != nil {
			fmt.Println(err)
			fmt.Println("Error Retrieving Catagories")
		}
		codeInt, _ := strconv.Atoi(code)
		exercises = append(exercises, Exercise{int(codeInt), cat, desc})
	}
	return exercises, err
}

func GetMetBurn(code int) (int, error) {
	query := fmt.Sprintf("SELECT met_units FROM mets where met_code = %v;", code)

	rows, err := db.Query(query)
	var metBurn int
	for rows.Next() {
		var burn string

		if err := rows.Scan(&burn); err != nil {
			fmt.Println(err)
			fmt.Println("Error Retrieving Catagories")
		}
		metBurn, _ = strconv.Atoi(burn)

	}
	return metBurn, err
}
