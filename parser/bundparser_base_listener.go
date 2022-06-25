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

// EnterDeclaration is called when production declaration is entered.
func (s *BaseBundParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseBundParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterNs_declaration is called when production ns_declaration is entered.
func (s *BaseBundParserListener) EnterNs_declaration(ctx *Ns_declarationContext) {}

// ExitNs_declaration is called when production ns_declaration is exited.
func (s *BaseBundParserListener) ExitNs_declaration(ctx *Ns_declarationContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BaseBundParserListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BaseBundParserListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterMacro_declaration is called when production macro_declaration is entered.
func (s *BaseBundParserListener) EnterMacro_declaration(ctx *Macro_declarationContext) {}

// ExitMacro_declaration is called when production macro_declaration is exited.
func (s *BaseBundParserListener) ExitMacro_declaration(ctx *Macro_declarationContext) {}

// EnterNsexpression is called when production nsexpression is entered.
func (s *BaseBundParserListener) EnterNsexpression(ctx *NsexpressionContext) {}

// ExitNsexpression is called when production nsexpression is exited.
func (s *BaseBundParserListener) ExitNsexpression(ctx *NsexpressionContext) {}

// EnterFexpression is called when production fexpression is entered.
func (s *BaseBundParserListener) EnterFexpression(ctx *FexpressionContext) {}

// ExitFexpression is called when production fexpression is exited.
func (s *BaseBundParserListener) ExitFexpression(ctx *FexpressionContext) {}

// EnterValue_ is called when production value_ is entered.
func (s *BaseBundParserListener) EnterValue_(ctx *Value_Context) {}

// ExitValue_ is called when production value_ is exited.
func (s *BaseBundParserListener) ExitValue_(ctx *Value_Context) {}

// EnterQ_ is called when production q_ is entered.
func (s *BaseBundParserListener) EnterQ_(ctx *Q_Context) {}

// ExitQ_ is called when production q_ is exited.
func (s *BaseBundParserListener) ExitQ_(ctx *Q_Context) {}

// EnterId_ is called when production id_ is entered.
func (s *BaseBundParserListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BaseBundParserListener) ExitId_(ctx *Id_Context) {}
