package flujo

import (
	"db_rust/analizador/ast/expresion"
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
)

type SentFor struct {
	Id     expresion.Identificador
	Inicio interfaces.Expresion
	Final  interfaces.Expresion
}

func NewSentFor(id expresion.Identificador, inicio interfaces.Expresion, final interfaces.Expresion) SentFor {
	return SentFor{
		Id:     id,
		Inicio: inicio,
		Final:  final,
	}
}

func (sentFor SentFor) Ejecutar(ent entorno.Entorno) interface{} {
	return nil
}
