package tc2

import (
  "github.com/antlr/antlr4/runtime/Go/antlr"
)

var VERSION = "2.00"

type TCExecListener struct {
  *BaseBundParserListener
  TC          *TCstate
}

type TCExtType    func(interface{}) int
type TCExtTypeStr func(interface{}) string


type TCstate struct {
  ID           string
}

type tcExecErrorListener struct {
	antlr.ErrorListener
  TC     *TCstate
	code   *string
	errors int
}

func Init() *TCstate {
  tc := new(TCstate)
  return tc
}
