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

type MemberRepository interface {
	Add(person *family.Member) error
	Change(person *family.Member) (*family.Member, error)
	Remove(id uuid.UUID) error
}
