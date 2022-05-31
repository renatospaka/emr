package familyInMemory_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
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

	id := uuid.NewV4()
	find, err := repoFam.FindById(id)
	require.NotNil(t, err)
	require.EqualError(t, err, family.ErrFamilyNotFound.Error())
	require.Equal(t, &family.Family{}, find)
}
