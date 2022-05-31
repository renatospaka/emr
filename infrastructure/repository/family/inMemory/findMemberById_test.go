package familyInMemory_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
)

func TestMember_FindById(t *testing.T) {
	_ = repoMember.Add(member)

	id := member.ID
	find, err := repoMember.FindById(id)
	require.Nil(t, err)
	require.True(t, find.Valid)
	require.Equal(t, id, find.ID)
}

func TestMember_FindById_NotFound(t *testing.T) {
	_ = repoMember.Add(member)

	id := uuid.NewV4()
	find, err := repoMember.FindById(id)
	require.NotNil(t, err)
	require.EqualError(t, err, family.ErrMemberNotFound.Error())
	require.Equal(t, &family.Member{}, find)
}
