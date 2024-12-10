package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonialucianapires/api-rest-crud-golang/models"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
