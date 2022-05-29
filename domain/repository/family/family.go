package familyRepository

import (
	"github.com/renatospaka/emr/domain/entity/family"
	uuid "github.com/satori/go.uuid"
)

type FamilyRepository interface {
	Create(family *family.Family) error 
	FindById(id uuid.UUID) (*family.Family, error)
	ChangeName(id uuid.UUID, newSurname string) (*family.Family, error)
}
