package familyRepository

import (
	uuid "github.com/satori/go.uuid"
	"github.com/renatospaka/emr/domain/entity/family"
)

type FamilyRepository interface {
	Create(family *family.Family) error 
	FindById(id uuid.UUID) (*family.Family, error)
	ChangeName(id uuid.UUID, newSurname string) (*family.Family, error)
}

type PersonRepository interface {
	Add(person *family.Person) error
	Change(person *family.Person) (*family.Person, error)
	Remove(id uuid.UUID) error
}
