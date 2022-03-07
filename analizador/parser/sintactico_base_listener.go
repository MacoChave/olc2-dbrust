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