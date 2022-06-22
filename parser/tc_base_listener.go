// Code generated from tc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // tc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetcListener is a complete listener for a parse tree produced by tcParser.
type BasetcListener struct{}

var _ tcListener = &BasetcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BasetcListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BasetcListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterRoot_term is called when production root_term is entered.
func (s *BasetcListener) EnterRoot_term(ctx *Root_termContext) {}

// ExitRoot_term is called when production root_term is exited.
func (s *BasetcListener) ExitRoot_term(ctx *Root_termContext) {}

// EnterNs is called when production ns is entered.
func (s *BasetcListener) EnterNs(ctx *NsContext) {}

// ExitNs is called when production ns is exited.
func (s *BasetcListener) ExitNs(ctx *NsContext) {}

// EnterFun is called when production fun is entered.
func (s *BasetcListener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BasetcListener) ExitFun(ctx *FunContext) {}

// EnterVal is called when production val is entered.
func (s *BasetcListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BasetcListener) ExitVal(ctx *ValContext) {}

// EnterNs_term is called when production ns_term is entered.
func (s *BasetcListener) EnterNs_term(ctx *Ns_termContext) {}

// ExitNs_term is called when production ns_term is exited.
func (s *BasetcListener) ExitNs_term(ctx *Ns_termContext) {}

// EnterFun_term is called when production fun_term is entered.
func (s *BasetcListener) EnterFun_term(ctx *Fun_termContext) {}

// ExitFun_term is called when production fun_term is exited.
func (s *BasetcListener) ExitFun_term(ctx *Fun_termContext) {}

// EnterPos_term is called when production pos_term is entered.
func (s *BasetcListener) EnterPos_term(ctx *Pos_termContext) {}

// ExitPos_term is called when production pos_term is exited.
func (s *BasetcListener) ExitPos_term(ctx *Pos_termContext) {}
