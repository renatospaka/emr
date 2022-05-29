package familyRepository

import (
	"errors"

	"github.com/renatospaka/emr/domain/entity/family"
	uuid "github.com/satori/go.uuid"
)


type FamilyRepositoryInMemory struct {
	family []family.Family
}

func NewFamilyRepositoryInMemory() *FamilyRepositoryInMemory {
	return &FamilyRepositoryInMemory{
		family: []family.Family{},
	}
}

// Creates a new family structure
// or returns an error
func (f *FamilyRepositoryInMemory) Create(family *family.Family) error {
	err := family.IsValid()
	if err != nil {
		return err
	}
	
	f.family = append(f.family, *family)
	if len(f.family) > 0 {
		return nil
	}
	return errors.New("ocorreu um erro na gravação da nova família")
}

// Find a family by its ID
// or returns an error
func (f *FamilyRepositoryInMemory) FindById(id uuid.UUID) (*family.Family, error) {
	for x, fam := range f.family {
		if fam.ID == id {
			return &f.family[x], nil
		}
	}
	return &family.Family{}, errors.New("família não encontrada")
}


// Allow user to change or update the surname of a specific family
// or returns an error
func (f *FamilyRepositoryInMemory) ChangeName(id uuid.UUID, newSurname string) (*family.Family, error) {
	fam, err := f.FindById(id)
	if err != nil {
		return &family.Family{}, errors.New("família não encontrada")
	}

	fam.Surname = newSurname
	err = fam.IsValid()
	if err != nil {
		return &family.Family{}, err
	}
	return fam, nil
}
