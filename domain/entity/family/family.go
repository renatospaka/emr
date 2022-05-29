package family

import (
	"strings"

	"github.com/satori/go.uuid"
)

type Family struct {
	ID      uuid.UUID
	Surname string
}

func NewFamily(surname string) *Family {
	return &Family{
		ID:      uuid.NewV4(),
		Surname: strings.TrimSpace(surname),
	}
}
