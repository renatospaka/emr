package familyInMemory_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
	"github.com/renatospaka/emr/infrastructure/utils"
)

func TestFamily_SetFamilyName(t *testing.T) {
	_ = repoFam.Create(fam)

	id := fam.ID
	updSurname := "MiddleName Lastname"
	upd, err := repoFam.SetFamilyName(id, updSurname)
	require.Nil(t, err)
	require.True(t, upd.Valid)
	require.Equal(t, "MiddleName Lastname", upd.Surname)
}

func TestFamily_ChangeName_NotFound(t *testing.T) {
	_ = repoFam.Create(fam)

	id := utils.GetID()
	updSurname := "MiddleName Lastname"
	upd, err := repoFam.SetFamilyName(id, updSurname)
	require.NotNil(t, err)
	require.NotEqual(t, "MiddleName Lastname", upd.Surname)
	require.EqualError(t, err, family.ErrFamilyNotFound.Error())
	require.Equal(t, &family.Family{}, upd)
}

func TestFamily_ChangeName_Surname_Missing(t *testing.T) {
	_ = repoFam.Create(fam)

	id := fam.ID
	updSurname := ""
	upd, err := repoFam.SetFamilyName(id, updSurname)
	require.NotNil(t, err)
	require.EqualError(t, err, family.ErrMissingFamilySurname.Error())
	require.False(t, upd.Valid)
	require.Equal(t, "", upd.Surname)
}
