package utils

type Errs struct {
	ErrType        error
	ErrDescription string
}

type AnalysisErrs struct {
	Analysis []*Errs
}

func NewAnalysisErrs() *AnalysisErrs {
	return &AnalysisErrs{
		Analysis: []*Errs{},
	}
}


// (Initialize and) add a new error into the slice 
func (a *AnalysisErrs) AddErr(err error) *AnalysisErrs {
	thisErr := Errs{err, err.Error()}
	a.Analysis = append(a.Analysis, &thisErr)
	return a
}

// Completely clear the slice
func (a *AnalysisErrs) RemoveAll() *AnalysisErrs {
	a.Analysis = nil
	return a
}

// Remove a specific item of the slice
func (a *AnalysisErrs) RemoveOne(index int) *AnalysisErrs {
		// perform bounds checking first to prevent a panic!
		if index >= len(a.Analysis) || index < 0 {
			return a
		}
	
		a.Analysis[index] = a.Analysis[len(a.Analysis)-1]
		// return a.Analysis[:len(a.Analysis)-1]
		return a
}


// Count how many errors there is in the current analysis
// it is not a zero-based function.
func (a *AnalysisErrs) Count() int {
	// if len(a.Analysis) > 0 {
	// 	return len(a.Analysis)
	// }
	return len(a.Analysis)
}