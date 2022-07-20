package family_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
)

func TestMemberBuilder_Build(t *testing.T) {
	member := createMember()

	require.IsTypef(t, &family.Member{}, member, "não é do tipo *Member{}")
	require.True(t, member.IsValid())
	require.Len(t, member.Err(), 0)
}

func TestMemberBuilder_Build_Invalid(t *testing.T) {
	member := createInvalidMember()

	require.IsTypef(t, &family.Member{}, member, family.ErrMemberError.Error())
	require.False(t, member.IsValid())
	require.NotEmpty(t, member.Err())
	require.Len(t, member.Err(), 4)
}
