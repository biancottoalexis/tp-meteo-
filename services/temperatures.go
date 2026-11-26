package services

import "tp-meteo-go/models"

func EchelleTemperatures(jours []models.Day) (float64, float64) {
	if len(jours) == 0 {
		return 0, 0
	}

	minimum := jours[0].TempMin
	maximum := jours[0].TempMax

	for _, jour := range jours {
		if jour.TempMin < minimum {
			minimum = jour.TempMin
		}

		if jour.TempMax > maximum {
			maximum = jour.TempMax
		}
	}

	return minimum, maximum
}

func PositionBarre(tempMin float64, minimum float64, maximum float64) float64 {
	if maximum <= minimum {
		return 0
	}

	position := (tempMin - minimum) / (maximum - minimum) * 100

	return position
}

func LargeurBarre(tempMin float64, tempMax float64, minimum float64, maximum float64) float64 {
	if maximum <= minimum {
		return 0
	}

	largeur := (tempMax - tempMin) / (maximum - minimum) * 100

	return largeur
}
