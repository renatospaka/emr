package utils

import (
	"math"
	"time"
)


// Calculate elapsed time (in seconds) from a given date
// and returns formated in months and years
func AgeFromToday(dateFrom time.Time) (years int, months int) {
	if dateFrom.IsZero() {
		return 0, 0
	}

	duration := time.Since(dateFrom)
	months = roundTime(duration.Seconds() / secondsInMonth)
	years = roundTime(duration.Seconds() / secondsInYear)

	return years, months
}

// Helps better calculate long periods of time from seconds 
// to months and years.
// There is rounding math to do
func roundTime(timeDuration float64) int {
	var result float64

	if timeDuration < 0 {
		result = math.Ceil(timeDuration - 0.5)
	} else {
		result = math.Floor(timeDuration + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)
	return int(i)
}