package flujo

import (
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"

	"github.com/colegno/arraylist"
)

type SentWhile struct {
	Condicion  interfaces.Expresion
	Sentencias *arraylist.List
}

func NewSentWhile(condicion interfaces.Expresion, sentencias *arraylist.List) SentWhile {
	return SentWhile{
		Condicion:  condicion,
		Sentencias: sentencias,
	}
}

func (sentWhile SentWhile) Ejecutar(ent entorno.Entorno) interface{} {
	resultadoCond := sentWhile.Condicion.GetValor(ent)
	if resultadoCond.Tipo != entorno.BOOLEAN {
		fmt.Println("Error semantico. Se esperaba tipo boolean en condicion de while")
		return nil
	}

	if resultadoCond.Valor.(bool) {
		condicion := true
		for condicion {
			nuevoEntorno := entorno.NewEntorno("WHILE", &ent)
			for i := 0; i < sentWhile.Sentencias.Len(); i++ {
				sentencia := sentWhile.Sentencias.GetValue(i).(interfaces.Instruccion)
				sentencia.Ejecutar(ent)
			}
			condicion = sentWhile.Condicion.GetValor(nuevoEntorno).Valor.(bool)
		}
	}
	return nil
}
