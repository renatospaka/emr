package family_test

import (
	"time"

	family "github.com/renatospaka/emr/family/domain/entity"
)

// member testing variables
var (
	// testMember *family.Member
	dobNewborn time.Time
	dobInfant  time.Time
	dobToddler time.Time
	dobChild   time.Time
	dobTeen    time.Time
	dobAdult   time.Time
	dobElderly time.Time
)

// // builder-member testing variables
// var (
// 	testMemberBuilder *family.MemberBuilder
// )

// // family testing variables
// var (
// 	testMemberHOF *family.Member
// )

// // builder-family testing variables
// var (
// 	testFamilyBuilder *family.FamilyBuilder
// )

// // builder-family member testing variables
// var (
// 	testFamilyMemberBuilder *family.FamilyMemberBuilder
// )

func createMemberBuilder() *family.MemberBuilder {
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

func createEmptyFamilyMember() *family.FamilyMember {
	return &family.FamilyMember{}
}