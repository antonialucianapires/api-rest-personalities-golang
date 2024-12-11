package routes

import (
	"log"
	"net/http"

	"github.com/antonialucianapires/api-rest-crud-golang/controllers"
	"github.com/gorilla/mux"
)

const prefix = "/api"

func HandleRequest() {
	r := mux.NewRouter()
    r.HandleFunc(prefix+"/ping", controllers.Ping)
    r.HandleFunc(prefix+"/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc(prefix+"/personalities/{id}", controllers.GetByPersonalityById).Methods("GET")
	r.HandleFunc(prefix+"/personalities", controllers.CreatePersonality).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", r))
}
