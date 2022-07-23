package family

const (
	Male   string = "Male"
	Female string = "Female"
	Other  string = "Other"
)

const (
	Newborn   string = "Newborn"
	Child     string = "Child"
	Toddler   string = "Toddler"
	Infant    string = "Infant"
	Teen      string = "Teenr"
	Adult     string = "Adult"
	Elderly   string = "Elderly"
	Undefined string = "Undefined"
)

const (
	RelFather        string = "Father"
	RelMother        string = "Mother"
	RelBrother       string = "Brother"
	RelSister        string = "Sister"
	RelDaughter      string = "Daughter"
	RelSon           string = "Son"
	RelHusband       string = "Husband"
	RelWife          string = "Wife"
	RelStepFather    string = "StepFather"
	RelStepMother    string = "StepMother"
	RelStepBrother   string = "StepBrother"
	RelStepSister    string = "StepSister"
	RelStepDaughter  string = "StepDaughter"
	RelStepSon       string = "StepSon"
	RelFatherInLaw   string = "FatherInLaw"
	RelMotherInLaw   string = "MotherInLaw"
	RelBrotherInLaw  string = "BrotherInLaw"
	RelSisterInLaw   string = "SisterInLaw"
	RelDaughterInLaw string = "DaughterInLaw"
	RelSonInLaw      string = "SonInLaw"
	RelTBDRelation   string = "ToBeDefined"
	RelSelf          string = "Self"
)

const (
	FreshMember    string = "fresh-member"
	ExistingMember string = "existing-member"
	RemovedMember  string = "removed-member"
	UpdatedMember  string = "updated-member"
)

var relations map[string]string

func init() {
	relations = make(map[string]string)
	relations[RelTBDRelation] = RelTBDRelation
	relations[RelSelf] = RelSelf
	relations[RelFather] = RelFather
	relations[RelMother] = RelMother
	relations[RelBrother] = RelBrother
	relations[RelSister] = RelSister
	relations[RelDaughter] = RelDaughter
	relations[RelSon] = RelSon
	relations[RelHusband] = RelHusband
	relations[RelWife] = RelWife
	relations[RelStepFather] = RelStepFather
	relations[RelStepMother] = RelStepMother
	relations[RelStepBrother] = RelStepBrother
	relations[RelStepSister] = RelStepSister
	relations[RelStepDaughter] = RelStepDaughter
	relations[RelStepSon] = RelStepSon
	relations[RelFatherInLaw] = RelFatherInLaw
	relations[RelMotherInLaw] = RelMotherInLaw
	relations[RelBrotherInLaw] = RelBrotherInLaw
	relations[RelSisterInLaw] = RelSisterInLaw
	relations[RelDaughterInLaw] = RelDaughterInLaw
	relations[RelSonInLaw] = RelSonInLaw
}
