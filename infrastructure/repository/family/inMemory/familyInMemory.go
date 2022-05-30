package familyInMemory

import "github.com/renatospaka/emr/domain/entity/family"

type FamilyRepositoryInMemory struct {
	family []family.Family
}

func NewFamilyRepositoryInMemory() *FamilyRepositoryInMemory {
	return &FamilyRepositoryInMemory{
		family: []family.Family{},
	}
}
