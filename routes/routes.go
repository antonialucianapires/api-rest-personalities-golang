package routes

import (
	"log"
	"net/http"

	"github.com/antonialucianapires/api-rest-crud-golang/controllers"
)

const prefix = "/api"

func HandleRequest() {
    http.HandleFunc(prefix+"/ping", controllers.Ping)
    http.HandleFunc(prefix+"/personalities", controllers.AllPersonalities)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
