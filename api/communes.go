package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"tp-meteo-go/models"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func appelerAPI(adresse string) (*http.Response, error) {
	requete, err := http.NewRequest("GET", adresse, nil)
	if err != nil {
		return nil, err
	}

	requete.Header.Set("User-Agent", "tp-meteo-go")
	return client.Do(requete)
}

func ChercherCommunes(nom string) ([]models.Commune, error) {
	adresse := "https://geo.api.gouv.fr/communes?nom=" + url.QueryEscape(nom)

	reponse, err := appelerAPI(adresse)
	if err != nil {
		return nil, err
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != 200 {
		return nil, fmt.Errorf("API indisponible")
	}

	var communes []models.Commune
	err = json.NewDecoder(reponse.Body).Decode(&communes)
	if err != nil {
		return nil, err
	}

	return communes, nil
}

func ChercherCommune(code string) (models.Commune, int, error) {
	var commune models.Commune
	adresse := "https://geo.api.gouv.fr/communes/" + url.PathEscape(code) + "?fields=nom,code,centre"

	reponse, err := appelerAPI(adresse)
	if err != nil {
		return commune, 503, err
	}
	defer reponse.Body.Close()

	if reponse.StatusCode == 404 {
		return commune, 404, nil
	}
	if reponse.StatusCode != 200 {
		return commune, 503, fmt.Errorf("API indisponible")
	}

	err = json.NewDecoder(reponse.Body).Decode(&commune)
	if err != nil {
		return commune, 503, err
	}

	return commune, 200, nil
}
