package family_test

import (
	"testing"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/stretchr/testify/require"
)

func TestFamilyMember_IsValid_HOF(t *testing.T) {
	hof := createHOFMember()
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder.
		AsHOF(hof).
		Build()

	require.EqualValues(t, family.Self, famMember.RelationType())
	require.True(t, famMember.IsHeadOfFamily())
	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
}

func TestFamilyMember_Inalid_HOF_Empty(t *testing.T) {
	empty := createEmptyMember()
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder.
		AsHOF(empty).
		Build()

	require.False(t, famMember.IsValid())
	allErrors := famMember.Err()

	require.Contains(t, allErrors, family.ErrMemberError.Error())
	require.Contains(t, allErrors, family.ErrFamilyMemberHOFError.Error())
	require.Len(t, allErrors, 2)
}

func TestFamilyMember_Inalid_HOF_Invalid_Age(t *testing.T) {
	teen := createTeenagerMember()
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder.
		AsHOF(teen).
		Build()

	require.False(t, famMember.IsValid())
	allErrors := famMember.Err()

	require.Contains(t, allErrors, family.ErrFamilyMemberHOFInvalidAge.Error())
	require.Len(t, allErrors, 1)
}

func TestFamilyMember_IsValid_Ordinary(t *testing.T) {
	ordinary := createMember()
	familyMemberBuilder := family.NewFamilyMemberBuilder()
	famMember := familyMemberBuilder.
		AsOrdinary(ordinary).
		RelatedAs(family.Father).
		Build()

	require.EqualValues(t, family.Father, famMember.RelationType())
	require.False(t, famMember.IsHeadOfFamily())
	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
}
