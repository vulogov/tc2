// Code generated from BundParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package tc2 // BundParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BundParserListener is a complete listener for a parse tree produced by BundParser.
type BundParserListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
