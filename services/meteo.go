package services

import (
	"fmt"
	"time"

	"tp-meteo-go/models"
)

func CreerJours(meteo models.Meteo) ([]models.Day, error) {
	var jours []models.Day

	nombre := len(meteo.Daily.Dates)

	if nombre == 0 {
		return nil, fmt.Errorf("aucune prévision")
	}

	if len(meteo.Daily.TempMax) != nombre {
		return nil, fmt.Errorf("températures max manquantes")
	}

	if len(meteo.Daily.TempMin) != nombre {
		return nil, fmt.Errorf("températures min manquantes")
	}

	if len(meteo.Daily.Codes) != nombre {
		return nil, fmt.Errorf("codes météo manquants")
	}

	for i := 0; i < nombre; i++ {
		var jour models.Day

		date, err := time.Parse("2006-01-02", meteo.Daily.Dates[i])

		if err != nil {
			return nil, err
		}

		jour.Date = date
		jour.TempMax = meteo.Daily.TempMax[i]
		jour.TempMin = meteo.Daily.TempMin[i]
		jour.Code = meteo.Daily.Codes[i]

		jours = append(jours, jour)
	}

	return jours, nil
}
