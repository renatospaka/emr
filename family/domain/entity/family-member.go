package family

import (
	"time"

	"github.com/renatospaka/emr/common/infrastructure/err"
)

// var (
// 	analysisErrsFamilyMembers = errs.NewErrors()
// )

type FamilyMember struct {
	*Member      `json:"member"`
	relationType string `json:"relation_type"`
	status       string `json:"status"`
	valid        bool   `json:"-"`
	headOfFamily bool   `json:"status"`
	lastChanged  int64  `json:"-"`
	err          *err.Errors
}

func newFamilyMember() *FamilyMember {
	// log.Println("FamilyMember.newFamilyMember()")
	return &FamilyMember{
		Member:       &Member{},
		relationType: TBDRelation,
		status:       FreshMember,
		valid:        false,
		headOfFamily: false,
		lastChanged:  time.Now().UnixNano(),
		err:          err.NewErrors().ClearAll(),
	}
}

// Set the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) SetHeadOfFamily() *FamilyMember {
	// log.Println("FamilyMember.SetHeadOfFamily()")
	fm.headOfFamily = true
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Unset the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) UnsetHeadOfFamily() *FamilyMember {
	// log.Println("FamilyMember.UnsetHeadOfFamily()")
	fm.headOfFamily = false
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Return if this member is the responsible of this family core
func (fm *FamilyMember) IsHeadOfFamily() bool {
	// log.Printf("FamilyMember.IsHeadOfFamily(%t)", fm.headOfFamily)
	return fm.headOfFamily
}

// Set relationship type of the person to the on who is the
// head of the family of this family core
func (fm *FamilyMember) ChangeRelationType(relationType string) *FamilyMember {
	// log.Println("FamilyMember.ChangeRelationType()")
	fm.relationType = relationType
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Return the  relationship type of this member to the
// head of the family
func (fm *FamilyMember) RelationType() string {
	// log.Printf("FamilyMember.RelationType(%s)", fm.relationType)
	return fm.relationType
}

// Return the current status of this member
func (fm *FamilyMember) Status() string {
	// log.Printf("FamilyMember.Status(%s)", fm.status)
	return fm.status
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (fm *FamilyMember) IsValid() bool {
	// log.Println("FamilyMember.IsValid() - starting")
	fm.validate()
	// log.Printf("FamilyMember.IsValid(%t)", fm.valid)
	return fm.valid
}

// Return all errors found during the validation process
// in an array
func (fm *FamilyMember) Err() []string {
	// log.Println("FamilyMember.Err()")
	toArray := []string{}
	if fm.err.Count() > 0 {
		for e := 0; e < len(fm.err.Err); e++ {
			toArray = append(toArray, fm.err.Err[e].Description)
		}
	}
	return toArray
}

// Set the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) add(member *Member) *FamilyMember {
	// log.Println("FamilyMember.add()")
	fm.Member = member
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// check whether the current member is able to assume
// the head of family position - only of age people can
func (fm *FamilyMember) hofReady() bool {
	hof := (fm.Member.IsAdult() || fm.Member.IsElderly())
	if !hof {
		fm.valid = false
	}
	// log.Printf("FamilyMember.hofReady(%t)", hof)
	return hof
}

// Check whenever the structure of this family member
// is intact and filled accordingly to the model rules
func (fm *FamilyMember) validate() {
	// log.Println("FamilyMember.validate()")
	fm.err.ClearAll()

	// check member validation
	memb := fm.Member
	if !memb.IsValid() {
		// invalid member (any reason)
		fm.err.Add(ErrMemberError)

		if fm.headOfFamily {
			// it is not possible an invalid HOF
			fm.err.Add(ErrFamilyMemberHOFError)
		}
	} else {
		// hof must be a valid member
		if fm.headOfFamily {
			if !fm.hofReady() {
				fm.err.Add(ErrFamilyMemberHOFInvalidAge)
			}
		}
	}

	// familiar relationship valid
	_, ok := relations[fm.relationType]
	if !ok {
		fm.err.Add(ErrFamilyMemberInvalidRelation)
	}
	if fm.relationType == "" {
		fm.err.Add(ErrFamilyMemberNotRelated)
	}
	if fm.relationType != Self && fm.headOfFamily {
		fm.err.Add(ErrFamilyMemberInvalidRelation)
	}

	fm.valid = (fm.err.Count() == 0 && memb.err.Count() == 0)
}
