package familyRepository

import (
	"github.com/renatospaka/emr/domain/entity/family"
	uuid "github.com/satori/go.uuid"
)

type FamilyRepository interface {
	Create(family *family.Family) error
	FindById(id uuid.UUID) (*family.Family, error)
	SetFamilyName(id uuid.UUID, newSurname string) (*family.Family, error)
}

type MemberRepository interface {
	Add(member *family.Member) error
	FindById(id uuid.UUID) (*family.Member, error)
	Change(member *family.Member) (*family.Member, error)
	Remove(id uuid.UUID) error
}
