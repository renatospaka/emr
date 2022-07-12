package family

import (
	"log"
	"strings"
	"time"
)

// Builder structure for new members
type memberActions func(*Member)
type MemberBuilder struct {
	actions []memberActions
}

// Initialize the new members builder
func NewMemberBuilder() *MemberBuilder {
	log.Println("MemberBuilder.NewMemberBuilder()")
	return &MemberBuilder{
		actions: []memberActions{},
	}
}

// Execute all actions, create the Member
// and return it to builder function
func (mb *MemberBuilder) Build() *Member {
	log.Println("MemberBuilder.Build()")
	member := newMember()
	for _, action := range mb.actions {
		action(member)
	}
	member.lastChanged = time.Now().UnixNano()
	member.valid = false
	member.validate()
	return member
}

// Set the full name during the construction of this member
// name + middle name + last name
func (mb *MemberBuilder) WithFullName(name string, middleName string, lastName string) *MemberBuilder {
	log.Println("MemberBuilder.WithFullName()")
	mb.actions = append(mb.actions, func(m *Member) {
		m.name = strings.TrimSpace(name)
		m.lastName = strings.TrimSpace(lastName)
		m.middleName = strings.TrimSpace(middleName)
	})
	return mb
}

// Set the day of birth during the construction of this member
func (mb *MemberBuilder) WithBirthDate(dob time.Time) *MemberBuilder {
	log.Println("MemberBuilder.WithBirthDate()")
	mb.actions = append(mb.actions, func(m *Member) {
		// m.SetBirthDate(dob)
		m.dob = dob
	})
	return mb
}

// Set the gender during the construction of this member
func (mb *MemberBuilder) WithGender(gender string) *MemberBuilder {
	log.Println("MemberBuilder.WithGender()")
	mb.actions = append(mb.actions, func(m *Member) {
		// m.SetGender(gender)
		m.gender = gender
	})
	return mb
}

// Set the nickname during the construction of this member
func (mb *MemberBuilder) WithNickname(nick string) *MemberBuilder {
	log.Println("MemberBuilder.WithNickname()")
	mb.actions = append(mb.actions, func(m *Member) {
		// m.SetNickname(nick)
		m.nickname = strings.TrimSpace(nick)
	})
	return mb
}
