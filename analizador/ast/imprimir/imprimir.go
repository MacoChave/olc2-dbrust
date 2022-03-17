package imprimir

import (
	"db_rust/analizador"
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"
)

type Imprimir struct {
	Expresiones interfaces.Expresion
}

func NewImprimir(lista interfaces.Expresion) Imprimir {
	return Imprimir{
		Expresiones: lista,
	}
}

func (impr Imprimir) Ejecutar(ent entorno.Entorno) interface{} {
	valor := impr.Expresiones.GetValor(ent)
	result := fmt.Sprintf("%v\n", valor.Valor)
	analizador.Consola += result

	return nil
}
