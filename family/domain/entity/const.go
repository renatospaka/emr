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
	Father        string = "Father"
	Mother        string = "Mother"
	Brother       string = "Brother"
	Sister        string = "Sister"
	Daughter      string = "Daughter"
	Son           string = "Son"
	Husband       string = "Husband"
	Wife          string = "Wife"
	StepFather    string = "StepFather"
	StepMother    string = "StepMother"
	StepBrother   string = "StepBrother"
	StepSister    string = "StepSister"
	StepDaughter  string = "StepDaughter"
	StepSon       string = "StepSon"
	FatherInLaw   string = "FatherInLaw"
	MotherInLaw   string = "MotherInLaw"
	BrotherInLaw  string = "BrotherInLaw"
	SisterInLaw   string = "SisterInLaw"
	DaughterInLaw string = "DaughterInLaw"
	SonInLaw      string = "SonInLaw"
	TBDRelation   string = "ToBeDefined"
	Self          string = "Self"
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
	relations[TBDRelation] = TBDRelation
	relations[Self] = Self
	relations[Father] = Father
	relations[Mother] = Mother
	relations[Brother] = Brother
	relations[Sister] = Sister
	relations[Daughter] = Daughter
	relations[Son] = Son
	relations[Husband] = Husband
	relations[Wife] = Wife
	relations[StepFather] = StepFather
	relations[StepMother] = StepMother
	relations[StepBrother] = StepBrother
	relations[StepSister] = StepSister
	relations[StepDaughter] = StepDaughter
	relations[StepSon] = StepSon
	relations[FatherInLaw] = FatherInLaw
	relations[MotherInLaw] = MotherInLaw
	relations[BrotherInLaw] = BrotherInLaw
	relations[SisterInLaw] = SisterInLaw
	relations[DaughterInLaw] = DaughterInLaw
	relations[SonInLaw] = SonInLaw
}
