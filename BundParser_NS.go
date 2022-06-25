package tc2

import (
  "fmt"
  "github.com/google/uuid"
  "github.com/vulogov/tc2/parser"
)

func (tc *TCExecListener) EnterNs_declaration(ctx *parser.Ns_declarationContext) {
  var nsname string

  if ctx.GetName() != nil {
    nsname = ctx.GetName().GetText()
  } else {
    nsname = uuid.New().String()
  }
  tc.TC.console.Section("NS=",nsname)
  if tc.TC.ResN.Contains(nsname) == false {
    tc.TC.AddNewStack(nsname)
  } else {
    err := tc.TC.PositionStack(nsname)
    tc.TC.console.IfError(err)
    if err != nil {
      tc.TC.SetError(fmt.Sprintf("error: %v", err))
    }
  }
}

func (tc *TCExecListener) ExitNs_declaration(ctx *parser.Ns_declarationContext) {
  nsname := tc.TC.ResNames.Front().(string)
  tc.TC.console.SectionEnd("NS=", nsname)
  tc.TC.StacksLeft(1)
}
