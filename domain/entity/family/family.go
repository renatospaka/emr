package family

import (
	"strings"

	"github.com/renatospaka/emr2/infrastructure/utils"
)

type Family struct {
	ID      string          `json:"family_id"`
	Surname string          `json:"surname"`
	Valid   bool            `json:"is_valid"`
	Members []*FamilyMember `json:"members"`
}

type FamilyMember struct {
	Member       *Member `json:"member"`
	RelationType string  `json:"relation_type"`
	Status       string  `json:"status"`
}

func NewFamily(surname string) *Family {
	newFamily := &Family{
		ID:      utils.GetID(),
		Surname: strings.TrimSpace(surname),
		Valid:   false,
		Members: []*FamilyMember{},
	}

	newFamily.IsValid()
	return newFamily
}

// Check whenever the family structure is intact
// and filled accordingly to the model rules
func (f *Family) IsValid() error {
	f.Valid = false
	if strings.TrimSpace(f.Surname) == "" {
		return ErrMissingFamilySurname
	}
	if strings.TrimSpace(f.ID) == "" {
		return ErrMissingFamilyID
	}

	f.Valid = true
	return nil
}

func NewFamilyMember(newMember *Member) *FamilyMember {
	return &FamilyMember{
		Member:       newMember,
		RelationType: TBDRelation,
		Status:       FreshMember,
	}
}
