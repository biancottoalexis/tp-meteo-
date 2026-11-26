package services

import "tp-meteo-go/models"

type JourBarre struct {
	Jour     models.Day
	Position float64
	Largeur  float64
}

func CreerBarres(jours []models.Day) []JourBarre {
	minimum, maximum := EchelleTemperatures(jours)

	var resultat []JourBarre

	for _, jour := range jours {
		var barre JourBarre

		barre.Jour = jour

		barre.Position = PositionBarre(
			jour.TempMin,
			minimum,
			maximum,
		)

		barre.Largeur = LargeurBarre(
			jour.TempMin,
			jour.TempMax,
			minimum,
			maximum,
		)

		resultat = append(resultat, barre)
	}

	return resultat
}
