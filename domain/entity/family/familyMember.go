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
	lastChanged  int64   `json:"-"`
	headOfFamily bool    `json:"status"`
}

func NewFamilyMember(member *Member) *FamilyMember {
	return &FamilyMember{
		member:       member,
		relationType: TBDRelation,
		status:       FreshMember,
		valid:        false,
		headOfFamily: false,
	}
}

// Set the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) SetHeadOfFamily() *FamilyMember {
	fm.headOfFamily = true
	fm.lastChanged = time.Now().UnixNano()
	return fm
}

// Unset the person who is the responsible for manage information
// of this family core
func (fm *FamilyMember) UnsetHeadOfFamily() *FamilyMember {
	fm.headOfFamily = false
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

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (fm *FamilyMember) validate() {
	analysisErrsFamilyMembers.RemoveAll()

	// member validation
	fm.member.validate()

	if !fm.member.IsAdult() && !fm.member.IsElderly() {
		analysisErrsFamilyMembers.AddErr(ErrFamilyMemberNotRelated)
	}

	if strings.TrimSpace(fm.relationType) == "" {
		analysisErrsFamilyMembers.AddErr(ErrFamilyMemberNotRelated)
	}

	fm.valid = (analysisErrsFamilyMembers.Count() == 0 && analysisErrsMembers.Count() == 0)
}
