package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	family "github.com/renatospaka/emr/family/domain/entity"
)

var testFamilyBuilder *family.FamilyBuilder

func init() {
	testFamilyBuilder = family.NewFamilyBuilder()
}

func TestFamilyBuilder_Build(t *testing.T) {
	member := testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Top Cat").
		Build()

	fam := testFamilyBuilder.
		WithSurname("Super Family").
		Add(member).
		Build()

	require.True(t, fam.IsValid())
	require.Empty(t, fam.Err())
	require.IsTypef(t, &family.Family{}, fam, "não é do tipo *Family{}")
}

func TestFamilyBuilder_Build_Invalid(t *testing.T) {
	fam := testFamilyBuilder.
		WithSurname("Super Family").
		Add(&family.Member{}).
		Build()

	require.False(t, fam.IsValid())
	require.NotEmpty(t, fam.Err())
}
