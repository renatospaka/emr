package family

import (
	"time"
	"strings"

	"github.com/satori/go.uuid"
)

type Person struct {
	ID         uuid.UUID
	Name       string
	LastName   string
	MiddleName string
	DOB        time.Time
	Gender     string
	valid      bool
}

func NewPerson(name string, lastName string, middleName string, gender string, dob time.Time) *Person {
	return &Person{
		ID:         uuid.NewV1(),
		Name:       strings.TrimSpace(name),
		LastName:   strings.TrimSpace(lastName),
		MiddleName: strings.TrimSpace(middleName),
		DOB:        dob,
		Gender:     strings.TrimSpace(gender),
		valid:      false,
	}
}
