package family_test

import (
	"testing"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/stretchr/testify/require"
)

func init() {
	testFamilyMemberBuilder = family.NewFamilyMemberBuilder()
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

	require.EqualValues(t, family.Self, famMember.RelationType())
	require.True(t, famMember.IsHeadOfFamily())
	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
}

func TestFamilyMember_Inalid_HOF_Empty(t *testing.T) {
	famMember := testFamilyMemberBuilder. 
		AsHOF(&family.Member{}).
		Build()

	require.False(t, famMember.IsValid())
	require.NotEmpty(t, famMember.Err())
	require.Len(t, famMember.Err(), 2)
}

func TestFamilyMember_Inalid_HOF_Invalid_Age(t *testing.T) {
	testMemberHOF = testMemberBuilder.
		WithFullName("Name", "Invalid Age", "Last").
		WithBirthDate(dobTeen).
		WithGender(family.Male).
		WithNickname("Teen").
		Build()
	
	famMember := testFamilyMemberBuilder. 
		AsHOF(testMemberHOF).
		Build()

	require.False(t, famMember.IsValid())
	require.NotEmpty(t, famMember.Err())
	require.Len(t, famMember.Err(), 1)
}

func TestFamilyMember_IsValid_Ordinary(t *testing.T) {
	testMemberOrdinary := testMemberBuilder.
		WithFullName("Name", "Ordinary", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
	
	famMember := testFamilyMemberBuilder. 
		AsOrdinary(testMemberOrdinary).
		RelatedAs(family.Father).
		Build()

	require.EqualValues(t, family.Father, famMember.RelationType())
	require.False(t, famMember.IsHeadOfFamily())
	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
}
