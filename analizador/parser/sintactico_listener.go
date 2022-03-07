// Code generated from sintactico.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // sintactico

import "github.com/antlr/antlr4/runtime/Go/antlr"

// sintacticoListener is a complete listener for a parse tree produced by sintactico.
type sintacticoListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

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
