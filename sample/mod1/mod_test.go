package mod1_test

import (
	"github.com/keanpedersen/go-jester/sample/mod1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFizzBuzz0(t *testing.T) {
	ret := mod1.FizzBuzz(0)
	require.Empty(t,ret)
}

func TestFizzBuzz3(t *testing.T) {
	ret := mod1.FizzBuzz(3)
	require.Len(t, ret, 3)
	assert.Equal(t, "1", ret[0])
	assert.Equal(t, "2", ret[1])
	assert.Equal(t, "fizz", ret[2])
}