package equality_test

import (
	"github.com/keanpedersen/go-jester/sample/equality"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEquality(t *testing.T) {
	require.Equal(t, 1, equality.Equality(6, 1))
	require.Equal(t, 1, equality.Equality(5, 1))
}
