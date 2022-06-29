package family_test

import (
	"testing"
	"time"

	"github.com/renatospaka/emr2/domain/entity/family"
	"github.com/stretchr/testify/require"
)

var (
	newMember = family.NewMember("Name", "Middle", "Last")
	dob       = time.Date(1980, 23, 8, 0, 0, 0, 0, time.UTC)
)

func TestMember_IsValid(t *testing.T) {
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.True(t, newMember.IsValid())
	require.Empty(t, newMember.Err())
}

func TestMember_IsNotValid(t *testing.T) {
	newMember.SetFullName("", "Middle", "Last")

	require.False(t, newMember.IsValid())
	require.NotEmpty(t, newMember.Err())
}

func TestMember_FullName(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.Equal(t, newMember.FullName(), "Name Middle Last")
	require.Empty(t, newMember.Err())
}

func TestMember_FullNameFormal(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")

	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.Equal(t, newMember.FullNameFormal(), "Sr. Name Middle Last")
	require.Empty(t, newMember.Err())
}

func TestMember_SetFullName_NameMissing(t *testing.T) {
	newMember.SetFullName("", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberName.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_NameTooShort(t *testing.T) {
	newMember.SetFullName("Na", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooShort.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_NameTooLong(t *testing.T) {
	newMember.SetFullName("NameNameNameNameNameNameNameNameName", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooLong.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_LastNameMissing(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberLastName.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_LastNameTooShort(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "La")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooShort.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_LastNameTooLong(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "LastLastLastLastLastLastLastLastLast")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooLong.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_MiddleNameTooLong(t *testing.T) {
	newMember.SetFullName("Name", "MiddleMiddleMiddleMiddleMiddleMiddleMiddle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberMiddleNameTooLong.Error()+"\n", newMember.Err())
}

func TestMember_SetGender(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dob)

	require.Equal(t, newMember.Gender(), family.Male)
	require.Empty(t, newMember.Err())
}

func TestMember_SetGenderMissing(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender("")
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberGender.Error()+"\n", newMember.Err())
}

func TestMember_SetGenderInvalid(t *testing.T) {
	newMember.SetGender("other")
	newMember.SetBirthDate(dob)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrInvalidMemberGender.Error()+"\n", newMember.Err())
}

func TestMember_SetBirthDate(t *testing.T) {
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(time.Time{})

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberDOB.Error()+"\n", newMember.Err())
}

func TestMember_MoreThanOneError(t *testing.T) {
	newMember.SetFullName("Na", "Middle", "LastLastLastLastLastLastLastLastLast")
	newMember.SetGender("other")
	newMember.SetBirthDate(time.Time{})

	allErrors := newMember.ErrToArray()
	require.False(t, newMember.IsValid())
	require.Contains(t, allErrors, family.ErrMemberNameTooShort.Error())
	require.Contains(t, allErrors, family.ErrMemberLastNameTooLong.Error())
	require.Contains(t, allErrors, family.ErrInvalidMemberGender.Error())
	require.Contains(t, allErrors, family.ErrMissingMemberDOB.Error())
	require.Equal(t, 4, len(allErrors))
}
