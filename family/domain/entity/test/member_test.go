package family_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/common/infrastructure/utils"
	family "github.com/renatospaka/emr/family/domain/entity"
)

func init() {
	today := time.Now()
	dobNewborn = today.Add(-15 * (utils.HoursDay) * time.Hour)
	dobInfant = today.Add(-11 * utils.HoursMonth * time.Hour)
	dobToddler = today.Add(-54 * utils.HoursMonth * time.Hour)
	dobChild = today.Add(-7 * utils.HoursYear * time.Hour)
	dobTeen = today.Add(-15 * utils.HoursYear * time.Hour)
	dobAdult = today.Add(-33 * utils.HoursYear * time.Hour)
	dobElderly = today.Add(-67 * utils.HoursYear * time.Hour)
}

func TestMember_IsValid(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobChild)

	require.True(t, member.IsValid())
	require.Len(t, member.Err(), 0)
}

func TestMember_Invalid_EmptyMember(t *testing.T) {
	emptyMember := createEmptyMember()
	require.False(t, emptyMember.IsValid())

	allErrors := emptyMember.Err()
	require.Contains(t, allErrors, family.ErrInvalidMember.Error())
	require.Len(t, allErrors, 1)
}

func TestMember_ID(t *testing.T) {
	member := createMember()
	id := member.ID()
	err := utils.IsVaalidUUID(id)
	valid := err == nil

	require.True(t, valid)
	require.NotEmpty(t, id)
	require.Nil(t, err)
}

func TestMember_FullName(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("Name2", "Middle2", "Last2").
		ChangeBirthDate(dobAdult).
		IsValid()

	require.Equal(t, member.FullName(), "Name2 Middle2 Last2")
	require.Len(t, member.Err(), 0)
}

func TestMember_FullNameFormal(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("Name2", "Middle2", "Last2").
		ChangeBirthDate(dobAdult).
		IsValid()

	require.Equal(t, member.FullNameFormal(), "Sr. Name2 Middle2 Last2")
	require.Len(t, member.Err(), 0)
}

func TestMember_ChangeFullName_NameMissing(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("", "Middle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMissingMemberName.Error())
}

func TestMember_ChangeFullName_NameTooShort(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("Na", "Middle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMemberNameTooShort.Error())
}

func TestMember_ChangeFullName_NameTooLong(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("NameNameNameNameNameNameNameNameName", "Middle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMemberNameTooLong.Error())
}

func TestMember_ChangeFullName_LastNameMissing(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("Name", "Middle", "").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMissingMemberLastName.Error())
}

func TestMember_ChangeFullName_LastNameTooShort(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("Name", "Middle", "La").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMemberLastNameTooShort.Error())
}

func TestMember_ChangeFullName_LastNameTooLong(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("Name", "Middle", "LastLastLastLastLastLastLastLastLast").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMemberLastNameTooLong.Error())
}

func TestMember_ChangeFullName_MiddleNameTooLong(t *testing.T) {
	member := createMember()
	member.
		ChangeFullName("Name", "MiddleMiddleMiddleMiddleMiddleMiddleMiddle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMemberMiddleNameTooLong.Error())
}

func TestMember_ChangeNickname(t *testing.T) {
	member := createMember()
	member.ChangeNickname("Nickname")

	require.True(t, member.IsValid())
}

func TestMember_Nickname(t *testing.T) {
	member := createMember()
	member.
		ChangeNickname("Nickname").
		IsValid()
	nick := member.Nickname()

	require.EqualValues(t, "Nickname", nick)
}

func TestMember_ChangeGender(t *testing.T) {
	member := createMember()
	member.
		ChangeGender(family.Female).
		ChangeBirthDate(dobAdult).
		IsValid()

	require.EqualValues(t, family.Female, member.Gender())
	require.Len(t, member.Err(), 0)
}

func TestMember_ChangeGender_Missing(t *testing.T) {
	member := createMember()
	member.
		ChangeGender("").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMissingMemberGender.Error())
}

func TestMember_ChangeGender_Invalid(t *testing.T) {
	member := createMember()
	member.
		ChangeGender("other").
		ChangeBirthDate(dobAdult)

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrInvalidMemberGender.Error())
}

func TestMember_Gender(t *testing.T) {
	member := createMember()

	require.EqualValues(t, family.Male, member.Gender())
}

func TestMember_ChangeBirthDate(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobElderly)

	require.True(t, member.IsValid())
	require.EqualValues(t, dobElderly, member.BirthDate())
}

func TestMember_ChangeBirthDate_Missing(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(time.Time{})

	require.False(t, member.IsValid())
	require.Contains(t, member.Err(), family.ErrMissingMemberDOB.Error())
}

func TestMember_BirthDate(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobToddler)
	dob := member.BirthDate()

	require.EqualValues(t, dobToddler, dob)
}

func TestMember_AgeInMonths(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobNewborn)
	ageInMonths := member.AgeInMonths()

	require.EqualValues(t, 0, ageInMonths)
}

func TestMember_AgeInYears(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobNewborn)
	ageInYears := member.AgeInYears()

	require.EqualValues(t, 0, ageInYears)
}

func TestMember_IsNewborn(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobNewborn)

	ageT := member.IsNewborn()
	ageF := member.IsInfant()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsInfant(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobInfant)

	ageT := member.IsInfant()
	ageF := member.IsToddler()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsToddler(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobToddler)

	ageT := member.IsToddler()
	ageF := member.IsChild()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsChild(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobChild)

	ageT := member.IsChild()
	ageF := member.IsTeen()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsTeen(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobTeen)

	ageT := member.IsTeen()
	ageF := member.IsAdult()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsAdult(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobAdult)

	ageT := member.IsAdult()
	ageF := member.IsElderly()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsElderly(t *testing.T) {
	member := createMember()
	member.ChangeBirthDate(dobElderly)

	ageT := member.IsElderly()
	ageF := member.IsNewborn()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_MoreThanOneError(t *testing.T) {
	member := createInvalidMember()
	require.False(t, member.IsValid())

	allErrors := member.Err()
	require.Contains(t, allErrors, family.ErrMemberNameTooShort.Error())
	require.Contains(t, allErrors, family.ErrMemberLastNameTooLong.Error())
	require.Contains(t, allErrors, family.ErrInvalidMemberGender.Error())
	require.Contains(t, allErrors, family.ErrMissingMemberDOB.Error())
	require.Equal(t, len(allErrors), 4)
}
