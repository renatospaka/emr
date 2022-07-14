package family

type FamilyEntityInterface interface {
	ID()
	SetSurname(surname string) *Family
	Surname() string
	HasMembers() bool
	HasHeadOfFamily() bool
}

type FamilyEntityBuilderInterface interface {
	Build() *Family
	WithSurname(surname string) *FamilyBuilder
	AsHOF(hof *Member) *FamilyBuilder
	WithMember(member *Member) *FamilyBuilder
}

type FamilyValidationEntityInterface interface {
	IsValid() bool
	Err() string
	ErrToArray() []string
}
