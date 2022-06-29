package utils_test

import (
	"errors"
	"testing"

	"github.com/renatospaka/emr2/infrastructure/utils"
	"github.com/stretchr/testify/require"
)

var (
	errorTest = errors.New("testing a new error")
	errorTest2 = errors.New("testing another error")
)

func TestErr_AddErr(t *testing.T) {
	analysisErrors := utils.NewAnalysisErrs()
	analysisErrors.AddErr(errorTest)

	require.EqualValues(t, errorTest.Error(), analysisErrors.Analysis[0].ErrDescription)
}

func TestErr_AddErr_List(t *testing.T) {
	analysisErrors := utils.NewAnalysisErrs()
	analysisErrors.AddErr(errorTest)
	analysisErrors.AddErr(errorTest2)

	require.EqualValues(t, errorTest.Error(), analysisErrors.Analysis[0].ErrDescription)
	require.EqualValues(t, errorTest2.Error(), analysisErrors.Analysis[1].ErrDescription)
	require.EqualValues(t, 2, analysisErrors.Count())
}

func TestErr_Count(t *testing.T) {
	analysisErrors := utils.NewAnalysisErrs()
	analysisErrors.AddErr(errorTest)

	require.EqualValues(t, 1, analysisErrors.Count())
}

func TestErr_Count_Empty(t *testing.T) {
	analysisErrors := utils.NewAnalysisErrs()

	require.EqualValues(t, 0, analysisErrors.Count())
}

func TestErr_RemoveAll(t *testing.T) {
	analysisErrors := utils.NewAnalysisErrs()
	analysisErrors.AddErr(errorTest)
	analysisErrors.RemoveAll()

	require.Nil(t, analysisErrors.Analysis)
}
