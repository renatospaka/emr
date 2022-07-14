package family_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	family "github.com/renatospaka/emr/family/domain/entity"
// )

// var (
// 	testHOF    = &family.Member{}
// )

// func init() {
// 	testMemberBuilder = family.NewMemberBuilder()
// 	testHOF = testMemberBuilder.
// 		WithFullName("Name", "Middle", "Last").
// 		WithBirthDate(dobAdult).
// 		WithGender(family.Male).
// 		WithNickname("Nick").
// 		Build()

// 	testFamilyBuilder = family.NewFamilyBuilder()
// 	testFamily = testFamilyBuilder.
// 		WithSurname("Super Family").
// 		Build()

// 	testFamilyMemberBuilder = family.NewFamilyMemberBuilder()
// }

// func TestFamilyMember_IsValid(t *testing.T) {
// 	familyMember := testFamilyMemberBuilder. 
// 		WithHeadOfFamily(testHOF).
// 		Build()
	
// 	require.EqualValues(t, family.Self, familyMember.relationType())
// 	require.True(t, familyMember.IsValid())
// 	require.Empty(t, familyMember.Err())
// }
