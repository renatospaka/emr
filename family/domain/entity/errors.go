package family

import "errors"

// Errors related to the Family
var (
	ErrInvalidFamily        = errors.New("todas as propriedades dessa família estão vazias")
	ErrInvalidFamilyID      = errors.New("o ID da família é inválido")
	ErrMissingFamilyID      = errors.New("o ID da família está em branco ou ausente")
	ErrMissingFamilySurname = errors.New("o nome de família está em branco ou ausente")
	ErrFamilyNotFound       = errors.New("família não encontrada")
	ErrFamilyError          = errors.New("não há informações suficientes para criar essa família")
)

// Errors related to Family Member
var (
	ErrInvalidFamilyMember         = errors.New("todas as propriedades desse membro  estão vazias")
	ErrFamilyMemberAlreadyLinked   = errors.New("a pessoa informada já pertence a essa família")
	ErrFamilyMemberNotLinked       = errors.New("a pessoa informada não pertence a essa família")
	ErrFamilyMemberNotRelated      = errors.New("a pessoa informada não tem vínculo familiar definido nessa família")
	ErrFamilyMemberMissing         = errors.New("não há membros nessa família")
	ErrFamilyMemberHOFMissing      = errors.New("não há uma pessoa responsável definida para essa família")
	ErrFamilyMemberHOFInvalidAge   = errors.New("a pessoa responsável definida para essa família deve ser maior de idade")
	ErrFamilyMemberHOFError        = errors.New("essa pessoa não pode possui os critérios para ser responsável por uma família")
	ErrFamilyMemberTooManyHOF      = errors.New("só pode haver um responsável por família")
	ErrFamilyMemberInvalidRelation = errors.New("o vínculo familiar é inválido ")
)

// Errors related to Member (one Member)
var (
	ErrInvalidMember           = errors.New("todas as propriedades dessa pessoa estão vazias")
	ErrInvalidMemberID         = errors.New("o ID da pessoa é inválido")
	ErrMissingMemberID         = errors.New("o ID da pessoa está em branco ou ausente")
	ErrMissingMemberName       = errors.New("o nome da pessoa está em branco ou ausente")
	ErrMemberNameTooShort      = errors.New("o nome da pessoa é muito pequeno")
	ErrMemberNameTooLong       = errors.New("o nome da pessoa é muito grande")
	ErrMissingMemberLastName   = errors.New("o sobrenome da pessoa está em branco ou ausente")
	ErrMemberLastNameTooShort  = errors.New("o sobrenome da pessoa é muito pequeno")
	ErrMemberLastNameTooLong   = errors.New("o sobrenome da pessoa é muito grande")
	ErrMemberMiddleNameTooLong = errors.New("o nome do meio da pessoa é muito grande")
	ErrMissingMemberDOB        = errors.New("a data de nascimento está em branco ou ausente")
	ErrInvalidMemberDOB        = errors.New("a data de nascimento é inválida")
	ErrMissingMemberGender     = errors.New("o gênero da pessoa está em branco ou ausente")
	ErrInvalidMemberGender     = errors.New("o gênero da pessoa é inválido")
	ErrMemberNotFound          = errors.New("pessoa não encontrada")
	ErrMemberError             = errors.New("não há informações suficientes para criar essa pessoa")
)

// Common errors
var (
	ErrNoChangesNeeded = errors.New("nenhuma alteração necessária")
	ErrIndexOutOfRange = errors.New("índice fora de escala")
)
