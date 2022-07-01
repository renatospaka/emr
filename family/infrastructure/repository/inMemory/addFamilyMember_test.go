package familyInMemory_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	family "github.com/renatospaka/emr/family/domain/entity"
// )

// // var (
// // 	repoFamMemb = repoFamily.NewFamilyMemberRepositoryInMemory()
// // )

// func TestFamily_AddFamilyMember(t *testing.T) {
// 	_ = repoFam.Create(fam)
// 	newMember := family.NewFamilyMember(member)
// 	newFamily, err := repoFam.AddFamilyMember(*newMember, fam.ID)

// 	require.Nil(t, err)
// 	require.Len(t, newFamily.Members, 1)
// }

// // func TestFamily_AddFamilyMember_MemberInvalid(t *testing.T) {
// // 	_ = repoFam.Create(fam)
// // 	newMember := family.NewFamilyMember()
// // 	newMember.Member.Name = ""
// // 	newFamily, err := repoFamMemb.AddFamilyMember(newMember, newFamily.ID)

// // 	require.NotNil(t, err)
// // 	require.Len(t, newFamily.Members, 0)
// // 	require.EqualError(t, err, family.ErrMissingMemberName.Error())
// // }
