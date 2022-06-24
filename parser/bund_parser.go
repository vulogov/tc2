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
		"", "'['", "']'",
	}
	staticData.symbolicNames = []string{
		"", "LSQ", "RSQ", "TOKEN", "INTEGER", "STRING", "RAW_STRING", "DECIMAL_INTEGER",
		"BTOKEN", "WS", "UNICODE_WS", "COMMENT", "TERMINATOR",
	}
	staticData.ruleNames = []string{
		"root", "expression", "value_",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 28, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 4, 0, 8, 8, 0,
		11, 0, 12, 0, 9, 1, 0, 3, 0, 13, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 18, 8, 1,
		10, 1, 12, 1, 21, 9, 1, 1, 1, 3, 1, 24, 8, 1, 1, 2, 1, 2, 1, 2, 0, 0, 3,
		0, 2, 4, 0, 0, 28, 0, 12, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 25, 1, 0, 0,
		0, 6, 8, 3, 2, 1, 0, 7, 6, 1, 0, 0, 0, 8, 9, 1, 0, 0, 0, 9, 7, 1, 0, 0,
		0, 9, 10, 1, 0, 0, 0, 10, 13, 1, 0, 0, 0, 11, 13, 5, 0, 0, 1, 12, 7, 1,
		0, 0, 0, 12, 11, 1, 0, 0, 0, 13, 1, 1, 0, 0, 0, 14, 23, 3, 4, 2, 0, 15,
		19, 5, 1, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 18, 21, 1, 0, 0,
		0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 22, 1, 0, 0, 0, 21, 19,
		1, 0, 0, 0, 22, 24, 5, 2, 0, 0, 23, 15, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0,
		24, 3, 1, 0, 0, 0, 25, 26, 5, 3, 0, 0, 26, 5, 1, 0, 0, 0, 4, 9, 12, 19,
		23,
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
	BundParserTOKEN           = 3
	BundParserINTEGER         = 4
	BundParserSTRING          = 5
	BundParserRAW_STRING      = 6
	BundParserDECIMAL_INTEGER = 7
	BundParserBTOKEN          = 8
	BundParserWS              = 9
	BundParserUNICODE_WS      = 10
	BundParserCOMMENT         = 11
	BundParserTERMINATOR      = 12
)

// BundParser rules.
const (
	BundParserRULE_root       = 0
	BundParserRULE_expression = 1
	BundParserRULE_value_     = 2
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
	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserTOKEN:
		p.SetState(7)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BundParserTOKEN {
			{
				p.SetState(6)
				p.Expression()
			}

			p.SetState(9)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case BundParserEOF:
		{
			p.SetState(11)
			p.Match(BundParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetVal sets the val rule contexts.
	SetVal(IValue_Context)

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

func (s *ExpressionContext) Get_expression() IExpressionContext { return s._expression }

func (s *ExpressionContext) SetVal(v IValue_Context) { s.val = v }

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

func (s *ExpressionContext) LSQ() antlr.TerminalNode {
	return s.GetToken(BundParserLSQ, 0)
}

func (s *ExpressionContext) RSQ() antlr.TerminalNode {
	return s.GetToken(BundParserRSQ, 0)
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
		p.SetState(14)

		var _x = p.Value_()

		localctx.(*ExpressionContext).val = _x
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BundParserLSQ {
		{
			p.SetState(15)
			p.Match(BundParserLSQ)
		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BundParserTOKEN {
			{
				p.SetState(16)

				var _x = p.Expression()

				localctx.(*ExpressionContext)._expression = _x
			}
			localctx.(*ExpressionContext).args = append(localctx.(*ExpressionContext).args, localctx.(*ExpressionContext)._expression)

			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(22)
			p.Match(BundParserRSQ)
		}

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
	p.EnterRule(localctx, 4, BundParserRULE_value_)

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
		p.SetState(25)
		p.Match(BundParserTOKEN)
	}

	return localctx
}
