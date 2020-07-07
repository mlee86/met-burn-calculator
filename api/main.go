package main

import (
	"log"
	"net/http"

	// "database/sql"
	"fmt"

	"github.com/gorilla/mux"
	// "github.com/mlee86/met-burn-calculator/database"
	"github.com/mlee86/met-burn-calculator/mets"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	landing := "Welcome"
	w.Write([]byte(landing))
}

func main() {
	fmt.Println("test")
	r := mux.NewRouter()
	mets.Init()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/catagories", mets.GetAllCategoriesHandler).Methods("GET")
	r.HandleFunc("/exercises/{category}", mets.GetExercisesHandler).Methods("GET")
	r.HandleFunc("/calculate/{code}", mets.GetMetBurnHandler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
