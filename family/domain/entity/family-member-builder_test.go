package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	family "github.com/renatospaka/emr/family/domain/entity"
)

var testFamilyMemberBuilder = &family.FamilyMemberBuilder{}

func init() {
	testFamilyMemberBuilder = family.NewFamilyMemberBuilder()
}

func TestFamilyMemberBuilder_Build(t *testing.T) {
	member := testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Top Cat").
		Build()

	famMemb := testFamilyMemberBuilder.
		AsHOF(member). 
		Build()

	require.IsTypef(t, &family.FamilyMember{}, famMemb, "não é do tipo *FamilyMember{}")
	// require.Len(t, famMemb.ErrToArray(), 2)
	require.True(t, famMemb.IsValid())
	require.Empty(t, famMemb.Err())
}
