package family

import (
	"time"
)

// Builder structure for new family member
type familyMemberActions func(*FamilyMember)
type FamilyMemberBuilder struct {
	actions []familyMemberActions
}

// Initialize the new family member builder
func NewFamilyMemberBuilder() *FamilyMemberBuilder {
	// log.Println("FamilyMemberBuilder.NewFamilyMemberBuilder()")
	return &FamilyMemberBuilder{
		actions: []familyMemberActions{},
	}
}

// Execute all actions, create the Family Member
// and return it to caller
func (fb *FamilyMemberBuilder) Build() *FamilyMember {
	// log.Println("FamilyMemberBuilder.Build()")
	famMemb := newFamilyMember()
	for _, action := range fb.actions {
		action(famMemb)
	}
	famMemb.valid = false
	famMemb.lastChanged = time.Now().UnixNano()
	famMemb.validate()
	return famMemb
}

// Set the relationship of this member to the HOF
func (fb *FamilyMemberBuilder) RelatedAs(relationType string) *FamilyMemberBuilder {
	// log.Println("FamilyMemberBuilder.SetRelationType()")
	fb.actions = append(fb.actions, func(fm *FamilyMember) {
		fm.SetRelationType(relationType)
	})
	return fb
}

// Add this member as the HOF
func (fb *FamilyMemberBuilder) AsHOF(hof *Member) *FamilyMemberBuilder {
	// log.Println("FamilyMemberBuilder.SetRelationType()")
	fb.actions = append(fb.actions, func(fm *FamilyMember) {
		fm.add(hof)
		fm.SetHeadOfFamily()
		fm.SetRelationType(Self)
	})
	return fb
}

// Add this member as the HOF
func (fb *FamilyMemberBuilder) WithMember(member *Member) *FamilyMemberBuilder {
	// log.Println("FamilyMemberBuilder.SetRelationType()")
	fb.actions = append(fb.actions, func(fm *FamilyMember) {
		fm.add(member)
		fm.UnsetHeadOfFamily()
	})
	return fb
}
