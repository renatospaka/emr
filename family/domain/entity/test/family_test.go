package family_test

import (
	"testing"
	"github.com/stretchr/testify/require"

	// family "github.com/renatospaka/emr/family/domain/entity"
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

// func TestFamily_ID(t *testing.T) {
// 	fam := familyBuilder.
// 		WithSurname("Super Family").
// 		Build()

// 	id := fam.ID()
// 	err := utils.IsVaalidUUID(id)

// 	require.IsType(t, "", id)
// 	require.Len(t, id, 36)
// 	require.Nil(t, err)
// }

// func TestFamily_IsValid_No(t *testing.T) {
// 	fam := familyBuilder.
// 		WithSurname("Super Family").
// 		Add(member).
// 		Build()
// 	fam.ChangeSurname("")

// 	require.False(t, fam.IsValid())
// 	require.NotEmpty(t, fam.Err())
// }

// func TestFamily_Surname(t *testing.T) {
// 	fam := familyBuilder.
// 		WithSurname("Super Family").
// 		Add(member).
// 		Build()
// 	fam.ChangeSurname("Another Surname")
// 	surname := fam.Surname()

// 	require.EqualValues(t, "Another Surname", surname)
// }

// func TestFamily_Size(t *testing.T) {
// 	fam := familyBuilder.
// 		WithSurname("Super Family").
// 		Add(member).
// 		Build()
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	size := fam.Size()

// 	require.EqualValues(t, 1, size)
// }

// func TestFamily_Size_Missing(t *testing.T) {
// 	fam := familyBuilder.
// 		WithSurname("Super Family").
// 		Add(member).
// 		Build()
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	fam.IsValid()
// 	size := fam.Size()

// 	allErrors := fam.Err()
// 	require.EqualValues(t, 1, size)
// 	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
// 	// require.Contains(t, allErrors, family.ErrFamilyMemberHOFMissing.Error())
// 	require.Equal(t, 1, len(allErrors))
// }

// func TestFamily_HasHeadOfFamily(t *testing.T) {
// 	fam := familyBuilder.
// 		WithSurname("Super Family").
// 		Add(member).
// 		Build()
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	hasHOF := fam.HasHeadOfFamily()

// 	require.True(t, hasHOF)
// }

// func TestFamily_HasHeadOfFamily_Missing(t *testing.T) {
// 	fam := familyBuilder.
// 		WithSurname("Super Family").
// 		Add(member).
// 		Build()
// 	//Need to prepare to include a member
// 	//This test will fail until then
// 	fam.IsValid()
// 	hasHOF := fam.HasHeadOfFamily()

// 	allErrors := fam.Err()
// 	require.False(t, hasHOF)
// 	require.Contains(t, allErrors, family.ErrFamilyMemberMissing.Error())
// 	// require.Contains(t, allErrors, family.ErrFamilyMemberHOFMissing.Error())
// 	require.Equal(t, 1, len(allErrors))
// }
