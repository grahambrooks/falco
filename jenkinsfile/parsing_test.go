package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJenkinsFileParsing(t *testing.T) {
	t.Run("Empty file", func(t *testing.T) {
		_, err := ParseJenkinsFile("testdata/empty")
		assert.NoError(t, err)
	})

	t.Run("Empty file", func(t *testing.T) {
		context, err := ParseJenkinsFile("testdata/library-ref")
		assert.NoError(t, err)
		assert.Equal(t, "'my-shared-library'", context.GetRef().GetLib().GetText())
	})
}

