package familyInMemory_test

import (
	"testing"

	"github.com/renatospaka/emr/domain/entity/family"
	"github.com/stretchr/testify/require"
)

func TestMember_Change(t *testing.T) {
	member = family.NewMember("Name", "", "Lastname", family.Male, dob)
	_ = repoMember.Add(member)

	member.MiddleName = "MiddleName"
	upd, err := repoMember.Change(member)
	require.Nil(t, err)
	require.True(t, upd.Valid)
	require.Equal(t, "MiddleName", upd.MiddleName)
}

func TestMember_Change_NotFound(t *testing.T) {
	member = family.NewMember("Name", "", "Lastname", family.Male, dob)
	_ = repoMember.Add(member)

	new := family.NewMember("Name", "", "Lastname", family.Male, dob)
	new.MiddleName = "MiddleName"
	upd, err := repoMember.Change(new)
	require.NotNil(t, err)
	require.True(t, upd.Valid)
	require.EqualError(t, err, family.ErrMemberNotFound.Error())
}

func TestMember_Change_NoChanges(t *testing.T) {
	member = family.NewMember("Name", "", "Lastname", family.Male, dob)
	_ = repoMember.Add(member)

	upd, err := repoMember.Change(member)
	require.NotNil(t, err)
	require.True(t, upd.Valid)
	require.EqualError(t, err, family.ErrNoChangesNeeded.Error())
	require.EqualValues(t, member, upd)
}
