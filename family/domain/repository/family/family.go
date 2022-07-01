package familyRepository

import (
	family "github.com/renatospaka/emr/family/domain/entity"
)

type FamilyRepository interface {
	Create(family *family.Family) error
	FindById(id string) (*family.Family, error)
	SetFamilyName(id string, surname string) (*family.Family, error)
	FindFamilyMemberById(idMember string, idFamily string) (*family.Member, error)
	AddFamilyMember(newMember *family.FamilyMember, toFamily string) (*family.Family, error)
}

type MemberRepository interface {
	Add(member *family.Member) error
	FindById(id string) (*family.Member, error)
	Change(member *family.Member) (*family.Member, error)
	Remove(id string) error
}
