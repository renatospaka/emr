package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
)

var (
	memberHOF       = family.NewMember("Name HOF", "", "Lastname")
	newFamilyMember = family.NewFamilyMember(memberHOF)
)

func TestFamilyMember_IsValid(t *testing.T) {
	memberHOF.SetBirthDate(dobAdult)
	memberHOF.SetGender(family.Male)
	// memberOK := memberHOF.IsValid()

	newFamily.SetSurname("Surname")
	// newFamilyMember.
		// familyOK := newFamily.IsValid()

	require.True(t, newFamilyMember.IsValid())
	require.Empty(t, newFamilyMember.Err())
}
