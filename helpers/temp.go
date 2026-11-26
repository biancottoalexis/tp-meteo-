package helpers

import "fmt"

func Temp(temperature float64) string {
    return fmt.Sprintf("%.1f °C", temperature)
}