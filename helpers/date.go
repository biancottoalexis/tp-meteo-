package helpers

import "time"

func FrDate(date time.Time) string {
    jours := []string{
        "dimanche",
        "lundi",
        "mardi",
        "mercredi",
        "jeudi",
        "vendredi",
        "samedi",
    }

    nom := jours[date.Weekday()]

    return nom + " " + date.Format("02/01/2006")
}