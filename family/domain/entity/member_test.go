package family_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/renatospaka/emr/infrastructure/utils"
)

var (
	testMember *family.Member
	dobNewborn time.Time
	dobInfant  time.Time
	dobToddler time.Time
	dobChild   time.Time
	dobTeen    time.Time
	dobAdult   time.Time
	dobElderly time.Time
)

func init() {
	testMemberBuilder = family.NewMemberBuilder()
	testMember = testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		Build()
	today := time.Now()
	dobNewborn = today.Add(-15 * (utils.HoursDay) * time.Hour)
	dobInfant = today.Add(-11 * utils.HoursMonth * time.Hour)
	dobToddler = today.Add(-54 * utils.HoursMonth * time.Hour)
	dobChild = today.Add(-7 * utils.HoursYear * time.Hour)
	dobTeen = today.Add(-15 * utils.HoursYear * time.Hour)
	dobAdult = today.Add(-33 * utils.HoursYear * time.Hour)
	dobElderly = today.Add(-67 * utils.HoursYear * time.Hour)
}

func TestMember_ID(t *testing.T) {
	id := testMember.ID()
	err := utils.IsVaalidUUID(id)
	valid := err == nil

	require.True(t, valid)
	require.NotEmpty(t, id)
	require.Nil(t, err)
}

func TestMember_IsValid(t *testing.T) {
	testMember.SetBirthDate(dobChild)

	require.True(t, testMember.IsValid())
	require.Empty(t, testMember.Err())
}

func TestMember_InValid_No(t *testing.T) {
	testMember.SetFullName("", "Middle", "Last")

	require.False(t, testMember.IsValid())
	require.NotEmpty(t, testMember.Err())
}

func TestMember_FullName(t *testing.T) {
	testMember.
		SetFullName("Name2", "Middle2", "Last2").
		SetBirthDate(dobAdult).
		IsValid()

	require.Equal(t, testMember.FullName(), "Name2 Middle2 Last2")
	require.Empty(t, testMember.Err())
}

func TestMember_FullNameFormal(t *testing.T) {
	testMember.
		SetFullName("Name2", "Middle2", "Last2").
		SetBirthDate(dobAdult).
		IsValid()

	require.Equal(t, testMember.FullNameFormal(), "Sr. Name2 Middle2 Last2")
	require.Empty(t, testMember.Err())
}

func TestMember_SetFullName_NameMissing(t *testing.T) {
	testMember.
		SetFullName("", "Middle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberName.Error()+"\n", testMember.Err())
}

func TestMember_SetFullName_NameTooShort(t *testing.T) {
	testMember.
		SetFullName("Na", "Middle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooShort.Error()+"\n", testMember.Err())
}

func TestMember_SetFullName_NameTooLong(t *testing.T) {
	testMember.
		SetFullName("NameNameNameNameNameNameNameNameName", "Middle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooLong.Error()+"\n", testMember.Err())
}

func TestMember_SetFullName_LastNameMissing(t *testing.T) {
	testMember.
		SetFullName("Name", "Middle", "").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberLastName.Error()+"\n", testMember.Err())
}

func TestMember_SetFullName_LastNameTooShort(t *testing.T) {
	testMember.
		SetFullName("Name", "Middle", "La").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooShort.Error()+"\n", testMember.Err())
}

func TestMember_SetFullName_LastNameTooLong(t *testing.T) {
	testMember.
		SetFullName("Name", "Middle", "LastLastLastLastLastLastLastLastLast").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooLong.Error()+"\n", testMember.Err())
}

func TestMember_SetFullName_MiddleNameTooLong(t *testing.T) {
	testMember.
		SetFullName("Name", "MiddleMiddleMiddleMiddleMiddleMiddleMiddle", "Last").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberMiddleNameTooLong.Error()+"\n", testMember.Err())
}

func TestMember_SetNickname(t *testing.T) {
	testMember.
		SetFullName("Name", "Middle", "Last").
		SetBirthDate(dobAdult).
		SetNickname("Nickname")

	require.True(t, testMember.IsValid())
}

func TestMember_Nickname(t *testing.T) {
	testMember = testMemberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
	testMember.
		SetNickname("Nickname").
		IsValid()
	nick := testMember.Nickname()

	require.EqualValues(t, "Nickname", nick)
}

func TestMember_SetGender(t *testing.T) {
	testMember.
		SetGender(family.Female).
		SetBirthDate(dobAdult).
		IsValid()

	require.EqualValues(t, family.Female, testMember.Gender())
	require.Empty(t, testMember.Err())
}

func TestMember_SetGender_Missing(t *testing.T) {
	testMember.
		SetGender("").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberGender.Error()+"\n", testMember.Err())
}

func TestMember_SetGender_Invalid(t *testing.T) {
	testMember.
		SetGender("other").
		SetBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrInvalidMemberGender.Error()+"\n", testMember.Err())
}

func TestMember_Gender(t *testing.T) {
	testMember.
		SetGender(family.Female).
		SetBirthDate(dobAdult).
		IsValid()
	gender := testMember.Gender()

	require.EqualValues(t, family.Female, gender)
}

func TestMember_SetBirthDate(t *testing.T) {
	testMember.SetBirthDate(dobElderly)

	require.True(t, testMember.IsValid())
	require.EqualValues(t, dobElderly, testMember.BirthDate())
}

func TestMember_SetBirthDate_Missing(t *testing.T) {
	testMember.SetBirthDate(time.Time{})

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberDOB.Error()+"\n", testMember.Err())
}

func TestMember_BirthDate(t *testing.T) {
	testMember.SetBirthDate(dobToddler)
	dob := testMember.BirthDate()

	require.EqualValues(t, dobToddler, dob)
}

func TestMember_AgeInMonths(t *testing.T) {
	testMember.SetBirthDate(dobNewborn)
	ageInMonths := testMember.AgeInMonths()

	require.EqualValues(t, 0, ageInMonths)
}

func TestMember_AgeInYears(t *testing.T) {
	testMember.SetBirthDate(dobNewborn)
	ageInYears := testMember.AgeInYears()

	require.EqualValues(t, 0, ageInYears)
}

func TestMember_IsNewborn(t *testing.T) {
	testMember.SetBirthDate(dobNewborn)

	ageT := testMember.IsNewborn()
	ageF := testMember.IsInfant()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsInfant(t *testing.T) {
	testMember.SetBirthDate(dobInfant)

	ageT := testMember.IsInfant()
	ageF := testMember.IsToddler()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsToddler(t *testing.T) {
	testMember.SetBirthDate(dobToddler)

	ageT := testMember.IsToddler()
	ageF := testMember.IsChild()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsChild(t *testing.T) {
	testMember.SetBirthDate(dobChild)

	ageT := testMember.IsChild()
	ageF := testMember.IsTeen()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsTeen(t *testing.T) {
	testMember.SetBirthDate(dobTeen)

	ageT := testMember.IsTeen()
	ageF := testMember.IsAdult()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsAdult(t *testing.T) {
	testMember.SetBirthDate(dobAdult)

	ageT := testMember.IsAdult()
	ageF := testMember.IsElderly()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsElderly(t *testing.T) {
	testMember.SetBirthDate(dobElderly)

	ageT := testMember.IsElderly()
	ageF := testMember.IsNewborn()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_MoreThanOneError(t *testing.T) {
	testMember = testMemberBuilder.
		WithFullName("Na", "Middle", "LastLastLastLastLastLastLastLastLast").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		Build()
	testMember.
		SetBirthDate(time.Time{}).
		SetGender("other").
		IsValid()

	allErrors := testMember.ErrToArray()
	require.False(t, testMember.IsValid())
	require.Contains(t, allErrors, family.ErrMemberNameTooShort.Error())
	require.Contains(t, allErrors, family.ErrMemberLastNameTooLong.Error())
	require.Contains(t, allErrors, family.ErrInvalidMemberGender.Error())
	require.Contains(t, allErrors, family.ErrMissingMemberDOB.Error())
	require.Equal(t, 4, len(allErrors))
}
