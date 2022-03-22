// Code generated from sintactico.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // sintactico

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesintacticoListener is a complete listener for a parse tree produced by sintactico.
type BasesintacticoListener struct{}

var _ sintacticoListener = &BasesintacticoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesintacticoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesintacticoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesintacticoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesintacticoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasesintacticoListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasesintacticoListener) ExitStart(ctx *StartContext) {}

// EnterProcedimientos is called when production procedimientos is entered.
func (s *BasesintacticoListener) EnterProcedimientos(ctx *ProcedimientosContext) {}

// ExitProcedimientos is called when production procedimientos is exited.
func (s *BasesintacticoListener) ExitProcedimientos(ctx *ProcedimientosContext) {}

// EnterProcedimiento is called when production procedimiento is entered.
func (s *BasesintacticoListener) EnterProcedimiento(ctx *ProcedimientoContext) {}

// ExitProcedimiento is called when production procedimiento is exited.
func (s *BasesintacticoListener) ExitProcedimiento(ctx *ProcedimientoContext) {}

// EnterPrincipal is called when production principal is entered.
func (s *BasesintacticoListener) EnterPrincipal(ctx *PrincipalContext) {}

// ExitPrincipal is called when production principal is exited.
func (s *BasesintacticoListener) ExitPrincipal(ctx *PrincipalContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BasesintacticoListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BasesintacticoListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BasesintacticoListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BasesintacticoListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterImprimir is called when production imprimir is entered.
func (s *BasesintacticoListener) EnterImprimir(ctx *ImprimirContext) {}

// ExitImprimir is called when production imprimir is exited.
func (s *BasesintacticoListener) ExitImprimir(ctx *ImprimirContext) {}

// EnterLista_exp is called when production lista_exp is entered.
func (s *BasesintacticoListener) EnterLista_exp(ctx *Lista_expContext) {}

// ExitLista_exp is called when production lista_exp is exited.
func (s *BasesintacticoListener) ExitLista_exp(ctx *Lista_expContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BasesintacticoListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BasesintacticoListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterTipo_dato is called when production tipo_dato is entered.
func (s *BasesintacticoListener) EnterTipo_dato(ctx *Tipo_datoContext) {}

// ExitTipo_dato is called when production tipo_dato is exited.
func (s *BasesintacticoListener) ExitTipo_dato(ctx *Tipo_datoContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BasesintacticoListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BasesintacticoListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterSent_if is called when production sent_if is entered.
func (s *BasesintacticoListener) EnterSent_if(ctx *Sent_ifContext) {}

// ExitSent_if is called when production sent_if is exited.
func (s *BasesintacticoListener) ExitSent_if(ctx *Sent_ifContext) {}

// EnterLista_elseif is called when production lista_elseif is entered.
func (s *BasesintacticoListener) EnterLista_elseif(ctx *Lista_elseifContext) {}

// ExitLista_elseif is called when production lista_elseif is exited.
func (s *BasesintacticoListener) ExitLista_elseif(ctx *Lista_elseifContext) {}

// EnterElseif is called when production elseif is entered.
func (s *BasesintacticoListener) EnterElseif(ctx *ElseifContext) {}

// ExitElseif is called when production elseif is exited.
func (s *BasesintacticoListener) ExitElseif(ctx *ElseifContext) {}

// EnterSent_match is called when production sent_match is entered.
func (s *BasesintacticoListener) EnterSent_match(ctx *Sent_matchContext) {}

// ExitSent_match is called when production sent_match is exited.
func (s *BasesintacticoListener) ExitSent_match(ctx *Sent_matchContext) {}

// EnterCasosMatch is called when production casosMatch is entered.
func (s *BasesintacticoListener) EnterCasosMatch(ctx *CasosMatchContext) {}

// ExitCasosMatch is called when production casosMatch is exited.
func (s *BasesintacticoListener) ExitCasosMatch(ctx *CasosMatchContext) {}

// EnterCasoMatch is called when production casoMatch is entered.
func (s *BasesintacticoListener) EnterCasoMatch(ctx *CasoMatchContext) {}

// ExitCasoMatch is called when production casoMatch is exited.
func (s *BasesintacticoListener) ExitCasoMatch(ctx *CasoMatchContext) {}

// EnterMatchDefecto is called when production matchDefecto is entered.
func (s *BasesintacticoListener) EnterMatchDefecto(ctx *MatchDefectoContext) {}

// ExitMatchDefecto is called when production matchDefecto is exited.
func (s *BasesintacticoListener) ExitMatchDefecto(ctx *MatchDefectoContext) {}

// EnterExp is called when production exp is entered.
func (s *BasesintacticoListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BasesintacticoListener) ExitExp(ctx *ExpContext) {}

// EnterLogica is called when production logica is entered.
func (s *BasesintacticoListener) EnterLogica(ctx *LogicaContext) {}

// ExitLogica is called when production logica is exited.
func (s *BasesintacticoListener) ExitLogica(ctx *LogicaContext) {}

// EnterRelacional is called when production relacional is entered.
func (s *BasesintacticoListener) EnterRelacional(ctx *RelacionalContext) {}

// ExitRelacional is called when production relacional is exited.
func (s *BasesintacticoListener) ExitRelacional(ctx *RelacionalContext) {}

// EnterAritmetica is called when production aritmetica is entered.
func (s *BasesintacticoListener) EnterAritmetica(ctx *AritmeticaContext) {}

// ExitAritmetica is called when production aritmetica is exited.
func (s *BasesintacticoListener) ExitAritmetica(ctx *AritmeticaContext) {}

// EnterExp_valor is called when production exp_valor is entered.
func (s *BasesintacticoListener) EnterExp_valor(ctx *Exp_valorContext) {}

// ExitExp_valor is called when production exp_valor is exited.
func (s *BasesintacticoListener) ExitExp_valor(ctx *Exp_valorContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BasesintacticoListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BasesintacticoListener) ExitPrimitivo(ctx *PrimitivoContext) {}
