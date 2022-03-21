parser grammar sintactico;

options {
	tokenVocab = lexico;
}

@header {
	import "db_rust/analizador/ast"
	import "db_rust/analizador/ast/expresion"
	import "db_rust/analizador/ast/funcion"
	import "db_rust/analizador/ast/imprimir"
	import "db_rust/analizador/ast/variables"
	import "db_rust/analizador/ast/flujo"
	import "db_rust/analizador/ast/interfaces"
	import "db_rust/analizador/entorno"
	import arrayList "github.com/colegno/arraylist"
}

start
	returns[ast.Ast root]:
	procedimientos {
		$root = ast.NewAst($procedimientos.lista)
	};

procedimientos
	returns[*arrayList.List lista]
	@init { $lista = arrayList.New() }:
	proc += procedimiento* {
		LISTA := localctx.(*ProcedimientosContext).GetProc()
		for _, i := range LISTA {
			$lista.Add(i.GetInstr())
		}
	};

procedimiento
	returns[interfaces.Instruccion instr]:
	principal { $instr = $principal.instr };
// | R_FN ID S_APAR S_CPAR S_ALLAV instrucciones S_CLLAV { params := arrayList.New() $instr =
// funcion.NewFuncion(entorno.VOID, $ID.text, params, $instrucciones.lista) } | R_FN ID S_APAR
// S_CPAR S_MOD_RET tipo_dato S_ALLAV instrucciones S_CLLAV { params := arrayList.New() $instr =
// funcion.NewFuncion($tipo_dato.tipo, $ID.text, params, $instrucciones.lista) } | R_FN ID S_APAR
// parametros S_CPAR S_ALLAV instrucciones S_CLLAV { $instr = funcion.NewFuncion(entorno.VOID,
// $ID.text, $parametros.lista, $instrucciones.lista) } | R_FN ID S_APAR parametros S_CPAR S_MOD_RET
// tipo_dato S_ALLAV instrucciones S_CLLAV { $instr = funcion.NewFuncion($tipo_dato.tipo, $ID.text,
// $parametros.lista, $instrucciones.lista) }

principal
	returns[interfaces.Instruccion instr]
	@init { params := arrayList.New() }:
	R_FN R_MAIN S_APAR S_CPAR S_ALLAV instrucciones S_CLLAV {
		$instr = funcion.NewFuncion(entorno.VOID, "main", params, $instrucciones.lista)
	};

// parametros returns[*arrayList.List lista] @init { $lista = arrayList.New() }: LISTA = parametros
// S_COMA parametro { $LISTA.lista.Add($parametro.instr) $lista = $LISTA.lista } | parametro {
// $lista.Add($parametro.instr) };

// parametro returns[interfaces.Instruccion instr]: ID S_DOSPT tipo_dato { id :=
// expresion.NewIdentificador($ID.text) $instr = variables.NewDeclaracion(false, id,
// $tipo_dato.tipo, nil) };

instrucciones
	returns[*arrayList.List lista]
	@init { $lista = arrayList.New() }:
	ins += instruccion+ {
		LISTA := localctx.(*InstruccionesContext).GetIns()
		for _, i := range LISTA {
			$lista.Add(i.GetInstr())
		}
	};

instruccion
	returns[interfaces.Instruccion instr]:
	imprimir S_PTCOMA { $instr = $imprimir.instr }
	| declaracion S_PTCOMA { $instr = $declaracion.instr }
	| asignacion S_PTCOMA { $instr = $asignacion.instr }
	| sent_if S_PTCOMA { $instr = $sent_if.instr };

imprimir
	returns[interfaces.Instruccion instr]:
	R_PRINTLN S_APAR lista_exp S_CPAR {
		$instr = imprimir.NewImprimir($lista_exp.lista)
	};

lista_exp
	returns[*arrayList.List lista]
	@init { $lista = arrayList.New() }:
	LISTA = lista_exp ',' exp { 
		$LISTA.lista.Add($exp.val) 
		$lista = $LISTA.lista 
	}
	| exp { $lista.Add($exp.val) };

declaracion
	returns[interfaces.Instruccion instr]:
	R_LET R_MUT ID S_DOSPT tipo_dato S_ASIGNAR exp {
		id := expresion.NewIdentificador($ID.text)
		$instr = variables.NewDeclaracion(true, id, $tipo_dato.tipo, $exp.val)
	}
	| R_LET R_MUT ID S_ASIGNAR exp {
		id := expresion.NewIdentificador($ID.text)
		$instr = variables.NewDeclaracion(true, id, entorno.NULL, $exp.val)
	}
	| R_LET ID S_DOSPT tipo_dato S_ASIGNAR exp {
		id := expresion.NewIdentificador($ID.text)
		$instr = variables.NewDeclaracion(false, id, $tipo_dato.tipo, $exp.val)
	}
	| R_LET ID S_ASIGNAR exp {
		id := expresion.NewIdentificador($ID.text)
		$instr = variables.NewDeclaracion(false, id, entorno.NULL, $exp.val)
	};

tipo_dato
	returns[entorno.TipoDato tipo]:
	R_INT { $tipo = entorno.INTEGER }
	| R_FLOAT { $tipo = entorno.FLOAT }
	| R_CHAR { $tipo = entorno.CHAR }
	| R_STRING { $tipo = entorno.STRING }
	| R_BOOL { $tipo = entorno.BOOLEAN };

asignacion
	returns[interfaces.Instruccion instr]:
	ID S_ASIGNAR exp {
		id := expresion.NewIdentificador($ID.text)
		$instr = variables.NewAsignacion(id, $exp.val)
	};

sent_if
	returns[interfaces.Instruccion instr]:
	R_IF exp S_ALLAV s_then = instrucciones S_CLLAV {
		$instr = flujo.NewIf($exp.val, $s_then.lista, nil, nil)
	}
	| R_IF exp S_ALLAV s_then = instrucciones S_CLLAV R_ELSE S_ALLAV s_else = instrucciones S_CLLAV
		{
		$instr = flujo.NewIf($exp.val, $s_then.lista, nil, $s_else.lista)
	}
	| R_IF exp S_ALLAV s_then = instrucciones S_CLLAV lista_elseif R_ELSE S_ALLAV s_else =
		instrucciones S_CLLAV {
		$instr = flujo.NewIf($exp.val, $s_then.lista, $lista_elseif.lista, $s_else.lista)
	};

lista_elseif
	returns[*arrayList.List lista]
	@init { $lista = arrayList.New() }:
	ins += elseif+ {
		LISTA := localctx.(*Lista_elseifContext).GetIns()
		for _, i := range LISTA {
			$lista.Add(i.GetInstr())
		}
	};

elseif
	returns[interfaces.Instruccion instr]:
	R_ELSE R_IF exp S_ALLAV instrucciones S_CLLAV {
		$instr = flujo.NewIf($exp.val, $instrucciones.lista, nil, nil)
	};

exp
	returns[interfaces.Expresion val]:
	logica { $val = $logica.val }
	| relacional { $val = $relacional.val }
	| aritmetica { $val = $aritmetica.val };

logica
	returns[interfaces.Expresion val]:
	izq = logica op = S_AND der = logica {
		$val = expresion.NewOperacion($izq.val, $der.val, $op.text, false)
	}
	| izq = logica op = S_OR der = logica {
		$val = expresion.NewOperacion($izq.val, $der.val, $op.text, false)
	}
	| relacional { $val = $relacional.val };

relacional
	returns[interfaces.Expresion val]:
	izq = relacional op = (
		S_MENOR
		| S_MAYOR
		| S_MENORI
		| S_MAYORI
		| S_IGUAL
		| S_DIFERENTE
	) der = relacional {
		$val = expresion.NewOperacion($izq.val, $der.val, $op.text, false)
	}
	| aritmetica { $val = $aritmetica.val };

aritmetica
	returns[interfaces.Expresion val]:
	'-' exp { $val = expresion.NewOperacion($exp.val, nil, "-", true) }
	| izq = aritmetica op = ('*' | '/' | '%') der = aritmetica {
		$val = expresion.NewOperacion($izq.val, $der.val, $op.text, false) 
	}
	| izq = aritmetica op = ('+' | '-') der = aritmetica {
		$val = expresion.NewOperacion($izq.val, $der.val, $op.text, false) 
	}
	| exp_valor { $val = $exp_valor.val }
	| S_APAR exp S_CPAR { $val = $exp.val };

exp_valor
	returns[interfaces.Expresion val]:
	primitivo { $val = $primitivo.val };

primitivo
	returns[interfaces.Expresion val]:
	NUMERO {
		val, err := strconv.Atoi($NUMERO.text)
		if err != nil {
			fmt.Println(err)
		}
		$val = expresion.NewPrimitivo(val, entorno.INTEGER)
	}
	| DECIMAL {
		val, err := strconv.ParseFloat($DECIMAL.text, 64)
		if err != nil {
			fmt.Println(err)
		}
		$val = expresion.NewPrimitivo(val, entorno.FLOAT)
	}
	| CADENA {
		val := $CADENA.text[1: len($CADENA.text) - 1]
		$val = expresion.NewPrimitivo(val, entorno.STRING)
	}
	| R_TRUE {
		$val = expresion.NewPrimitivo(true, entorno.BOOLEAN)
	}
	| R_FALSE {
		$val = expresion.NewPrimitivo(false, entorno.BOOLEAN)
	}
	| ID {
		$val = expresion.NewIdentificador($ID.text)
	};
