package familyInMemory

import (
	"github.com/renatospaka/emr/domain/entity/family"
	uuid "github.com/satori/go.uuid"
)

// Remove (completely) a member
// or returns an error
func (m *MemberRepositoryInMemory) Remove(id uuid.UUID) error {
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
