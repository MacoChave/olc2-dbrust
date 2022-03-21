package flujo

import (
	"db_rust/analizador/ast/interfaces"
	"db_rust/analizador/entorno"
	"fmt"

	"github.com/colegno/arraylist"
)

type SentIf struct {
	Condicion       interfaces.Expresion
	Sentencias_then *arraylist.List
	Lista_elseif    *arraylist.List
	Sentencias_else *arraylist.List
}

func NewIf(condicion interfaces.Expresion, sentencias_then *arraylist.List, lista_elseif *arraylist.List, sentencias_else *arraylist.List) SentIf {
	return SentIf{
		Condicion:       condicion,
		Sentencias_then: sentencias_then,
		Lista_elseif:    lista_elseif,
		Sentencias_else: sentencias_else,
	}
}

func (sentIf SentIf) Ejecutar(ent entorno.Entorno) interface{} {
	valorCondicion := sentIf.Condicion.GetValor(ent)
	fmt.Printf("If %v\n", sentIf)

	if valorCondicion.Tipo != entorno.BOOLEAN {
		fmt.Println("Error semantico. Condición de if debe ser boolean")
		return nil
	}

	if valorCondicion.Valor.(bool) {
		nuevoEntorno := entorno.NewEntorno("IF", &ent)

		for i := 0; i < sentIf.Sentencias_then.Len(); i++ {
			instr := sentIf.Sentencias_then.GetValue(i).(interfaces.Instruccion)
			instr.Ejecutar(nuevoEntorno)
		}
	} else {
		if sentIf.Lista_elseif != nil {
			for _, element := range sentIf.Lista_elseif.ToArray() {
				elseIf := element.(SentIf)
				valorNuevaCond := elseIf.Condicion.GetValor(ent)
				if valorNuevaCond.Tipo == entorno.BOOLEAN {
					fmt.Println("Error semantico. Condición else if debe ser boolean")
					return nil
				}

				if valorNuevaCond.Valor.(bool) {
					nuevoEntorno := entorno.NewEntorno("ELSE IF", &ent)

					for j := 0; j < elseIf.Sentencias_then.Len(); j++ {
						instr := elseIf.Sentencias_then.GetValue(j).(interfaces.Instruccion)
						instr.Ejecutar(nuevoEntorno)
					}

					return nil
				}
			}
		}
		if sentIf.Sentencias_else != nil {
			entornoIf := entorno.NewEntorno("ELSE", &ent)
			for i := 0; i < sentIf.Sentencias_else.Len(); i++ {
				instr := sentIf.Sentencias_else.GetValue(i).(interfaces.Instruccion)
				instr.Ejecutar(entornoIf)
			}
		}
	}
	return nil
}
