package familyInMemory

import (
	"github.com/renatospaka/emr/domain/entity/family"
)

// Remove (completely) a member
// or returns an error
func (m *MemberRepositoryInMemory) Remove(id string) error {
	for x, memb := range m.member {
		if memb.ID == id {
			var err error
			m.member, err = removeMember(m.member, x)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return family.ErrMemberNotFound
}
