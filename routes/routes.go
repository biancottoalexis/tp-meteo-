package routes

import (
    "fmt"
    "net/http"
)

func InitRoutes() {
    http.HandleFunc("GET /{$}", accueil)
    http.HandleFunc("GET /ville", ville)
    http.HandleFunc("GET /meteo", meteo)

    http.Handle("/static/",
        http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))

    http.HandleFunc("/", pageIntrouvable)
}

func accueil(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Accueil météo")
}

func ville(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Recherche de communes")
}

func meteo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Prévisions météo")
}

func pageIntrouvable(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Cette page n'existe pas.", http.StatusNotFound)
}