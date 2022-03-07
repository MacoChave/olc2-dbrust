parser grammar sintactico;

options {
	tokenVocab = lexico;
}

@header {
	import "db_rust/analizador/ast/expresion"
	import "db_rust/analizador/ast/interfaces"
	import "db_rust/analizador/entorno"
}

start
	returns[interfaces.Expresion res]: exp { $res = $exp.res};

exp
	returns[interfaces.Expresion res]:
	logica { $res = $logica.res }
	| relacional { $res = $relacional.res }
	| aritmetica { $res = $aritmetica.res };

logica
	returns[interfaces.Expresion res]:
	izq = logica op = S_AND der = logica {
		$res = expresion.NewOperacion($izq.res, $der.res, $op.text, false)
	}
	| izq = logica op = S_OR der = logica {
		$res = expresion.NewOperacion($izq.res, $der.res, $op.text, false)
	}
	| relacional { $res = $relacional.res };

relacional
	returns[interfaces.Expresion res]:
	izq = relacional op = (
		S_MENOR
		| S_MAYOR
		| S_MENORI
		| S_MAYORI
		| S_IGUAL
		| S_DIFERENTE
	) der = relacional {
		$res = expresion.NewOperacion($izq.res, $der.res, $op.text, false)
	}
	| aritmetica { $res = $aritmetica.res };

aritmetica
	returns[interfaces.Expresion res]:
	'-' exp { $res = expresion.NewOperacion($exp.res, nil, "-", true) }
	| izq = aritmetica op = ('*' | '/' | '%') der = aritmetica {
		$res = expresion.NewOperacion($izq.res, $der.res, $op.text, false) 
	}
	| izq = aritmetica op = ('+' | '-') der = aritmetica {
		$res = expresion.NewOperacion($izq.res, $der.res, $op.text, false) 
	}
	| exp_valor { $res = $exp_valor.res }
	| S_APAR exp S_CPAR { $res = $exp.res };

exp_valor
	returns[interfaces.Expresion res]:
	primitivo { $res = $primitivo.res };

primitivo
	returns[interfaces.Expresion res]:
	R_INT {
		val, err := strconv.Atoi($R_INT.text)
		if err != nil {
			fmt.Println(err)
		}
		$res = expresion.NewPrimitivo(val, entorno.INTEGER)
	}
	| R_FLOAT {
		val, err := strconv.ParseFloat($R_FLOAT.text, 64)
		if err != nil {
			fmt.Println(err)
		}
		$res = expresion.NewPrimitivo(val, entorno.FLOAT)
	}
	| R_STRING {
		val := $R_STRING.text[1: len($R_STRING.text) - 1]
		$res = expresion.NewPrimitivo(val, entorno.STRING)
	}
	| R_TRUE {
		$res = expresion.NewPrimitivo(true, entorno.BOOLEAN)
	}
	| R_FALSE {
		$res = expresion.NewPrimitivo(false, entorno.BOOLEAN)
	}
	| ID {
		$res = expresion.NewIdentificador($ID.text)
	};
