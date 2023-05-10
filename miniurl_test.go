package miniurl_test

import (
	"github.com/JarnoLahti/miniurl-fork"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashLength(t *testing.T) {
	const (
		input          = "test_input"
		expectedLength = 32
	)

	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}
