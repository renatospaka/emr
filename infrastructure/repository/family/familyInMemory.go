package familyRepository

import (
	"errors"

	"github.com/renatospaka/emr/domain/entity/family"
)


type FamilyRepositoryInMemory struct {
	family []family.Family
}

func (f *FamilyRepositoryInMemory) Create(family *family.Family) error {
	f.family = append(f.family, *family)
	if len(f.family) > 0 {
		return nil
	}
	return errors.New("ocorreu um erro na gravação da nova família")
}