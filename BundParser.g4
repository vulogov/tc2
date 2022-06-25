parser grammar BundParser;

options {
	tokenVocab = BundLexer;
}

root: ((expression)*| (declaration)* EOF) ;

expression:
      val=value_ (LRQ q=q_ RRQ)? (LSQ (args+=expression)* RSQ)? ;

declaration
	: function_declaration
	| macro_declaration
	| ns_declaration
	;

ns_declaration:
			LSQ (name=id_)? RSQ (args+=nsexpression)* THEEND ;

function_declaration:
			LAMBDA_HDR (name=id_)? LAMBDA_HDR (args+=fexpression)* THEEND ;

macro_declaration:
			MACRO_HDR name=id_ MACRO_HDR (args+=fexpression)* THEEND ;

nsexpression
	: expression
	| function_declaration
	;

fexpression
	: expression
	;

value_
	: TOKEN
	;

q_
	: Q
	;

id_
	: TOKEN
	;
