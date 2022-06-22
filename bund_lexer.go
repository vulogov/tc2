// Code generated from BundLexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package tc2

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BundLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var bundlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bundlexerLexerInit() {
	staticData := &bundlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.symbolicNames = []string{
		"", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 48, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 4, 0, 11, 8, 0, 11, 0, 12, 0, 12, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 21, 8, 1, 10, 1, 12, 1, 24, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 4, 2, 32, 8, 2, 11, 2, 12, 2, 33, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 42, 8, 3, 10, 3, 12, 3, 45, 9, 3, 1, 3, 1, 3, 1, 22, 0, 4, 1, 1,
		3, 2, 5, 3, 7, 4, 1, 0, 2, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 51,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		1, 10, 1, 0, 0, 0, 3, 16, 1, 0, 0, 0, 5, 31, 1, 0, 0, 0, 7, 37, 1, 0, 0,
		0, 9, 11, 7, 0, 0, 0, 10, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 10, 1,
		0, 0, 0, 12, 13, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 15, 6, 0, 0, 0, 15,
		2, 1, 0, 0, 0, 16, 17, 5, 47, 0, 0, 17, 18, 5, 42, 0, 0, 18, 22, 1, 0,
		0, 0, 19, 21, 9, 0, 0, 0, 20, 19, 1, 0, 0, 0, 21, 24, 1, 0, 0, 0, 22, 23,
		1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 25, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0,
		25, 26, 5, 42, 0, 0, 26, 27, 5, 47, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 6,
		1, 0, 0, 29, 4, 1, 0, 0, 0, 30, 32, 7, 1, 0, 0, 31, 30, 1, 0, 0, 0, 32,
		33, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 1, 0, 0,
		0, 35, 36, 6, 2, 0, 0, 36, 6, 1, 0, 0, 0, 37, 38, 5, 47, 0, 0, 38, 39,
		5, 47, 0, 0, 39, 43, 1, 0, 0, 0, 40, 42, 8, 1, 0, 0, 41, 40, 1, 0, 0, 0,
		42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 1,
		0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 6, 3, 0, 0, 47, 8, 1, 0, 0, 0, 5,
		0, 12, 22, 33, 43, 1, 0, 1, 0,
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

// BundLexerInit initializes any static state used to implement BundLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBundLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BundLexerInit() {
	staticData := &bundlexerLexerStaticData
	staticData.once.Do(bundlexerLexerInit)
}

// NewBundLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBundLexer(input antlr.CharStream) *BundLexer {
	BundLexerInit()
	l := new(BundLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &bundlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "BundLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BundLexer tokens.
const (
	BundLexerWS           = 1
	BundLexerCOMMENT      = 2
	BundLexerTERMINATOR   = 3
	BundLexerLINE_COMMENT = 4
)