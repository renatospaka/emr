package familyInMemory

// import (
// 	"errors"

// 	family "github.com/renatospaka/emr/family/domain/entity"
// )

// // Creates a new family structure
// // or returns an error
// func (f *FamilyRepositoryInMemory) Create(family *family.Family) error {
// 	err := family.IsValid()
// 	if err != nil {
// 		return err
// 	}

// 	f.family = append(f.family, *family)
// 	if len(f.family) > 0 {
// 		return nil
// 	}
// 	return errors.New("ocorreu um erro na gravação da nova família")
// }
