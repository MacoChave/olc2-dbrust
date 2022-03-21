package variables

import (
	"db_rust/analizador/ast/expresion"
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"
)

/* TIPO ASIGNACION
		| int   | float | char  | bool   | string | null  |
int		| int   | null  | null  | null   | null   | null  |
float	| null  | float | null  | null   | null   | null  |
char	| null 	| null  | char  | null   | null   | null  |
bool	| null 	| null  | null  | bool   | null   | null  |
string	| null 	| null  | null  | null   | string | null  |
null	| int 	| float | char  | bool   | string | null  |
*/

var tipoDef = [6][6]entorno.TipoDato{
	{entorno.INTEGER, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.CHAR, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL},
	{entorno.INTEGER, entorno.FLOAT, entorno.CHAR, entorno.BOOLEAN, entorno.STRING, entorno.NULL},
}

type Declaracion struct {
	Mutable bool
	Id      interfaces.Expresion
	Tipo    entorno.TipoDato
	Valor   interfaces.Expresion
}

func NewDeclaracion(mutable bool, nombre interfaces.Expresion, tipo entorno.TipoDato, valor interfaces.Expresion) Declaracion {
	return Declaracion{
		Mutable: mutable,
		Id:      nombre,
		Tipo:    tipo,
		Valor:   valor,
	}
}

func (dec Declaracion) Ejecutar(ent entorno.Entorno) interface{} {
	valorInicial := dec.Valor.GetValor(ent)
	tipoResultante := tipoDef[dec.Tipo][valorInicial.Tipo]

	if tipoResultante == entorno.NULL {
		fmt.Println("Error semantico. La asignación no es del tipo de la variable")
		return nil
	}

	variable := dec.Id.(expresion.Identificador)

	if ent.ExistSimbolo(variable.Identificador) {
		fmt.Printf("Error semantico. Variable %s ya declarada", variable.Identificador)
	} else {
		simbolo := entorno.NewSimboloId(0, 0, variable.Identificador, dec.Mutable, valorInicial.Valor, tipoResultante)
		ent.AddSimbolo(variable.Identificador, simbolo)
	}

	// data, err := json.MarshalIndent(ent, "", " ")
	// if err != nil {
	// 	panic(err)
	// }

	// stringQuery := string(data)
	// fmt.Println(stringQuery)
	// fmt.Printf("%v", dec.Id)
	return nil
}
