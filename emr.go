package main

import (
	"log"
	"time"

	"github.com/renatospaka/emr/infrastructure/utils"
	// 	family "github.com/renatospaka/emr/family/domain/entity"
	// 	familyRepository "github.com/renatospaka/emr/infrastructure/repository/family/inMemory"
)

func main() {
	// 	repoFamily := familyRepository.NewFamilyRepositoryInMemory()
	// 	fam1 := family.NewFamily("Familyname")
	// 	err := repoFamily.Create(fam1)
	// 	if err != nil {
	// 		log.Println("Nova família criada com erro:", err)
	// 	} else {
	// 		log.Println("Nova família criada -", fam1)
	// 	}

	// 	repoMember := familyRepository.NewMemberRepositoryInMemory()
	// 	memb := family.NewMember("Name", "MiddleName", "Lastname", family.Male, time.Date(2006, 8, 5, 0, 0, 0, 0, time.UTC))
	// 	err = repoMember.Add(memb)
	// 	if err != nil {
	// 		log.Println("Novo membro criado com erro:", err)
	// 	} else {
	// 		log.Println("Novo membro criado! Full name:(", memb.FullNameFormal(), memb.ID, ")")
	// 	}

	// 	id := memb.ID
	// 	findMemb, err := repoMember.FindById(id)
	// 	if err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Println("Achou! Full name:", findMemb.FullNameFormal())
	// 	}

	// 	famMember := family.NewFamilyMember(memb)
	// 	// log.Println(famMember.Member.FullName(), famMember.RelationType, famMember.Status)
	// 	fam2, err := repoFamily.AddFamilyMember(*famMember, fam1.ID)
	// 	if err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Println("Adicionou o novo integrante:", fam2.Members[0].Member.FullNameFormal(), fam2.Members[0].Member.ID)
	// 	}

	// 	// err = repoMember.Remove(id)
	// 	// if err != nil {
	// 	// 	log.Println(err)
	// 	// }

	// 	// err = repoMember.Remove(id)
	// 	// if err != nil {
	// 	// 	log.Println(err)
	// 	// }

	today := time.Now()
	ageInYears, ageInMonths := utils.AgeFromToday(today)
	log.Println("today", today, ageInYears, ageInMonths)

	dobNewborn := time.Date(today.Year(), today.Month(), today.Day()-33, today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths = utils.AgeFromToday(dobNewborn)
	log.Println("dobNewborn", dobNewborn, ageInYears, ageInMonths)

	dobInfant := time.Date(today.Year(), today.Month()-11, today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths = utils.AgeFromToday(dobInfant)
	log.Println("dobInfant", dobInfant, ageInYears, ageInMonths)

	dobToddler := time.Date(today.Year(), today.Month()-54, today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths = utils.AgeFromToday(dobToddler)
	log.Println("dobToddler", dobToddler, ageInYears, ageInMonths)

	dobChild := time.Date(today.Year()-7, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths = utils.AgeFromToday(dobChild)
	log.Println("dobChild", dobChild, ageInYears, ageInMonths)

	dobTeen := time.Date(today.Year()-15, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths = utils.AgeFromToday(dobTeen)
	log.Println("dobTeen", dobTeen, ageInYears, ageInMonths)

	dobAdult := time.Date(today.Year()-33, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths = utils.AgeFromToday(dobAdult)
	log.Println("dobAdult", dobAdult, ageInYears, ageInMonths)

	dobElderly := time.Date(today.Year()-67, today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), time.UTC)
	ageInYears, ageInMonths = utils.AgeFromToday(dobElderly)
	log.Println("dobElderly", dobElderly, ageInYears, ageInMonths)
}
