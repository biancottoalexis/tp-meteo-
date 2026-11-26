package controllers

import (
	"net/http"
	"strings"
	"unicode/utf8"

	"tp-meteo-go/api"
	"tp-meteo-go/services"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	afficherAccueil(w, "", "", 200)
}

func Ville(w http.ResponseWriter, r *http.Request) {
	nom := strings.TrimSpace(r.URL.Query().Get("nom"))
	taille := utf8.RuneCountInString(nom)

	if taille < 2 {
		afficherAccueil(w, nom, "Nom trop court", 400)
		return
	}
	if taille > 60 {
		afficherAccueil(w, nom, "Nom trop long", 400)
		return
	}

	communes, err := api.ChercherCommunes(nom)
	if err != nil {
		afficherErreur(w, "API indisponible", 503, err)
		return
	}
	if len(communes) == 0 {
		afficherErreur(w, "Aucune commune trouvée", 404, nil)
		return
	}

	afficherCommunes(w, communes)
}

func Meteo(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		afficherErreur(w, "Code manquant", 400, nil)
		return
	}

	commune, statut, err := api.ChercherCommune(code)
	if err != nil {
		afficherErreur(w, "API indisponible", 503, err)
		return
	}
	if statut == 404 {
		afficherErreur(w, "Commune introuvable", 404, nil)
		return
	}
	if len(commune.Centre.Coordinates) < 2 {
		afficherErreur(w, "Coordonnées indisponibles", 503, nil)
		return
	}

	longitude := commune.Centre.Coordinates[0]
	latitude := commune.Centre.Coordinates[1]

	meteo, err := api.ChercherMeteo(latitude, longitude)
	if err != nil {
		afficherErreur(w, "Météo indisponible", 503, err)
		return
	}

	jours, err := services.CreerJours(meteo)
	if err != nil {
		afficherErreur(w, "Prévisions indisponibles", 503, err)
		return
	}

	afficherMeteo(w, commune, jours)
}

func PageIntrouvable(w http.ResponseWriter, r *http.Request) {
	afficherErreur(w, "Page introuvable", 404, nil)
}
