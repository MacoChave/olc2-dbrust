package imprimir

import (
	"db_rust/analizador"
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"
	"strings"

	"github.com/colegno/arraylist"
)

type Imprimir struct {
	Expresiones *arraylist.List
}

func NewImprimir(lista *arraylist.List) Imprimir {
	return Imprimir{
		Expresiones: lista,
	}
}

func (impr Imprimir) Ejecutar(ent entorno.Entorno) interface{} {
	exp := impr.Expresiones.GetValue(0).(interfaces.Expresion)
	valorFormato := exp.GetValor(ent)
	result := fmt.Sprintf("%v\n", valorFormato.Valor)
	if valorFormato.Tipo != entorno.STRING {
		if impr.Expresiones.Len() > 1 {
			fmt.Println("Error semántico. No se esperaba expresiones")
			return nil
		} else if valorFormato.Tipo == entorno.NULL {
			fmt.Println("Error semántico. No se ha encontrado el valor")
			return nil
		}
		analizador.Consola += result
	} else {
		cantPrimitivos := strings.Count(result, "{}")
		// cantObjs := strings.Count(result, "{:?}")
		if cantPrimitivos != impr.Expresiones.Len()-1 {
			fmt.Println("Error semántico. Cantidad incorrecta de expresiones")
			return nil
		}
		for i := 1; i < impr.Expresiones.Len(); i++ {
			exp := impr.Expresiones.GetValue(i).(interfaces.Expresion)
			valor := exp.GetValor(ent)
			if valor.Tipo == entorno.NULL {
				fmt.Println("Error semántico. No se ha encontrado el valor")
			} else {
				strResult := fmt.Sprintf("%v", valor.Valor)
				result = strings.Replace(result, "{}", strResult, 1)
			}
		}
		analizador.Consola += result
	}

	return nil
}
