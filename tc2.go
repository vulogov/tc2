package tc2

import (
  "os"
  "fmt"
  "github.com/gammazero/deque"
  "github.com/google/uuid"
  "github.com/vulogov/tconsole"
  "github.com/lrita/cmap"
  "github.com/deckarep/golang-set"
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/vulogov/tc2/parser"
)

var log *tconsole.TConsole = nil
var VERSION = "2.00"

type TCExecListener struct {
  *parser.BaseBundParserListener
  TC          *TCstate
}

type TCExtType    func(interface{}) int
type TCExtTypeStr func(interface{}) string


type TCstate struct {
  ts          *TCExecListener
  ID           string
  cfg          cmap.Cmap
  errors       int
  console      *tconsole.TConsole
  Res          *TwoStack                  // main stack
  ResNames     deque.Deque[interface{}]   // stack of stack names
  ResN         mapset.Set                 // set of stack names
  StackList    cmap.Cmap                  // Reference to stacks
  StackChan    cmap.Cmap                  // Reference to stack channels
  Vars         cmap.Cmap                  // variables
  Functions    cmap.Cmap                  // Functions
  UFunctions   cmap.Cmap                  // User Functions
}

type tcExecErrorListener struct {
	antlr.ErrorListener
  TC     *TCstate
	code   *string
	errors int
}

func Init() *TCstate {
  tc := new(TCstate)

  tc.ts     = nil
  tc.ID     = uuid.New().String()
  tc.Res    = InitTS(tc)
  tc.ResN   = mapset.NewSet()

  // Configure environment
  tc.SetVariable("tc.Version", VERSION)
  tc.SetVariable("tc.Maxfilesize", 16777216)
  tc.SetVariable("tc.Chancapacity", 102400)
  tc.SetVariable("tc.Poolsize", 25)
  tc.SetVariable("tc.Debuglevel", "info")
  tc.SetVariable("tc.Logoutput", os.Stderr)

  // Configure console
  tc.SetVariable("console.Enable", true)
  tc.SetVariable("console.Color", true)
  tc.SetVariable("console.Debug", true)
  tc.SetVariable("console.Info", true)
  tc.SetVariable("console.Warning", true)
  tc.SetVariable("console.Error", true)
  tc.SetVariable("console.Function", true)
  tc.SetVariable("console.Critical", true)
  tc.SetVariable("console.Msg", true)

  // And finally
  tc.console, _ = tconsole.New(&tc.Vars)
  if log == nil  {
    log = tc.console
  }
  tc.console.Debug("BUND version", VERSION)
  tc.console.Debug("BUND instance", tc.ID, "initialized")
  tc.AddNewStack(uuid.NewString())
  return tc
}

func (tc *TCstate) Debug(msg ...interface{}) {
  if tc.console == nil {
    return
  }
  tc.console.Debug(msg...)
}

func (l *tcExecErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
  l.TC.console.Error("Syntax error line=",line,"column=",column, msg)
	l.errors += 1
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
