package main

import (
	"log"
	"net/http"

	"tp-meteo-go/routes"
)

func main() {
	routes.InitRoutes()

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
