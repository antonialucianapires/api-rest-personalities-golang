package main

import (
	"github.com/antonialucianapires/api-rest-crud-golang/models"
	"github.com/antonialucianapires/api-rest-crud-golang/routes"
)


func main() {
	models.Personalities = []models.Personality {
		{Name: "Ada Lovelace", History: "Pioneira da programação e autora do primeiro algoritmo."},
        {Name: "Alan Turing", History: "Fundador da ciência da computação e da criptoanálise moderna."},
	}
	routes.HandleRequest()
}