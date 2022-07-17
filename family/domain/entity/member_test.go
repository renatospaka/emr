package family_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	family "github.com/renatospaka/emr/family/domain/entity"
	"github.com/renatospaka/emr/common/infrastructure/utils"
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

func TestMember_IsValid(t *testing.T) {
	testMember.ChangeBirthDate(dobChild)

	require.True(t, testMember.IsValid())
	require.Empty(t, testMember.Err())
}

func TestMember_Invalid_EmptyMember(t *testing.T) {
	emptyMember := &family.Member{}

	require.False(t, emptyMember.IsValid())

	allErrors := emptyMember.Err()
	require.Contains(t, allErrors, family.ErrInvalidMember.Error())
	require.Len(t, allErrors, 1)
}

func TestMember_ID(t *testing.T) {
	id := testMember.ID()
	err := utils.IsVaalidUUID(id)
	valid := err == nil

	require.True(t, valid)
	require.NotEmpty(t, id)
	require.Nil(t, err)
}

func TestMember_FullName(t *testing.T) {
	testMember.
		ChangeFullName("Name2", "Middle2", "Last2").
		ChangeBirthDate(dobAdult).
		IsValid()

	require.Equal(t, testMember.FullName(), "Name2 Middle2 Last2")
	require.Empty(t, testMember.Err())
}

func TestMember_FullNameFormal(t *testing.T) {
	testMember.
		ChangeFullName("Name2", "Middle2", "Last2").
		ChangeBirthDate(dobAdult).
		IsValid()

	require.Equal(t, testMember.FullNameFormal(), "Sr. Name2 Middle2 Last2")
	require.Empty(t, testMember.Err())
}

func TestMember_ChangeFullName_NameMissing(t *testing.T) {
	testMember.
		ChangeFullName("", "Middle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberName.Error()+"\n", testMember.Err())
}

func TestMember_ChangeFullName_NameTooShort(t *testing.T) {
	testMember.
		ChangeFullName("Na", "Middle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooShort.Error()+"\n", testMember.Err())
}

func TestMember_ChangeFullName_NameTooLong(t *testing.T) {
	testMember.
		ChangeFullName("NameNameNameNameNameNameNameNameName", "Middle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberNameTooLong.Error()+"\n", testMember.Err())
}

func TestMember_ChangeFullName_LastNameMissing(t *testing.T) {
	testMember.
		ChangeFullName("Name", "Middle", "").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberLastName.Error()+"\n", testMember.Err())
}

func TestMember_ChangeFullName_LastNameTooShort(t *testing.T) {
	testMember.
		ChangeFullName("Name", "Middle", "La").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooShort.Error()+"\n", testMember.Err())
}

func TestMember_ChangeFullName_LastNameTooLong(t *testing.T) {
	testMember.
		ChangeFullName("Name", "Middle", "LastLastLastLastLastLastLastLastLast").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberLastNameTooLong.Error()+"\n", testMember.Err())
}

func TestMember_ChangeFullName_MiddleNameTooLong(t *testing.T) {
	testMember.
		ChangeFullName("Name", "MiddleMiddleMiddleMiddleMiddleMiddleMiddle", "Last").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMemberMiddleNameTooLong.Error()+"\n", testMember.Err())
}

func TestMember_ChangeNickname(t *testing.T) {
	testMember.
		ChangeFullName("Name", "Middle", "Last").
		ChangeBirthDate(dobAdult).
		ChangeNickname("Nickname")

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
		ChangeNickname("Nickname").
		IsValid()
	nick := testMember.Nickname()

	require.EqualValues(t, "Nickname", nick)
}

func TestMember_ChangeGender(t *testing.T) {
	testMember.
		ChangeGender(family.Female).
		ChangeBirthDate(dobAdult).
		IsValid()

	require.EqualValues(t, family.Female, testMember.Gender())
	require.Empty(t, testMember.Err())
}

func TestMember_ChangeGender_Missing(t *testing.T) {
	testMember.
		ChangeGender("").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberGender.Error()+"\n", testMember.Err())
}

func TestMember_ChangeGender_Invalid(t *testing.T) {
	testMember.
		ChangeGender("other").
		ChangeBirthDate(dobAdult)

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrInvalidMemberGender.Error()+"\n", testMember.Err())
}

func TestMember_Gender(t *testing.T) {
	testMember.
		ChangeGender(family.Female).
		ChangeBirthDate(dobAdult).
		IsValid()
	gender := testMember.Gender()

	require.EqualValues(t, family.Female, gender)
}

func TestMember_ChangeBirthDate(t *testing.T) {
	testMember.ChangeBirthDate(dobElderly)

	require.True(t, testMember.IsValid())
	require.EqualValues(t, dobElderly, testMember.BirthDate())
}

func TestMember_ChangeBirthDate_Missing(t *testing.T) {
	testMember.ChangeBirthDate(time.Time{})

	require.False(t, testMember.IsValid())
	require.EqualValues(t, family.ErrMissingMemberDOB.Error()+"\n", testMember.Err())
}

func TestMember_BirthDate(t *testing.T) {
	testMember.ChangeBirthDate(dobToddler)
	dob := testMember.BirthDate()

	require.EqualValues(t, dobToddler, dob)
}

func TestMember_AgeInMonths(t *testing.T) {
	testMember.ChangeBirthDate(dobNewborn)
	ageInMonths := testMember.AgeInMonths()

	require.EqualValues(t, 0, ageInMonths)
}

func TestMember_AgeInYears(t *testing.T) {
	testMember.ChangeBirthDate(dobNewborn)
	ageInYears := testMember.AgeInYears()

	require.EqualValues(t, 0, ageInYears)
}

func TestMember_IsNewborn(t *testing.T) {
	testMember.ChangeBirthDate(dobNewborn)

	ageT := testMember.IsNewborn()
	ageF := testMember.IsInfant()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsInfant(t *testing.T) {
	testMember.ChangeBirthDate(dobInfant)

	ageT := testMember.IsInfant()
	ageF := testMember.IsToddler()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsToddler(t *testing.T) {
	testMember.ChangeBirthDate(dobToddler)

	ageT := testMember.IsToddler()
	ageF := testMember.IsChild()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsChild(t *testing.T) {
	testMember.ChangeBirthDate(dobChild)

	ageT := testMember.IsChild()
	ageF := testMember.IsTeen()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsTeen(t *testing.T) {
	testMember.ChangeBirthDate(dobTeen)

	ageT := testMember.IsTeen()
	ageF := testMember.IsAdult()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsAdult(t *testing.T) {
	testMember.ChangeBirthDate(dobAdult)

	ageT := testMember.IsAdult()
	ageF := testMember.IsElderly()
	require.True(t, ageT)
	require.False(t, ageF)
}

func TestMember_IsElderly(t *testing.T) {
	testMember.ChangeBirthDate(dobElderly)

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
		ChangeBirthDate(time.Time{}).
		ChangeGender("other")

	require.False(t, testMember.IsValid())

	allErrors := testMember.Err()
	require.Contains(t, allErrors, family.ErrMemberNameTooShort.Error())
	require.Contains(t, allErrors, family.ErrMemberLastNameTooLong.Error())
	require.Contains(t, allErrors, family.ErrInvalidMemberGender.Error())
	require.Contains(t, allErrors, family.ErrMissingMemberDOB.Error())
	require.Equal(t, len(allErrors), 4)
}
