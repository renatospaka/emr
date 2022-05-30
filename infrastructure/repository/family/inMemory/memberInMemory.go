package familyInMemory

import "github.com/renatospaka/emr/domain/entity/family"

type MemberRepositoryInMemory struct {
	member []family.Member
}

func NewMemberRepositoryInMemory() *MemberRepositoryInMemory {
	return &MemberRepositoryInMemory{
		member: []family.Member{},
	}
}
