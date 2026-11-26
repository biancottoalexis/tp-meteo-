package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"tp-meteo-go/helpers"
	"tp-meteo-go/models"
	"tp-meteo-go/services"
)

var accueilHTML *template.Template
var communesHTML *template.Template
var meteoHTML *template.Template
var erreurHTML *template.Template

type Formulaire struct {
	Nom    string
	Erreur string
}

type PageMeteo struct {
	Commune models.Commune
	Jours   []models.Day
	Carte   string
	Barres  []services.JourBarre
}

type PageErreur struct {
	Code    int
	Message string
}

func InitTemplates() {
	accueilHTML = template.Must(template.ParseFiles("templates/layout.html", "templates/accueil.html"))
	communesHTML = template.Must(template.ParseFiles("templates/communes.html"))
	erreurHTML = template.Must(template.ParseFiles("templates/erreur.html"))

	fonctions := template.FuncMap{
		"frDate":       helpers.FrDate,
		"weatherLabel": helpers.WeatherLabel,
		"weatherIcon":  helpers.WeatherIcon,
		"temp":         helpers.Temp,
	}

	meteoHTML = template.Must(template.New("meteo.html").Funcs(fonctions).ParseFiles("templates/meteo.html"))
}

func afficherErreur(w http.ResponseWriter, message string, code int, err error) {
	if err != nil {
		fmt.Println(err)
	}

	infos := PageErreur{Code: code, Message: message}
	afficherPage(w, erreurHTML, "erreur.html", infos, code)
}

func afficherPage(w http.ResponseWriter, fichier *template.Template, nom string, infos any, code int) {
	var page bytes.Buffer

	err := fichier.ExecuteTemplate(&page, nom, infos)

	if err != nil {
		if fichier == erreurHTML {
			fmt.Println(err)
			http.Error(w, "Erreur du serveur", 500)
			return
		}

		afficherErreur(w, "Erreur du serveur", 500, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	page.WriteTo(w)
}

func afficherAccueil(w http.ResponseWriter, nom string, erreur string, code int) {
	infos := Formulaire{Nom: nom, Erreur: erreur}
	afficherPage(w, accueilHTML, "layout", infos, code)
}

func afficherCommunes(w http.ResponseWriter, communes []models.Commune) {
	afficherPage(w, communesHTML, "communes.html", communes, 200)
}

func afficherMeteo(w http.ResponseWriter, commune models.Commune, jours []models.Day) {
	infos := PageMeteo{
		Commune: commune,
		Jours:   jours,
		Barres:  services.CreerBarres(jours),
	}

	if len(commune.Centre.Coordinates) >= 2 {
		longitude := commune.Centre.Coordinates[0]
		latitude := commune.Centre.Coordinates[1]
		infos.Carte = helpers.CarteURL(latitude, longitude)
	}

	afficherPage(w, meteoHTML, "meteo.html", infos, 200)
}
