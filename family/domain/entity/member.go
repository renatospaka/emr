package family

import (
	"strings"
	"time"

	"github.com/renatospaka/emr/infrastructure/utils"
)

var (
	analysisErrsMembers   = utils.NewAnalysisErrs()
	clearErrsOnValidation = true
)

type Member struct {
	id           string    `json:"family_member_id"`
	name         string    `json:"name"`
	lastName     string    `json:"last_name"`
	middleName   string    `json:"middle_name"`
	nickname     string    `json:"nick_name"`
	gender       string    `json:"gender"`
	age          string    `json:"age_of"`
	ageInMonths  int64     `json:"age_in_months"`
	ageInYears   int64     `json:"age_in_years"`
	lastChanged  int64     `json:"-"`
	valid        bool      `json:"-"`
	headOfFamily bool      `json:"head_family"`
	dob          time.Time `json:"day_of_birth"`
}

func newMember() *Member {
	// log.Println("Member.newMember()")
	clearErrsOnValidation = true
	member := &Member{
		id:           utils.GetID(),
		name:         "",
		lastName:     "",
		middleName:   "",
		nickname:     "",
		gender:       "",
		valid:        false,
		headOfFamily: false,
		dob:          time.Time{},
		lastChanged:  time.Now().UnixNano(),
	}
	return member
}

// Return the ID of this member
func (m *Member) ID() string {
	return m.id
}

// Set the full name of this member
// name + middle name + last name
func (m *Member) ChangeFullName(name string, middleName string, lastName string) *Member {
	// log.Println("Member.ChangeFullName()")
	m.name = strings.TrimSpace(name)
	m.lastName = strings.TrimSpace(lastName)
	m.middleName = strings.TrimSpace(middleName)
	m.lastChanged = time.Now().UnixNano()
	m.valid = false
	return m
}

// Return the full name of this member
// in casual mode
func (m *Member) FullName() string {
	// log.Println("Member.FullName()")
	builder := strings.Builder{}
	fullName := ""

	if len(m.middleName) > 0 {
		builder.WriteString(m.name)
		builder.WriteString(" ")
		builder.WriteString(m.middleName)
		builder.WriteString(" ")
		builder.WriteString(m.lastName)
		fullName = builder.String()
	} else {
		builder.WriteString(m.name)
		builder.WriteString(" ")
		builder.WriteString(m.lastName)
		fullName = builder.String()
	}
	return strings.TrimSpace(fullName)
}

// Return the full name of this member
// in formal mode
func (m *Member) FullNameFormal() string {
	// log.Println("Member.FullNameFormal()")
	builder := strings.Builder{}
	fullNameTmp := m.FullName()
	fullName := ""
	if len(fullNameTmp) > 0 {
		if m.gender == Male {
			builder.WriteString("Sr. ")
			builder.WriteString(fullNameTmp)
			fullName = builder.String()
		} else if m.gender == Female {
			builder.WriteString("Sra. ")
			builder.WriteString(fullNameTmp)
			fullName = builder.String()
		} else {
			fullName = fullNameTmp
		}
	}
	return strings.TrimSpace(fullName)
}

// Set a nickname for this member
func (m *Member) ChangeNickname(nick string) *Member {
	// log.Println("Member.ChangeNickname()")
	m.nickname = strings.TrimSpace(nick)
	m.lastChanged = time.Now().UnixNano()
	m.valid = false
	return m
}

// Return the nickname for this member, if any
func (m *Member) Nickname() string {
	// log.Println("Member.Nickname()")
	return m.nickname
}

// Set the day of birth of this member
func (m *Member) ChangeBirthDate(dob time.Time) *Member {
	// log.Println("Member.ChangeBirthDate()")
	m.dob = dob
	m.lastChanged = time.Now().UnixNano()
	m.valid = false
	m.setAge()
	return m
}

// Return the day of birth of this member
func (m *Member) BirthDate() time.Time {
	// log.Println("Member.BirthDate()")
	return m.dob
}

// Set the gender of this member
func (m *Member) ChangeGender(gender string) *Member {
	// log.Println("Member.ChangeGender()")
	m.gender = gender
	m.lastChanged = time.Now().UnixNano()
	m.valid = false
	return m
}

// Return the gender of this member
func (m *Member) Gender() string {
	// log.Println("Member.Gender()")
	return m.gender
}

// Return the age of the member in years since birth
func (m *Member) AgeInYears() int64 {
	// log.Println("Member.AgeInYears()")
	return m.ageInYears
}

// Return the age of the member in months since birth
func (m *Member) AgeInMonths() int64 {
	// log.Println("Member.AgeInMonths()")
	return m.ageInMonths
}

// Return if the member is a newborn
func (m *Member) IsNewborn() bool {
	// log.Println("Member.IsNewborn()")
	return m.age == Newborn
}

// Return if the member is a infant
func (m *Member) IsInfant() bool {
	// log.Println("Member.IsInfant()")
	return m.age == Infant
}

// Return if the member is a toddler
func (m *Member) IsToddler() bool {
	// log.Println("Member.IsToddler()")
	return m.age == Toddler
}

// Return if the member is a child
func (m *Member) IsChild() bool {
	// log.Println("Member.IsChild()")
	return m.age == Child
}

// Return if the member is a teen
func (m *Member) IsTeen() bool {
	// log.Println("Member.IsTeen()")
	return m.age == Teen
}

// Return if the member is a adult
func (m *Member) IsAdult() bool {
	// log.Println("Member.IsAdult()")
	return m.age == Adult
}

// Return if the member is a elderly
func (m *Member) IsElderly() bool {
	// log.Println("Member.IsElderly()")
	return m.age == Elderly
}

// Return all errors found during the validation process
// in a single string with a \n segregating each error
func (m *Member) Err() string {
	// log.Println("Member.Err()")
	analysis := ""
	if analysisErrsMembers.Count() > 0 {
		builder := strings.Builder{}
		for e := 0; e < len(analysisErrsMembers.Analysis); e++ {
			builder.WriteString(analysisErrsMembers.Analysis[e].ErrDescription)
			builder.WriteString("\n")
		}
		analysis = builder.String()
	}
	return analysis
}

// Return all errors found during the validation process
// in an array
func (m *Member) ErrToArray() []string {
	// log.Println("Member.ErrToArray()")
	analysis := m.Err()
	toArray := []string{}
	if len(analysis) > 0 {
		newAnalisys := strings.Split(analysis, "\n")
		for e := 0; e < len(newAnalisys)-1; e++ {
			toArray = append(toArray, newAnalisys[e])
		}
	}
	return toArray
}

// Return if the member is validated
// Use this whenever you want to guarantee the integrity of the structure
func (m *Member) IsValid() bool {
	// log.Println("Member.IsValid() - starting")
	clearErrsOnValidation = true
	m.validate()
	// log.Println("Member.IsValid(", m.valid, ")")
	return m.valid
}

// Calculate the age of this member in months, year
// and classify his/her age accordingly
// every time the day of birth changes
func (m *Member) setAge() {
	// log.Println("Member.setAge()")
	m.age = Undefined
	m.ageInMonths = 0
	m.ageInYears = 0
	m.lastChanged = time.Now().UnixNano()
	if m.dob.IsZero() {
		return
	}

	// Calculate age in months and years based on the DOB
	ageInYears, ageInMonths := utils.AgeFromToday(m.dob)
	ageOf := ""

	// uses this artcle to identify which phase of life (in a young phase) is today
	// https://www.verywellfamily.com/difference-between-baby-newborn-infant-toddler-113848
	if ageInMonths <= 2 {
		ageOf = Newborn
	} else if ageInMonths > 2 && ageInMonths <= 24 {
		ageOf = Infant
	} else if ageInMonths > 24 && ageInMonths < 60 {
		ageOf = Toddler
	} else if ageInYears >= 5 && ageInYears < 12 {
		ageOf = Child
	} else if ageInYears >= 12 && ageInYears < 18 {
		ageOf = Teen
	} else if ageInYears >= 18 && ageInYears < 65 {
		ageOf = Adult
	} else if ageInYears >= 65 {
		ageOf = Elderly
	} else {
		ageOf = Undefined
	}

	m.age = ageOf
	m.ageInYears = ageInYears
	m.ageInMonths = ageInMonths
	m.lastChanged = time.Now().UnixNano()
	m.valid = false
}

// Check whenever the member structure is intact
// and filled accordingly to the model rules
func (m *Member) validate() {
	// log.Printf("Member.validate() -> clearErrsOnValidation: %t", clearErrsOnValidation)
	if clearErrsOnValidation {
		analysisErrsMembers.RemoveAll()
	}

	// test if all properties are nil or empty
	if m.id == "" &&
		m.name == "" &&
		m.lastName == "" &&
		m.gender == "" &&
		m.dob.IsZero() {
		analysisErrsMembers.AddErr(ErrInvalidMember)
		m.valid = (analysisErrsMembers.Count() == 0)
		// log.Printf("Member.validate(%t)", m.valid)
		return
	}

	// test each property individually

	if m.name == "" {
		analysisErrsMembers.AddErr(ErrMissingMemberName)
	} else if len(m.name) < 3 {
		analysisErrsMembers.AddErr(ErrMemberNameTooShort)
	} else if len(m.name) > 20 {
		analysisErrsMembers.AddErr(ErrMemberNameTooLong)
	}

	if len(m.middleName) > 20 {
		analysisErrsMembers.AddErr(ErrMemberMiddleNameTooLong)
	}

	if m.lastName == "" {
		analysisErrsMembers.AddErr(ErrMissingMemberLastName)
	} else if len(m.lastName) < 3 {
		analysisErrsMembers.AddErr(ErrMemberLastNameTooShort)
	} else if len(m.lastName) > 20 {
		analysisErrsMembers.AddErr(ErrMemberLastNameTooLong)
	}

	gender := m.gender
	if gender == "" {
		analysisErrsMembers.AddErr(ErrMissingMemberGender)
	} else if gender != Male && gender != Female && gender != Other {
		analysisErrsMembers.AddErr(ErrInvalidMemberGender)
	}

	if m.dob.IsZero() {
		analysisErrsMembers.AddErr(ErrMissingMemberDOB)
	}

	if m.id == "" {
		analysisErrsMembers.AddErr(ErrMissingMemberID)
	}
	m.valid = (analysisErrsMembers.Count() == 0)
	// log.Printf("Member.validate(%t)", m.valid)
}
