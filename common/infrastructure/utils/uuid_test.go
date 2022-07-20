package utils_test

import (
	"strings"
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/renatospaka/emr/common/infrastructure/utils"
)

func Test_GetID(t *testing.T) {
	id := utils.GetID()
	_, err := uuid.Parse(id)

	require.Nil(t, err)
}

func Test_GetID_InvalidLength(t *testing.T) {
	id := utils.GetID()

	builder := strings.Builder{}
	builder.WriteString(id)
	builder.WriteString("d")
	id2 := builder.String()
	_, err := uuid.Parse(id2)

	require.NotNil(t, err)
	require.EqualError(t, err, "invalid UUID length: 37")
}

func Test_GetID_InvalidFormat(t *testing.T) {
	id := utils.GetID()

	builder := strings.Builder{}
	builder.WriteString(id)
	builder.WriteString("dd")
	id2 := builder.String()
	_, err := uuid.Parse(id2)

	require.NotNil(t, err)
	require.EqualError(t, err, "invalid UUID format")
}

func Test_IsVaalidUUID(t *testing.T) {
	id := utils.GetID()
	err := utils.IsVaalidUUID(id)

	require.Nil(t, err)
}

func Test_IsVaalidUUID_InvalidLength(t *testing.T) {
	id := utils.GetID()

	builder := strings.Builder{}
	builder.WriteString(id)
	builder.WriteString("d")
	id2 := builder.String()
	err := utils.IsVaalidUUID(id2)

	require.NotNil(t, err)
	require.EqualError(t, err, "invalid UUID length: 37")
}

func Test_IsVaalidUUID_InvalidFormat(t *testing.T) {
	id := utils.GetID()

	builder := strings.Builder{}
	builder.WriteString(id)
	builder.WriteString("dd")
	id2 := builder.String()
	err := utils.IsVaalidUUID(id2)

	require.NotNil(t, err)
	require.EqualError(t, err, "invalid UUID format")
}
