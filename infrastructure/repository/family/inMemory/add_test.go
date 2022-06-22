package familyInMemory_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
)

func TestMember_Add(t *testing.T) {
	err := repoMember.Add(member)
	require.Nil(t, err)
	require.True(t, member.Valid)
}

func TestMember_Add_MissingName(t *testing.T) {
	member = family.NewMember("", "MiddleName", "Lastname", family.Male, dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberName.Error())
}

func TestMember_Add_NameTooShort(t *testing.T) {
	member = family.NewMember("Re", "MiddleName", "Lastname", family.Male, dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberNameTooShort.Error())
}

func TestMember_Add_NameTooLong(t *testing.T) {
	member = family.NewMember("Name12345678901234567890", "MiddleName", "Lastname", family.Male, dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberNameTooLong.Error())
}

func TestMember_Add_MissingLastName(t *testing.T) {
	member = family.NewMember("Name", "MiddleName", "", family.Male, dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberLastName.Error())
}

func TestMember_Add_LastNameTooShort(t *testing.T) {
	member = family.NewMember("Name", "MiddleName", "Sp", family.Male, dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberLastNameTooShort.Error())
}

func TestMember_Add_LastNameTooLong(t *testing.T) {
	member = family.NewMember("Name", "MiddleName", "Lastname12345678901234567890", family.Male, dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberLastNameTooLong.Error())
}

func TestMember_Add_MissingDOB(t *testing.T) {
	member = family.NewMember("Name", "MiddleName", "Lastname", family.Male, time.Time{})
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberDOB.Error())
}

func TestMember_Add_MissingGender(t *testing.T) {
	member = family.NewMember("Name", "MiddleName", "Lastname", "", dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMembeGender.Error())
}

func TestMember_Add_InvalidGender(t *testing.T) {
	member = family.NewMember("Name", "MiddleName", "Lastname", "Outro", dob)
	err := repoMember.Add(member)

	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrInvalidMembeGender.Error())
}

func TestMember_FullName(t *testing.T) {
	_ = repoMember.Add(member)

	fullName := member.FullName()
	require.EqualValues(t, "Name MiddleName Lastname", fullName)
}

func TestMember_FullName_Formal(t *testing.T) {
	_ = repoMember.Add(member)

	fullNameFormal := member.FullNameFormal()
	require.EqualValues(t, "Sr. Name MiddleName Lastname", fullNameFormal)
}
