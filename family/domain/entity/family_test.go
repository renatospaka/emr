package family_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	family "github.com/renatospaka/emr/family/domain/entity"
// 	"github.com/renatospaka/emr/infrastructure/utils"
// )

// var (
// 	fam = family.newFamily()
// )

// func TestFamily_ID(t *testing.T) {
// 	id := fam.ID()
// 	err := utils.IsVaalidUUID(id)

// 	require.IsType(t, "", id)
// 	require.Len(t, id, 36)
// 	require.Nil(t, err)
// }

// func TestFamily_IsValid(t *testing.T) {
// 	fam.SetSurname("Surname")

// 	require.True(t, fam.IsValid())
// 	require.Empty(t, fam.Err())
// }

// func TestFamily_IsValid_No(t *testing.T) {
// 	fam.SetSurname("")

// 	require.False(t, fam.IsValid())
// 	require.NotEmpty(t, fam.Err())
// }

// func TestFamily_Surname(t *testing.T) {
// 	fam.SetSurname("Another Surname")
// 	surname := fam.Surname()

// 	require.EqualValues(t, "Another Surname", surname)
// }

// func TestFamily_HasMembers(t *testing.T) {
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	hasMembers := fam.HasMembers()

// 	require.True(t, hasMembers)
// }

// func TestFamily_HasMembers_Missing(t *testing.T) {
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	fam.IsValid()
// 	hasMembers := fam.HasMembers()

// 	allErrors := fam.ErrToArray()
// 	require.False(t, hasMembers)
// 	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
// 	require.Contains(t, allErrors, family.ErrFamilyMemberHOFMissing.Error())
// 	require.Equal(t, 2, len(allErrors))
// }

// func TestFamily_HasHeadOfFamily(t *testing.T) {
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	hasHOF := fam.HasHeadOfFamily()

// 	require.True(t, hasHOF)
// }

// func TestFamily_HasHeadOfFamily_Missing(t *testing.T) {
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	fam.IsValid()
// 	hasHOF := fam.HasHeadOfFamily()

// 	allErrors := fam.ErrToArray()
// 	require.False(t, hasHOF)
// 	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
// 	require.Contains(t, allErrors, family.ErrFamilyMemberHOFMissing.Error())
// 	require.Equal(t, 2, len(allErrors))
// }
