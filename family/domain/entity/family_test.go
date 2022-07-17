package family_test

import (
	"testing"
	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/renatospaka/emr/common/infrastructure/utils"
)

var (
	testFamily *family.Family
)

func init() {
	testFamilyBuilder = family.NewFamilyBuilder()
	testMember = testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
}

func TestFamily_ID(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Build()

	id := testFamily.ID()
	err := utils.IsVaalidUUID(id)

	require.IsType(t, "", id)
	require.Len(t, id, 36)
	require.Nil(t, err)
}

func TestFamily_IsValid(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Add(testMember).
		Build()
	testFamily.ChangeSurname("Surname")

	require.True(t, testFamily.IsValid())
	require.Empty(t, testFamily.Err())
}

func TestFamily_IsValid_No(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Add(testMember).
		Build()
	testFamily.ChangeSurname("")

	require.False(t, testFamily.IsValid())
	require.NotEmpty(t, testFamily.Err())
}

func TestFamily_Surname(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Add(testMember).
		Build()
	testFamily.ChangeSurname("Another Surname")
	surname := testFamily.Surname()

	require.EqualValues(t, "Another Surname", surname)
}

func TestFamily_Size(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Add(testMember).
		Build()
	//Need to prepare to include a member
	//This test will fail until then
	size := testFamily.Size()

	require.EqualValues(t, 1, size)
}

func TestFamily_Size_Missing(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Add(testMember).
		Build()
	//Need to prepare to include a member
	//This test will fail until then
	testFamily.IsValid()
	size := testFamily.Size()

	allErrors := testFamily.Err()
	require.EqualValues(t, 1, size)
	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
	// require.Contains(t, allErrors, family.ErrFamilyMemberHOFMissing.Error())
	require.Equal(t, 1, len(allErrors))
}

func TestFamily_HasHeadOfFamily(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Add(testMember).
		Build()
	//Need to prepare to include a member
	//This test will fail until then
	hasHOF := testFamily.HasHeadOfFamily()

	require.True(t, hasHOF)
}

func TestFamily_HasHeadOfFamily_Missing(t *testing.T) {
	testFamily := testFamilyBuilder.
		WithSurname("Super Family").
		Add(testMember).
		Build()
	//Need to prepare to include a member
	//This test will fail until then
	testFamily.IsValid()
	hasHOF := testFamily.HasHeadOfFamily()

	allErrors := testFamily.Err()
	require.False(t, hasHOF)
	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
	// require.Contains(t, allErrors, family.ErrFamilyMemberHOFMissing.Error())
	require.Equal(t, 1, len(allErrors))
}
