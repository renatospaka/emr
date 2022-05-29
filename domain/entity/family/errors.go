package family

import "errors"

// Errors related to the Family
var (
	ErrMissingFamilySurname = errors.New("o nome de família está em branco ou ausente")
	ErrMissingFamilyID      = errors.New("o ID da família está em branco ou ausente")
)
