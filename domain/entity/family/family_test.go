package family_test

import (
	"testing"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr2/domain/entity/family"
)

var (
	newFamily = family.NewFamily("Surname")
)

func TestFamily_IsValid(t *testing.T) {
	newFamily.SetSurname("Surname")

	require.True(t, newFamily.IsValid())
	require.Empty(t, newFamily.Err())
}