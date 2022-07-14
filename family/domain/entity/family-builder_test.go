package family_test

import (
	"testing"
	"time"

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
		WithHOF(member).
		Build()

	require.IsTypef(t, &family.Family{}, fam, "não é do tipo *Family{}")
	require.True(t, fam.IsValid())
	require.Empty(t, fam.Err())
}

func TestFamilyBuilder_InvalidMember(t *testing.T) {
	invalidMember := testMemberBuilder.
		WithFullName("Na", "", "LastLastLastLastLastLastLast").
		WithBirthDate(time.Time{}).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()

	fam := testFamilyBuilder.
		WithSurname("Super Family").
		WithHOF(invalidMember).
		Build()

	require.IsTypef(t, &family.Family{}, fam, family.ErrFamilyError.Error())
	require.False(t, fam.IsValid())
	require.NotEmpty(t, fam.Err())
	require.Len(t, fam.ErrToArray(), 4)
}

func TestFamilyBuilder_InvalidFamily(t *testing.T) {
	missingMember := &family.Member{}
	fam := testFamilyBuilder.
		WithSurname("Super Family").
		WithHOF(missingMember).
		Build()

	require.IsTypef(t, &family.Family{}, fam, family.ErrFamilyError.Error())
	require.False(t, fam.IsValid())
	require.NotEmpty(t, fam.Err())
	require.Len(t, fam.ErrToArray(), 10)
}
