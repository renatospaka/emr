package familyInMemory_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	family "github.com/renatospaka/emr/family/domain/entity"
// 	"github.com/renatospaka/emr/infrastructure/utils"
// )

// func TestMember_FindById(t *testing.T) {
// 	_ = repoMember.Add(member)

// 	id := member.ID
// 	find, err := repoMember.FindById(id)
// 	require.Nil(t, err)
// 	require.True(t, find.Valid)
// 	require.Equal(t, id, find.ID)
// }

// func TestMember_FindById_NotFound(t *testing.T) {
// 	_ = repoMember.Add(member)

// 	id := utils.GetID()
// 	find, err := repoMember.FindById(id)
// 	require.NotNil(t, err)
// 	require.EqualError(t, err, family.ErrMemberNotFound.Error())
// 	require.Equal(t, &family.Member{}, find)
// }
