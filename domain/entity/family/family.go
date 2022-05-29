package family

import (
	"errors"
	"strings"

	"github.com/satori/go.uuid"
)

type Family struct {
	ID      uuid.UUID
	Surname string
	valid   bool
}

func NewFamily(surname string) *Family {
	return &Family{
		ID:      uuid.NewV4(),
		Surname: strings.TrimSpace(surname),
		valid: false,
	}
}

// Check whenever the family structure is intact and filled
func (f *Family) IsValid() error {
	f.valid = false
	if strings.TrimSpace(f.Surname) == "" {
		return errors.New("o nome de família está em branco ou ausente")
	}
	if strings.TrimSpace(f.ID.String()) == "" {
		return errors.New("o ID da família está em branco ou ausente")
	}

	f.valid = true
	return nil
}
