package main

import (
	// "fmt"
	"log"

	"github.com/renatospaka/emr/domain/entity/family"
	familyRepository "github.com/renatospaka/emr/infrastructure/repository/family"
)

func main() {
	famRepo := familyRepository.NewFamilyRepositoryInMemory()
	fam1 := family.NewFamily("Spakauskas")
	err := famRepo.Create(fam1)
	if err != nil {
		log.Println("Nova família criada com erro: ", err)
	} else {
		log.Println("Nova família criada - ID: ", fam1.ID, ", Surname: ", fam1.Surname)
	}


	fam2 := family.NewFamily("")
	err = famRepo.Create(fam2)
	if err != nil {
		log.Println("Nova família criada com erro: ", err)
	} else {
		log.Println("Nova família criada - ID: ", fam2.ID, ", Surname: ", fam2.Surname)
	}

}
