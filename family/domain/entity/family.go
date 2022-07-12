package family

import (
	"log"
	"strings"
	"time"

	"github.com/renatospaka/emr/infrastructure/utils"
)

var (
	analysisErrsFamily = utils.NewAnalysisErrs()
)

type Family struct {
	id          string          `json:"family_id"`
	surname     string          `json:"surname"`
	valid       bool            `json:"-"`
	lastChanged int64           `json:"-"`
	members     []*FamilyMember `json:"members"`
}

func newFamily() *Family {
	log.Println("Family.newFamily()")
	fam := &Family{
		id:          utils.GetID(),
		surname:     "",
		valid:       false,
		lastChanged: time.Now().UnixNano(),
		members:     []*FamilyMember{},
	}

	return fam
}

// Return the ID of this family
func (f *Family) ID() string {
	log.Println("Family.ID()")
	return f.id
}

// Set the surname of the family
func (f *Family) SetSurname(surname string) *Family {
	log.Println("Family.SetSurname()")
	f.surname = strings.TrimSpace(surname)
	f.lastChanged = time.Now().UnixNano()
	f.valid = false
	return f
}

// Return the surname of the family
func (f *Family) Surname() string {
	log.Println("Family.Surname()")
	return f.surname
}

// A valid family must have at least one member
func (f *Family) HasMembers() bool {
	log.Println("Family.HasMembers(", len(f.members) > 0, ")")
	return len(f.members) > 0
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
	log.Println("Family.HasHeadOfFamily(", hasHOF, ")")
	return hasHOF
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (f *Family) IsValid() bool {
	f.validate()
	log.Println("Family.IsValid(", f.valid, ")")
	return f.valid
}

// Return all errors found during the validation process
// in a single string with a \n segregating each error
func (f *Family) Err() string {
	log.Println("Family.Err()")
	analysis := ""
	builder := strings.Builder{}

	// errors from family validation
	if analysisErrsFamily.Count() > 0 {
		for e := 0; e < len(analysisErrsFamily.Analysis); e++ {
			builder.WriteString(analysisErrsFamily.Analysis[e].ErrDescription)
			builder.WriteString("\n")
		}
	}

	// errors from member validation
	if analysisErrsMembers.Count() > 0 {
		for e := 0; e < len(analysisErrsMembers.Analysis); e++ {
			builder.WriteString(analysisErrsMembers.Analysis[e].ErrDescription)
			builder.WriteString("\n")
		}
	}

	analysis = builder.String()
	return analysis
}

// Return all errors found during the validation process
// in an array
func (f *Family) ErrToArray() []string {
	log.Println("Family.ErrToArray()")
	analysis := f.Err()
	toArray := []string{}
	if len(analysis) > 0 {
		newAnalisys := strings.Split(analysis, "\n")
		for e := 0; e < len(newAnalisys)-1; e++ {
			toArray = append(toArray, newAnalisys[e])
		}
	}
	return toArray
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (f *Family) validate() {
	log.Println("Family.validate() -> clearErrsOnValidation:", clearErrsOnValidation)
	analysisErrsFamily.RemoveAll()

	if strings.TrimSpace(f.surname) == "" {
		analysisErrsFamily.AddErr(ErrMissingFamilySurname)
	}

	// Validate if all members, if any, are filled
	// accordingly to the model rules
	if !f.HasMembers() {
		analysisErrsFamily.AddErr(ErrFamilyMemberMissing)
	} else {
		f.validateHeadOfFamily()
	}

	if strings.TrimSpace(f.id) == "" {
		analysisErrsFamily.AddErr(ErrMissingFamilyID)
	}

	log.Println("Family.validate(", analysisErrsFamily.Count() == 0, ")")
	f.valid = (analysisErrsFamily.Count() == 0 && analysisErrsMembers.Count() == 0)
}

// There must be at least one member in a family
// All members must be valid
// There must be one HOF defined
// Only ONE HOF defined per family
func (f *Family) validateHeadOfFamily() {
	log.Println("Family.validateHeadOfFamily()")
	hasHOF := false
	thisHOF := false
	countHOF := 0
	for m := 0; m < len(f.members); m++ {
		clearErrsOnValidation = false
		thisHOF = f.members[m].headOfFamily
		f.members[m].member.validate()

		// how many HOF are there?
		if thisHOF {
			countHOF++
			hasHOF = true
		}

		// invalid member (any reason)
		if !f.members[m].member.valid {
			analysisErrsMembers.AddErr(ErrMemberError)

			// there is not possible an invalid HOF
			if thisHOF {
				analysisErrsMembers.AddErr(ErrFamilyMemberHOFError)
			}
		}
	}

	// there is no HOF defined
	if !hasHOF {
		analysisErrsFamily.AddErr(ErrFamilyMemberHOFMissing)
	} else
	// there are more than one HOF defined
	if hasHOF && countHOF > 1 {
		analysisErrsFamily.AddErr(ErrFamilyMemberTooManyHOF)
	}
	clearErrsOnValidation = true
}
