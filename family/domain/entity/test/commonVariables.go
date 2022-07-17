package family_test

import (
	"time"

	family "github.com/renatospaka/emr/family/domain/entity"
)

// member testing variables
var (
	testMember *family.Member
	dobNewborn time.Time
	dobInfant  time.Time
	dobToddler time.Time
	dobChild   time.Time
	dobTeen    time.Time
	dobAdult   time.Time
	dobElderly time.Time
)

// builder-member testing variables
var (
	testMemberBuilder *family.MemberBuilder
)

// family testing variables
var (
	testMemberHOF *family.Member
)

// builder-family testing variables
var (
	testFamilyBuilder *family.FamilyBuilder
)

// builder-family member testing variables
var (
	testFamilyMemberBuilder *family.FamilyMemberBuilder
)
