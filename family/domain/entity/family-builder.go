package family

import (
	"log"
	"time"
)

// Builder structure for new families
type familyActions func(*Family)
type FamilyBuilder struct {
	actions []familyActions
}

// Initialize the new families builder
func NewFamilyBuilder() *FamilyBuilder {
	return &FamilyBuilder{
		actions: []familyActions{},
	}
}

// Execute all actions, create the Family
// and return it to caller
func (fb *FamilyBuilder) Build() *Family {
	log.Println("FamilyBuilder.Build()")
	fam := newFamily()
	for _, action := range fb.actions {
		action(fam)
	}
	fam.lastChanged = time.Now().UnixNano()
	fam.validate()

	return fam
}

// Set the surname of the family
func (fb *FamilyBuilder) WithSurname(surname string) *FamilyBuilder {
	fb.actions = append(fb.actions, func(f *Family) {
		f.surname = surname
	})
	return fb
}

// Set the person who is the responsible for manage information
// of this family core
func (fb *FamilyBuilder) HasHeadOfFamily(headOfFamily *Member) *FamilyBuilder {
	fb.actions = append(fb.actions, func(f *Family) {
		hof := newFamilyMember(headOfFamily)
		hof.SetHeadOfFamily().
			SetRelationType(Self)
		f.members = append(f.members, hof)
	})
	return fb
}
