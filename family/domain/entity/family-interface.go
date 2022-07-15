package family

type FamilyEntityInterface interface {
	ID()
	ChangeSurname(surname string) *Family
	Surname() string
	Size() int
	HasHeadOfFamily() bool
}

type FamilyEntityBuilderInterface interface {
	Build() *Family
	WithSurname(surname string) *FamilyBuilder
	Add(member *FamilyMember) *FamilyBuilder
}

type FamilyValidationEntityInterface interface {
	IsValid() bool
	Err() string
	ErrToArray() []string
}
