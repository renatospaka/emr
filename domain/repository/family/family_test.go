package familyRepository_test

import (
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

type Family struct {
	ID      uuid.UUID
	Surname string
}

type familyRepositoryInMemory struct {
	family []Family
}

func NewFamily(surname string) *Family {
	return &Family{
		ID: uuid.NewV4(),
		Surname: surname,	
	}
}

func (f *familyRepositoryInMemory) Create(family *Family) error {
	f.family = append(f.family, *family)
	if len(f.family) > 0 {return nil}
	return errors.New("Ocorreu um erro na gravação da nova família")
}


func TestFamily_Create(t *testing.T) {
	var repo familyRepositoryInMemory 
	fam := NewFamily("Essa Família")

	err := repo.Create(fam)
	require.Nil(t, err)
}
