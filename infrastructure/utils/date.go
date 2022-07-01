package utils

import (
	"time"
)

// Calculate elapsed time (in seconds) from a given date
// and returns formated in months and years
func AgeFromToday(dateFrom time.Time) (years int64, months int64) {
	if dateFrom.IsZero() {
		return 0, 0
	}

	today := time.Now()
	duration := today.Sub(dateFrom)
	years = int64(duration.Hours() / 24 / 365)
	months = int64(duration.Hours() / 24 / 30)
	return years, months
}
