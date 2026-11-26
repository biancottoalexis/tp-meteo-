package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"tp-meteo-go/models"
)

func ChercherMeteo(latitude float64, longitude float64) (models.Meteo, error) {
	var meteo models.Meteo

	parametres := url.Values{}

	parametres.Set("latitude", strconv.FormatFloat(latitude, 'f', -1, 64))
	parametres.Set("longitude", strconv.FormatFloat(longitude, 'f', -1, 64))
	parametres.Set("daily", "temperature_2m_max,temperature_2m_min,weather_code")
	parametres.Set("timezone", "Europe/Paris")
	parametres.Set("forecast_days", "7")

	adresse := "https://api.open-meteo.com/v1/forecast?" + parametres.Encode()

	requete, err := http.NewRequest("GET", adresse, nil)

	if err != nil {
		return meteo, err
	}

	requete.Header.Set("User-Agent", "tp-meteo-go")

	reponse, err := client.Do(requete)

	if err != nil {
		return meteo, err
	}

	defer reponse.Body.Close()

	if reponse.StatusCode != 200 {
		return meteo, fmt.Errorf("erreur météo : %d", reponse.StatusCode)
	}

	err = json.NewDecoder(reponse.Body).Decode(&meteo)

	if err != nil {
		return meteo, err
	}

	return meteo, nil
}
