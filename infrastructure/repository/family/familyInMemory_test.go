package familyRepository_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
	repoFamily "github.com/renatospaka/emr/infrastructure/repository/family"
)

func TestFamily_Create(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("Dino da Silva")

	err := repo.Create(fam)
	require.Nil(t, err)
	require.Equal(t, "Dino da Silva", fam.Surname)
	
	err2 := fam.IsValid()
	require.Nil(t, err2)
}

func TestFamily_Create_Surname_Missing(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("")
	_ = repo.Create(fam)

	err := fam.IsValid()
	require.NotNil(t, err)
	require.Equal(t, "", fam.Surname)
	require.Error(t, err, "o nome de família está em branco ou ausente")
}

func TestFamily_FindById(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := fam.ID
	find, err := repo.FindById(id)
	require.Nil(t, err)
	require.Equal(t, id, find.ID)
}

func TestFamily_FindById_NotFound(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := uuid.NewV4()
	find, err := repo.FindById(id)
	require.NotNil(t, err)
	require.Error(t, err, "família não encontradaa")
	require.Equal(t, &family.Family{}, find)
}

func TestFamily_ChangeName(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := fam.ID
	updSurname := "Dino da Silva Sauro"
	upd, err := repo.ChangeName(id, updSurname)
	require.Nil(t, err)
	require.Equal(t, "Dino da Silva Sauro", upd.Surname)
}

func TestFamily_ChangeName_NotFound(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := uuid.NewV4()
	updSurname := "Dino da Silva Sauro"
	upd, err := repo.ChangeName(id, updSurname)
	require.NotNil(t, err)
	require.NotEqual(t, "Dino da Silva Sauro", upd.Surname)
	require.Error(t, err, "família não encontrada")
	require.Equal(t, &family.Family{}, upd)
}

func TestFamily_ChangeName_Surname_Missing(t *testing.T) {
	var repo repoFamily.FamilyRepositoryInMemory 
	fam := family.NewFamily("Dino da Silva")
	_ = repo.Create(fam)

	id := fam.ID
	updSurname := ""
	upd, _ := repo.ChangeName(id, updSurname)
	err := upd.IsValid()
	require.NotNil(t, err)
	require.Error(t, err, "o nome de família está em branco ou ausente")
	require.Equal(t, "", upd.Surname)
}
