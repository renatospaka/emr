package familyRepository

import (
	"errors"
	// "fmt"
	// "log"

	"github.com/renatospaka/emr/domain/entity/family"
	uuid "github.com/satori/go.uuid"
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

	for x, memb := range m.member {
		if memb.ID == id {
			return &m.member[x], nil
		}
	}
	return &family.Member{}, family.ErrMemberNotFound
}

// Allow user to change or update attributes of a specific family member
// or returns an error
func (m *MemberRepositoryInMemory) Change(member *family.Member) (*family.Member, error) {
	err := member.IsValid()
	if err != nil {
		return member, err
	}

	id := member.ID
	current, err := m.FindById(id)
	if err != nil {
		return member, err
	}

	// Check if there is any change to make.
	// If there isn't, no action would be taken
	if current.Name == member.Name &&
			current.MiddleName == member.MiddleName &&
			current.LastName == member.LastName &&
			current.DOB == member.DOB &&
			current.Gender == member.Gender {
		return member, family.ErrNoChangesNeeded
	}

	// apply the changes
	current.Name = member.Name
	current.MiddleName = member.MiddleName
	current.LastName = member.LastName
	current.DOB = member.DOB
	current.Gender = member.Gender
	current.IsValid()

	return current, nil
}

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
