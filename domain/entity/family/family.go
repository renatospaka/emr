package family

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Family struct {
	ID      uuid.UUID
	Surname string
	Valid   bool
}

func NewFamily(surname string) *Family {
	newFamily := &Family{
		ID:      uuid.NewV4(),
		Surname: strings.TrimSpace(surname),
		Valid:   false,
	}

	newFamily.IsValid()
	return newFamily
}

// Check whenever the family structure is intact and filled
func (f *Family) IsValid() error {
	f.Valid = false
	if strings.TrimSpace(f.Surname) == "" {
		return ErrMissingFamilySurname
	}
	if strings.TrimSpace(f.ID.String()) == "" {
		return ErrMissingFamilyID
	}

	f.Valid = true
	return nil
}
