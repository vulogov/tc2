parser grammar BundParser;

options {
	tokenVocab = BundLexer;
}

root: ((expression)+| EOF) ;

expression:
      val=value_ (LSQ (args+=expression)* RSQ)? ;

value_
	: TOKEN
	;
