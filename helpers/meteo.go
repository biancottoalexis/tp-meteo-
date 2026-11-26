package helpers

func WeatherLabel(code int) string {
	if code == 0 {
		return "Ciel dégagé"
	}

	if code == 1 {
		return "Peu nuageux"
	}

	if code == 2 {
		return "Partiellement nuageux"
	}

	if code == 3 {
		return "Couvert"
	}

	if code == 45 || code == 48 {
		return "Brouillard"
	}

	if code == 51 || code == 53 || code == 55 {
		return "Bruine"
	}

	if code == 56 || code == 57 {
		return "Bruine verglaçante"
	}

	if code == 61 || code == 63 || code == 65 {
		return "Pluie"
	}

	if code == 66 || code == 67 {
		return "Pluie verglaçante"
	}

	if code == 71 || code == 73 || code == 75 || code == 77 {
		return "Neige"
	}

	if code == 80 || code == 81 || code == 82 {
		return "Averses"
	}

	if code == 85 || code == 86 {
		return "Averses de neige"
	}

	if code == 95 {
		return "Orage"
	}

	if code == 96 || code == 99 {
		return "Orage avec grêle"
	}

	return "Météo inconnue"
}

func WeatherIcon(code int) string {
	if code == 0 {
		return "☀️"
	}

	if code == 1 || code == 2 {
		return "🌤️"
	}

	if code == 3 {
		return "☁️"
	}

	if code == 45 || code == 48 {
		return "🌫️"
	}

	if code == 51 || code == 53 || code == 55 {
		return "🌦️"
	}

	if code == 56 || code == 57 {
		return "🌧️"
	}

	if code == 61 || code == 63 || code == 65 {
		return "🌧️"
	}

	if code == 66 || code == 67 {
		return "🌧️"
	}

	if code == 71 || code == 73 || code == 75 || code == 77 {
		return "🌨️"
	}

	if code == 80 || code == 81 || code == 82 {
		return "🌧️"
	}

	if code == 85 || code == 86 {
		return "🌨️"
	}

	if code == 95 || code == 96 || code == 99 {
		return "⛈️"
	}

	return "❓"
}
