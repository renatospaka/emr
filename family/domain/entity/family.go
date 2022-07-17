package family

import (
	"strings"
	"time"

	"github.com/renatospaka/emr/common/infrastructure/err"
	"github.com/renatospaka/emr/common/infrastructure/utils"
)

// var (
// 	analysisErrsFamily = errs.NewErrors()
// )

type Family struct {
	id          string          `json:"family_id"`
	surname     string          `json:"surname"`
	valid       bool            `json:"-"`
	lastChanged int64           `json:"-"`
	members     []*FamilyMember `json:"members"`
	err          *err.Errors
}

func newFamily() *Family {
	// log.Println("Family.newFamily()")
	fam := &Family{
		id:          utils.GetID(),
		surname:     "",
		valid:       false,
		lastChanged: time.Now().UnixNano(),
		members:     []*FamilyMember{},
		err:          err.NewErrors().ClearAll(),
	}

	return fam
}

// Return the ID of this family
func (f *Family) ID() string {
	// log.Println("Family.ID()")
	return f.id
}

// Set the surname of the family
func (f *Family) ChangeSurname(surname string) *Family {
	// log.Println("Family.ChangeSurname()")
	f.surname = strings.TrimSpace(surname)
	f.lastChanged = time.Now().UnixNano()
	f.valid = false
	return f
}

// Return the surname of the family
func (f *Family) Surname() string {
	// log.Println("Family.Surname()")
	return f.surname
}

// The size of the family
//(count of family members)
func (f *Family) Size() int {
	// log.Println("Family.Size(", len(f.members) > 0, ")")
	return len(f.members)
}

// A valid family must have one head of family
func (f *Family) HasHeadOfFamily() bool {
	hasHOF := false
	for m := 0; m < len(f.members); m++ {
		if f.members[m].IsHeadOfFamily() {
			hasHOF = true
			break
		}
	}
	// log.Println("Family.HasHeadOfFamily(", hasHOF, ")")
	return hasHOF
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (f *Family) IsValid() bool {
	// log.Println("Family.IsValid() - starting")
	clearErrsOnValidation = true
	f.validate()
	// log.Printf("Family.IsValid(%t)", f.valid)
	return f.valid
}

// Return all errors found during the validation process
// in a single array
func (f *Family) Err() []string {
	// log.Println("Family.Err()")
	toArray := []string{}
	if f.err.Count() > 0 {
		for e := 0; e < len(f.err.Err); e++ {
			toArray = append(toArray, f.err.Err[e].Description)
		}
	}

	// returns all errors found during the Family Member process
	// in the same array
	if len(f.members) > 0 {
		if f.Size() > 0 {
			for m := 0; m < len(f.members); m++ {
				// family member level
				famMemb := f.members[m]
				for e := 0; e < len(famMemb.err.Err); e++ {
					toArray = append(toArray, famMemb.err.Err[e].Description)
				}

				// member level
				memb := famMemb.Member
				for e := 0; e < len(memb.err.Err); e++ {
					toArray = append(toArray, memb.err.Err[e].Description)
				}
			}
		}
	}
	
	return toArray
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (f *Family) validate() {
	// log.Println("Family.validate() -> clearErrsOnValidation:", clearErrsOnValidation)
	f.err.ClearAll()

	err := utils.IsVaalidUUID(f.id)
	if err != nil {
		f.err.Add(ErrInvalidFamilyID)
		f.err.Add(err)
	}

	if f.surname == "" {
		f.err.Add(ErrMissingFamilySurname)
	}

	// validate if all members, if any, are filled
	// accordingly to the model rules
	hasHOF, hasErros := false, false
	countHOF, countErrors := 0, 0
	if f.Size() < 1 {
		f.err.Add(ErrFamilyMemberMissing)
	}

	for fm := 0; fm < len(f.members); fm++ {
		fMemb := f.members[fm]
		isValid := !fMemb.IsValid()
		isHOF := fMemb.IsHeadOfFamily()

		if !isValid {
			countErrors += fMemb.err.Count()
			hasErros = true
		}
		
		if isHOF {
			countHOF ++
			hasHOF = true
		}
	}

	// there is no HOF defined
	if !hasHOF {
		f.err.Add(ErrFamilyMemberHOFMissing)
		hasErros = true
		countErrors ++
	}

	// there are more than one HOF defined
	if hasHOF && countHOF > 1 {
		f.err.Add(ErrFamilyMemberTooManyHOF)
		hasErros = true
		countErrors ++
	}

	f.valid = (f.err.Count() == 0 && !hasErros)
}