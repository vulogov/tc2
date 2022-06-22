// Code generated from BundParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package tc2 // BundParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBundParserListener is a complete listener for a parse tree produced by BundParser.
type BaseBundParserListener struct{}

var _ BundParserListener = &BaseBundParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBundParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBundParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBundParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBundParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBundParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBundParserListener) ExitExpression(ctx *ExpressionContext) {}
