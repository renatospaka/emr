package utils

import (
	"time"
)

// Calculate elapsed time (in seconds) from a given date
// (today - day of birth)
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

// Calculate elapsed time (in seconds) from two given dates,
// (recent date - date)
// and returns formated in months and years
func AgeBetweenDates(recentDate time.Time, date time.Time) (years int64, months int64) {
	if recentDate.IsZero() || date.IsZero() {
		return 0, 0
	}

	duration := recentDate.Sub(date)
	years = int64(duration.Hours() / 24 / 365)
	months = int64(duration.Hours() / 24 / 30)
	return years, months
}
