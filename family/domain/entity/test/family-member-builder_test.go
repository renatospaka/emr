package family_test

import (
	// "log"
	"testing"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/stretchr/testify/require"
)

func TestFamilyMember_Build_HOF(t *testing.T) {
	// log.Println("TestFamilyMember_Build_HOF.createHOFMember()")
	hof := createHOFMember()

	// log.Println("TestFamilyMember_Build_HOF.NewFamilyMemberBuilder()")
	famMembBuilder := family.NewFamilyMemberBuilder()
	// log.Println("TestFamilyMember_Build_HOF.famMembBuilder()")
	famMember := famMembBuilder.
		AsHOF(hof).
		Build()

	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
	require.IsTypef(t, &family.Member{}, famMember.Member, "não é do tipo *FamilyMember{}")
}

func TestFamilyMember_Build_HOF_Invalid(t *testing.T) {
	empty := createEmptyMember()
	famMembBuilder := family.NewFamilyMemberBuilder()
	famMember := famMembBuilder.
		AsHOF(empty).
		Build()

	require.False(t, famMember.IsValid())
	require.NotEmpty(t, famMember.Err())
}

func TestFamilyMember_Build_Ordinary(t *testing.T) {
	member := createMember()
	famMembBuilder := family.NewFamilyMemberBuilder()	
	famMember := famMembBuilder.
		AsOrdinary(member).
		RelatedAs(family.Wife).
		Build()

	require.True(t, famMember.IsValid())
	require.Empty(t, famMember.Err())
	require.IsTypef(t, &family.Member{}, famMember.Member, "não é do tipo *FamilyMember{}")
}
