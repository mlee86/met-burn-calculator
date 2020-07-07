package mets

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// GetAllCategoriesHandler - handler to retrieve list of all catagories
func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	catagories, err := GetAllCategories()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsn, err := json.Marshal(catagories)

	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

func GetExercisesHandler(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/exercises/")
	exercises, err := GetExercisesByCatagory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsn, err := json.Marshal(exercises)

	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

func GetMetBurnHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/calculate/")
	code, err := strconv.Atoi(id)
	metBurn, err := GetMetBurn(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsn, err := json.Marshal(metBurn)

	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}
