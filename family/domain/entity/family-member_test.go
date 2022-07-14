package family_test

import (
	"testing"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/stretchr/testify/require"
)

var (
	testMemberHOF    *family.Member
)

func init() {
	// testMemberBuilder = family.NewMemberBuilder()
	testFamilyMemberBuilder = family.NewFamilyMemberBuilder()

	// testMember = testMemberBuilder.
	// 	WithFullName("Name2", "Ordinary", "Last2").
	// 	WithBirthDate(dobAdult).
	// 	WithGender(family.Female).
	// 	WithNickname("Nick2").
	// 	Build()

}

func TestFamilyMember_IsValid_HOF(t *testing.T) {
	testMemberHOF = testMemberBuilder.
		WithFullName("Name", "HOF", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
	
	famMember := testFamilyMemberBuilder. 
		AsHOF(testMemberHOF).
		Build()


	require.True(t, famMember.IsHeadOfFamily())
	require.True(t, testMemberHOF.IsAdult())
	require.EqualValues(t, family.Self, famMember.RelationType())
	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
	// require.Len(t, famMember.ErrToArray(), 3)
}

func TestFamilyMember_Inalid_HOF_Empty(t *testing.T) {
	famMember := testFamilyMemberBuilder. 
		AsHOF(&family.Member{}).
		Build()

	require.False(t, famMember.IsValid())
	require.NotEmpty(t, famMember.Err())
	require.Len(t, famMember.ErrToArray(), 3)
}
