package family

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Family struct {
	ID      uuid.UUID `json:"family_id"`
	Surname string    `json:"surname"`
	Valid   bool      `json:"is_valid"`
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

// Check whenever the family structure is intact 
// and filled accordingly to the model rules
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
