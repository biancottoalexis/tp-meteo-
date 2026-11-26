package models

type Commune struct {
    Nom string `json:"nom"`
    Code string `json:"code"`
    Centre Centre `json:"centre"`
}

type Centre struct {
    Coordinates []float64 `json:"coordinates"`
}