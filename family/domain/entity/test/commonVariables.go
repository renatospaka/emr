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
	// log.Println("createMember()")
	memberBuilder := createMemberBuilder()
	return memberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		Build()
}

func createInvalidMember() *family.Member {
	// log.Println("createInvalidMember()")
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
	// log.Println("createTeenagerMember()")
	memberBuilder := createMemberBuilder()
	return memberBuilder.
		WithFullName("Name", "Teenager", "Last").
		WithBirthDate(dobTeen).
		WithGender(family.Male).
		WithNickname("Teen").
		Build()
}

func createEmptyMember() *family.Member {
	// log.Println("createEmptyMember()")
	return &family.Member{}
}

func createFamilyMemberBuilder() *family.FamilyMemberBuilder {
	// log.Println("createFamilyMemberBuilder()")
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
	// log.Println("createEmptyFamilyMember()")
	return &family.FamilyMember{}
}

func createOrdinaryFamilyMember() *family.FamilyMember {
	// log.Println("createOrdinaryFamilyMember()")
	member := createMember()
	famMembBuilder := createFamilyMemberBuilder()
	return famMembBuilder.
		AsOrdinary(member).
		Build()
}

func createEmptyOrdinaryFamilyMember() *family.FamilyMember {
	// log.Println("createEmptyOrdinaryFamilyMember()")
	member := createEmptyMember()
	famMembBuilder := createFamilyMemberBuilder()
	return famMembBuilder.
		AsOrdinary(member).
		Build()
}

func createEmptyMemberFamilyMember() *family.FamilyMember {
	// log.Println("createEmptyMemberFamilyMember()")
	member := createEmptyMember()
	famMembBuilder := createFamilyMemberBuilder()
	return famMembBuilder.
		AsHOF(member).
		Build()
}

func createFamilyBuilder() *family.FamilyBuilder {
	// log.Println("createFamilyBuilder()")
	return family.NewFamilyBuilder()
}

func createFamily() *family.Family {
	// log.Println("createFamily()")
	familyMember := createFamilyMember()
	familyBuilder := createFamilyBuilder()
	return familyBuilder.
		WithSurname("Super Family").
		Add(familyMember).
		Build()
}

func createInvalidFamily() *family.Family {
	// log.Println("createInvalidFamily()")
	missingFamilyMember := createEmptyMemberFamilyMember()
	familyBuilder := createFamilyBuilder()
	return familyBuilder.
		WithSurname("Super Family").
		Add(missingFamilyMember).
		Build()
}

func createMissingHOFFamily() *family.Family {
	// log.Println("createMissingHOFFamily()")
	ordinaryFamilyMember := createOrdinaryFamilyMember()
	familyBuilder := createFamilyBuilder()
	return familyBuilder.
		WithSurname("Super Family").
		Add(ordinaryFamilyMember).
		Build()
}

func createEmptyFamily() *family.Family {
	// log.Println("createEmptyFamily()")
	return &family.Family{}
}
