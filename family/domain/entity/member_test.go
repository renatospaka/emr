package family_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/renatospaka/emr/infrastructure/utils"
)

var (
	member     = &family.Member{}
	dobNewborn = time.Time{}
	dobInfant  = time.Time{}
	dobToddler = time.Time{}
	dobChild   = time.Time{}
	dobTeen    = time.Time{}
	dobAdult   = time.Time{}
	dobElderly = time.Time{}
)

func init() {
	testMemberBuilder = family.NewMemberBuilder()
	member = testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		Build()
	today := time.Now()
	dobNewborn = today.Add(-15 * (utils.HoursDay) * time.Hour)
	// dobInfant = time.Date(today.Year(), today.Month()-11, today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobInfant = today.Add(-11 * utils.HoursMonth * time.Hour)
	// dobToddler = time.Date(today.Year(), today.Month()-54, today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobToddler = today.Add(-54 * utils.HoursMonth * time.Hour)
	// dobChild = time.Date(today.Year()-7, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobChild = today.Add(-7 * utils.HoursYear * time.Hour)
	// dobTeen = time.Date(today.Year()-15, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobTeen = today.Add(-15 * utils.HoursYear * time.Hour)
	// dobAdult = time.Date(today.Year()-33, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobAdult = today.Add(-33 * utils.HoursYear * time.Hour)
	// dobElderly = time.Date(today.Year()-67, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobElderly = today.Add(-67 * utils.HoursYear * time.Hour)
}

func TestMember_ID(t *testing.T) {
	id := member.ID()
	err := utils.IsVaalidUUID(id)
	valid := err == nil

	require.True(t, valid)
	require.NotEmpty(t, id)
	require.Nil(t, err)
}

func TestMember_IsValid(t *testing.T) {
	member.SetBirthDate(dobChild)

	require.True(t, member.IsValid())
	require.Empty(t, member.Err())
}

func TestMember_InValid_No(t *testing.T) {
	member.SetFullName("", "Middle", "Last")

	require.False(t, member.IsValid())
	require.NotEmpty(t, member.Err())
}

func TestMember_FullName(t *testing.T) {
	member.
		SetFullName("Name2", "Middle2", "Last2").
		SetBirthDate(dobAdult).
		IsValid()

	require.Equal(t, member.FullName(), "Name2 Middle2 Last2")
	require.Empty(t, member.Err())
}

func TestMember_FullNameFormal(t *testing.T) {
	member.
		SetFullName("Name2", "Middle2", "Last2").
		SetBirthDate(dobAdult).
		IsValid()

	require.Equal(t, member.FullNameFormal(), "Sr. Name2 Middle2 Last2")
	require.Empty(t, member.Err())
}

func TestMember_SetFullName_NameMissing(t *testing.T) {
	member.
		SetFullName("", "Middle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMissingMemberName.Error()+"\n", member.Err())
}

func TestMember_SetFullName_NameTooShort(t *testing.T) {
	member.
		SetFullName("Na", "Middle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooShort.Error()+"\n", member.Err())
}

func TestMember_SetFullName_NameTooLong(t *testing.T) {
	member.
		SetFullName("NameNameNameNameNameNameNameNameName", "Middle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooLong.Error()+"\n", member.Err())
}

func TestMember_SetFullName_LastNameMissing(t *testing.T) {
	member.
		SetFullName("Name", "Middle", "").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMissingMemberLastName.Error()+"\n", member.Err())
}

func TestMember_SetFullName_LastNameTooShort(t *testing.T) {
	member.
		SetFullName("Name", "Middle", "La").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooShort.Error()+"\n", member.Err())
}

func TestMember_SetFullName_LastNameTooLong(t *testing.T) {
	member.
		SetFullName("Name", "Middle", "LastLastLastLastLastLastLastLastLast").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooLong.Error()+"\n", member.Err())
}

func TestMember_SetFullName_MiddleNameTooLong(t *testing.T) {
	member.
		SetFullName("Name", "MiddleMiddleMiddleMiddleMiddleMiddleMiddle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMemberMiddleNameTooLong.Error()+"\n", member.Err())
}

func TestMember_SetNickname(t *testing.T) {
	member.
		SetFullName("Name", "Middle", "Last").
		SetBirthDate(dobAdult).
		SetNickname("Nickname")

	require.True(t, member.IsValid())
}

func TestMember_Nickname(t *testing.T) {
	member = testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
	member.
		SetNickname("Nickname").
		IsValid()
	nick := member.Nickname()

	require.EqualValues(t, "Nickname", nick)
}

func TestMember_SetGender(t *testing.T) {
	member.
		SetGender(family.Female).
		SetBirthDate(dobAdult).
		IsValid()

	require.EqualValues(t, family.Female, member.Gender())
	require.Empty(t, member.Err())
}

func TestMember_SetGender_Missing(t *testing.T) {
	member.
		SetGender("").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMissingMemberGender.Error()+"\n", member.Err())
}

func TestMember_SetGender_Invalid(t *testing.T) {
	member.
		SetGender("other").
		SetBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrInvalidMemberGender.Error()+"\n", member.Err())
}

func TestMember_Gender(t *testing.T) {
	member.
		SetGender(family.Female).
		SetBirthDate(dobAdult).
		IsValid()
	gender := member.Gender()

	require.EqualValues(t, family.Female, gender)
}

func TestMember_SetBirthDate(t *testing.T) {
	member.SetBirthDate(dobElderly)

	require.True(t, member.IsValid())
	require.EqualValues(t, dobElderly, member.BirthDate())
}

func TestMember_SetBirthDate_Missing(t *testing.T) {
	member.SetBirthDate(time.Time{})

	require.False(t, member.IsValid())
	require.EqualValues(t, family.ErrMissingMemberDOB.Error()+"\n", member.Err())
}

func TestMember_BirthDate(t *testing.T) {
	member.SetBirthDate(dobToddler)
	dob := member.BirthDate()

	require.EqualValues(t, dobToddler, dob)
}

func TestMember_AgeInMonths(t *testing.T) {
	member.SetBirthDate(dobNewborn)
	ageInMonths := member.AgeInMonths()

	require.EqualValues(t, 0, ageInMonths)
}

func TestMember_AgeInYears(t *testing.T) {
	member.SetBirthDate(dobNewborn)
	ageInYears := member.AgeInYears()

	require.EqualValues(t, 0, ageInYears)
}

func TestMember_IsNewborn(t *testing.T) {
	member.SetBirthDate(dobNewborn)

	ageT := member.IsNewborn()
	ageF := member.IsInfant()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsInfant(t *testing.T) {
	member.SetBirthDate(dobInfant)

	ageT := member.IsInfant()
	ageF := member.IsToddler()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsToddler(t *testing.T) {
	member.SetBirthDate(dobToddler)

	ageT := member.IsToddler()
	ageF := member.IsChild()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsChild(t *testing.T) {
	member.SetBirthDate(dobChild)

	ageT := member.IsChild()
	ageF := member.IsTeen()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsTeen(t *testing.T) {
	member.SetBirthDate(dobTeen)

	ageT := member.IsTeen()
	ageF := member.IsAdult()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsAdult(t *testing.T) {
	member.SetBirthDate(dobAdult)

	ageT := member.IsAdult()
	ageF := member.IsElderly()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsElderly(t *testing.T) {
	member.	SetBirthDate(dobElderly)

	ageT := member.IsElderly()
	ageF := member.IsNewborn()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_MoreThanOneError(t *testing.T) {
	member = testMemberBuilder.
		WithFullName("Na", "Middle", "LastLastLastLastLastLastLastLastLast").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		Build()
	member.
		SetBirthDate(time.Time{}).
		SetGender("other").
		IsValid()

	allErrors := member.ErrToArray()
	require.False(t, member.IsValid())
	require.Contains(t, allErrors, family.ErrMemberNameTooShort.Error())
	require.Contains(t, allErrors, family.ErrMemberLastNameTooLong.Error())
	require.Contains(t, allErrors, family.ErrInvalidMemberGender.Error())
	require.Contains(t, allErrors, family.ErrMissingMemberDOB.Error())
	require.Equal(t, 4, len(allErrors))
}
