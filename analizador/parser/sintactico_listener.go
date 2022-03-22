// Code generated from sintactico.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // sintactico

import "github.com/antlr/antlr4/runtime/Go/antlr"

// sintacticoListener is a complete listener for a parse tree produced by sintactico.
type sintacticoListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterProcedimientos is called when entering the procedimientos production.
	EnterProcedimientos(c *ProcedimientosContext)

	// EnterProcedimiento is called when entering the procedimiento production.
	EnterProcedimiento(c *ProcedimientoContext)

	// EnterPrincipal is called when entering the principal production.
	EnterPrincipal(c *PrincipalContext)

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

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterSentIf is called when entering the sentIf production.
	EnterSentIf(c *SentIfContext)

	// EnterListaElseIf is called when entering the listaElseIf production.
	EnterListaElseIf(c *ListaElseIfContext)

	// EnterElseIf is called when entering the elseIf production.
	EnterElseIf(c *ElseIfContext)

	// EnterSentMatch is called when entering the sentMatch production.
	EnterSentMatch(c *SentMatchContext)

	// EnterCasosMatch is called when entering the casosMatch production.
	EnterCasosMatch(c *CasosMatchContext)

	// EnterCasoMatch is called when entering the casoMatch production.
	EnterCasoMatch(c *CasoMatchContext)

	// EnterMatchDefecto is called when entering the matchDefecto production.
	EnterMatchDefecto(c *MatchDefectoContext)

	// EnterSentWhile is called when entering the sentWhile production.
	EnterSentWhile(c *SentWhileContext)

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

	// ExitProcedimientos is called when exiting the procedimientos production.
	ExitProcedimientos(c *ProcedimientosContext)

	// ExitProcedimiento is called when exiting the procedimiento production.
	ExitProcedimiento(c *ProcedimientoContext)

	// ExitPrincipal is called when exiting the principal production.
	ExitPrincipal(c *PrincipalContext)

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

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitSentIf is called when exiting the sentIf production.
	ExitSentIf(c *SentIfContext)

	// ExitListaElseIf is called when exiting the listaElseIf production.
	ExitListaElseIf(c *ListaElseIfContext)

	// ExitElseIf is called when exiting the elseIf production.
	ExitElseIf(c *ElseIfContext)

	// ExitSentMatch is called when exiting the sentMatch production.
	ExitSentMatch(c *SentMatchContext)

	// ExitCasosMatch is called when exiting the casosMatch production.
	ExitCasosMatch(c *CasosMatchContext)

	// ExitCasoMatch is called when exiting the casoMatch production.
	ExitCasoMatch(c *CasoMatchContext)

	// ExitMatchDefecto is called when exiting the matchDefecto production.
	ExitMatchDefecto(c *MatchDefectoContext)

	// ExitSentWhile is called when exiting the sentWhile production.
	ExitSentWhile(c *SentWhileContext)

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
