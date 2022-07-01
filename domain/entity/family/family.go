package family

import (
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

func NewFamily(surname string) *Family {
	newFamily := &Family{
		id:      utils.GetID(),
		surname: strings.TrimSpace(surname),
		valid:   false,
		members: []*FamilyMember{},
	}

	newFamily.IsValid()
	return newFamily
}

// Return the ID of this family
func (f *Family) ID() string {
	return f.id
}

// Set the surname of the family
func (f *Family) SetSurname(surname string) *Family {
	f.surname = strings.TrimSpace(surname)
	f.lastChanged = time.Now().UnixNano()
	f.validate()
	return f
}

// Return the surname of the family
func (f *Family) Surname() string {
	return f.surname
}

// A valid family must have at least one member
func (f *Family) HasMembers() bool {
	return len(f.members) > 0
}

// A valid family must have one head of family
func (f *Family) HasHeadOfFamily() bool {
	hasHead := false
	for m := 0; m < len(f.members); m++ {
		if f.members[m].IsHeadOfFamily() {
			hasHead = true
			break
		}
	}
	return hasHead
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (f *Family) IsValid() bool {
	f.validate()
	return f.valid
}

// Return all errors found during the validation process
// in a single string with a \n segregating each error
func (f *Family) Err() string {
	analysis := ""
	if analysisErrsFamily.Count() > 0 {
		builder := strings.Builder{}
		for e := 0; e < len(analysisErrsFamily.Analysis); e++ {
			builder.WriteString(analysisErrsFamily.Analysis[e].ErrDescription)
			builder.WriteString("\n")
		}
		analysis = builder.String()
	}
	return analysis
}

// Return all errors found during the validation process
// in an array
func (f *Family) ErrToArray() []string {
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
	analysisErrsFamily.RemoveAll()

	if strings.TrimSpace(f.surname) == "" {
		analysisErrsFamily.AddErr(ErrMissingFamilySurname)
	}

	if !f.HasMembers() {
		analysisErrsFamily.AddErr(ErrFamilyMemberMissing)
	}

	if !f.HasHeadOfFamily() {
		analysisErrsFamily.AddErr(ErrFamilyMemberHeadMissing)
	}

	if strings.TrimSpace(f.id) == "" {
		analysisErrsFamily.AddErr(ErrMissingFamilyID)
	}

	f.valid = (analysisErrsFamily.Count() == 0)
}
