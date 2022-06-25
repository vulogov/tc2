package tc2

import (
  "github.com/vulogov/tc2/parser"
)

func (tc *TCExecListener) EnterExpression(ctx *parser.ExpressionContext) {
  if tc.TC.Errors() > 0 {
    return
  }
  tc.TC.console.Func("EnterExpression", ctx.GetVal().GetText())
}

func (tc *TCExecListener) ExitExpression(ctx *parser.ExpressionContext) {
  if tc.TC.Errors() > 0 {
    return
  }
  tc.TC.console.Fclose("No more")
}
