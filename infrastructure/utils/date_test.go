package utils_test

import (
	"testing"
	"time"

	"github.com/renatospaka/emr/infrastructure/utils"
	"github.com/stretchr/testify/require"
)

func TestDate_AgeFromToday(t *testing.T) {
	today := time.Now()
	dayOfBirth := time.Date(today.Year()-46, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths := utils.AgeFromToday(dayOfBirth)

	require.EqualValues(t, int64(46), ageInYears)
	require.EqualValues(t, int64(560), ageInMonths)
}

func TestDate_AgeFromToday_Zero(t *testing.T) {
	dayOfBirth := time.Time{}
	ageInYears, ageInMonths := utils.AgeFromToday(dayOfBirth)

	require.EqualValues(t, int64(0), ageInYears)
	require.EqualValues(t, int64(0), ageInMonths)
}

func TestDate_AgeBetweenDates(t *testing.T) {
	today := time.Now()
	onePastMonth := time.Date(today.Year(), today.Month()-1, today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths := utils.AgeBetweenDates(today, onePastMonth )

	require.EqualValues(t, int64(0), ageInYears)
	require.EqualValues(t, int64(1), ageInMonths)
}
