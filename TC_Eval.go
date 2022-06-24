package tc2

import (
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/vulogov/tc2/parser"
)

func (tc *TCstate) Eval(code string) *TCstate {
  tc.console.Debug("Eval:", code)
  errorListener := new(tcExecErrorListener)
  errorListener.code = &code
  _input := antlr.NewInputStream(code)
  lexer := parser.NewBundLexer(_input)
  stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
  p := parser.NewBundParser(stream)
  p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
  listener := new(TCExecListener)
  listener.TC = tc
  tc.ts       = listener
  errorListener.TC = tc
  if errorListener.errors > 0 {
    tc.errors = errorListener.errors
  }
  if errorListener.errors > 0 {
		tc.console.Critical("Lexer errors detected:", errorListener.errors)
		return nil
	}
  antlr.ParseTreeWalkerDefault.Walk(listener, p.Root())
  if errorListener.errors > 0 {
    tc.errors = errorListener.errors
  }
  if errorListener.errors > 0 {
		tc.console.Critical("Errors detected:", errorListener.errors)
		return nil
	}
  return tc
}
