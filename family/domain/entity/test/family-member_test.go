package family_test

import (
	"testing"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/stretchr/testify/require"
)

// func init() {
// 	testFamilyMemberBuilder = family.NewFamilyMemberBuilder()
// }

func TestFamilyMember_IsValid_HOF(t *testing.T) {
	memberBuilder := family.NewMemberBuilder()
	memberHOF := memberBuilder.
		WithFullName("Name", "HOF", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
	
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder. 
		AsHOF(memberHOF).
		Build()

	require.EqualValues(t, family.Self, famMember.RelationType())
	require.True(t, famMember.IsHeadOfFamily())
	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
}

func TestFamilyMember_Inalid_HOF_Empty(t *testing.T) {
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder. 
		AsHOF(&family.Member{}).
		Build()

	require.False(t, famMember.IsValid())
	require.NotEmpty(t, famMember.Err())
	require.Len(t, famMember.Err(), 2)
}

func TestFamilyMember_Inalid_HOF_Invalid_Age(t *testing.T) {
	memberBuilder := family.NewMemberBuilder()
	memberHOF := memberBuilder.
		WithFullName("Name", "Invalid Age", "Last").
		WithBirthDate(dobTeen).
		WithGender(family.Male).
		WithNickname("Teen").
		Build()
	
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder. 
		AsHOF(memberHOF).
		Build()

	require.False(t, famMember.IsValid())
	require.NotEmpty(t, famMember.Err())
	require.Len(t, famMember.Err(), 1)
}

func TestFamilyMember_IsValid_Ordinary(t *testing.T) {
	memberBuilder := family.NewMemberBuilder()
	testMemberOrdinary := memberBuilder.
		WithFullName("Name", "Ordinary", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
	
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder. 
		AsOrdinary(testMemberOrdinary).
		RelatedAs(family.Father).
		Build()

	require.EqualValues(t, family.Father, famMember.RelationType())
	require.False(t, famMember.IsHeadOfFamily())
	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
}
