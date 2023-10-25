package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	c := require.New(t)

	c.True(true, "Test")
}
