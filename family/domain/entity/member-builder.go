package family

import (
	"time"
)

// Builder structure for new members
type memberActions func(*Member)
type MemberBuilder struct {
	actions []memberActions
}

// Initialize the new members builder
func NewMemberBuilder() *MemberBuilder {
	return &MemberBuilder{
		actions: []memberActions{},
	}
}

// Execute all actions, create the Member
// and return it to builder function
func (mb *MemberBuilder) Build() *Member {
	member := newMember()
	for _, action := range mb.actions {
		action(member)
	}
	member.lastChanged = time.Now().UnixNano()
	member.validate()

	// if !member.valid {
	// 	return &Member{}
	// }
	return member
}

// Set the full name during the construction of this member
// name + middle name + last name
func (mb *MemberBuilder) WithFullName(name string, middleName string, lastName string) *MemberBuilder {
	mb.actions = append(mb.actions, func(m *Member) {
		m.SetFullName(name, middleName, lastName)
	})
	return mb
}

// Set the day of birth during the construction of this member
func (mb *MemberBuilder) WithBirthDate(dob time.Time) *MemberBuilder {
	mb.actions = append(mb.actions, func(m *Member) {
		m.SetBirthDate(dob)
	})
	return mb
}

// Set the gender during the construction of this member
func (mb *MemberBuilder) WithGender(gender string) *MemberBuilder {
	mb.actions = append(mb.actions, func(m *Member) {
		m.SetGender(gender)
	})
	return mb
}

// Set the nickname during the construction of this member
func (mb *MemberBuilder) WithNickname(nickname string) *MemberBuilder {
	mb.actions = append(mb.actions, func(m *Member) {
		m.SetNickname(nickname)
	})
	return mb
}
