package familyInMemory

import (
	"github.com/renatospaka/emr/domain/entity/family"
)

// Find a member by his/her ID
// or returns an error
func (m *MemberRepositoryInMemory) FindById(id string) (*family.Member, error) {
	if len(m.member) == 0 {
		return &family.Member{}, family.ErrMemberNotFound
	}

	for x, memb := range m.member {
		if memb.ID == id {
			return &m.member[x], nil
		}
	}
	return &family.Member{}, family.ErrMemberNotFound
}
