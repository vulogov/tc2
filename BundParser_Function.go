package tc2

import (
  "github.com/vulogov/tc2/parser"
)

func (tc *TCExecListener) EnterFunction_declaration(ctx *parser.Function_declarationContext) {
  if ctx.GetName() != nil {
    tc.TC.console.Section("EnterFunction", ctx.GetName().GetText())
  } else {
    tc.TC.console.Section("EnterFunction", "anonymous")
  }
}

func (tc *TCExecListener) ExitFunction_declaration(ctx *parser.Function_declarationContext) {
  tc.TC.console.SectionEnd("No more")
}
