package main

import (
	"github.com/antonialucianapires/api-rest-crud-golang/database"
	"github.com/antonialucianapires/api-rest-crud-golang/models"
	"github.com/antonialucianapires/api-rest-crud-golang/routes"
)


func main() {
	models.Personalities = []models.Personality {
		{Id: 1, Name: "Ada Lovelace", History: "Pioneira da programação e autora do primeiro algoritmo."},
        {Id: 2, Name: "Alan Turing", History: "Fundador da ciência da computação e da criptoanálise moderna."},
	}
	database.ConnectDatabase()
	routes.HandleRequest()
}