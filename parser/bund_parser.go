// Code generated from BundParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BundParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BundParser struct {
	*antlr.BaseParser
}

var bundparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bundparserParserInit() {
	staticData := &bundparserParserStaticData
	staticData.literalNames = []string{
		"", "'['", "']'", "'('", "')'", "';;'", "':'", "';'", "'#'", "'$'",
	}
	staticData.symbolicNames = []string{
		"", "LSQ", "RSQ", "LRQ", "RRQ", "THEEND", "LAMBDA_HDR", "MACRO_HDR",
		"HASH", "DOLL", "TOKEN", "DATA", "Q", "INTEGER", "STRING", "RAW_STRING",
		"DECIMAL_INTEGER", "BTOKEN", "ID", "WS", "UNICODE_WS", "COMMENT", "TERMINATOR",
	}
	staticData.ruleNames = []string{
		"root", "expression", "declaration", "ns_declaration", "function_declaration",
		"macro_declaration", "nsexpression", "fexpression", "value_", "q_",
		"id_",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 110, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 0, 5, 0, 30, 8, 0,
		10, 0, 12, 0, 33, 9, 0, 1, 0, 3, 0, 36, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 43, 8, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1,
		1, 1, 3, 1, 53, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 58, 8, 2, 1, 3, 1, 3, 3,
		3, 62, 8, 3, 1, 3, 1, 3, 5, 3, 66, 8, 3, 10, 3, 12, 3, 69, 9, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 3, 4, 75, 8, 4, 1, 4, 1, 4, 5, 4, 79, 8, 4, 10, 4, 12,
		4, 82, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 90, 8, 5, 10, 5,
		12, 5, 93, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 100, 8, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 0, 0, 113, 0, 35, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0,
		4, 57, 1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8, 72, 1, 0, 0, 0, 10, 85, 1, 0,
		0, 0, 12, 99, 1, 0, 0, 0, 14, 101, 1, 0, 0, 0, 16, 103, 1, 0, 0, 0, 18,
		105, 1, 0, 0, 0, 20, 107, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0, 23, 22, 1, 0,
		0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 36,
		1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 30, 3, 4, 2, 0, 29, 28, 1, 0, 0, 0,
		30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 34, 1,
		0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 36, 5, 0, 0, 1, 35, 25, 1, 0, 0, 0, 35,
		31, 1, 0, 0, 0, 36, 1, 1, 0, 0, 0, 37, 42, 3, 16, 8, 0, 38, 39, 5, 3, 0,
		0, 39, 40, 3, 18, 9, 0, 40, 41, 5, 4, 0, 0, 41, 43, 1, 0, 0, 0, 42, 38,
		1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 52, 1, 0, 0, 0, 44, 48, 5, 1, 0, 0,
		45, 47, 3, 2, 1, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51,
		53, 5, 2, 0, 0, 52, 44, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 3, 1, 0, 0,
		0, 54, 58, 3, 8, 4, 0, 55, 58, 3, 10, 5, 0, 56, 58, 3, 6, 3, 0, 57, 54,
		1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 5, 1, 0, 0, 0,
		59, 61, 5, 1, 0, 0, 60, 62, 3, 20, 10, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 67, 5, 2, 0, 0, 64, 66, 3, 12, 6, 0, 65,
		64, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0,
		0, 68, 70, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 5, 5, 0, 0, 71, 7, 1,
		0, 0, 0, 72, 74, 5, 6, 0, 0, 73, 75, 3, 20, 10, 0, 74, 73, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 80, 5, 6, 0, 0, 77, 79, 3,
		14, 7, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80,
		81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 84, 5, 5, 0,
		0, 84, 9, 1, 0, 0, 0, 85, 86, 5, 7, 0, 0, 86, 87, 3, 20, 10, 0, 87, 91,
		5, 7, 0, 0, 88, 90, 3, 14, 7, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0,
		91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1,
		0, 0, 0, 94, 95, 5, 5, 0, 0, 95, 11, 1, 0, 0, 0, 96, 100, 3, 2, 1, 0, 97,
		100, 3, 6, 3, 0, 98, 100, 3, 8, 4, 0, 99, 96, 1, 0, 0, 0, 99, 97, 1, 0,
		0, 0, 99, 98, 1, 0, 0, 0, 100, 13, 1, 0, 0, 0, 101, 102, 3, 2, 1, 0, 102,
		15, 1, 0, 0, 0, 103, 104, 5, 10, 0, 0, 104, 17, 1, 0, 0, 0, 105, 106, 5,
		12, 0, 0, 106, 19, 1, 0, 0, 0, 107, 108, 5, 10, 0, 0, 108, 21, 1, 0, 0,
		0, 13, 25, 31, 35, 42, 48, 52, 57, 61, 67, 74, 80, 91, 99,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BundParserInit initializes any static state used to implement BundParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBundParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BundParserInit() {
	staticData := &bundparserParserStaticData
	staticData.once.Do(bundparserParserInit)
}

// NewBundParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBundParser(input antlr.TokenStream) *BundParser {
	BundParserInit()
	this := new(BundParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bundparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BundParser.g4"

	return this
}

// BundParser tokens.
const (
	BundParserEOF             = antlr.TokenEOF
	BundParserLSQ             = 1
	BundParserRSQ             = 2
	BundParserLRQ             = 3
	BundParserRRQ             = 4
	BundParserTHEEND          = 5
	BundParserLAMBDA_HDR      = 6
	BundParserMACRO_HDR       = 7
	BundParserHASH            = 8
	BundParserDOLL            = 9
	BundParserTOKEN           = 10
	BundParserDATA            = 11
	BundParserQ               = 12
	BundParserINTEGER         = 13
	BundParserSTRING          = 14
	BundParserRAW_STRING      = 15
	BundParserDECIMAL_INTEGER = 16
	BundParserBTOKEN          = 17
	BundParserID              = 18
	BundParserWS              = 19
	BundParserUNICODE_WS      = 20
	BundParserCOMMENT         = 21
	BundParserTERMINATOR      = 22
)

// BundParser rules.
const (
	BundParserRULE_root                 = 0
	BundParserRULE_expression           = 1
	BundParserRULE_declaration          = 2
	BundParserRULE_ns_declaration       = 3
	BundParserRULE_function_declaration = 4
	BundParserRULE_macro_declaration    = 5
	BundParserRULE_nsexpression         = 6
	BundParserRULE_fexpression          = 7
	BundParserRULE_value_               = 8
	BundParserRULE_q_                   = 9
	BundParserRULE_id_                  = 10
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(BundParserEOF, 0)
}

func (s *RootContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RootContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RootContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *RootContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *BundParser) Root() (localctx IRootContext) {
	this := p
	_ = this

	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BundParserRULE_root)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BundParserTOKEN {
			{
				p.SetState(22)
				p.Expression()
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserLSQ)|(1<<BundParserLAMBDA_HDR)|(1<<BundParserMACRO_HDR))) != 0 {
			{
				p.SetState(28)
				p.Declaration()
			}

			p.SetState(33)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(34)
			p.Match(BundParserEOF)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val rule contexts.
	GetVal() IValue_Context

	// GetQ returns the q rule contexts.
	GetQ() IQ_Context

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetVal sets the val rule contexts.
	SetVal(IValue_Context)

	// SetQ sets the q rule contexts.
	SetQ(IQ_Context)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IExpressionContext

	// SetArgs sets the args rule context list.
	SetArgs([]IExpressionContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	val         IValue_Context
	q           IQ_Context
	_expression IExpressionContext
	args        []IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetVal() IValue_Context { return s.val }

func (s *ExpressionContext) GetQ() IQ_Context { return s.q }

func (s *ExpressionContext) Get_expression() IExpressionContext { return s._expression }

func (s *ExpressionContext) SetVal(v IValue_Context) { s.val = v }

func (s *ExpressionContext) SetQ(v IQ_Context) { s.q = v }

func (s *ExpressionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ExpressionContext) GetArgs() []IExpressionContext { return s.args }

func (s *ExpressionContext) SetArgs(v []IExpressionContext) { s.args = v }

func (s *ExpressionContext) Value_() IValue_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_Context)
}

func (s *ExpressionContext) LRQ() antlr.TerminalNode {
	return s.GetToken(BundParserLRQ, 0)
}

func (s *ExpressionContext) RRQ() antlr.TerminalNode {
	return s.GetToken(BundParserRRQ, 0)
}

func (s *ExpressionContext) LSQ() antlr.TerminalNode {
	return s.GetToken(BundParserLSQ, 0)
}

func (s *ExpressionContext) RSQ() antlr.TerminalNode {
	return s.GetToken(BundParserRSQ, 0)
}

func (s *ExpressionContext) Q_() IQ_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQ_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQ_Context)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BundParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BundParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)

		var _x = p.Value_()

		localctx.(*ExpressionContext).val = _x
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserLRQ {
		{
			p.SetState(38)
			p.Match(BundParserLRQ)
		}
		{
			p.SetState(39)

			var _x = p.Q_()

			localctx.(*ExpressionContext).q = _x
		}
		{
			p.SetState(40)
			p.Match(BundParserRRQ)
		}

	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(44)
			p.Match(BundParserLSQ)
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BundParserTOKEN {
			{
				p.SetState(45)

				var _x = p.Expression()

				localctx.(*ExpressionContext)._expression = _x
			}
			localctx.(*ExpressionContext).args = append(localctx.(*ExpressionContext).args, localctx.(*ExpressionContext)._expression)

			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(51)
			p.Match(BundParserRSQ)
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Function_declaration() IFunction_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *DeclarationContext) Macro_declaration() IMacro_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacro_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacro_declarationContext)
}

func (s *DeclarationContext) Ns_declaration() INs_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INs_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INs_declarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *BundParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BundParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserLAMBDA_HDR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Function_declaration()
		}

	case BundParserMACRO_HDR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Macro_declaration()
		}

	case BundParserLSQ:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Ns_declaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INs_declarationContext is an interface to support dynamic dispatch.
type INs_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IId_Context

	// Get_nsexpression returns the _nsexpression rule contexts.
	Get_nsexpression() INsexpressionContext

	// SetName sets the name rule contexts.
	SetName(IId_Context)

	// Set_nsexpression sets the _nsexpression rule contexts.
	Set_nsexpression(INsexpressionContext)

	// GetArgs returns the args rule context list.
	GetArgs() []INsexpressionContext

	// SetArgs sets the args rule context list.
	SetArgs([]INsexpressionContext)

	// IsNs_declarationContext differentiates from other interfaces.
	IsNs_declarationContext()
}

type Ns_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	name          IId_Context
	_nsexpression INsexpressionContext
	args          []INsexpressionContext
}

func NewEmptyNs_declarationContext() *Ns_declarationContext {
	var p = new(Ns_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ns_declaration
	return p
}

func (*Ns_declarationContext) IsNs_declarationContext() {}

func NewNs_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ns_declarationContext {
	var p = new(Ns_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ns_declaration

	return p
}

func (s *Ns_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Ns_declarationContext) GetName() IId_Context { return s.name }

func (s *Ns_declarationContext) Get_nsexpression() INsexpressionContext { return s._nsexpression }

func (s *Ns_declarationContext) SetName(v IId_Context) { s.name = v }

func (s *Ns_declarationContext) Set_nsexpression(v INsexpressionContext) { s._nsexpression = v }

func (s *Ns_declarationContext) GetArgs() []INsexpressionContext { return s.args }

func (s *Ns_declarationContext) SetArgs(v []INsexpressionContext) { s.args = v }

func (s *Ns_declarationContext) LSQ() antlr.TerminalNode {
	return s.GetToken(BundParserLSQ, 0)
}

func (s *Ns_declarationContext) RSQ() antlr.TerminalNode {
	return s.GetToken(BundParserRSQ, 0)
}

func (s *Ns_declarationContext) THEEND() antlr.TerminalNode {
	return s.GetToken(BundParserTHEEND, 0)
}

func (s *Ns_declarationContext) Id_() IId_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *Ns_declarationContext) AllNsexpression() []INsexpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INsexpressionContext); ok {
			len++
		}
	}

	tst := make([]INsexpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INsexpressionContext); ok {
			tst[i] = t.(INsexpressionContext)
			i++
		}
	}

	return tst
}

func (s *Ns_declarationContext) Nsexpression(i int) INsexpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INsexpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INsexpressionContext)
}

func (s *Ns_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ns_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ns_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterNs_declaration(s)
	}
}

func (s *Ns_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitNs_declaration(s)
	}
}

func (p *BundParser) Ns_declaration() (localctx INs_declarationContext) {
	this := p
	_ = this

	localctx = NewNs_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BundParserRULE_ns_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(BundParserLSQ)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserTOKEN {
		{
			p.SetState(60)

			var _x = p.Id_()

			localctx.(*Ns_declarationContext).name = _x
		}

	}
	{
		p.SetState(63)
		p.Match(BundParserRSQ)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserLSQ)|(1<<BundParserLAMBDA_HDR)|(1<<BundParserTOKEN))) != 0 {
		{
			p.SetState(64)

			var _x = p.Nsexpression()

			localctx.(*Ns_declarationContext)._nsexpression = _x
		}
		localctx.(*Ns_declarationContext).args = append(localctx.(*Ns_declarationContext).args, localctx.(*Ns_declarationContext)._nsexpression)

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(BundParserTHEEND)
	}

	return localctx
}

// IFunction_declarationContext is an interface to support dynamic dispatch.
type IFunction_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IId_Context

	// Get_fexpression returns the _fexpression rule contexts.
	Get_fexpression() IFexpressionContext

	// SetName sets the name rule contexts.
	SetName(IId_Context)

	// Set_fexpression sets the _fexpression rule contexts.
	Set_fexpression(IFexpressionContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IFexpressionContext

	// SetArgs sets the args rule context list.
	SetArgs([]IFexpressionContext)

	// IsFunction_declarationContext differentiates from other interfaces.
	IsFunction_declarationContext()
}

type Function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	name         IId_Context
	_fexpression IFexpressionContext
	args         []IFexpressionContext
}

func NewEmptyFunction_declarationContext() *Function_declarationContext {
	var p = new(Function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_function_declaration
	return p
}

func (*Function_declarationContext) IsFunction_declarationContext() {}

func NewFunction_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declarationContext {
	var p = new(Function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_function_declaration

	return p
}

func (s *Function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declarationContext) GetName() IId_Context { return s.name }

func (s *Function_declarationContext) Get_fexpression() IFexpressionContext { return s._fexpression }

func (s *Function_declarationContext) SetName(v IId_Context) { s.name = v }

func (s *Function_declarationContext) Set_fexpression(v IFexpressionContext) { s._fexpression = v }

func (s *Function_declarationContext) GetArgs() []IFexpressionContext { return s.args }

func (s *Function_declarationContext) SetArgs(v []IFexpressionContext) { s.args = v }

func (s *Function_declarationContext) AllLAMBDA_HDR() []antlr.TerminalNode {
	return s.GetTokens(BundParserLAMBDA_HDR)
}

func (s *Function_declarationContext) LAMBDA_HDR(i int) antlr.TerminalNode {
	return s.GetToken(BundParserLAMBDA_HDR, i)
}

func (s *Function_declarationContext) THEEND() antlr.TerminalNode {
	return s.GetToken(BundParserTHEEND, 0)
}

func (s *Function_declarationContext) Id_() IId_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *Function_declarationContext) AllFexpression() []IFexpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFexpressionContext); ok {
			len++
		}
	}

	tst := make([]IFexpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFexpressionContext); ok {
			tst[i] = t.(IFexpressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_declarationContext) Fexpression(i int) IFexpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFexpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFexpressionContext)
}

func (s *Function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterFunction_declaration(s)
	}
}

func (s *Function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitFunction_declaration(s)
	}
}

func (p *BundParser) Function_declaration() (localctx IFunction_declarationContext) {
	this := p
	_ = this

	localctx = NewFunction_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BundParserRULE_function_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(BundParserLAMBDA_HDR)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserTOKEN {
		{
			p.SetState(73)

			var _x = p.Id_()

			localctx.(*Function_declarationContext).name = _x
		}

	}
	{
		p.SetState(76)
		p.Match(BundParserLAMBDA_HDR)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserTOKEN {
		{
			p.SetState(77)

			var _x = p.Fexpression()

			localctx.(*Function_declarationContext)._fexpression = _x
		}
		localctx.(*Function_declarationContext).args = append(localctx.(*Function_declarationContext).args, localctx.(*Function_declarationContext)._fexpression)

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(83)
		p.Match(BundParserTHEEND)
	}

	return localctx
}

// IMacro_declarationContext is an interface to support dynamic dispatch.
type IMacro_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IId_Context

	// Get_fexpression returns the _fexpression rule contexts.
	Get_fexpression() IFexpressionContext

	// SetName sets the name rule contexts.
	SetName(IId_Context)

	// Set_fexpression sets the _fexpression rule contexts.
	Set_fexpression(IFexpressionContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IFexpressionContext

	// SetArgs sets the args rule context list.
	SetArgs([]IFexpressionContext)

	// IsMacro_declarationContext differentiates from other interfaces.
	IsMacro_declarationContext()
}

type Macro_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	name         IId_Context
	_fexpression IFexpressionContext
	args         []IFexpressionContext
}

func NewEmptyMacro_declarationContext() *Macro_declarationContext {
	var p = new(Macro_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_macro_declaration
	return p
}

func (*Macro_declarationContext) IsMacro_declarationContext() {}

func NewMacro_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Macro_declarationContext {
	var p = new(Macro_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_macro_declaration

	return p
}

func (s *Macro_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Macro_declarationContext) GetName() IId_Context { return s.name }

func (s *Macro_declarationContext) Get_fexpression() IFexpressionContext { return s._fexpression }

func (s *Macro_declarationContext) SetName(v IId_Context) { s.name = v }

func (s *Macro_declarationContext) Set_fexpression(v IFexpressionContext) { s._fexpression = v }

func (s *Macro_declarationContext) GetArgs() []IFexpressionContext { return s.args }

func (s *Macro_declarationContext) SetArgs(v []IFexpressionContext) { s.args = v }

func (s *Macro_declarationContext) AllMACRO_HDR() []antlr.TerminalNode {
	return s.GetTokens(BundParserMACRO_HDR)
}

func (s *Macro_declarationContext) MACRO_HDR(i int) antlr.TerminalNode {
	return s.GetToken(BundParserMACRO_HDR, i)
}

func (s *Macro_declarationContext) THEEND() antlr.TerminalNode {
	return s.GetToken(BundParserTHEEND, 0)
}

func (s *Macro_declarationContext) Id_() IId_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *Macro_declarationContext) AllFexpression() []IFexpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFexpressionContext); ok {
			len++
		}
	}

	tst := make([]IFexpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFexpressionContext); ok {
			tst[i] = t.(IFexpressionContext)
			i++
		}
	}

	return tst
}

func (s *Macro_declarationContext) Fexpression(i int) IFexpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFexpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFexpressionContext)
}

func (s *Macro_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Macro_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Macro_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterMacro_declaration(s)
	}
}

func (s *Macro_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitMacro_declaration(s)
	}
}

func (p *BundParser) Macro_declaration() (localctx IMacro_declarationContext) {
	this := p
	_ = this

	localctx = NewMacro_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BundParserRULE_macro_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(BundParserMACRO_HDR)
	}
	{
		p.SetState(86)

		var _x = p.Id_()

		localctx.(*Macro_declarationContext).name = _x
	}
	{
		p.SetState(87)
		p.Match(BundParserMACRO_HDR)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserTOKEN {
		{
			p.SetState(88)

			var _x = p.Fexpression()

			localctx.(*Macro_declarationContext)._fexpression = _x
		}
		localctx.(*Macro_declarationContext).args = append(localctx.(*Macro_declarationContext).args, localctx.(*Macro_declarationContext)._fexpression)

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(BundParserTHEEND)
	}

	return localctx
}

// INsexpressionContext is an interface to support dynamic dispatch.
type INsexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNsexpressionContext differentiates from other interfaces.
	IsNsexpressionContext()
}

type NsexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNsexpressionContext() *NsexpressionContext {
	var p = new(NsexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_nsexpression
	return p
}

func (*NsexpressionContext) IsNsexpressionContext() {}

func NewNsexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NsexpressionContext {
	var p = new(NsexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_nsexpression

	return p
}

func (s *NsexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NsexpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NsexpressionContext) Ns_declaration() INs_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INs_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INs_declarationContext)
}

func (s *NsexpressionContext) Function_declaration() IFunction_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *NsexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NsexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NsexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterNsexpression(s)
	}
}

func (s *NsexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitNsexpression(s)
	}
}

func (p *BundParser) Nsexpression() (localctx INsexpressionContext) {
	this := p
	_ = this

	localctx = NewNsexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BundParserRULE_nsexpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserTOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Expression()
		}

	case BundParserLSQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Ns_declaration()
		}

	case BundParserLAMBDA_HDR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Function_declaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFexpressionContext is an interface to support dynamic dispatch.
type IFexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFexpressionContext differentiates from other interfaces.
	IsFexpressionContext()
}

type FexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFexpressionContext() *FexpressionContext {
	var p = new(FexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_fexpression
	return p
}

func (*FexpressionContext) IsFexpressionContext() {}

func NewFexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FexpressionContext {
	var p = new(FexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_fexpression

	return p
}

func (s *FexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FexpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterFexpression(s)
	}
}

func (s *FexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitFexpression(s)
	}
}

func (p *BundParser) Fexpression() (localctx IFexpressionContext) {
	this := p
	_ = this

	localctx = NewFexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BundParserRULE_fexpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Expression()
	}

	return localctx
}

// IValue_Context is an interface to support dynamic dispatch.
type IValue_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_Context differentiates from other interfaces.
	IsValue_Context()
}

type Value_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_Context() *Value_Context {
	var p = new(Value_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_value_
	return p
}

func (*Value_Context) IsValue_Context() {}

func NewValue_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_Context {
	var p = new(Value_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_value_

	return p
}

func (s *Value_Context) GetParser() antlr.Parser { return s.parser }

func (s *Value_Context) TOKEN() antlr.TerminalNode {
	return s.GetToken(BundParserTOKEN, 0)
}

func (s *Value_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterValue_(s)
	}
}

func (s *Value_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitValue_(s)
	}
}

func (p *BundParser) Value_() (localctx IValue_Context) {
	this := p
	_ = this

	localctx = NewValue_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BundParserRULE_value_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(BundParserTOKEN)
	}

	return localctx
}

// IQ_Context is an interface to support dynamic dispatch.
type IQ_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQ_Context differentiates from other interfaces.
	IsQ_Context()
}

type Q_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQ_Context() *Q_Context {
	var p = new(Q_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_q_
	return p
}

func (*Q_Context) IsQ_Context() {}

func NewQ_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Q_Context {
	var p = new(Q_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_q_

	return p
}

func (s *Q_Context) GetParser() antlr.Parser { return s.parser }

func (s *Q_Context) Q() antlr.TerminalNode {
	return s.GetToken(BundParserQ, 0)
}

func (s *Q_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Q_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Q_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterQ_(s)
	}
}

func (s *Q_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitQ_(s)
	}
}

func (p *BundParser) Q_() (localctx IQ_Context) {
	this := p
	_ = this

	localctx = NewQ_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BundParserRULE_q_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(BundParserQ)
	}

	return localctx
}

// IId_Context is an interface to support dynamic dispatch.
type IId_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsId_Context differentiates from other interfaces.
	IsId_Context()
}

type Id_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_Context() *Id_Context {
	var p = new(Id_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_id_
	return p
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) TOKEN() antlr.TerminalNode {
	return s.GetToken(BundParserTOKEN, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundParserListener); ok {
		listenerT.ExitId_(s)
	}
}

func (p *BundParser) Id_() (localctx IId_Context) {
	this := p
	_ = this

	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BundParserRULE_id_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(BundParserTOKEN)
	}

	return localctx
}
