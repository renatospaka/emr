package familyInMemory

import "github.com/renatospaka/emr/domain/entity/family"

func removeMember(m []family.Member, index int) ([]family.Member, error) {
	// perform bounds checking first to prevent a panic!
	if index >= len(m) || index < 0 {
		return nil, family.ErrIndexOutOfRange
	}

	m[index] = m[len(m)-1]
	return m[:len(m)-1], nil
}
