package interfaces

import "db_rust/analizador/entorno"

type Expresion interface {
	GetValor(ent entorno.Entorno) entorno.Valor
}

type Instruccion interface {
	Ejecutar(ent entorno.Entorno) interface{}
}
