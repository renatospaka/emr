package familyInMemory

import (
	family "github.com/renatospaka/emr/family/domain/entity"
)

// Family stuff
type FamilyRepositoryInMemory struct {
	family []family.Family
}

func NewFamilyRepositoryInMemory() *FamilyRepositoryInMemory {
	return &FamilyRepositoryInMemory{
		family: []family.Family{},
	}
}

// Family members stuff
type FamilyMemberRepositoryInMemory struct {
	familyMember family.FamilyMember
}

func NewFamilyMemberRepositoryInMemory() *FamilyMemberRepositoryInMemory {
	return &FamilyMemberRepositoryInMemory{
		familyMember: family.FamilyMember{},
	}
}
