package routes

import (
	"log"
	"net/http"

	"github.com/antonialucianapires/api-rest-crud-golang/controllers"
	"github.com/antonialucianapires/api-rest-crud-golang/middleware"
	"github.com/gorilla/mux"
)

const prefix = "/api"

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
    r.HandleFunc(prefix+"/ping", controllers.Ping)
    r.HandleFunc(prefix+"/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc(prefix+"/personalities/{id}", controllers.GetByPersonalityById).Methods("GET")
	r.HandleFunc(prefix+"/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc(prefix+"/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc(prefix+"/personalities/{id}", controllers.UpdtadePersonality).Methods("PUT")
    log.Fatal(http.ListenAndServe(":8000", r))
}
