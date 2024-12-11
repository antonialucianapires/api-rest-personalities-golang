package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonialucianapires/api-rest-crud-golang/database"
	"github.com/antonialucianapires/api-rest-crud-golang/models"
	"github.com/gorilla/mux"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func GetByPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}