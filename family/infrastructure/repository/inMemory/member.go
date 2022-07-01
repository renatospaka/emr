package familyInMemory

import family "github.com/renatospaka/emr/family/domain/entity"

type MemberRepositoryInMemory struct {
	member []family.Member
}

func NewMemberRepositoryInMemory() *MemberRepositoryInMemory {
	return &MemberRepositoryInMemory{
		member: []family.Member{},
	}
}
