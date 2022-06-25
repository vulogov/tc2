// Code generated from BundParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BundParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BundParserListener is a complete listener for a parse tree produced by BundParser.
type BundParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterNs_declaration is called when entering the ns_declaration production.
	EnterNs_declaration(c *Ns_declarationContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterMacro_declaration is called when entering the macro_declaration production.
	EnterMacro_declaration(c *Macro_declarationContext)

	// EnterNsexpression is called when entering the nsexpression production.
	EnterNsexpression(c *NsexpressionContext)

	// EnterFexpression is called when entering the fexpression production.
	EnterFexpression(c *FexpressionContext)

	// EnterValue_ is called when entering the value_ production.
	EnterValue_(c *Value_Context)

	// EnterQ_ is called when entering the q_ production.
	EnterQ_(c *Q_Context)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitNs_declaration is called when exiting the ns_declaration production.
	ExitNs_declaration(c *Ns_declarationContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitMacro_declaration is called when exiting the macro_declaration production.
	ExitMacro_declaration(c *Macro_declarationContext)

	// ExitNsexpression is called when exiting the nsexpression production.
	ExitNsexpression(c *NsexpressionContext)

	// ExitFexpression is called when exiting the fexpression production.
	ExitFexpression(c *FexpressionContext)

	// ExitValue_ is called when exiting the value_ production.
	ExitValue_(c *Value_Context)

	// ExitQ_ is called when exiting the q_ production.
	ExitQ_(c *Q_Context)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)
}
