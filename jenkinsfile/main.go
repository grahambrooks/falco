package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	fmt.Printf("jenkinsfile")
	// Setup the input
	is := antlr.NewInputStream("@Library('my-shared-library') _")

	// Create the Lexer
	lexer := NewjenkinsLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := NewjenkinsParser(stream)


	context := p.Start()

	fmt.Printf("Parser response %#v\n", context)

	fmt.Printf("Library reference %s\n", context.GetRef().GetLib().GetText())


	// Finally parse the expression

	//antlr.ParseTreeWalkerDefault.Walk(&jenkinsListener{}, p.Start())
}
