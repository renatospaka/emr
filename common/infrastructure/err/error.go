package err

type Errs struct {
	Err         error
	Description string
	Severity		string
}

type Errors struct {
	Err []*Errs
}

// Initiate the Errs structure
func NewErrors() *Errors {
	return &Errors{
		Err: []*Errs{},
	}
}

// (Initialize and) add a new error into the slice
func (a *Errors) Add(err error) *Errors {
	thisErr := Errs{err, err.Error(), ""}
	a.Err = append(a.Err, &thisErr)
	return a
}

// Clearing the slice completely
func (a *Errors) ClearAll() *Errors {
	a.Err = nil
	return a
}

// Count how many errors there is in the current Err
func (a *Errors) Count() int {
	return len(a.Err)
}
