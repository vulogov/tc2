// Code generated from BundParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BundParser

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

// EnterRoot is called when production root is entered.
func (s *BaseBundParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseBundParserListener) ExitRoot(ctx *RootContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBundParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBundParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterValue_ is called when production value_ is entered.
func (s *BaseBundParserListener) EnterValue_(ctx *Value_Context) {}

// ExitValue_ is called when production value_ is exited.
func (s *BaseBundParserListener) ExitValue_(ctx *Value_Context) {}
