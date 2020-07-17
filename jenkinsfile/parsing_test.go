package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJenkinsFileParsing(t *testing.T) {
	t.Run("Empty file", func(t *testing.T) {
		context, err := ParseJenkinsFile("testdata/empty")
		assert.NoError(t, err)
		assert.NotNil(t, context)
	})

	t.Run("Simple library reference", func(t *testing.T) {
		context, err := ParseJenkinsFile("testdata/library-ref")
		assert.NoError(t, err)
		assert.Equal(t, "'my-shared-library'", context.GetRef().GetLib().GetText())
		antlr.ParseTreeWalkerDefault.Walk(&jenkinsListener{}, context)
	})
	t.Run("Basic pipeline", func(t *testing.T) {
		_, _ = fmt.Fprintf(os.Stderr, "Hello from teh test")
		context, err := ParseJenkinsFile("testdata/basic")
		assert.NoError(t, err)
		assert.NotNil(t, context)

		listener := jenkinsListener{}

		antlr.ParseTreeWalkerDefault.Walk(&listener, context)
		assert.NotNil(t, context.GetPipe())
		assert.Equal(t, "any", context.GetPipe().GetAgent().GetAgentName().GetText())
		assert.Fail(t, "You asked for it")
	})
}
