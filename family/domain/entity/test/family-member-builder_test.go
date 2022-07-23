package family_test

import (
	"testing"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/stretchr/testify/require"
)

func TestFamilyMember_Build_HOF(t *testing.T) {
	hof := createHOFMember()

	famMembBuilder := family.NewFamilyMemberBuilder()
	famMember := famMembBuilder.
		AsHOF(hof).
		Build()

	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
	require.IsTypef(t, &family.Member{}, famMember.Member, "não é do tipo *FamilyMember{}")
}

func TestFamilyMember_Build_HOF_Invalid(t *testing.T) {
	empty := createEmptyMember()
	famMembBuilder := family.NewFamilyMemberBuilder()
	famMember := famMembBuilder.
		AsHOF(empty).
		Build()

	require.False(t, famMember.IsValid())
	allErrors := famMember.Err()

	require.Contains(t, allErrors, family.ErrMemberError.Error())
	require.Contains(t, allErrors, family.ErrFamilyMemberHOFError.Error())
	require.Len(t, allErrors, 2)
}

func TestFamilyMember_Build_Ordinary(t *testing.T) {
	member := createMember()
	famMembBuilder := family.NewFamilyMemberBuilder()
	famMember := famMembBuilder.
		AsOrdinary(member).
		RelatedAs(family.RelWife).
		Build()

	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
	require.IsTypef(t, &family.Member{}, famMember.Member, "não é do tipo *FamilyMember{}")
}

func TestFamilyMember_Empty(t *testing.T) {
	famMember := createEmptyFamilyMember()

	require.False(t, famMember.IsValid())
	allErrors := famMember.Err()

	require.Contains(t, allErrors, family.ErrInvalidFamilyMember.Error())
	require.Len(t, allErrors, 1)
}
