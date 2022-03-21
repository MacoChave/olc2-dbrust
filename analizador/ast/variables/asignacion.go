package variables

import (
	"db_rust/analizador/ast/expresion"
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"
)

type Asignacion struct {
	Id    interfaces.Expresion
	Valor interfaces.Expresion
}

func NewAsignacion(id interfaces.Expresion, valor interfaces.Expresion) Asignacion {
	return Asignacion{
		Id:    id,
		Valor: valor,
	}
}

func (asign Asignacion) Ejecutar(ent entorno.Entorno) interface{} {
	id := asign.Id.(expresion.Identificador)
	if !ent.ExistSimbolo(id.Identificador) {
		fmt.Printf("Error semantico. Variable %s no ha sido declarado", id.Identificador)
		return nil
	}

	variable := ent.GetSimbolo(id.Identificador)

	if variable.IsConst {
		fmt.Println("Error semantico. No se puede asignar valor a una constante")
		return nil
	}

	valor := asign.Valor.GetValor(ent)
	tipoResultante := tipoDef[variable.Tipo][valor.Tipo]
	if tipoResultante == entorno.NULL {
		fmt.Println("Error semantico. La asignación no es del tipo de la variable")
		return nil
	}

	variable.Valor = valor.Valor
	variable.Tipo = tipoResultante
	ent.UpdateSimbolo(id.Identificador, variable)

	return nil
}
