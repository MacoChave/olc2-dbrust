package expresion

import (
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

/* TIPO RESULTANTE - SUMA
		| int   | float | char  | bool   | string | null  |
int		| int   | null  | null  | null   | null   | null  |
float	| null  | float | null  | null   | null   | null  |
char	| null 	| null  | null  | null   | null   | null  |
bool	| null 	| null  | null  | null   | null   | null  |
string	| null 	| null  | null  | null   | string | null  |
null	| null 	| null  | null  | null   | null   | null  |
*/

/* TIPO RESULTANTE - NEGATIVO
		| int   | float | char  | bool   | string | null  |
int		| int   | null  | null  | null   | null   | null  |
float	| null  | float | null  | null   | null   | null  |
char	| null 	| null  | null  | null   | null   | null  |
bool	| null 	| null  | null  | null   | null   | null  |
string	| null 	| null  | null  | null   | null   | null  |
null	| null 	| null  | null  | null   | null   | null  |
*/

/* TIPO RESULTANTE - RESTA, MULTIPLICACION, DIVISION, MODULO
		| int   | float | char  | bool   | string | null  |
int		| int   | null  | null  | null   | null   | null  |
float	| null  | float | null  | null   | null   | null  |
char	| null 	| null  | null  | null   | null   | null  |
bool	| null 	| null  | null  | null   | null   | null  |
string	| null 	| null  | null  | null   | null   | null  |
null	| null 	| null  | null  | null   | null   | null  |
*/

/* TIPO RESULTANTE - RELACIONAL
		| int   | float | char  | bool   | string | null  |
int		| bool  | null  | null  | null   | null   | null  |
float	| null  | bool  | null  | null   | null   | null  |
char	| null 	| null  | null  | null   | null   | null  |
bool	| null 	| null  | null  | null   | null   | null  |
string	| null 	| null  | null  | null   | bool   | null  |
null	| null 	| null  | null  | null   | null   | null  |
*/

/* TIPO RESULTANTE - LOGICA
		| int   | float | char  | bool   | string | null  |
int		| null  | null  | null  | null   | null   | null  |
float	| null  | null  | null  | null   | null   | null  |
char	| null 	| null  | null  | null   | null   | null  |
bool	| null 	| null  | null  | bool   | null   | null  |
string	| null 	| null  | null  | null   | null   | null  |
null	| null 	| null  | null  | null   | null   | null  |
*/

var tipoResult_suma = [6][6]entorno.TipoDato{
	{entorno.INTEGER, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var tipoResult_neg = [6]entorno.TipoDato{
	entorno.INTEGER, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL,
}

var tipoResult_arit = [6][6]entorno.TipoDato{
	{entorno.INTEGER, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var tipoResult_rel = [6][6]entorno.TipoDato{
	{entorno.BOOLEAN, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var tipoResult_log = [6][6]entorno.TipoDato{
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var tipoResult_not = [6]entorno.TipoDato{
	entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL,
}

type Operacion struct {
	Izquierdo interfaces.Expresion
	Derecho   interfaces.Expresion
	Operador  string
	IsUnario  bool
}

func NewOperacion(izq interfaces.Expresion, der interfaces.Expresion, op string, isUnario bool) Operacion {
	operacion := Operacion{
		Izquierdo: izq,
		Derecho:   der,
		Operador:  op,
		IsUnario:  isUnario,
	}
	return operacion
}

func (op Operacion) GetValor(ent entorno.Entorno) entorno.Valor {
	var valorIzq entorno.Valor
	var valorDer entorno.Valor

	if op.IsUnario {
		valorIzq = op.Izquierdo.GetValor(ent)
	} else {
		if reflect.TypeOf(op.Izquierdo).Name() == "Identificador" {
			existId := ent.ExistSimbolo(op.Izquierdo.(Identificador).Identificador)
			if !existId {
				return entorno.Valor{
					Tipo:  entorno.NULL,
					Valor: nil,
				}
			}
		}
		if reflect.TypeOf(op.Derecho).Name() == "Identificador" {
			existId := ent.ExistSimbolo(op.Derecho.(Identificador).Identificador)
			if !existId {
				return entorno.Valor{
					Tipo:  entorno.NULL,
					Valor: nil,
				}
			}
		}

		valorIzq = op.Izquierdo.GetValor(ent)
		valorDer = op.Derecho.GetValor(ent)
	}

	var tipoResult entorno.TipoDato

	switch op.Operador {
	case "+":
		{
			tipoResult = tipoResult_suma[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.INTEGER {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: valorIzq.Valor.(int) + valorDer.Valor.(int),
				}
			} else if tipoResult == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: val1 + val2,
				}
			} else if tipoResult == entorno.STRING {
				val1 := fmt.Sprintf("%v", valorIzq.Valor)
				val2 := fmt.Sprintf("%v", valorDer.Valor)
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: val1 + val2,
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "-":
		{
			if op.IsUnario {
				tipoResult = tipoResult_neg[valorIzq.Tipo]
				if tipoResult == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) * -1,
					}
				} else if tipoResult == entorno.FLOAT {
					valU, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valU * -1,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: nil,
					}
				}
			} else {
				tipoResult = tipoResult_arit[valorIzq.Tipo][valorDer.Tipo]
				if tipoResult == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) - valorDer.Valor.(int),
					}
				} else if tipoResult == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: val1 - val2,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: nil,
					}
				}
			}
		}
	case "*":
		{
			tipoResult = tipoResult_arit[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.INTEGER {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: valorIzq.Valor.(int) * valorDer.Valor.(int),
				}
			} else if tipoResult == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: val1 * val2,
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "/":
		{
			tipoResult = tipoResult_arit[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.INTEGER {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: valorIzq.Valor.(int) / valorDer.Valor.(int),
				}
			} else if tipoResult == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: val1 / val2,
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "%":
		{
			tipoResult = tipoResult_arit[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.INTEGER {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: valorIzq.Valor.(int) % valorDer.Valor.(int),
				}
			} else if tipoResult == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: math.Mod(val1, val2),
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case ">":
		{
			tipoResult = tipoResult_rel[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				if valorIzq.Tipo == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) > valorDer.Valor.(int),
					}
				} else if valorIzq.Tipo == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: val1 > val2,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(string) > valorDer.Valor.(string),
					}
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "<":
		{
			tipoResult = tipoResult_rel[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				if valorIzq.Tipo == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) < valorDer.Valor.(int),
					}
				} else if valorIzq.Tipo == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: val1 < val2,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(string) < valorDer.Valor.(string),
					}
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case ">=":
		{
			tipoResult = tipoResult_rel[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				if valorIzq.Tipo == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) >= valorDer.Valor.(int),
					}
				} else if valorIzq.Tipo == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: val1 >= val2,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(string) >= valorDer.Valor.(string),
					}
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "<=":
		{
			tipoResult = tipoResult_rel[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				if valorIzq.Tipo == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) >= valorDer.Valor.(int),
					}
				} else if valorIzq.Tipo == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: val1 <= val2,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(string) <= valorDer.Valor.(string),
					}
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "==":
		{
			tipoResult = tipoResult_rel[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				if valorIzq.Tipo == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) == valorDer.Valor.(int),
					}
				} else if valorIzq.Tipo == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: val1 == val2,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(string) == valorDer.Valor.(string),
					}
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "!=":
		{
			tipoResult = tipoResult_rel[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				if valorIzq.Tipo == entorno.INTEGER {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(int) != valorDer.Valor.(int),
					}
				} else if valorIzq.Tipo == entorno.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valorDer.Valor), 64)
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: val1 != val2,
					}
				} else {
					return entorno.Valor{
						Tipo:  tipoResult,
						Valor: valorIzq.Valor.(string) != valorDer.Valor.(string),
					}
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "!":
		{
			tipoResult = tipoResult_log[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: !valorIzq.Valor.(bool),
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "&&":
		{
			tipoResult = tipoResult_log[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: valorIzq.Valor.(bool) && valorDer.Valor.(bool),
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	case "||":
		{
			tipoResult = tipoResult_log[valorIzq.Tipo][valorDer.Tipo]
			if tipoResult == entorno.BOOLEAN {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: valorIzq.Valor.(bool) || valorDer.Valor.(bool),
				}
			} else {
				return entorno.Valor{
					Tipo:  tipoResult,
					Valor: nil,
				}
			}
		}
	}

	return entorno.Valor{
		Tipo:  entorno.NULL,
		Valor: nil,
	}
}
