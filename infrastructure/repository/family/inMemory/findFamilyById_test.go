package familyInMemory_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
	"github.com/renatospaka/emr/infrastructure/utils"
)

func TestFamily_FindById(t *testing.T) {
	_ = repoFam.Create(fam)

	id := fam.ID
	find, err := repoFam.FindById(id)
	require.Nil(t, err)
	require.Equal(t, id, find.ID)
}

func TestFamily_FindById_NotFound(t *testing.T) {
	_ = repoFam.Create(fam)

	id := utils.GetID()
	find, err := repoFam.FindById(id)
	require.NotNil(t, err)
	require.EqualError(t, err, family.ErrFamilyNotFound.Error())
	require.Equal(t, &family.Family{}, find)
}
