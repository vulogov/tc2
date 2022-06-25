package tc2

import (
  "os"
  "fmt"
  "sync"
  "github.com/pterm/pterm"
  "github.com/gammazero/deque"
  "github.com/google/uuid"
  "github.com/vulogov/tconsole"
  "github.com/lrita/cmap"
  "github.com/deckarep/golang-set"
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/vulogov/tc2/parser"
  "github.com/loveleshsharma/gohive"
  "github.com/srfrog/dict"
)

var Functions           cmap.Cmap
var ExpressionCBChain   deque.Deque[interface{}]
var Callbacks           cmap.Cmap
var extType             TCExtType       // Functions for external types generation
var extTypeStr          TCExtTypeStr    // and conversion
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
  errmsg       string
  HandleErr    bool
  ErrStack     deque.Deque[interface{}]  // Stack of errors
  console      *tconsole.TConsole
  stack        []pterm.LeveledListItem
  Res          *TwoStack                  // main stack
  ResNames     deque.Deque[interface{}]   // stack of stack names
  ResN         mapset.Set                 // set of stack names
  StackList    cmap.Cmap                  // Reference to stacks
  StackChan    cmap.Cmap                  // Reference to stack channels
  Vars         cmap.Cmap                  // variables
  Functions    cmap.Cmap                  // Functions
  UFunctions   cmap.Cmap                  // User Functions
  UMacros      cmap.Cmap                  // User Macros
  FCStack      deque.Deque[string]        // Function call stack
  CodeStack    deque.Deque[string]        // Accumulation for codeblocks
  FCodeStack   deque.Deque[string]        // Accumulation for function code
  MCodeStack   deque.Deque[string]        // Accumulation for macro code
  Ctx          cmap.Cmap                  // Global context
  CtxStack     deque.Deque[interface{}]   // Stack of local contexts
  Wg           sync.WaitGroup             // Global wait group
  Pool        *gohive.PoolService         // Global execution pool
  ExReq       chan bool                   // Exit request channel
  IsExitReq   bool                        // Exit flag
  ExitCb      *dict.Dict                  // Exit callbacks
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
  tc.stack      = make([]pterm.LeveledListItem, 0)
  if log == nil  {
    log = tc.console
  }
  tc.console.Debug("BUND version", VERSION)
  tc.console.Debug("BUND instance", tc.ID, "initialized")
  tc.AddNewStack(uuid.NewString())
  return tc
}

func (tc *TCstate) ClearErrors() {
  tc.Debug("Clearing errors")
  tc.errors = 0
  tc.errmsg = ""
}

func (tc *TCstate) Errors() int {
  if tc.HandleErr {
    return 0
  }
  return tc.errors
}

func (tc *TCstate) TrueErrors() int {
  return tc.errors
}

func (tc *TCstate) Error() string {
  return tc.errmsg
}

func (tc *TCstate) Debug(msg ...interface{}) {
  if tc.console == nil {
    return
  }
  tc.console.Debug(msg...)
}

func (tc *TCstate) Debugf(msg string, args ...interface{}) {
  if tc.console == nil {
    return
  }
  tc.console.Debug(fmt.Sprintf(msg, args...))
}

func (tc *TCstate) GoEval(code string) *TCstate {
  tc.Wg.Add(1)
  task := func() {
    defer tc.Wg.Done()
    tc.Eval(code)
  }
  tc.Pool.Submit(task)
  return tc
}

func (tc *TCstate) SetError(msg string, args ...interface{}) {
  tc.errmsg = fmt.Sprintf(msg, args...)
  tc.errors += 1
  tc.console.Error(tc.errmsg)
}

func (l *TCExecListener) SetError(msg string, args ...interface{}) {
  l.TC.SetError(msg, args...)
}
