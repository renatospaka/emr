package main

import (
	"log"
	"time"

	"github.com/renatospaka/emr/domain/entity/family"
	familyRepository "github.com/renatospaka/emr/infrastructure/repository/family/inMemory"
)

func main() {
	repoFamily := familyRepository.NewFamilyRepositoryInMemory()
	fam1 := family.NewFamily("Familyname")
	err := repoFamily.Create(fam1)
	if err != nil {
		log.Println("Nova família criada com erro:", err)
	} else {
		log.Println("Nova família criada -", fam1)
	}

	repoMember := familyRepository.NewMemberRepositoryInMemory()
	memb := family.NewMember("Name", "MiddleName", "Lastname", family.Male, time.Date(2006, 8, 5, 0, 0, 0, 0, time.UTC))
	err = repoMember.Add(memb)
	if err != nil {
		log.Println("Novo membro criado com erro:", err)
	} else {
		log.Println("Novo membro criado! Full name:(", memb.FullNameFormal(), memb.ID, ")")
	}

	id := memb.ID
	findMemb, err := repoMember.FindById(id)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Achou! Full name:", findMemb.FullNameFormal())
	}

	famMember := family.NewFamilyMember(memb)
	// log.Println(famMember.Member.FullName(), famMember.RelationType, famMember.Status)
	fam2, err := repoFamily.AddFamilyMember(*famMember, fam1.ID)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Adicionou o novo integrante:", fam2.Members[0].Member.FullNameFormal(), fam2.Members[0].Member.ID)
	}

	// err = repoMember.Remove(id)
	// if err != nil {
	// 	log.Println(err)
	// }

	// err = repoMember.Remove(id)
	// if err != nil {
	// 	log.Println(err)
	// }
}
