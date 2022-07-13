package family

import (
	"strings"
	"time"

	"github.com/renatospaka/emr/infrastructure/utils"
)

var (
	analysisErrsFamilyMembers = utils.NewAnalysisErrs()
)

type FamilyMember struct {
	member       *Member `json:"member"`
	relationType string  `json:"relation_type"`
	status       string  `json:"status"`
	valid        bool    `json:"-"`
	headOfFamily bool    `json:"status"`
	lastChanged  int64   `json:"-"`
}

func newFamilyMember(member *Member) *FamilyMember {
	return &FamilyMember{
		member:       member,
		relationType: TBDRelation,
		status:       FreshMember,
		valid:        false,
		headOfFamily: false,
		lastChanged: time.Now().UnixNano(),
	}
}

// Set the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) Add(member *Member) *FamilyMember {
	fm.member = member
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Set the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) SetHeadOfFamily() *FamilyMember {
	fm.headOfFamily = fm.canBeHOF()
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Unset the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) UnsetHeadOfFamily() *FamilyMember {
	fm.headOfFamily = false
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Return if this member is the responsible of this family core
func (fm *FamilyMember) IsHeadOfFamily() bool {
	return fm.headOfFamily
}

// Set relationship type of the person to the on who is the
// head of the family of this family core
func (fm *FamilyMember) SetRelationType(relationType string) *FamilyMember {
	fm.relationType = relationType
	fm.valid = false
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Return the  relationship type of this member to the
// head of the family
func (fm *FamilyMember) RelationType() string {
	return fm.relationType
}

// Return the current status of this member
func (fm *FamilyMember) Status() string {
	return fm.status
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (fm *FamilyMember) IsValid() bool {
	fm.validate()
	return fm.valid
}

// Return all errors found during the validation process
// in a single string with a \n segregating each error
func (fm *FamilyMember) Err() string {
	analysis := ""
	builder := strings.Builder{}

	// errors from familyMember validation
	if analysisErrsFamilyMembers.Count() > 0 {
		for e := 0; e < len(analysisErrsFamilyMembers.Analysis); e++ {
			builder.WriteString(analysisErrsFamilyMembers.Analysis[e].ErrDescription)
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
func (fm *FamilyMember) ErrToArray() []string {
	analysis := fm.Err()
	toArray := []string{}
	if len(analysis) > 0 {
		newAnalisys := strings.Split(analysis, "\n")
		for e := 0; e < len(newAnalisys)-1; e++ {
			toArray = append(toArray, newAnalisys[e])
		}
	}
	return toArray
}

// check whether the current member is able to assume
// the head of family position - only of age people can
func (fm *FamilyMember) canBeHOF() bool {
	hof := (fm.member.IsAdult() || fm.member.IsElderly())
	if !hof {
		fm.valid = false
	}
	return hof
}

// Check whenever the structure of this family member
// is intact and filled accordingly to the model rules
func (fm *FamilyMember) validate() {
	analysisErrsFamilyMembers.RemoveAll()

	// check member validation
	fm.member.validate()

	if !fm.canBeHOF() {
		analysisErrsFamilyMembers.AddErr(ErrFamilyMemberHOFInvalidAge)
	}

	if strings.TrimSpace(fm.relationType) == "" {
		analysisErrsFamilyMembers.AddErr(ErrFamilyMemberNotRelated)
	} 
	if (fm.relationType != Self && fm.headOfFamily) {
		analysisErrsFamilyMembers.AddErr(ErrFamilyMemberInvalidRelation)
	} 

	fm.valid = (analysisErrsFamilyMembers.Count() == 0 && analysisErrsMembers.Count() == 0)
}
