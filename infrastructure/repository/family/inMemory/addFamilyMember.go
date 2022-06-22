package familyInMemory

import (
	"github.com/renatospaka/emr/domain/entity/family"
)

func (f *FamilyRepositoryInMemory) AddFamilyMember(newMember family.FamilyMember, familyId string) (*family.Family, error) {
	member := newMember.Member
	err := member.IsValid()
	if err != nil {
		return &family.Family{}, err
	}

	toFamily, err := f.FindById(familyId)
	if err != nil {
		return &family.Family{}, err
	}

	// Verify if the member already exists
	// If s/he does, then raises an error
	repoMem := NewMemberRepositoryInMemory()
	_, err = repoMem.FindById(member.ID)
	if err == family.ErrMemberNotFound {
		err = repoMem.Add(member)
		if err != nil {
			return &family.Family{}, err
		}
	} else if err != nil {
		return &family.Family{}, err
	}

	// Verify if the member is already linked to the family
	// If s/he isn't, then link her/him
	// _, err = FindFamilyMemberById(member.ID, toFamily.ID)
	// if err != nil {
	// 	return &family.Family{}, err
	// }

	// Add the new member to the family
	toFamily.Members = append(toFamily.Members, &newMember)
	return toFamily, nil
}
