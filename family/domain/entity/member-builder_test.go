package family_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
)

var testMemberBuilder *family.MemberBuilder

func init() {
	testMemberBuilder = family.NewMemberBuilder()
}

func TestMemberBuilder_Build(t *testing.T) {
	member := testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()

	require.IsTypef(t, &family.Member{}, member, "não é do tipo *Member{}")
	require.True(t, member.IsValid())
	require.Empty(t, member.Err())
}

func TestMemberBuilder_Build_Invalid(t *testing.T) {
	member := testMemberBuilder.
		WithFullName("Na", "Middle", "LastLastLastLastLastLastLast").
		WithBirthDate(time.Time{}).
		WithGender("other").
		WithNickname("Nick").
		Build()

	require.IsTypef(t, &family.Member{}, member, family.ErrMemberError.Error())
	require.False(t, member.IsValid())
	require.NotEmpty(t, member.Err())
	require.Len(t, member.Err(), 4)
}
