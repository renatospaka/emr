package family

type FamilyEntityInterface interface {
	ID()
	ChangeSurname(surname string) *Family
	Surname() string
	Size() int
	HasHeadOfFamily() bool
	AddMember(member *FamilyMember) *Family
	RemoveMember(memberId string)
}

type FamilyEntityBuilderInterface interface {
	Build() *Family
	WithSurname(surname string) *FamilyBuilder
	Add(member *FamilyMember) *FamilyBuilder
}
