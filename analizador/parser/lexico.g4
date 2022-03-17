lexer grammar lexico;

// TOKENS

S_PUNTO: '.';
S_COMA: ',';
S_PTCOMA: ';';
S_ASIGNAR: '=';
S_AMP: '&';
S_MATCH_OR: '|';
S_MATCH_RET: '=>';
S_MATCH_DEFAULT: '_';
S_SUMA: '+';
S_RESTA: '-';
S_POR: '*';
S_DIVISION: '/';
S_MODULO: '%';
S_MAYOR: '>';
S_MENOR: '<';
S_MAYORI: '>=';
S_MENORI: '<=';
S_IGUAL: '==';
S_DIFERENTE: '!=';
S_OR: '||';
S_AND: '&&';
S_MOD_RET: '->';
S_APAR: '(';
S_CPAR: ')';
S_DOSPT: ':';
S_ALIAS: '::';
S_ACOR: '[';
S_CCOR: ']';
S_ALLAV: '{';
S_CLLAV: '}';
S_RANGO: '..';

R_TRUE: 'true';
R_FALSE: 'false';
R_LET: 'let';
R_MUT: 'mut';
R_INT: 'i64';
R_FLOAT: 'f64';
R_BOOL: 'bool';
R_CHAR: 'char';
R_STRING: 'string';
R_STR: '&str';
R_FN: 'fn';
R_MAIN: 'main';
R_AS: 'as';
R_TO_OWNED: 'to_owned';
R_TO_STRING: 'to_string';
R_POW: 'pow';
R_POWF: 'powF';
R_PRINTLN: 'println!';
R_ABS: 'abs';
R_SQRT: 'sqrt';
R_CLONE: 'clone';
R_NEW: 'new';
R_LEN: 'len';
R_PUSH: 'push';
R_REMOVE: 'remove';
R_CONTAINS: 'contains';
R_INSERT: 'insert';
R_CAPACITY: 'capacity';
R_WITH_CAPACITY: 'with_capacity';
R_IF: 'if';
R_ELSE: 'else';
R_MATCH: 'match';
R_LOOP: 'loop';
R_BREAK: 'break';
R_WHILE: 'while';
R_FOR: 'for';
R_IN: 'in';
R_CONTINUE: 'continue';
R_RETURN: 'return';
R_STRUCT: 'struct';
R_VECTOR: 'vector';
R_VEC: 'vec';
R_MOD: 'mod';
R_PUB: 'pub';

NUMERO: [0-9]+;
DECIMAL: [0-9]+ [.] [0-9]+;
CARACTER: ['] . ['];
CADENA: '"' ~["]* '"';
ID: [a-zA-Z_][a-zA-Z0-9_]*;

COMENTARIO: ('//' [^.]* [\n]) -> skip;
BLANCOS: [ \r\n\t]+ -> skip;

fragment ESC_SEQ:
	'\\' (
		'\\'
		| '@'
		| '['
		| ']'
		| '.'
		| '#'
		| '+'
		| '-'
		| '!'
		| ':'
		| ' '
	);