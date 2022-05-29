package familyRepository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
	familyRepository "github.com/renatospaka/emr/infrastructure/repository/family"
)

var (
	dob = time.Date(1970, 14, 11, 0, 0, 0, 0, time.UTC)
)

func TestMember_Add(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, dob)

	err := famMember.Add(member)
	require.Nil(t, err)
	require.True(t, member.Valid)
}

func TestMember_FullName(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, dob)

	_ = famMember.Add(member)
	fullName := member.FullName(false)
	fullNameComplete := member.FullName(true)
	require.Equal(t, "Renato Costa Spakauskas", fullName)
	require.Equal(t, "Sr. Renato Costa Spakauskas", fullNameComplete)
	require.NotEqual(t, "Sra. Renato Costa Spakauskas", fullNameComplete)
}

func TestMember_Add_MissingName(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("", "Costa", "Spakauskas", family.Male, dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberName.Error())
}

func TestMember_Add_NameTooShort(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Re", "Costa", "Spakauskas", family.Male, dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberNameTooShort.Error())
}

func TestMember_Add_NameTooLong(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato12345678901234567890", "Costa", "Spakauskas", family.Male, dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberNameTooLong.Error())
}

func TestMember_Add_MissingLastName(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "", family.Male, dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberLastName.Error())
}

func TestMember_Add_LastNameTooShort(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Sp", family.Male, dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberLastNameTooShort.Error())
}

func TestMember_Add_LastNameTooLong(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas12345678901234567890", family.Male, dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberLastNameTooLong.Error())
}

func TestMember_Add_MissingDOB(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, time.Time{})

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberDOB.Error())
}

func TestMember_Add_MissingGender(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", "", dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMembeGender.Error())
}

func TestMember_Add_InvalidGender(t *testing.T) {
	famMember := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", "Outro", dob)

	err := famMember.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrInvalidMembeGender.Error())
}
