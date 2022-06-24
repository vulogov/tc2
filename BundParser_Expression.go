package tc2

import (
  "github.com/vulogov/tc2/parser"
)

func (tc *TCExecListener) EnterExpression(ctx *parser.ExpressionContext) {
  tc.TC.console.Func("EnterExpression", ctx.GetVal().GetText())
}

func (tc *TCExecListener) ExitExpression(ctx *parser.ExpressionContext) {
  tc.TC.console.Fclose("No more")
}
