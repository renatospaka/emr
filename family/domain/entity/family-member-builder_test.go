package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	family "github.com/renatospaka/emr/family/domain/entity"
)

var (
	testFamilyMemberBuilder *family.FamilyMemberBuilder
)

func init() {
	testMemberBuilder = family.NewMemberBuilder()
	testFamilyBuilder = family.NewFamilyBuilder()
	testFamilyMemberBuilder = family.NewFamilyMemberBuilder()
}

func TestFamilyMember_Build(t *testing.T) {
	hof := testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
		
	famMember := testFamilyMemberBuilder.
		AsHOF(hof).
		Build()

	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
	require.IsTypef(t, &family.Member{}, famMember.Member, "não é do tipo *FamilyMember{}")
}

func TestFamilyMember_Build_Invalid(t *testing.T) {
	famMember := testFamilyMemberBuilder.
		AsHOF( &family.Member{}).
		Build()

	require.False(t, famMember.IsValid())
	require.NotEmpty(t, famMember.Err())
}
