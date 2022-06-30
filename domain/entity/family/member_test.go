package family_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr2/domain/entity/family"
)

var (
	newMember  = &family.Member{}
	dobNewborn = time.Time{}
	dobInfant  = time.Time{}
	dobToddler = time.Time{}
	dobChild   = time.Time{}
	dobTeen    = time.Time{}
	dobAdult   = time.Time{}
	dobElderly = time.Time{}
)

func init() {
	newMember = family.NewMember("Name", "Middle", "Last")

	today := time.Now()

	dobNewborn = time.Date(today.Year(), today.Month(), today.Day()-33, 0, 0, 0, 0, time.UTC)
	dobInfant = time.Date(today.Year(), today.Month()-11, today.Day(), 0, 0, 0, 0, time.UTC)
	dobToddler = time.Date(today.Year(), today.Month()-54, today.Day(), 0, 0, 0, 0, time.UTC)
	dobChild = time.Date(today.Year()-7, today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	dobTeen = time.Date(today.Year()-15, today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	dobAdult = time.Date(today.Year()-33, today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	dobElderly = time.Date(today.Year()-67, today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
}

func TestMember_IsValid(t *testing.T) {
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

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
	newMember.SetBirthDate(dobAdult)

	require.Equal(t, newMember.FullName(), "Name Middle Last")
	require.Empty(t, newMember.Err())
}

func TestMember_FullNameFormal(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")

	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.Equal(t, newMember.FullNameFormal(), "Sr. Name Middle Last")
	require.Empty(t, newMember.Err())
}

func TestMember_SetFullName_NameMissing(t *testing.T) {
	newMember.SetFullName("", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberName.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_NameTooShort(t *testing.T) {
	newMember.SetFullName("Na", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooShort.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_NameTooLong(t *testing.T) {
	newMember.SetFullName("NameNameNameNameNameNameNameNameName", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooLong.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_LastNameMissing(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberLastName.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_LastNameTooShort(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "La")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooShort.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_LastNameTooLong(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "LastLastLastLastLastLastLastLastLast")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooLong.Error()+"\n", newMember.Err())
}

func TestMember_SetFullName_MiddleNameTooLong(t *testing.T) {
	newMember.SetFullName("Name", "MiddleMiddleMiddleMiddleMiddleMiddleMiddle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMemberMiddleNameTooLong.Error()+"\n", newMember.Err())
}

func TestMember_SetGender(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.Equal(t, newMember.Gender(), family.Male)
	require.Empty(t, newMember.Err())
}

func TestMember_SetGenderMissing(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender("")
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberGender.Error()+"\n", newMember.Err())
}

func TestMember_SetGenderInvalid(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender("other")
	newMember.SetBirthDate(dobAdult)

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrInvalidMemberGender.Error()+"\n", newMember.Err())
}

func TestMember_SetBirthDate(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	require.True(t, newMember.IsValid())
	require.EqualValues(t, dobAdult, newMember.BirthDate())
}

func TestMember_SetBirthDateMissing(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(time.Time{})

	require.False(t, newMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberDOB.Error()+"\n", newMember.Err())
}

func TestMember_IsNewborn(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobNewborn)

	ageInMonths := newMember.AgeInMonths()
	require.EqualValues(t, 1, ageInMonths)

	ageInYears := newMember.AgeInYears()
	require.EqualValues(t, 0, ageInYears)

	ageT := newMember.IsNewborn()
	ageF := newMember.IsInfant()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsInfant(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobInfant)

	ageInMonths := newMember.AgeInMonths()
	require.EqualValues(t, 11, ageInMonths)

	ageInYears := newMember.AgeInYears()
	require.EqualValues(t, 1, ageInYears)

	ageT := newMember.IsInfant()
	ageF := newMember.IsToddler()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsToddler(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobToddler)

	ageInMonths := newMember.AgeInMonths()
	require.EqualValues(t, 55, ageInMonths)

	ageInYears := newMember.AgeInYears()
	require.EqualValues(t, 5, ageInYears)

	ageT := newMember.IsToddler()
	ageF := newMember.IsChild()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsChild(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobChild)

	ageInMonths := newMember.AgeInMonths()
	require.EqualValues(t, 85, ageInMonths)

	ageInYears := newMember.AgeInYears()
	require.EqualValues(t, 7, ageInYears)

	ageT := newMember.IsChild()
	ageF := newMember.IsTeen()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsTeen(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobTeen)

	ageInMonths := newMember.AgeInMonths()
	require.EqualValues(t, 183, ageInMonths)

	ageInYears := newMember.AgeInYears()
	require.EqualValues(t, 15, ageInYears)

	ageT := newMember.IsTeen()
	ageF := newMember.IsAdult()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsAdult(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobAdult)

	ageInMonths := newMember.AgeInMonths()
	require.EqualValues(t, 402, ageInMonths)

	ageInYears := newMember.AgeInYears()
	require.EqualValues(t, 33, ageInYears)

	ageT := newMember.IsAdult()
	ageF := newMember.IsElderly()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsElderly(t *testing.T) {
	newMember.SetFullName("Name", "Middle", "Last")
	newMember.SetGender(family.Male)
	newMember.SetBirthDate(dobElderly)

	ageInMonths := newMember.AgeInMonths()
	require.EqualValues(t, 816, ageInMonths)

	ageInYears := newMember.AgeInYears()
	require.EqualValues(t, 68, ageInYears)

	ageT := newMember.IsElderly()
	ageF := newMember.IsNewborn()
	require.True(t, ageT)
	require.False(t, ageF)
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