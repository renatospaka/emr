package familyRepository_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
	repoFamily "github.com/renatospaka/emr/infrastructure/repository/family"
)

func TestFamily_Create(t *testing.T) {
	repo := repoFamily.NewFamilyRepositoryInMemory()
	fam := family.NewFamily("Dino da Silva")

	err := repo.Create(fam)
	require.Nil(t, err)
	require.Equal(t, "Dino da Silva", fam.Surname)
	require.True(t, fam.Valid)

	err2 := fam.IsValid()
	require.Nil(t, err2)
}

func TestFamily_Create_Surname_Missing(t *testing.T) {
	repo := repoFamily.NewFamilyRepositoryInMemory()
	fam := family.NewFamily("")

	err := repo.Create(fam)
	require.NotNil(t, err)
	require.Equal(t, "", fam.Surname)
	require.False(t, fam.Valid)
	require.EqualError(t, err, family.ErrMissingFamilySurname.Error())
}

func TestFamily_FindById(t *testing.T) {
	repo := repoFamily.NewFamilyRepositoryInMemory()
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := fam.ID
	find, err := repo.FindById(id)
	require.Nil(t, err)
	require.Equal(t, id, find.ID)
}

func TestFamily_FindById_NotFound(t *testing.T) {
	repo := repoFamily.NewFamilyRepositoryInMemory()
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := uuid.NewV4()
	find, err := repo.FindById(id)
	require.NotNil(t, err)
	require.EqualError(t, err, family.ErrFamilyNotFound.Error())
	require.Equal(t, &family.Family{}, find)
}

func TestFamily_SetFamilyName(t *testing.T) {
	repo := repoFamily.NewFamilyRepositoryInMemory()
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := fam.ID
	updSurname := "Dino da Silva Sauro"
	upd, err := repo.SetFamilyName(id, updSurname)
	require.Nil(t, err)
	require.True(t, upd.Valid)
	require.Equal(t, "Dino da Silva Sauro", upd.Surname)
}

func TestFamily_SetFamilyName_NotFound(t *testing.T) {
	repo := repoFamily.NewFamilyRepositoryInMemory()
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := uuid.NewV4()
	updSurname := "Dino da Silva Sauro"
	upd, err := repo.SetFamilyName(id, updSurname)
	require.NotNil(t, err)
	require.NotEqual(t, "Dino da Silva Sauro", upd.Surname)
	require.EqualError(t, err, family.ErrFamilyNotFound.Error())
	require.Equal(t, &family.Family{}, upd)
}

func TestFamily_SetFamilyName_Surname_Missing(t *testing.T) {
	repo := repoFamily.NewFamilyRepositoryInMemory()
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := fam.ID
	updSurname := ""
	upd, err := repo.SetFamilyName(id, updSurname)
	require.NotNil(t, err)
	require.EqualError(t, err, family.ErrMissingFamilySurname.Error())
	require.False(t, upd.Valid)
	require.Equal(t, "", upd.Surname)
}
