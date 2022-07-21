package family_test

import (
	"time"

	family "github.com/renatospaka/emr/family/domain/entity"
)

// member testing variables
var (
	dobNewborn time.Time
	dobInfant  time.Time
	dobToddler time.Time
	dobChild   time.Time
	dobTeen    time.Time
	dobAdult   time.Time
	dobElderly time.Time
)

func createMemberBuilder() *family.MemberBuilder {
	// log.Println("createMemberBuilder()")
	return family.NewMemberBuilder()
}

func createMember() *family.Member {
	memberBuilder := createMemberBuilder()
	return  memberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		Build()
}

func createInvalidMember() *family.Member {
	memberBuilder := createMemberBuilder()
	return memberBuilder.
		WithFullName("Na", "Middle", "LastLastLastLastLastLastLast").
		WithBirthDate(time.Time{}).
		WithGender("other").
		WithNickname("Nick").
		Build()
}

func createHOFMember() *family.Member {
	// log.Println("createHOFMember()")
	memberBuilder := createMemberBuilder()
	return memberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("HOF").
		Build()
}

func createTeenagerMember() *family.Member {
	memberBuilder := createMemberBuilder()
	return memberBuilder.
		WithFullName("Name", "Teenager", "Last").
		WithBirthDate(dobTeen).
		WithGender(family.Male).
		WithNickname("Teen").
		Build()
}

func createEmptyMember() *family.Member {
	return  &family.Member{}
}

func createFamilyMemberBuilder() *family.FamilyMemberBuilder {
	return family.NewFamilyMemberBuilder()
}

func createFamilyMember() *family.FamilyMember {
	// log.Println("createFamilyMember()")
	member := createHOFMember()
	famMembBuilder := createFamilyMemberBuilder()
	return famMembBuilder.
		AsHOF(member).
		Build()
}

func createEmptyFamilyMember() *family.FamilyMember {
	return &family.FamilyMember{}
}

func createFamilyBuilder() *family.FamilyBuilder {
	// log.Println("createFamilyBuilder()")
	return family.NewFamilyBuilder()
}

func createEmptyFamily() *family.Family {
	return &family.Family{}
}
