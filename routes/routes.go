package routes

import (
	"net/http"

	"tp-meteo-go/controllers"
)

func InitRoutes() {
	controllers.InitTemplates()

	http.HandleFunc("GET /{$}", controllers.Accueil)
	http.HandleFunc("GET /ville", controllers.Ville)
	http.HandleFunc("GET /meteo", controllers.Meteo)

	fichiers := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fichiers))

	http.HandleFunc("/", controllers.PageIntrouvable)
}
