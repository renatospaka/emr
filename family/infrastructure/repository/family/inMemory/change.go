package familyInMemory

// import family "github.com/renatospaka/emr/family/domain/entity"

// // Allow user to change or update attributes of a specific family member
// // or returns an error
// func (m *MemberRepositoryInMemory) Change(member *family.Member) (*family.Member, error) {
// 	err := member.IsValid()
// 	if err != nil {
// 		return member, err
// 	}

// 	id := member.ID
// 	current, err := m.FindById(id)
// 	if err != nil {
// 		return member, err
// 	}

// 	// Check if there is any change to make.
// 	// If there isn't, no action would be taken
// 	if current.Name == member.Name &&
// 		current.MiddleName == member.MiddleName &&
// 		current.LastName == member.LastName &&
// 		current.DOB == member.DOB &&
// 		current.Gender == member.Gender {
// 		return member, family.ErrNoChangesNeeded
// 	}

// 	// apply the changes
// 	current.Name = member.Name
// 	current.MiddleName = member.MiddleName
// 	current.LastName = member.LastName
// 	current.DOB = member.DOB
// 	current.Gender = member.Gender
// 	current.IsValid()

// 	return current, nil
// }
