package family

type FamilyMemberEntityInterface interface {
	SetHeadOfFamily() *FamilyMember
	UnsetHeadOfFamily() *FamilyMember
	IsHeadOfFamily() bool
	ChangeRelationType(relationType string) *FamilyMember
	RelationType() string
	Status() string
}

type FamilyMemberBuilderEntityInterface interface {
	Build() *FamilyMember
	RelatedAs(relationType string) *FamilyMemberBuilder
	AsHOF(member *Member) *FamilyMemberBuilder
	AsOrdinary(member *Member) *FamilyMemberBuilder
}
