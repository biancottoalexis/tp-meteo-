package models

import "time"

type Day struct {
	Date    time.Time
	TempMax float64
	TempMin float64
	Code    int
}
