package familyInMemory_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	family "github.com/renatospaka/emr/family/domain/entity"
// )

// func TestFamily_Create(t *testing.T) {
// 	err := repoFam.Create(fam)
// 	require.Nil(t, err)
// 	require.Equal(t, "Lastname", fam.Surname)
// 	require.True(t, fam.Valid)

// 	err2 := fam.IsValid()
// 	require.Nil(t, err2)
// }

// func TestFamily_Create_Surname_Missing(t *testing.T) {
// 	famBlank := family.NewFamily("")
// 	err := repoFam.Create(famBlank)
// 	require.NotNil(t, err)
// 	require.Equal(t, "", famBlank.Surname)
// 	require.False(t, famBlank.Valid)
// 	require.EqualError(t, err, family.ErrMissingFamilySurname.Error())
// }
