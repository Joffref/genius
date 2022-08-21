package uuid

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	u, err := New()
	require.NoError(t, err)
	require.IsType(t, "string", u)
}

func TestGenerate(t *testing.T) {
	u, err := generate()
	require.NoError(t, err)
	require.IsType(t, &uuid.UUID{}, u)
}
