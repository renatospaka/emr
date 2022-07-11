package family_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
)

var testFamilyBuilder = &family.FamilyBuilder{}

func init() {
	testFamilyBuilder = family.NewFamilyBuilder()
}

func TestFamilyBuilder_Build(t *testing.T) {
	member := testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()

	fam := testFamilyBuilder.
		WithSurname("Middle Last").
		WithHeadOfFamily(member).
		Build()

	require.IsTypef(t, &family.Family{}, fam, "não é do tipo *Family{}")
	require.True(t, fam.IsValid())
	require.Empty(t, fam.Err())
}

func TestFamilyBuilder_Invalid(t *testing.T) {
	member := testMemberBuilder.
		WithFullName("Na", "Middle", "LastLastLastLastLastLastLast").
		WithBirthDate(time.Time{}).
		WithGender("other").
		WithNickname("Nick").
		Build()

	fam := testFamilyBuilder.
		WithSurname("Middle Last").
		WithHeadOfFamily(member).
		Build()

	require.IsTypef(t, &family.Family{}, fam, family.ErrFamilyError.Error())
	require.False(t, fam.IsValid())
	require.NotEmpty(t, fam.Err())
	require.Len(t, fam.ErrToArray(), 6)
}
