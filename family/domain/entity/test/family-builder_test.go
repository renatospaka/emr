package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	family "github.com/renatospaka/emr/family/domain/entity"
)

func init() {
	testFamilyBuilder = family.NewFamilyBuilder()
}

func TestFamilyBuilder_Build(t *testing.T) {
	memberBuilder := family.NewMemberBuilder()
	member := memberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Top Cat").
		Build()

	famMembBuilder := family.NewFamilyMemberBuilder()
	famMember := famMembBuilder.
		AsHOF(member).
		Build()
	
	famBuilder := family.NewFamilyBuilder()
	fam := famBuilder.
		WithSurname("Super Family").
		Add(famMember).
		Build()

	require.True(t, fam.IsValid())
	require.Empty(t, fam.Err())
	require.IsTypef(t, &family.Family{}, fam, "não é do tipo *Family{}")
}

// func TestFamilyBuilder_Build_Invalid(t *testing.T) {
// 	fam := testFamilyBuilder.
// 		WithSurname("Super Family").
// 		Add(&family.Member{}).
// 		Build()

// 	require.False(t, fam.IsValid())
// 	require.NotEmpty(t, fam.Err())
// }
