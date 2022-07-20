package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	family "github.com/renatospaka/emr/family/domain/entity"
)


func TestFamilyBuilder_Build(t *testing.T) {
	hof := createHOFMember()

	famMembBuilder := family.NewFamilyMemberBuilder()
	famMember := famMembBuilder.
		AsHOF(hof).
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
