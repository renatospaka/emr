package familyInMemory

// import (
// 	family "github.com/renatospaka/emr/family/domain/entity"
// )

// func (f *FamilyRepositoryInMemory) AddFamilyMember(newMember family.FamilyMember, familyId string) (*family.Family, error) {
// 	member := newMember.Member
// 	err := member.IsValid()
// 	if err != nil {
// 		return &family.Family{}, err
// 	}

// 	toFamily, err := f.FindById(familyId)
// 	if err != nil {
// 		return &family.Family{}, err
// 	}
