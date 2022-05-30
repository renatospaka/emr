package familyInMemory

import (
	"github.com/renatospaka/emr/domain/entity/family"
	uuid "github.com/satori/go.uuid"
)

// Find a family by its ID
// or returns an error
func (f *FamilyRepositoryInMemory) FindById(id uuid.UUID) (*family.Family, error) {
	if len(f.family) == 0 {
		return &family.Family{}, family.ErrFamilyNotFound
	}

	for x, fam := range f.family {
		if fam.ID == id {
			return &f.family[x], nil
		}
	}
	return &family.Family{}, family.ErrFamilyNotFound
}
