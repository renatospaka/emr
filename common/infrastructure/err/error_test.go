package err_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/renatospaka/emr/common/infrastructure/err"
)

var (
	errorTest  = errors.New("testing a new error")
	errorTest2 = errors.New("testing another error")
)

func TestErr_Add(t *testing.T) {
	analysisErrors := err.NewErrors()
	analysisErrors.Add(errorTest)

	require.EqualValues(t, errorTest.Error(), analysisErrors.Err[0].Description)
}

func TestErr_Add_List(t *testing.T) {
	analysisErrors := err.NewErrors()
	analysisErrors.Add(errorTest)
	analysisErrors.Add(errorTest2)

	require.EqualValues(t, errorTest.Error(), analysisErrors.Err[0].Description)
	require.EqualValues(t, errorTest2.Error(), analysisErrors.Err[1].Description)
	require.EqualValues(t, 2, analysisErrors.Count())
}

func TestErr_Count(t *testing.T) {
	analysisErrors := err.NewErrors()
	analysisErrors.Add(errorTest)

	require.EqualValues(t, 1, analysisErrors.Count())
}

func TestErr_Count_Empty(t *testing.T) {
	analysisErrors := err.NewErrors()

	require.EqualValues(t, 0, analysisErrors.Count())
}

func TestErr_ClearAll(t *testing.T) {
	analysisErrors := err.NewErrors()
	analysisErrors.Add(errorTest)
	analysisErrors.ClearAll()

	require.Nil(t, analysisErrors.Err)
}
