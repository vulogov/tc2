lexer grammar BundLexer;

LSQ: '[' ;
RSQ: ']' ;

LRQ: '(' ;
RRQ: ')' ;

THEEND:     ';;' ;
LAMBDA_HDR: ':' ;
MACRO_HDR:  ';' ;

HASH: '#' ;
DOLL: '$' ;

TOKEN
  : INTEGER
  | STRING
  | EXPONENT_OR_POINT_FLOAT
  | BTOKEN
  ;

DATA
  : INTEGER
  | STRING
  | EXPONENT_OR_POINT_FLOAT
  ;

Q
  : INTEGER
  | EXPONENT_OR_POINT_FLOAT
  ;

INTEGER
  :  (SIGN)? DECIMAL_INTEGER
  ;

STRING
  : RAW_STRING
  | SHORT_STRING
  ;

RAW_STRING
  : '`' ~'`'* '`'
  ;

DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  | '0'+
  ;

BTOKEN
  : (LETTER)+ (LETTER)*
  ;

ID
  : (UNICODE_LETTER)+ (UNICODE_LETTER|UNICODE_DIGIT)*
  ;

fragment OP
  : '+'
  | '-'
  | '*'
  | '/'
  | '|'
  | '^'
  | '&'
  | '!'
  | '↑'
  | '↓'
  | '×'
  | '÷'
  | '∆'
  | '∇'
  | ','
  | '_'
  | ';'
  | '<'
  | '>'
  | '='
  | '¬'
  | '∨'
  | '∧'
  | 'λ'
  | '?'
  | '`'
  | '~'
  | '@'
  | '$'
  | '∘'
  | '∀'
  | '∃'
  | '∄'
  | '□'
  | '#'
  | '.'
  ;

fragment LETTER
  : UNICODE_LETTER
  | UNICODE_DIGIT
  | OP
  ;

fragment NON_ZERO_DIGIT
  : [1-9]
  ;

fragment DIGIT
  : [0-9]
  ;

fragment SIGN
  : ('+' | '-')
  ;

fragment RN
  : '\r'? '\n'
  ;
fragment SHORT_STRING
  : '\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | '"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment EXPONENT_OR_POINT_FLOAT
  : (SIGN)? ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | (SIGN)? POINT_FLOAT
  ;

fragment POINT_FLOAT
  : [0-9]* '.' [0-9]+
  | [0-9]+ '.'
  ;

fragment UNICODE_LETTER
  : [\p{L}]
  ;

fragment UNICODE_DIGIT
  : [\p{Nd}]
  ;

// Hidden tokens
WS                     : [ \t]+             -> skip;
UNICODE_WS             : [\p{White_Space}]  -> skip;
COMMENT                : '/*' .*? '*/'      -> skip;
TERMINATOR             : [\r\n]+            -> skip;
