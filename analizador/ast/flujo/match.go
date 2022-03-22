package flujo

import (
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

/* TIPO RESULTANTE - RELACIONAL
		| int   | float | char  | bool   | string | null  |
int		| bool  | null  | null  | null   | null   | null  |
float	| null  | bool  | null  | null   | null   | null  |
char	| null 	| null  | null  | null   | null   | null  |
bool	| null 	| null  | null  | null   | null   | null  |
string	| null 	| null  | null  | null   | bool   | null  |
null	| null 	| null  | null  | null   | null   | null  |
*/

var tipoResult = [6][6]entorno.TipoDato{
	{entorno.BOOLEAN, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

type CaseMatch struct {
	Condicion interfaces.Expresion
	Sentencia interfaces.Instruccion
}

func NewCaseMatch(condicion interfaces.Expresion, sentencia interfaces.Instruccion) CaseMatch {
	return CaseMatch{
		Condicion: condicion,
		Sentencia: sentencia,
	}
}

type SentMatch struct {
	Expresion interfaces.Expresion
	Casos     *arraylist.List
	Defecto   CaseMatch
}

func NewSentMatch(expresion interfaces.Expresion, casos *arraylist.List, defecto CaseMatch) SentMatch {
	return SentMatch{
		Expresion: expresion,
		Casos:     casos,
		Defecto:   defecto,
	}
}

func CompareMatch(ent entorno.Entorno, expresion interfaces.Expresion, caso CaseMatch) interface{} {
	resultadoExp := expresion.GetValor(ent)
	resultadoComp := caso.Condicion.GetValor(ent)

	tipoResultado := tipoResult[resultadoExp.Tipo][resultadoComp.Tipo]
	if tipoResultado == entorno.NULL {
		return entorno.Valor{
			Tipo:  entorno.NULL,
			Valor: nil,
		}
	}

	comparacion := false
	if resultadoComp.Tipo == entorno.INTEGER {
		comparacion = resultadoExp.Valor.(int) == resultadoComp.Valor.(int)
	} else if resultadoComp.Tipo == entorno.FLOAT {
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultadoExp.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultadoComp.Valor), 64)
		comparacion = val1 == val2
	} else if resultadoComp.Tipo == entorno.STRING {
		comparacion = resultadoExp.Valor.(string) == resultadoComp.Valor.(string)
	}

	if comparacion {
		caso.Sentencia.Ejecutar(ent)
		return entorno.Valor{
			Tipo:  entorno.VOID,
			Valor: nil,
		}
	}
	return entorno.Valor{
		Tipo:  entorno.NULL,
		Valor: nil,
	}
}

func (sentMatch SentMatch) Ejecutar(ent entorno.Entorno) interface{} {
	for i := 0; i < sentMatch.Casos.Len(); i++ {
		caso := sentMatch.Casos.GetValue(i).(CaseMatch)
		resultado := CompareMatch(ent, sentMatch.Expresion, caso).(entorno.Valor)
		if resultado.Tipo != entorno.NULL {
			return resultado
		}
	}
	sentMatch.Defecto.Sentencia.Ejecutar(ent)

	return nil
}
