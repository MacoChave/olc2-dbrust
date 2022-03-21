// Code generated from sintactico.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // sintactico

import "github.com/antlr/antlr4/runtime/Go/antlr"

// sintacticoListener is a complete listener for a parse tree produced by sintactico.
type sintacticoListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterFunciones is called when entering the funciones production.
	EnterFunciones(c *FuncionesContext)

	// EnterFuncion is called when entering the funcion production.
	EnterFuncion(c *FuncionContext)

	// EnterFuncMain is called when entering the funcMain production.
	EnterFuncMain(c *FuncMainContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterImprimir is called when entering the imprimir production.
	EnterImprimir(c *ImprimirContext)

	// EnterLista_exp is called when entering the lista_exp production.
	EnterLista_exp(c *Lista_expContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterTipo_dato is called when entering the tipo_dato production.
	EnterTipo_dato(c *Tipo_datoContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterLogica is called when entering the logica production.
	EnterLogica(c *LogicaContext)

	// EnterRelacional is called when entering the relacional production.
	EnterRelacional(c *RelacionalContext)

	// EnterAritmetica is called when entering the aritmetica production.
	EnterAritmetica(c *AritmeticaContext)

	// EnterExp_valor is called when entering the exp_valor production.
	EnterExp_valor(c *Exp_valorContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitFunciones is called when exiting the funciones production.
	ExitFunciones(c *FuncionesContext)

	// ExitFuncion is called when exiting the funcion production.
	ExitFuncion(c *FuncionContext)

	// ExitFuncMain is called when exiting the funcMain production.
	ExitFuncMain(c *FuncMainContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitImprimir is called when exiting the imprimir production.
	ExitImprimir(c *ImprimirContext)

	// ExitLista_exp is called when exiting the lista_exp production.
	ExitLista_exp(c *Lista_expContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitTipo_dato is called when exiting the tipo_dato production.
	ExitTipo_dato(c *Tipo_datoContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitLogica is called when exiting the logica production.
	ExitLogica(c *LogicaContext)

	// ExitRelacional is called when exiting the relacional production.
	ExitRelacional(c *RelacionalContext)

	// ExitAritmetica is called when exiting the aritmetica production.
	ExitAritmetica(c *AritmeticaContext)

	// ExitExp_valor is called when exiting the exp_valor production.
	ExitExp_valor(c *Exp_valorContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}
