package helpers

import (
	"fmt"
	"net/url"
)

func CarteURL(latitude float64, longitude float64) string {
	gauche := longitude - 0.05
	droite := longitude + 0.05
	bas := latitude - 0.05
	haut := latitude + 0.05

	bbox := fmt.Sprintf(
		"%f,%f,%f,%f",
		gauche, bas, droite, haut,
	)

	marker := fmt.Sprintf(
		"%f,%f",
		latitude, longitude,
	)

	parametres := url.Values{}

	parametres.Set("bbox", bbox)
	parametres.Set("layer", "mapnik")
	parametres.Set("marker", marker)

	return "https://www.openstreetmap.org/export/embed.html?" + parametres.Encode()
}
