package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
)

func TestFamily_IsValid(t *testing.T) {
	familyMember := createFamilyMember()
	familyBuilder := createFamilyBuilder()
	fam := familyBuilder.
		WithSurname("Super Family").
		Add(familyMember).
		Build()

	require.Empty(t, fam.Err())
	require.True(t, fam.IsValid())
}

func TestFamily_Invalid(t *testing.T) {
	familyMember := createFamilyMember()
	familyBuilder := createFamilyBuilder()
	fam := familyBuilder.
		WithSurname("").
		Add(familyMember).
		Build()

	allErrors := fam.Err()
	require.False(t, fam.IsValid())
	require.Contains(t, allErrors, family.ErrMissingFamilySurname.Error())
	require.Len(t, allErrors, 1)
}

func TestFamily_Invalid_EmptyFamily(t *testing.T) {
	fam := createEmptyFamily()

	require.False(t, fam.IsValid())
	allErrors := fam.Err()
	require.Contains(t, allErrors, family.ErrInvalidFamily.Error())
	require.Len(t, allErrors, 1)
}

func TestFamily_Invalid_InvalidMember(t *testing.T) {
	fam := createInvalidFamily()
	size := fam.Size()

	allErrors := fam.Err()
	require.False(t, fam.IsValid())
	require.EqualValues(t, 1, size)
	require.Contains(t, allErrors, family.ErrInvalidMember.Error())
	require.Equal(t, 3, len(allErrors))
}

func TestFamily_ID(t *testing.T) {
	fam := createFamily()
	id := fam.ID()

	require.True(t, fam.IsValid())
	require.IsType(t, "", id)
}

func TestFamily_ChangeSurname(t *testing.T) {
	fam := createFamily()
	fam.ChangeSurname("Another Surname")
	surname := fam.Surname()

	require.True(t, fam.IsValid())
	require.EqualValues(t, "Another Surname", surname)
}

func TestFamily_Size(t *testing.T) {
	fam := createFamily()
	size := fam.Size()

	require.True(t, fam.IsValid())
	require.EqualValues(t, 1, size)
}

func TestFamily_HasHeadOfFamily(t *testing.T) {
	fam := createFamily()
	hasHOF := fam.HasHeadOfFamily()

	require.True(t, fam.IsValid())
	require.True(t, hasHOF)
}

func TestFamily_HasHeadOfFamily_Missing(t *testing.T) {
	fam := createMissingHOFFamily()
	hasHOF := fam.HasHeadOfFamily()
	allErrors := fam.Err()

	require.False(t, fam.IsValid())
	require.False(t, hasHOF)
	require.Contains(t, allErrors, family.ErrFamilyMemberHOFMissing.Error())
	require.Equal(t, 1, len(allErrors))
}

func TestFamily_AddMember(t *testing.T) {
	fam := createFamily()
	newMember := createOrdinaryFamilyMember()
	newMember.Member.ChangeBirthDate(dobChild)
	newMember.ChangeGender(family.Female)
	newMember.ChangeRelationType(family.RelDaughter)
	fam.AddMember(newMember)

	require.Empty(t, fam.Err())
	require.True(t, fam.IsValid())
}

func TestFamily_Members(t *testing.T) {
	fam := createFamily()
	newMember := createOrdinaryFamilyMember()
	newMember.Member.ChangeBirthDate(dobChild)
	newMember.ChangeGender(family.Female)
	newMember.ChangeRelationType(family.RelDaughter)
	fam.AddMember(newMember)
	allMembers := fam.Members()

	require.Empty(t, fam.Err())
	require.Equal(t, 2, len(allMembers))
	require.True(t, fam.IsValid())
}

func TestFamily_AddMember_InvalidMember(t *testing.T) {
	fam := createFamily()
	invalidNewMember := createEmptyOrdinaryFamilyMember()
	invalidNewMember.Member.ChangeBirthDate(dobChild)
	invalidNewMember.ChangeGender(family.Female)
	invalidNewMember.ChangeRelationType(family.RelDaughter)
	fam.AddMember(invalidNewMember)
	allErrors := fam.Err()

	require.False(t, fam.IsValid())
	// require.Contains(t, allErrors, family.ErrInvalidMember.Error())
	require.Contains(t, allErrors, family.ErrMemberError.Error())
	require.Contains(t, allErrors, family.ErrInvalidMemberID.Error())
	require.Equal(t, 4, len(allErrors))
}
