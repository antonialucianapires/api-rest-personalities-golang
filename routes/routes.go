package routes

import (
	"log"
	"net/http"

	"github.com/antonialucianapires/api-rest-crud-golang/controllers"
)

func HandleRequest() {
	http.HandleFunc("/ping", controllers.Ping)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
