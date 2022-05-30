package familyRepository_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/domain/entity/family"
	familyRepository "github.com/renatospaka/emr/infrastructure/repository/family"
)

var (
	dob = time.Date(1970, 14, 11, 0, 0, 0, 0, time.UTC)
)

func TestMember_Add(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, dob)

	err := memberRepo.Add(member)
	require.Nil(t, err)
	require.True(t, member.Valid)
}

func TestMember_FullName(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, dob)

	_ = memberRepo.Add(member)
	fullName := member.FullName(false)
	fullNameComplete := member.FullName(true)
	require.Equal(t, "Renato Costa Spakauskas", fullName)
	require.Equal(t, "Sr. Renato Costa Spakauskas", fullNameComplete)
	require.NotEqual(t, "Sra. Renato Costa Spakauskas", fullNameComplete)
}

func TestMember_Add_MissingName(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("", "Costa", "Spakauskas", family.Male, dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberName.Error())
}

func TestMember_Add_NameTooShort(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Re", "Costa", "Spakauskas", family.Male, dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberNameTooShort.Error())
}

func TestMember_Add_NameTooLong(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato12345678901234567890", "Costa", "Spakauskas", family.Male, dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberNameTooLong.Error())
}

func TestMember_Add_MissingLastName(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "", family.Male, dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberLastName.Error())
}

func TestMember_Add_LastNameTooShort(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Sp", family.Male, dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberLastNameTooShort.Error())
}

func TestMember_Add_LastNameTooLong(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas12345678901234567890", family.Male, dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMemberLastNameTooLong.Error())
}

func TestMember_Add_MissingDOB(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, time.Time{})

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMemberDOB.Error())
}

func TestMember_Add_MissingGender(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", "", dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrMissingMembeGender.Error())
}

func TestMember_Add_InvalidGender(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", "Outro", dob)

	err := memberRepo.Add(member)
	require.NotNil(t, err)
	require.False(t, member.Valid)
	require.EqualError(t, err, family.ErrInvalidMembeGender.Error())
}

func TestMember_FullName_FindById(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, dob)
	_ = memberRepo.Add(member)

	id := member.ID
	find, err := memberRepo.FindById(id)
	require.Nil(t, err)
	require.True(t, find.Valid)
	require.Equal(t, id, find.ID)
}

func TestMember_FullName_FindById_NotFound(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "Costa", "Spakauskas", family.Male, dob)
	_ = memberRepo.Add(member)

	id :=uuid.NewV4()
	find, err := memberRepo.FindById(id)
	require.NotNil(t, err)
	require.EqualError(t, err, family.ErrMemberNotFound.Error())
	require.Equal(t, &family.Member{}, find)
}

func TestMember_Chance(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "", "Spakauskas", family.Male, dob)
	_ = memberRepo.Add(member)

	member.MiddleName = "Costa"
	upd, err := memberRepo.Change(member)	
	require.Nil(t, err)
	require.True(t, upd.Valid)
	require.Equal(t, "Costa", upd.MiddleName)
}

func TestMember_Chance_NotFound(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "", "Spakauskas", family.Male, dob)
	_ = memberRepo.Add(member)

	new := family.NewMember("Renato", "", "Spakauskas", family.Male, dob)

	new.MiddleName = "Costa"
	upd, err := memberRepo.Change(new)	
	require.NotNil(t, err)
	require.True(t, upd.Valid)
	require.EqualError(t, err, family.ErrMemberNotFound.Error())
}

func TestMember_Chance_NoChanges(t *testing.T) {
	memberRepo := familyRepository.NewMemberRepositoryInMemory()
	member := family.NewMember("Renato", "", "Spakauskas", family.Male, dob)
	_ = memberRepo.Add(member)

	upd, err := memberRepo.Change(member)	
	require.NotNil(t, err)
	require.True(t, upd.Valid)
	require.EqualError(t, err, family.ErrNoChangesNeeded.Error())
	require.EqualValues(t, member, upd)
}
