package main

import (
	"log"
	"time"

	"github.com/renatospaka/emr/common/infrastructure/utils"
	family "github.com/renatospaka/emr/family/domain/entity"
)

var (
	today    = time.Now()
	dobAdult = today.Add(-33 * utils.HoursYear * time.Hour)
	dobTeen = today.Add(-15 * utils.HoursYear * time.Hour)
)

func createMemberBuilder() *family.MemberBuilder {
	// log.Println("createMemberBuilder()")
	return family.NewMemberBuilder()
}

func createFamilyMemberBuilder() *family.FamilyMemberBuilder {
	// log.Println("createFamilyMemberBuilder()")
	return family.NewFamilyMemberBuilder()
}

func createFamilyBuilder() *family.FamilyBuilder {
	// log.Println("createFamilyBuilder()")
	return family.NewFamilyBuilder()
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
		WithGender(family.Female).
		WithNickname("Teen").
		Build()
}

func createFamilyMember() *family.FamilyMember {
	// log.Println("createFamilyMember()")
	member := createHOFMember()
	famMembBuilder := createFamilyMemberBuilder()
	return famMembBuilder.
		AsHOF(member).
		Build()
}

func createTeenFamilyMember() *family.FamilyMember {
	// log.Println("createFamilyMember()")
	member := createTeenagerMember()
	famMembBuilder := createFamilyMemberBuilder()
	return famMembBuilder.
		AsOrdinary(member).
		Build()
}

func main() {
	log.Println("Testing Family Member")
	testFamilyMember()

	log.Println("Testing Family ")
	testFamily()

	// log.Println("Testing Dates")
	// testDates()
}

func testDates() {
	today := time.Now()
	dobNewborn1 := time.Date(today.Year(), today.Month(), today.Day()-33, today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobNewborn2 := today.Add(-33 * utils.HoursDay * time.Hour)

	log.Println("dobNewborn1", dobNewborn1)
	log.Println("dobNewborn2", dobNewborn2)

	dobTeen1 := time.Date(today.Year()-15, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	dobTeen2 := today.Add(-15 * (utils.HoursYear) * time.Hour)

	log.Println("dobTeen1", dobTeen1)
	log.Println("dobTeen2", dobTeen2)
}

func memberNoErr() *family.Member {
	memberBuilder := createMemberBuilder()
	member := memberBuilder.
		WithFullName("Name", "Middle", "Last").
		WithBirthDate(dobAdult).
		WithGender(family.Male).
		WithNickname("Nick").
		Build()
	return member
}

func memberErr() *family.Member {
	memberBuilder := createMemberBuilder()
	member := memberBuilder.
		WithFullName("Na", "Middle", "LastLastLastLastLastLastLast").
		WithBirthDate(time.Time{}).
		WithGender("other").
		WithNickname("Nick").
		Build()
	return member
}

func testFamilyMember() {
	member := createHOFMember()
	log.Println("Member :=", member)

	famMemberBuilder := createFamilyMemberBuilder()
	famMember := famMemberBuilder.
		AsHOF(member).
		Build()
	log.Println("famMember :=", famMember)
}

func testFamily() {
	famMember := createFamilyMember()
	familyBuilder := createFamilyBuilder()
	fam := familyBuilder.
		WithSurname("Super Family").
		Add(famMember).
		Build()

	// 	teenFamMember := createTeenFamilyMember()
	// 	teenFamMember.Member.ChangeGender(family.Female)
	// 	teenFamMember.ChangeRelationType(family.RelDaughter)
	// 	// adding a new member
	// fam.AddMember(teenFamMember)

	log.Println("fam :=", fam)
}
