package models

type Meteo struct {
	Daily Daily `json:"daily"`
}

type Daily struct {
	Dates   []string  `json:"time"`
	TempMax []float64 `json:"temperature_2m_max"`
	TempMin []float64 `json:"temperature_2m_min"`
	Codes   []int     `json:"weather_code"`
}
