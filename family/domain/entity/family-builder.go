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
	log.Println("FamilyBuilder.NewFamilyBuilder()")
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
	fam.valid = false
	fam.validate()
	return fam
}

// Set the surname of the family
func (fb *FamilyBuilder) WithSurname(surname string) *FamilyBuilder {
	log.Println("FamilyBuilder.WithSurname()")
	fb.actions = append(fb.actions, func(f *Family) {
		f.SetSurname(surname)
	})
	return fb
}

// Set the Head of Family (hof), a person who is the responsible
// for manage information of this family core
func (fb *FamilyBuilder) WithHOF(headOfFamily *Member) *FamilyBuilder {
	log.Println("FamilyBuilder.WithHOF()")
	fb.actions = append(fb.actions, func(f *Family) {
		hof := newFamilyMember(headOfFamily)
		hof.
			SetHeadOfFamily().
			SetRelationType(Self)
		f.members = append(f.members, hof)
	})
	return fb
}

// add a family member to the family core
func (fb *FamilyBuilder) WithMember(member *Member) *FamilyBuilder {
	// log.Println("FamilyBuilder.WithMember()")
	// fb.actions = append(fb.actions, func(f *Family) {
	// 	memb := newFamilyMember(member)
	// 	memb.SetRelationType(Self)
	// 	// .SetMember()
	// 	f.members = append(f.members, memb)
	// })
	return fb
}
