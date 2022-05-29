package familyRepository_test

import (
	// "errors"
	"testing"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
	repoFamily "github.com/renatospaka/emr/infrastructure/repository/family"
)

func TestFamily_Create(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("Essa Família")

	err := repo.Create(fam)
	require.Nil(t, err)
	require.Equal(t, fam.Surname, "Essa Família")
}
