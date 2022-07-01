package familyInMemory

// import (
// 	"errors"

// 	family "github.com/renatospaka/emr/family/domain/entity"
// )

// // Add a member so s/he may be late related to a family
// // or returns an error
// func (m *MemberRepositoryInMemory) Add(member *family.Member) error {
// 	err := member.IsValid()
// 	if err != nil {
// 		return err
// 	}

// 	m.member = append(m.member, *member)
// 	if len(m.member) > 0 {
// 		return nil
// 	}
// 	return errors.New("ocorreu um erro na inclusão de um novo membro da família")
// }
