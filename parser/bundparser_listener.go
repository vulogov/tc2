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

	// EnterValue_ is called when entering the value_ production.
	EnterValue_(c *Value_Context)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitValue_ is called when exiting the value_ production.
	ExitValue_(c *Value_Context)
}
