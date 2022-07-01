package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/renatospaka/emr/infrastructure/utils"
)

var (
	newFamily = family.NewFamily("Surname")
)

func TestFamily_ID(t *testing.T) {
	id := newFamily.ID()
	err := utils.IsVaalidUUID(id)

	require.IsType(t, "", id)
	require.Len(t, id, 36)
	require.Nil(t, err)
}

func TestFamily_IsValid(t *testing.T) {
	newFamily.SetSurname("Surname")

	require.True(t, newFamily.IsValid())
	require.Empty(t, newFamily.Err())
}

func TestFamily_IsValid_No(t *testing.T) {
	newFamily.SetSurname("")

	require.False(t, newFamily.IsValid())
	require.NotEmpty(t, newFamily.Err())
}

func TestFamily_Surname(t *testing.T) {
	newFamily.SetSurname("Another Surname")
	surname := newFamily.Surname()

	require.EqualValues(t, "Another Surname", surname)
}

func TestFamily_HasMembers(t *testing.T) {
	//Need to prepare to include a member
	//This test will fail until then
	hasMembers := newFamily.HasMembers()

	require.True(t, hasMembers)
}

func TestFamily_HasMembers_Missing(t *testing.T) {
	//Need to prepare to include a member
	//This test will fail until then
	newFamily.IsValid()
	hasMembers := newFamily.HasMembers()

	allErrors := newFamily.ErrToArray()
	require.False(t, hasMembers)
	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
	require.Contains(t, allErrors, family.ErrFamilyMemberHeadMissing.Error())
	require.Equal(t, 2, len(allErrors))
}

func TestFamily_HasHeadOfFamily(t *testing.T) {
	//Need to prepare to include a member
	//This test will fail until then
	hasHOF := newFamily.HasHeadOfFamily()

	require.True(t, hasHOF)
}

func TestFamily_HasHeadOfFamily_Missing(t *testing.T) {
	//Need to prepare to include a member
	//This test will fail until then
	newFamily.IsValid()
	hasHOF := newFamily.HasHeadOfFamily()

	allErrors := newFamily.ErrToArray()
	require.False(t, hasHOF)
	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
	require.Contains(t, allErrors, family.ErrFamilyMemberHeadMissing.Error())
	require.Equal(t, 2, len(allErrors))
}
