package familyInMemory

import (
	"github.com/renatospaka/emr/domain/entity/family"
	uuid "github.com/satori/go.uuid"
)

// Allow user to change or update the surname of a specific family
// or returns an error
func (f *FamilyRepositoryInMemory) SetFamilyName(id uuid.UUID, newSurname string) (*family.Family, error) {
	fam, err := f.FindById(id)
	if err != nil {
		return &family.Family{}, family.ErrFamilyNotFound
	}

	fam.Surname = newSurname
	err = fam.IsValid()
	if err != nil {
		return &family.Family{}, err
	}
	return fam, nil
}
