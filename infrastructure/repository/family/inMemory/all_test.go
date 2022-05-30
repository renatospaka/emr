package familyInMemory_test

import (
	"time"

	"github.com/renatospaka/emr/domain/entity/family"
	repoFamily "github.com/renatospaka/emr/infrastructure/repository/family/inMemory"
)

var (
	repoFam = repoFamily.NewFamilyRepositoryInMemory()
	fam     = family.NewFamily("Lastname")
)

var (
	repoMember = repoFamily.NewMemberRepositoryInMemory()
	dob        = time.Date(1980, 23, 8, 0, 0, 0, 0, time.UTC)
	member     = family.NewMember("Name", "MiddleName", "Lastname", family.Male, dob)
)
