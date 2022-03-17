package funcion

import (
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"log"
	"reflect"

	"github.com/colegno/arraylist"
)

/* TIPO RESULTANTE - RETORNO
		| int   | float | char  | bool   | string | null  | void |
int		| int   | null  | null  | null   | null   | null  | null |
float	| null  | float | null  | null   | null   | null  | null |
char	| null 	| null  | char  | null   | null   | null  | null |
bool	| null 	| null  | null  | bool   | null   | null  | null |
string	| null 	| null  | null  | null   | string | null  | null |
null	| null 	| null  | null  | null   | null   | null  | null |
void	| null 	| null  | null  | null   | null   | null  | void |
*/

var tipoResult_retorno = [7][7]entorno.TipoDato{
	{entorno.INTEGER, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.CHAR, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.VOID},
}

type Funcion struct {
	entorno.Simbolo
	Parametros    *arraylist.List
	Instrucciones *arraylist.List
}

func NewFuncion(tipo entorno.TipoDato, nombre string, parametros *arraylist.List, instrucciones *arraylist.List) Funcion {
	simbolo := entorno.NewSimboloFunc(0, 0, nombre, tipo, parametros)

	return Funcion{
		Simbolo:       simbolo,
		Parametros:    parametros,
		Instrucciones: instrucciones,
	}
}

func (funcion Funcion) EjecutarParams(ent entorno.Entorno, expresiones *arraylist.List) bool {
	declaraciones := funcion.Parametros.Clone()

	if declaraciones.Len() != expresiones.Len() {
		log.Println("Error en variables")
		return false
	}

	for i := 0; i < declaraciones.Len(); i++ {
		// pivote := declaraciones.GetValue(i).(*definicion.Declaracion)
		// pivote.ValorInicializacion = expresiones.GetValue(i).(interfaces.Expresion)
		// pivote.Ejecutar(ent)
	}
	return true
}

func (funcion Funcion) Ejecutar(ent entorno.Entorno) interface{} {
	for i := 0; i < funcion.Instrucciones.Len(); i++ {
		instruccionActual := funcion.Instrucciones.GetValue(i).(interfaces.Instruccion)
		valorInstruccion := instruccionActual.Ejecutar(ent)
		tipoFuncion := funcion.Tipo

		if valorInstruccion != nil {
			if reflect.TypeOf(valorInstruccion) != reflect.TypeOf(entorno.Valor{}) {
				log.Println("Error en función, se esperaba un retorno válido")
				return entorno.Valor{
					Tipo:  entorno.NULL,
					Valor: -1,
				}
			}

			tipoRetorno := valorInstruccion.(entorno.Valor)
			validarTipos := tipoResult_retorno[tipoFuncion][tipoRetorno.Tipo]
			if validarTipos == entorno.NULL {
				log.Println("Error de tipos, se esperaba un retorno de igual tipo")
				return entorno.Valor{
					Tipo:  entorno.NULL,
					Valor: -1,
				}
			}

			return tipoRetorno
		}
	}
	return entorno.Valor{
		Tipo:  entorno.NULL,
		Valor: -1,
	}
}
