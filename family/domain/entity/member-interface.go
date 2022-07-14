package family

import (
	"time"
)

type MemberEntityInterface interface {
	ID()
	SetFullName(name string, middleName string, lastName string) *Member
	FullName() string
	FullNameFormal() string
	SetNickname(nick string) *Member
	Nickname() string
	SetGender(gender string) *Member
	Gender() string
	SetBirthDate(dob time.Time) *Member
	BirthDate() time.Time
	AgeInYears() int64
}

type MemberAgeEntityInterface interface {
	AgeInMonths() int64
	IsNewborn() bool
	IsInfant() bool
	IsToddler() bool
	IsChild() bool
	IsTeen() bool
	IsAdult() bool
	IsElderly() bool
}

type MemberBuilderEntityInterface interface {
	Build() *Member
	WithFullName(name string, middleName string, lastName string) *MemberBuilder
	WithBirthDate(dob time.Time) *MemberBuilder
	WithGender(gender string) *MemberBuilder
	WithNickname(nickname string) *MemberBuilder
}
