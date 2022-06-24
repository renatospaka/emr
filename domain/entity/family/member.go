package family

import (
	"strings"
	"time"

	"github.com/renatospaka/emr/infrastructure/utils"
)

type Member struct {
	ID         string    `json:"family_member_id"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
	DOB        time.Time `json:"day_of_birth"`
	Gender     string    `json:"gender"`
	Valid      bool      `json:"is_valid"`
}

func NewMember(name string, middleName string, lastName string, gender string, dob time.Time) *Member {
	return &Member{
		ID:         utils.GetID(),
		Name:       strings.TrimSpace(name),
		LastName:   strings.TrimSpace(lastName),
		MiddleName: strings.TrimSpace(middleName),
		DOB:        dob,
		Gender:     strings.TrimSpace(gender),
		Valid:      false,
	}
}

// Check whenever the member structure is intact
// and filled accordingly to the model rules
func (m *Member) IsValid() error {
	m.Valid = false
	if strings.TrimSpace(m.Name) == "" {
		return ErrMissingMemberName
	}
	if len(strings.TrimSpace(m.Name)) < 3 {
		return ErrMemberNameTooShort
	}
	if len(strings.TrimSpace(m.Name)) > 20 {
		return ErrMemberNameTooLong
	}
	if strings.TrimSpace(m.LastName) == "" {
		return ErrMissingMemberLastName
	}
	if len(strings.TrimSpace(m.LastName)) < 3 {
		return ErrMemberLastNameTooShort
	}
	if len(strings.TrimSpace(m.LastName)) > 20 {
		return ErrMemberLastNameTooLong
	}
	if m.DOB.IsZero() {
		return ErrMissingMemberDOB
	}

	gender := strings.TrimSpace(m.Gender)
	if gender == "" {
		return ErrMissingMembeGender
	}
	if gender != Male && gender != Female && gender != Other {
		return ErrInvalidMembeGender
	}

	if strings.TrimSpace(m.ID) == "" {
		return ErrMissingMemberID
	}

	m.Valid = true
	return nil
}

// Return the full name of this member
// in casual mode
func (m *Member) FullName() string {
	builder := strings.Builder{}
	fullName := ""

	if len(m.MiddleName) > 0 {
		builder.WriteString(m.Name)
		builder.WriteString(" ")
		builder.WriteString(m.MiddleName)
		builder.WriteString(" ")
		builder.WriteString(m.LastName)
		fullName = builder.String()

	} else {
		builder.WriteString(m.Name)
		builder.WriteString(" ")
		builder.WriteString(m.LastName)
		fullName = builder.String()
	}

	return strings.TrimSpace(fullName)
}

// Return the full name of this member
// in formal mode
func (m *Member) FullNameFormal() string {
	builder := strings.Builder{}
	fullNameTmp := m.FullName()
	fullName := ""
	if len(fullNameTmp) > 0 {

		if m.Gender == Male {
			builder.WriteString("Sr. ")
			builder.WriteString(fullNameTmp)
			fullName = builder.String()

		}
		if m.Gender == Female {
			builder.WriteString("Sra. ")
			builder.WriteString(fullNameTmp)
			fullName = builder.String()
		}

		// if len(fullName) <= 5 {
		// 	fullName = ""
		// }
	} else {
		fullName = "Palmeiras Campeão"
	}
	return strings.TrimSpace(fullName)
}
