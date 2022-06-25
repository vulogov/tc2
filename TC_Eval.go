package tc2

import (
  "fmt"
  "strings"
  "github.com/pterm/pterm"
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/vulogov/tc2/parser"
)

func (tc *TCstate) Eval(code string) *TCstate {
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
  if tc.errors > 0 || errorListener.errors > 0 {
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

func (l *tcExecErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
  l.TC.console.Error("Syntax error line=",line,"column=",column, msg)
  clines := splitLines(*l.code)
  l.TC.console.Print(pterm.FgLightGreen.Sprint(clines[line-1]))
  pointer := strings.Repeat(" ", column) + "^"
  l.TC.console.Print(pterm.FgLightRed.Sprint(pointer))
  l.errors += 1
  l.TC.errors += 1
}
func (l *tcExecErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
  msgout := fmt.Sprintf("Ambiguity Error")
  l.TC.console.Error(msgout)
	l.errors += 1
}
func (l *tcExecErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
  msgout := fmt.Sprintf("Attempting in Full Context")
  l.TC.console.Error(msgout)
	l.errors += 1
}
func (l *tcExecErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
  msgout := fmt.Sprintf("Context sensitivity error")
  l.TC.console.Error(msgout)
	l.errors += 1
}
