package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	fmt.Printf("jenkinsfile")
}

type jenkinsListener struct {
	*BaseJenkinsListener
}

func (jenkinsListener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("Visiting error node %v\n", node)
}

func (s *jenkinsListener) EnterLibrary_ref(ctx *Library_refContext) {
	fmt.Println("Enter lib ref")
}

// ExitLibrary_ref is called when production library_ref is exited.
func (s *jenkinsListener) ExitLibrary_ref(ctx *Library_refContext) {
	fmt.Println("Exit lib ref")
}

type JenkinsfileErrorStrategy struct {
	*antlr.ErrorStrategy
	ErrorCount int
}

func (j *JenkinsfileErrorStrategy) reset(antlr.Parser) {
	panic("implement me")
}

func (j *JenkinsfileErrorStrategy) RecoverInline(antlr.Parser) antlr.Token {
	panic("implement me")
}

func (j *JenkinsfileErrorStrategy) Recover(antlr.Parser, antlr.RecognitionException) {
	panic("implement me")
}

func (j *JenkinsfileErrorStrategy) Sync(antlr.Parser) {
	panic("implement me")
}

func (j *JenkinsfileErrorStrategy) inErrorRecoveryMode(antlr.Parser) bool {
	panic("implement me")
}

func (j *JenkinsfileErrorStrategy) ReportMatch(antlr.Parser) {
	panic("implement me")
}

func (j *JenkinsfileErrorStrategy) ReportError(parser antlr.Parser, e antlr.RecognitionException) {
	j.ErrorStrategy.ReportError(parser, e)
	j.ErrorCount++
}

func (j *JenkinsfileErrorStrategy) Errors() error {
	if j.ErrorCount > 0 {
		return fmt.Errorf("%d errors encountered during parsing", j.ErrorCount)
	}
	return nil
}

func ParseJenkinsFile(filename string) (IStartContext, error) {
	is, err := antlr.NewFileStream(filename)

	if err != nil {
		return nil, err
	}

	// Create the Lexer
	lexer := NewJenkinsLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewJenkinsParser(stream)

	return p.Start(), strategy.Errors()
}
