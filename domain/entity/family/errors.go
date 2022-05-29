package family

import "errors"

// Errors related to the Family
var (
	ErrMissingFamilyID      = errors.New("o ID da família está em branco ou ausente")
	ErrMissingFamilySurname = errors.New("o nome de família está em branco ou ausente")
)

// Errors related to the Family Member (one Member)
var (
	ErrMissingMemberID    = errors.New("o ID da pessoa está em branco ou ausente")
	ErrMissingMemberName = errors.New("o nome da pessoa está em branco ou ausente")
	ErrMemberNameTooShort = errors.New("o nome da pessoa é muito pequeno")
	ErrMemberNameTooLong = errors.New("o nome da pessoa é muito grande")
	ErrMissingMemberLastName = errors.New("o sobrenome da pessoa está em branco ou ausente")
	ErrMemberLastNameTooShort = errors.New("o sobrenome da pessoa é muito pequeno")
	ErrMemberLastNameTooLong = errors.New("o sobrenomenome da pessoa é muito grande")
	ErrMissingMemberDOB = errors.New("a data de nascimento da pessoa está em branco ou ausente")
	ErrInvalidMemberDOB = errors.New("a data de nascimento da pessoa é inválido")
)
