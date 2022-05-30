package familyRepository

import (
	"errors"
	
	uuid "github.com/satori/go.uuid"
	"github.com/renatospaka/emr/domain/entity/family"
)

type MemberRepositoryInMemory struct {
	member []family.Member
}

func NewMemberRepositoryInMemory() *MemberRepositoryInMemory {
	return &MemberRepositoryInMemory{
		member: []family.Member{},
	}
}

// Add a member so s/he may be late related to a family
// or returns an error
func (m *MemberRepositoryInMemory) Add(member *family.Member) error {
	err := member.IsValid()
	if err != nil {
		return err
	}

	m.member = append(m.member, *member)
	if len(m.member) > 0 {
		return nil
	}
	return errors.New("ocorreu um erro na inclusão de um novo membro da família")
}

// Find a member by his/her ID
// or returns an error
func (m *MemberRepositoryInMemory) FindById(id uuid.UUID) (*family.Member, error) {
	if len(m.member) == 0 {
		return &family.Member{}, family.ErrMemberNotFound
	}

	for x, mem := range m.member {
		if mem.ID == id {
			return &m.member[x], nil
		}
	}
	return &family.Member{}, family.ErrMemberNotFound
}

// Allow user to change or update attributes of a specific family member
// or returns an error
func (m *MemberRepositoryInMemory) Change(member *family.Member) (*family.Member, error) {
	panic("Not Implemented")
}

// Remove (completely) a member
// or returns an error
func (m *MemberRepositoryInMemory) Remove(id uuid.UUID) error {
	panic("Not Implemented")
}
