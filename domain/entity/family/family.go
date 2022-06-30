package family

import (
	"strings"
	"time"

	"github.com/renatospaka/emr2/infrastructure/utils"
)

var (
	analysisErrsFamily = utils.NewAnalysisErrs()
)

type Family struct {
	id          string `json:"family_id"`
	surname     string `json:"surname"`
	valid       bool   `json:"is_valid"`
	lastChanged int64
	members     []*FamilyMember `json:"members"`
}

type FamilyMember struct {
	member       *Member `json:"member"`
	relationType string  `json:"relation_type"`
	status       string  `json:"status"`
	headOfFamily bool    `json:"status"`
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

// Set the surname of the family
func (f *Family) SetSurname(surname string) *Family {
	f.surname = strings.TrimSpace(surname)
	f.lastChanged = time.Now().UnixNano()
	f.validate()
	return f
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

	if strings.TrimSpace(f.id) == "" {
		analysisErrsFamily.AddErr(ErrMissingFamilyID)
	}

	f.valid = (analysisErrsFamily.Count() == 0)
}

func NewFamilyMember(member *Member) *FamilyMember {
	return &FamilyMember{
		member:       member,
		relationType: TBDRelation,
		status:       FreshMember,
		headOfFamily: false,
	}
}

// Set the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) SetHeadOfFamily() *FamilyMember {
	fm.headOfFamily = true
	return fm
}

// Return if this member is the responsible of this family core
func (fm *FamilyMember) IsHeadOfFamily() bool {
	return fm.headOfFamily
}
