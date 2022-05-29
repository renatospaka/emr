package familyRepository

import "github.com/renatospaka/emr/domain/entity/family"

type FamilyRepository interface {
	Create(family *family.Family) error 
}
