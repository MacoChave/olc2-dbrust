// Code generated from sintactico.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // sintactico

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "db_rust/analizador/ast"
import "db_rust/analizador/ast/expresion"
import "db_rust/analizador/ast/funcion"
import "db_rust/analizador/ast/imprimir"
import "db_rust/analizador/ast/variables"
import "db_rust/analizador/ast/interfaces"
import "db_rust/analizador/entorno"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 84, 230,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 3, 7, 3, 39, 10, 3, 12, 3, 14, 3, 42, 11, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 6, 6, 59, 10, 6, 13, 6, 14, 6, 60, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 73, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 90,
	10, 9, 12, 9, 14, 9, 93, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 125, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 137, 10, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 148, 10, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 7, 13, 164, 10, 13, 12, 13, 14, 13, 167, 11, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 178,
	10, 14, 12, 14, 14, 14, 181, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 196, 10,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	7, 15, 208, 10, 15, 12, 15, 14, 15, 211, 11, 15, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 5, 17, 228, 10, 17, 3, 17, 2, 6, 16, 24, 26, 28, 18, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 5, 3, 2, 16, 21, 3,
	2, 13, 15, 3, 2, 11, 12, 2, 238, 2, 34, 3, 2, 2, 2, 4, 40, 3, 2, 2, 2,
	6, 45, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 58, 3, 2, 2, 2, 12, 72, 3, 2,
	2, 2, 14, 74, 3, 2, 2, 2, 16, 80, 3, 2, 2, 2, 18, 124, 3, 2, 2, 2, 20,
	136, 3, 2, 2, 2, 22, 147, 3, 2, 2, 2, 24, 149, 3, 2, 2, 2, 26, 168, 3,
	2, 2, 2, 28, 195, 3, 2, 2, 2, 30, 212, 3, 2, 2, 2, 32, 227, 3, 2, 2, 2,
	34, 35, 5, 4, 3, 2, 35, 36, 8, 2, 1, 2, 36, 3, 3, 2, 2, 2, 37, 39, 5, 6,
	4, 2, 38, 37, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41,
	3, 2, 2, 2, 41, 43, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 8, 3, 1, 2,
	44, 5, 3, 2, 2, 2, 45, 46, 5, 8, 5, 2, 46, 47, 8, 4, 1, 2, 47, 7, 3, 2,
	2, 2, 48, 49, 7, 44, 2, 2, 49, 50, 7, 45, 2, 2, 50, 51, 7, 25, 2, 2, 51,
	52, 7, 26, 2, 2, 52, 53, 7, 31, 2, 2, 53, 54, 5, 10, 6, 2, 54, 55, 7, 32,
	2, 2, 55, 56, 8, 5, 1, 2, 56, 9, 3, 2, 2, 2, 57, 59, 5, 12, 7, 2, 58, 57,
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2,
	61, 62, 3, 2, 2, 2, 62, 63, 8, 6, 1, 2, 63, 11, 3, 2, 2, 2, 64, 65, 5,
	14, 8, 2, 65, 66, 7, 5, 2, 2, 66, 67, 8, 7, 1, 2, 67, 73, 3, 2, 2, 2, 68,
	69, 5, 18, 10, 2, 69, 70, 7, 5, 2, 2, 70, 71, 8, 7, 1, 2, 71, 73, 3, 2,
	2, 2, 72, 64, 3, 2, 2, 2, 72, 68, 3, 2, 2, 2, 73, 13, 3, 2, 2, 2, 74, 75,
	7, 51, 2, 2, 75, 76, 7, 25, 2, 2, 76, 77, 5, 16, 9, 2, 77, 78, 7, 26, 2,
	2, 78, 79, 8, 8, 1, 2, 79, 15, 3, 2, 2, 2, 80, 81, 8, 9, 1, 2, 81, 82,
	5, 22, 12, 2, 82, 83, 8, 9, 1, 2, 83, 91, 3, 2, 2, 2, 84, 85, 12, 4, 2,
	2, 85, 86, 7, 4, 2, 2, 86, 87, 5, 22, 12, 2, 87, 88, 8, 9, 1, 2, 88, 90,
	3, 2, 2, 2, 89, 84, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2,
	91, 92, 3, 2, 2, 2, 92, 17, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 7,
	36, 2, 2, 95, 96, 7, 37, 2, 2, 96, 97, 7, 82, 2, 2, 97, 98, 7, 27, 2, 2,
	98, 99, 5, 20, 11, 2, 99, 100, 7, 6, 2, 2, 100, 101, 5, 22, 12, 2, 101,
	102, 8, 10, 1, 2, 102, 125, 3, 2, 2, 2, 103, 104, 7, 36, 2, 2, 104, 105,
	7, 37, 2, 2, 105, 106, 7, 82, 2, 2, 106, 107, 7, 6, 2, 2, 107, 108, 5,
	22, 12, 2, 108, 109, 8, 10, 1, 2, 109, 125, 3, 2, 2, 2, 110, 111, 7, 36,
	2, 2, 111, 112, 7, 82, 2, 2, 112, 113, 7, 27, 2, 2, 113, 114, 5, 20, 11,
	2, 114, 115, 7, 6, 2, 2, 115, 116, 5, 22, 12, 2, 116, 117, 8, 10, 1, 2,
	117, 125, 3, 2, 2, 2, 118, 119, 7, 36, 2, 2, 119, 120, 7, 82, 2, 2, 120,
	121, 7, 6, 2, 2, 121, 122, 5, 22, 12, 2, 122, 123, 8, 10, 1, 2, 123, 125,
	3, 2, 2, 2, 124, 94, 3, 2, 2, 2, 124, 103, 3, 2, 2, 2, 124, 110, 3, 2,
	2, 2, 124, 118, 3, 2, 2, 2, 125, 19, 3, 2, 2, 2, 126, 127, 7, 38, 2, 2,
	127, 137, 8, 11, 1, 2, 128, 129, 7, 39, 2, 2, 129, 137, 8, 11, 1, 2, 130,
	131, 7, 41, 2, 2, 131, 137, 8, 11, 1, 2, 132, 133, 7, 42, 2, 2, 133, 137,
	8, 11, 1, 2, 134, 135, 7, 40, 2, 2, 135, 137, 8, 11, 1, 2, 136, 126, 3,
	2, 2, 2, 136, 128, 3, 2, 2, 2, 136, 130, 3, 2, 2, 2, 136, 132, 3, 2, 2,
	2, 136, 134, 3, 2, 2, 2, 137, 21, 3, 2, 2, 2, 138, 139, 5, 24, 13, 2, 139,
	140, 8, 12, 1, 2, 140, 148, 3, 2, 2, 2, 141, 142, 5, 26, 14, 2, 142, 143,
	8, 12, 1, 2, 143, 148, 3, 2, 2, 2, 144, 145, 5, 28, 15, 2, 145, 146, 8,
	12, 1, 2, 146, 148, 3, 2, 2, 2, 147, 138, 3, 2, 2, 2, 147, 141, 3, 2, 2,
	2, 147, 144, 3, 2, 2, 2, 148, 23, 3, 2, 2, 2, 149, 150, 8, 13, 1, 2, 150,
	151, 5, 26, 14, 2, 151, 152, 8, 13, 1, 2, 152, 165, 3, 2, 2, 2, 153, 154,
	12, 5, 2, 2, 154, 155, 7, 23, 2, 2, 155, 156, 5, 24, 13, 6, 156, 157, 8,
	13, 1, 2, 157, 164, 3, 2, 2, 2, 158, 159, 12, 4, 2, 2, 159, 160, 7, 22,
	2, 2, 160, 161, 5, 24, 13, 5, 161, 162, 8, 13, 1, 2, 162, 164, 3, 2, 2,
	2, 163, 153, 3, 2, 2, 2, 163, 158, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165,
	163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 25, 3, 2, 2, 2, 167, 165, 3,
	2, 2, 2, 168, 169, 8, 14, 1, 2, 169, 170, 5, 28, 15, 2, 170, 171, 8, 14,
	1, 2, 171, 179, 3, 2, 2, 2, 172, 173, 12, 4, 2, 2, 173, 174, 9, 2, 2, 2,
	174, 175, 5, 26, 14, 5, 175, 176, 8, 14, 1, 2, 176, 178, 3, 2, 2, 2, 177,
	172, 3, 2, 2, 2, 178, 181, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 180,
	3, 2, 2, 2, 180, 27, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 182, 183, 8, 15,
	1, 2, 183, 184, 7, 12, 2, 2, 184, 185, 5, 22, 12, 2, 185, 186, 8, 15, 1,
	2, 186, 196, 3, 2, 2, 2, 187, 188, 5, 30, 16, 2, 188, 189, 8, 15, 1, 2,
	189, 196, 3, 2, 2, 2, 190, 191, 7, 25, 2, 2, 191, 192, 5, 22, 12, 2, 192,
	193, 7, 26, 2, 2, 193, 194, 8, 15, 1, 2, 194, 196, 3, 2, 2, 2, 195, 182,
	3, 2, 2, 2, 195, 187, 3, 2, 2, 2, 195, 190, 3, 2, 2, 2, 196, 209, 3, 2,
	2, 2, 197, 198, 12, 6, 2, 2, 198, 199, 9, 3, 2, 2, 199, 200, 5, 28, 15,
	7, 200, 201, 8, 15, 1, 2, 201, 208, 3, 2, 2, 2, 202, 203, 12, 5, 2, 2,
	203, 204, 9, 4, 2, 2, 204, 205, 5, 28, 15, 6, 205, 206, 8, 15, 1, 2, 206,
	208, 3, 2, 2, 2, 207, 197, 3, 2, 2, 2, 207, 202, 3, 2, 2, 2, 208, 211,
	3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 29, 3, 2,
	2, 2, 211, 209, 3, 2, 2, 2, 212, 213, 5, 32, 17, 2, 213, 214, 8, 16, 1,
	2, 214, 31, 3, 2, 2, 2, 215, 216, 7, 78, 2, 2, 216, 228, 8, 17, 1, 2, 217,
	218, 7, 79, 2, 2, 218, 228, 8, 17, 1, 2, 219, 220, 7, 81, 2, 2, 220, 228,
	8, 17, 1, 2, 221, 222, 7, 34, 2, 2, 222, 228, 8, 17, 1, 2, 223, 224, 7,
	35, 2, 2, 224, 228, 8, 17, 1, 2, 225, 226, 7, 82, 2, 2, 226, 228, 8, 17,
	1, 2, 227, 215, 3, 2, 2, 2, 227, 217, 3, 2, 2, 2, 227, 219, 3, 2, 2, 2,
	227, 221, 3, 2, 2, 2, 227, 223, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228,
	33, 3, 2, 2, 2, 16, 40, 60, 72, 91, 124, 136, 147, 163, 165, 179, 195,
	207, 209, 227,
}
var literalNames = []string{
	"", "'.'", "','", "';'", "'='", "'&'", "'|'", "'=>'", "'_'", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'||'",
	"'&&'", "'->'", "'('", "')'", "':'", "'::'", "'['", "']'", "'{'", "'}'",
	"'..'", "'true'", "'false'", "'let'", "'mut'", "'i64'", "'f64'", "'bool'",
	"'char'", "'string'", "'&str'", "'fn'", "'main'", "'as'", "'to_owned'",
	"'to_string'", "'pow'", "'powF'", "'println!'", "'abs'", "'sqrt'", "'clone'",
	"'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'",
	"'with_capacity'", "'if'", "'else'", "'match'", "'loop'", "'break'", "'while'",
	"'for'", "'in'", "'continue'", "'return'", "'struct'", "'vector'", "'vec'",
	"'mod'", "'pub'",
}
var symbolicNames = []string{
	"", "S_PUNTO", "S_COMA", "S_PTCOMA", "S_ASIGNAR", "S_AMP", "S_MATCH_OR",
	"S_MATCH_RET", "S_MATCH_DEFAULT", "S_SUMA", "S_RESTA", "S_POR", "S_DIVISION",
	"S_MODULO", "S_MAYOR", "S_MENOR", "S_MAYORI", "S_MENORI", "S_IGUAL", "S_DIFERENTE",
	"S_OR", "S_AND", "S_MOD_RET", "S_APAR", "S_CPAR", "S_DOSPT", "S_ALIAS",
	"S_ACOR", "S_CCOR", "S_ALLAV", "S_CLLAV", "S_RANGO", "R_TRUE", "R_FALSE",
	"R_LET", "R_MUT", "R_INT", "R_FLOAT", "R_BOOL", "R_CHAR", "R_STRING", "R_STR",
	"R_FN", "R_MAIN", "R_AS", "R_TO_OWNED", "R_TO_STRING", "R_POW", "R_POWF",
	"R_PRINTLN", "R_ABS", "R_SQRT", "R_CLONE", "R_NEW", "R_LEN", "R_PUSH",
	"R_REMOVE", "R_CONTAINS", "R_INSERT", "R_CAPACITY", "R_WITH_CAPACITY",
	"R_IF", "R_ELSE", "R_MATCH", "R_LOOP", "R_BREAK", "R_WHILE", "R_FOR", "R_IN",
	"R_CONTINUE", "R_RETURN", "R_STRUCT", "R_VECTOR", "R_VEC", "R_MOD", "R_PUB",
	"NUMERO", "DECIMAL", "CARACTER", "CADENA", "ID", "COMENTARIO", "BLANCOS",
}

var ruleNames = []string{
	"start", "funciones", "funcion", "funcMain", "instrucciones", "instruccion",
	"imprimir", "lista_exp", "declaracion", "tipo_dato", "exp", "logica", "relacional",
	"aritmetica", "exp_valor", "primitivo",
}

type sintactico struct {
	*antlr.BaseParser
}

// Newsintactico produces a new parser instance for the optional input antlr.TokenStream.
//
// The *sintactico instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newsintactico(input antlr.TokenStream) *sintactico {
	this := new(sintactico)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "sintactico.g4"

	return this
}

// sintactico tokens.
const (
	sintacticoEOF             = antlr.TokenEOF
	sintacticoS_PUNTO         = 1
	sintacticoS_COMA          = 2
	sintacticoS_PTCOMA        = 3
	sintacticoS_ASIGNAR       = 4
	sintacticoS_AMP           = 5
	sintacticoS_MATCH_OR      = 6
	sintacticoS_MATCH_RET     = 7
	sintacticoS_MATCH_DEFAULT = 8
	sintacticoS_SUMA          = 9
	sintacticoS_RESTA         = 10
	sintacticoS_POR           = 11
	sintacticoS_DIVISION      = 12
	sintacticoS_MODULO        = 13
	sintacticoS_MAYOR         = 14
	sintacticoS_MENOR         = 15
	sintacticoS_MAYORI        = 16
	sintacticoS_MENORI        = 17
	sintacticoS_IGUAL         = 18
	sintacticoS_DIFERENTE     = 19
	sintacticoS_OR            = 20
	sintacticoS_AND           = 21
	sintacticoS_MOD_RET       = 22
	sintacticoS_APAR          = 23
	sintacticoS_CPAR          = 24
	sintacticoS_DOSPT         = 25
	sintacticoS_ALIAS         = 26
	sintacticoS_ACOR          = 27
	sintacticoS_CCOR          = 28
	sintacticoS_ALLAV         = 29
	sintacticoS_CLLAV         = 30
	sintacticoS_RANGO         = 31
	sintacticoR_TRUE          = 32
	sintacticoR_FALSE         = 33
	sintacticoR_LET           = 34
	sintacticoR_MUT           = 35
	sintacticoR_INT           = 36
	sintacticoR_FLOAT         = 37
	sintacticoR_BOOL          = 38
	sintacticoR_CHAR          = 39
	sintacticoR_STRING        = 40
	sintacticoR_STR           = 41
	sintacticoR_FN            = 42
	sintacticoR_MAIN          = 43
	sintacticoR_AS            = 44
	sintacticoR_TO_OWNED      = 45
	sintacticoR_TO_STRING     = 46
	sintacticoR_POW           = 47
	sintacticoR_POWF          = 48
	sintacticoR_PRINTLN       = 49
	sintacticoR_ABS           = 50
	sintacticoR_SQRT          = 51
	sintacticoR_CLONE         = 52
	sintacticoR_NEW           = 53
	sintacticoR_LEN           = 54
	sintacticoR_PUSH          = 55
	sintacticoR_REMOVE        = 56
	sintacticoR_CONTAINS      = 57
	sintacticoR_INSERT        = 58
	sintacticoR_CAPACITY      = 59
	sintacticoR_WITH_CAPACITY = 60
	sintacticoR_IF            = 61
	sintacticoR_ELSE          = 62
	sintacticoR_MATCH         = 63
	sintacticoR_LOOP          = 64
	sintacticoR_BREAK         = 65
	sintacticoR_WHILE         = 66
	sintacticoR_FOR           = 67
	sintacticoR_IN            = 68
	sintacticoR_CONTINUE      = 69
	sintacticoR_RETURN        = 70
	sintacticoR_STRUCT        = 71
	sintacticoR_VECTOR        = 72
	sintacticoR_VEC           = 73
	sintacticoR_MOD           = 74
	sintacticoR_PUB           = 75
	sintacticoNUMERO          = 76
	sintacticoDECIMAL         = 77
	sintacticoCARACTER        = 78
	sintacticoCADENA          = 79
	sintacticoID              = 80
	sintacticoCOMENTARIO      = 81
	sintacticoBLANCOS         = 82
)

// sintactico rules.
const (
	sintacticoRULE_start         = 0
	sintacticoRULE_funciones     = 1
	sintacticoRULE_funcion       = 2
	sintacticoRULE_funcMain      = 3
	sintacticoRULE_instrucciones = 4
	sintacticoRULE_instruccion   = 5
	sintacticoRULE_imprimir      = 6
	sintacticoRULE_lista_exp     = 7
	sintacticoRULE_declaracion   = 8
	sintacticoRULE_tipo_dato     = 9
	sintacticoRULE_exp           = 10
	sintacticoRULE_logica        = 11
	sintacticoRULE_relacional    = 12
	sintacticoRULE_aritmetica    = 13
	sintacticoRULE_exp_valor     = 14
	sintacticoRULE_primitivo     = 15
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funciones returns the _funciones rule contexts.
	Get_funciones() IFuncionesContext

	// Set_funciones sets the _funciones rule contexts.
	Set_funciones(IFuncionesContext)

	// GetRoot returns the root attribute.
	GetRoot() ast.Ast

	// SetRoot sets the root attribute.
	SetRoot(ast.Ast)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	root       ast.Ast
	_funciones IFuncionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_funciones() IFuncionesContext { return s._funciones }

func (s *StartContext) Set_funciones(v IFuncionesContext) { s._funciones = v }

func (s *StartContext) GetRoot() ast.Ast { return s.root }

func (s *StartContext) SetRoot(v ast.Ast) { s.root = v }

func (s *StartContext) Funciones() IFuncionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *sintactico) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, sintacticoRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)

		var _x = p.Funciones()

		localctx.(*StartContext)._funciones = _x
	}

	localctx.(*StartContext).root = ast.NewAst(localctx.(*StartContext).Get_funciones().GetLista())

	return localctx
}

// IFuncionesContext is an interface to support dynamic dispatch.
type IFuncionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funcion returns the _funcion rule contexts.
	Get_funcion() IFuncionContext

	// Set_funcion sets the _funcion rule contexts.
	Set_funcion(IFuncionContext)

	// GetFun returns the fun rule context list.
	GetFun() []IFuncionContext

	// SetFun sets the fun rule context list.
	SetFun([]IFuncionContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsFuncionesContext differentiates from other interfaces.
	IsFuncionesContext()
}

type FuncionesContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	lista    *arrayList.List
	_funcion IFuncionContext
	fun      []IFuncionContext
}

func NewEmptyFuncionesContext() *FuncionesContext {
	var p = new(FuncionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_funciones
	return p
}

func (*FuncionesContext) IsFuncionesContext() {}

func NewFuncionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionesContext {
	var p = new(FuncionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_funciones

	return p
}

func (s *FuncionesContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionesContext) Get_funcion() IFuncionContext { return s._funcion }

func (s *FuncionesContext) Set_funcion(v IFuncionContext) { s._funcion = v }

func (s *FuncionesContext) GetFun() []IFuncionContext { return s.fun }

func (s *FuncionesContext) SetFun(v []IFuncionContext) { s.fun = v }

func (s *FuncionesContext) GetLista() *arrayList.List { return s.lista }

func (s *FuncionesContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *FuncionesContext) AllFuncion() []IFuncionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncionContext)(nil)).Elem())
	var tst = make([]IFuncionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncionContext)
		}
	}

	return tst
}

func (s *FuncionesContext) Funcion(i int) IFuncionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncionContext)
}

func (s *FuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterFunciones(s)
	}
}

func (s *FuncionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitFunciones(s)
	}
}

func (p *sintactico) Funciones() (localctx IFuncionesContext) {
	this := p
	_ = this

	localctx = NewFuncionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sintacticoRULE_funciones)
	localctx.(*FuncionesContext).lista = arrayList.New()
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sintacticoR_FN {
		{
			p.SetState(35)

			var _x = p.Funcion()

			localctx.(*FuncionesContext)._funcion = _x
		}
		localctx.(*FuncionesContext).fun = append(localctx.(*FuncionesContext).fun, localctx.(*FuncionesContext)._funcion)

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	LISTA := localctx.(*FuncionesContext).GetFun()
	for _, i := range LISTA {
		localctx.(*FuncionesContext).lista.Add(i.GetInstr())
	}

	return localctx
}

// IFuncionContext is an interface to support dynamic dispatch.
type IFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funcMain returns the _funcMain rule contexts.
	Get_funcMain() IFuncMainContext

	// Set_funcMain sets the _funcMain rule contexts.
	Set_funcMain(IFuncMainContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsFuncionContext differentiates from other interfaces.
	IsFuncionContext()
}

type FuncionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	instr     interfaces.Instruccion
	_funcMain IFuncMainContext
}

func NewEmptyFuncionContext() *FuncionContext {
	var p = new(FuncionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_funcion
	return p
}

func (*FuncionContext) IsFuncionContext() {}

func NewFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionContext {
	var p = new(FuncionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_funcion

	return p
}

func (s *FuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionContext) Get_funcMain() IFuncMainContext { return s._funcMain }

func (s *FuncionContext) Set_funcMain(v IFuncMainContext) { s._funcMain = v }

func (s *FuncionContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *FuncionContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *FuncionContext) FuncMain() IFuncMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncMainContext)
}

func (s *FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterFuncion(s)
	}
}

func (s *FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitFuncion(s)
	}
}

func (p *sintactico) Funcion() (localctx IFuncionContext) {
	this := p
	_ = this

	localctx = NewFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sintacticoRULE_funcion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)

		var _x = p.FuncMain()

		localctx.(*FuncionContext)._funcMain = _x
	}
	localctx.(*FuncionContext).instr = localctx.(*FuncionContext).Get_funcMain().GetInstr()

	return localctx
}

// IFuncMainContext is an interface to support dynamic dispatch.
type IFuncMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsFuncMainContext differentiates from other interfaces.
	IsFuncMainContext()
}

type FuncMainContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruccion
	_instrucciones IInstruccionesContext
}

func NewEmptyFuncMainContext() *FuncMainContext {
	var p = new(FuncMainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_funcMain
	return p
}

func (*FuncMainContext) IsFuncMainContext() {}

func NewFuncMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncMainContext {
	var p = new(FuncMainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_funcMain

	return p
}

func (s *FuncMainContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncMainContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *FuncMainContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *FuncMainContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *FuncMainContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *FuncMainContext) R_FN() antlr.TerminalNode {
	return s.GetToken(sintacticoR_FN, 0)
}

func (s *FuncMainContext) R_MAIN() antlr.TerminalNode {
	return s.GetToken(sintacticoR_MAIN, 0)
}

func (s *FuncMainContext) S_APAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_APAR, 0)
}

func (s *FuncMainContext) S_CPAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CPAR, 0)
}

func (s *FuncMainContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *FuncMainContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *FuncMainContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *FuncMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncMainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterFuncMain(s)
	}
}

func (s *FuncMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitFuncMain(s)
	}
}

func (p *sintactico) FuncMain() (localctx IFuncMainContext) {
	this := p
	_ = this

	localctx = NewFuncMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sintacticoRULE_funcMain)
	params := arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(sintacticoR_FN)
	}
	{
		p.SetState(47)
		p.Match(sintacticoR_MAIN)
	}
	{
		p.SetState(48)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(49)
		p.Match(sintacticoS_CPAR)
	}
	{
		p.SetState(50)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(51)

		var _x = p.Instrucciones()

		localctx.(*FuncMainContext)._instrucciones = _x
	}
	{
		p.SetState(52)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*FuncMainContext).instr = funcion.NewFuncion(entorno.VOID, "main", params, localctx.(*FuncMainContext).Get_instrucciones().GetLista())

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstruccionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstruccionContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        *arrayList.List
	_instruccion IInstruccionContext
	ins          []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetIns() []IInstruccionContext { return s.ins }

func (s *InstruccionesContext) SetIns(v []IInstruccionContext) { s.ins = v }

func (s *InstruccionesContext) GetLista() *arrayList.List { return s.lista }

func (s *InstruccionesContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *sintactico) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, sintacticoRULE_instrucciones)
	localctx.(*InstruccionesContext).lista = arrayList.New()
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == sintacticoR_LET || _la == sintacticoR_PRINTLN {
		{
			p.SetState(55)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).ins = append(localctx.(*InstruccionesContext).ins, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	LISTA := localctx.(*InstruccionesContext).GetIns()
	for _, i := range LISTA {
		localctx.(*InstruccionesContext).lista.Add(i.GetInstr())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_imprimir returns the _imprimir rule contexts.
	Get_imprimir() IImprimirContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruccion
	_imprimir    IImprimirContext
	_declaracion IDeclaracionContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_imprimir() IImprimirContext { return s._imprimir }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *InstruccionContext) Imprimir() IImprimirContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImprimirContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImprimirContext)
}

func (s *InstruccionContext) S_PTCOMA() antlr.TerminalNode {
	return s.GetToken(sintacticoS_PTCOMA, 0)
}

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *sintactico) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, sintacticoRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)

			var _x = p.Imprimir()

			localctx.(*InstruccionContext)._imprimir = _x
		}
		{
			p.SetState(63)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_imprimir().GetInstr()

	case sintacticoR_LET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(67)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracion().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImprimirContext is an interface to support dynamic dispatch.
type IImprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_lista_exp returns the _lista_exp rule contexts.
	Get_lista_exp() ILista_expContext

	// Set_lista_exp sets the _lista_exp rule contexts.
	Set_lista_exp(ILista_expContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsImprimirContext differentiates from other interfaces.
	IsImprimirContext()
}

type ImprimirContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	instr      interfaces.Instruccion
	_lista_exp ILista_expContext
}

func NewEmptyImprimirContext() *ImprimirContext {
	var p = new(ImprimirContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_imprimir
	return p
}

func (*ImprimirContext) IsImprimirContext() {}

func NewImprimirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImprimirContext {
	var p = new(ImprimirContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_imprimir

	return p
}

func (s *ImprimirContext) GetParser() antlr.Parser { return s.parser }

func (s *ImprimirContext) Get_lista_exp() ILista_expContext { return s._lista_exp }

func (s *ImprimirContext) Set_lista_exp(v ILista_expContext) { s._lista_exp = v }

func (s *ImprimirContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ImprimirContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ImprimirContext) R_PRINTLN() antlr.TerminalNode {
	return s.GetToken(sintacticoR_PRINTLN, 0)
}

func (s *ImprimirContext) S_APAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_APAR, 0)
}

func (s *ImprimirContext) Lista_exp() ILista_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILista_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILista_expContext)
}

func (s *ImprimirContext) S_CPAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CPAR, 0)
}

func (s *ImprimirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImprimirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImprimirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterImprimir(s)
	}
}

func (s *ImprimirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitImprimir(s)
	}
}

func (p *sintactico) Imprimir() (localctx IImprimirContext) {
	this := p
	_ = this

	localctx = NewImprimirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, sintacticoRULE_imprimir)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(sintacticoR_PRINTLN)
	}
	{
		p.SetState(73)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(74)

		var _x = p.lista_exp(0)

		localctx.(*ImprimirContext)._lista_exp = _x
	}
	{
		p.SetState(75)
		p.Match(sintacticoS_CPAR)
	}

	localctx.(*ImprimirContext).instr = imprimir.NewImprimir(localctx.(*ImprimirContext).Get_lista_exp().GetLista())

	return localctx
}

// ILista_expContext is an interface to support dynamic dispatch.
type ILista_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLISTA returns the LISTA rule contexts.
	GetLISTA() ILista_expContext

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// SetLISTA sets the LISTA rule contexts.
	SetLISTA(ILista_expContext)

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsLista_expContext differentiates from other interfaces.
	IsLista_expContext()
}

type Lista_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  *arrayList.List
	LISTA  ILista_expContext
	_exp   IExpContext
}

func NewEmptyLista_expContext() *Lista_expContext {
	var p = new(Lista_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_lista_exp
	return p
}

func (*Lista_expContext) IsLista_expContext() {}

func NewLista_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_expContext {
	var p = new(Lista_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_lista_exp

	return p
}

func (s *Lista_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_expContext) GetLISTA() ILista_expContext { return s.LISTA }

func (s *Lista_expContext) Get_exp() IExpContext { return s._exp }

func (s *Lista_expContext) SetLISTA(v ILista_expContext) { s.LISTA = v }

func (s *Lista_expContext) Set_exp(v IExpContext) { s._exp = v }

func (s *Lista_expContext) GetLista() *arrayList.List { return s.lista }

func (s *Lista_expContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *Lista_expContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Lista_expContext) S_COMA() antlr.TerminalNode {
	return s.GetToken(sintacticoS_COMA, 0)
}

func (s *Lista_expContext) Lista_exp() ILista_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILista_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILista_expContext)
}

func (s *Lista_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterLista_exp(s)
	}
}

func (s *Lista_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitLista_exp(s)
	}
}

func (p *sintactico) Lista_exp() (localctx ILista_expContext) {
	return p.lista_exp(0)
}

func (p *sintactico) lista_exp(_p int) (localctx ILista_expContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLista_expContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILista_expContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, sintacticoRULE_lista_exp, _p)
	localctx.(*Lista_expContext).lista = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)

		var _x = p.Exp()

		localctx.(*Lista_expContext)._exp = _x
	}
	localctx.(*Lista_expContext).lista.Add(localctx.(*Lista_expContext).Get_exp().GetVal())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLista_expContext(p, _parentctx, _parentState)
			localctx.(*Lista_expContext).LISTA = _prevctx
			p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_lista_exp)
			p.SetState(82)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(83)
				p.Match(sintacticoS_COMA)
			}
			{
				p.SetState(84)

				var _x = p.Exp()

				localctx.(*Lista_expContext)._exp = _x
			}

			localctx.(*Lista_expContext).GetLISTA().GetLista().Add(localctx.(*Lista_expContext).Get_exp().GetVal())
			localctx.(*Lista_expContext).lista = localctx.(*Lista_expContext).GetLISTA().GetLista()

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_tipo_dato returns the _tipo_dato rule contexts.
	Get_tipo_dato() ITipo_datoContext

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// Set_tipo_dato sets the _tipo_dato rule contexts.
	Set_tipo_dato(ITipo_datoContext)

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	instr      interfaces.Instruccion
	_ID        antlr.Token
	_tipo_dato ITipo_datoContext
	_exp       IExpContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclaracionContext) Get_tipo_dato() ITipo_datoContext { return s._tipo_dato }

func (s *DeclaracionContext) Get_exp() IExpContext { return s._exp }

func (s *DeclaracionContext) Set_tipo_dato(v ITipo_datoContext) { s._tipo_dato = v }

func (s *DeclaracionContext) Set_exp(v IExpContext) { s._exp = v }

func (s *DeclaracionContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *DeclaracionContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *DeclaracionContext) R_LET() antlr.TerminalNode {
	return s.GetToken(sintacticoR_LET, 0)
}

func (s *DeclaracionContext) R_MUT() antlr.TerminalNode {
	return s.GetToken(sintacticoR_MUT, 0)
}

func (s *DeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(sintacticoID, 0)
}

func (s *DeclaracionContext) S_DOSPT() antlr.TerminalNode {
	return s.GetToken(sintacticoS_DOSPT, 0)
}

func (s *DeclaracionContext) Tipo_dato() ITipo_datoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_datoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_datoContext)
}

func (s *DeclaracionContext) S_ASIGNAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ASIGNAR, 0)
}

func (s *DeclaracionContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *sintactico) Declaracion() (localctx IDeclaracionContext) {
	this := p
	_ = this

	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, sintacticoRULE_declaracion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(93)
			p.Match(sintacticoR_MUT)
		}
		{
			p.SetState(94)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(95)
			p.Match(sintacticoS_DOSPT)
		}
		{
			p.SetState(96)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(97)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(98)

			var _x = p.Exp()

			localctx.(*DeclaracionContext)._exp = _x
		}

		id := expresion.NewIdentificador((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()))
		localctx.(*DeclaracionContext).instr = variables.NewDeclaracion(true, id, localctx.(*DeclaracionContext).Get_tipo_dato().GetTipo(), localctx.(*DeclaracionContext).Get_exp().GetVal())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(102)
			p.Match(sintacticoR_MUT)
		}
		{
			p.SetState(103)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(104)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(105)

			var _x = p.Exp()

			localctx.(*DeclaracionContext)._exp = _x
		}

		id := expresion.NewIdentificador((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()))
		localctx.(*DeclaracionContext).instr = variables.NewDeclaracion(true, id, entorno.NULL, localctx.(*DeclaracionContext).Get_exp().GetVal())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(109)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(110)
			p.Match(sintacticoS_DOSPT)
		}
		{
			p.SetState(111)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(112)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(113)

			var _x = p.Exp()

			localctx.(*DeclaracionContext)._exp = _x
		}

		id := expresion.NewIdentificador((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()))
		localctx.(*DeclaracionContext).instr = variables.NewDeclaracion(false, id, localctx.(*DeclaracionContext).Get_tipo_dato().GetTipo(), localctx.(*DeclaracionContext).Get_exp().GetVal())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(117)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(118)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(119)

			var _x = p.Exp()

			localctx.(*DeclaracionContext)._exp = _x
		}

		id := expresion.NewIdentificador((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()))
		localctx.(*DeclaracionContext).instr = variables.NewDeclaracion(false, id, entorno.NULL, localctx.(*DeclaracionContext).Get_exp().GetVal())

	}

	return localctx
}

// ITipo_datoContext is an interface to support dynamic dispatch.
type ITipo_datoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo returns the tipo attribute.
	GetTipo() entorno.TipoDato

	// SetTipo sets the tipo attribute.
	SetTipo(entorno.TipoDato)

	// IsTipo_datoContext differentiates from other interfaces.
	IsTipo_datoContext()
}

type Tipo_datoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tipo   entorno.TipoDato
}

func NewEmptyTipo_datoContext() *Tipo_datoContext {
	var p = new(Tipo_datoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_tipo_dato
	return p
}

func (*Tipo_datoContext) IsTipo_datoContext() {}

func NewTipo_datoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_datoContext {
	var p = new(Tipo_datoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_tipo_dato

	return p
}

func (s *Tipo_datoContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_datoContext) GetTipo() entorno.TipoDato { return s.tipo }

func (s *Tipo_datoContext) SetTipo(v entorno.TipoDato) { s.tipo = v }

func (s *Tipo_datoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(sintacticoR_INT, 0)
}

func (s *Tipo_datoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(sintacticoR_FLOAT, 0)
}

func (s *Tipo_datoContext) R_CHAR() antlr.TerminalNode {
	return s.GetToken(sintacticoR_CHAR, 0)
}

func (s *Tipo_datoContext) R_STRING() antlr.TerminalNode {
	return s.GetToken(sintacticoR_STRING, 0)
}

func (s *Tipo_datoContext) R_BOOL() antlr.TerminalNode {
	return s.GetToken(sintacticoR_BOOL, 0)
}

func (s *Tipo_datoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_datoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_datoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterTipo_dato(s)
	}
}

func (s *Tipo_datoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitTipo_dato(s)
	}
}

func (p *sintactico) Tipo_dato() (localctx ITipo_datoContext) {
	this := p
	_ = this

	localctx = NewTipo_datoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, sintacticoRULE_tipo_dato)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(sintacticoR_INT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.INTEGER

	case sintacticoR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Match(sintacticoR_FLOAT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.FLOAT

	case sintacticoR_CHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Match(sintacticoR_CHAR)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.CHAR

	case sintacticoR_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.Match(sintacticoR_STRING)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.STRING

	case sintacticoR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Match(sintacticoR_BOOL)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.BOOLEAN

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_logica returns the _logica rule contexts.
	Get_logica() ILogicaContext

	// Get_relacional returns the _relacional rule contexts.
	Get_relacional() IRelacionalContext

	// Get_aritmetica returns the _aritmetica rule contexts.
	Get_aritmetica() IAritmeticaContext

	// Set_logica sets the _logica rule contexts.
	Set_logica(ILogicaContext)

	// Set_relacional sets the _relacional rule contexts.
	Set_relacional(IRelacionalContext)

	// Set_aritmetica sets the _aritmetica rule contexts.
	Set_aritmetica(IAritmeticaContext)

	// GetVal returns the val attribute.
	GetVal() interfaces.Expresion

	// SetVal sets the val attribute.
	SetVal(interfaces.Expresion)

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	val         interfaces.Expresion
	_logica     ILogicaContext
	_relacional IRelacionalContext
	_aritmetica IAritmeticaContext
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Get_logica() ILogicaContext { return s._logica }

func (s *ExpContext) Get_relacional() IRelacionalContext { return s._relacional }

func (s *ExpContext) Get_aritmetica() IAritmeticaContext { return s._aritmetica }

func (s *ExpContext) Set_logica(v ILogicaContext) { s._logica = v }

func (s *ExpContext) Set_relacional(v IRelacionalContext) { s._relacional = v }

func (s *ExpContext) Set_aritmetica(v IAritmeticaContext) { s._aritmetica = v }

func (s *ExpContext) GetVal() interfaces.Expresion { return s.val }

func (s *ExpContext) SetVal(v interfaces.Expresion) { s.val = v }

func (s *ExpContext) Logica() ILogicaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicaContext)
}

func (s *ExpContext) Relacional() IRelacionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelacionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelacionalContext)
}

func (s *ExpContext) Aritmetica() IAritmeticaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAritmeticaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAritmeticaContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *sintactico) Exp() (localctx IExpContext) {
	this := p
	_ = this

	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, sintacticoRULE_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)

			var _x = p.logica(0)

			localctx.(*ExpContext)._logica = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_logica().GetVal()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)

			var _x = p.relacional(0)

			localctx.(*ExpContext)._relacional = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_relacional().GetVal()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(142)

			var _x = p.aritmetica(0)

			localctx.(*ExpContext)._aritmetica = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_aritmetica().GetVal()

	}

	return localctx
}

// ILogicaContext is an interface to support dynamic dispatch.
type ILogicaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetIzq returns the izq rule contexts.
	GetIzq() ILogicaContext

	// Get_relacional returns the _relacional rule contexts.
	Get_relacional() IRelacionalContext

	// GetDer returns the der rule contexts.
	GetDer() ILogicaContext

	// SetIzq sets the izq rule contexts.
	SetIzq(ILogicaContext)

	// Set_relacional sets the _relacional rule contexts.
	Set_relacional(IRelacionalContext)

	// SetDer sets the der rule contexts.
	SetDer(ILogicaContext)

	// GetVal returns the val attribute.
	GetVal() interfaces.Expresion

	// SetVal sets the val attribute.
	SetVal(interfaces.Expresion)

	// IsLogicaContext differentiates from other interfaces.
	IsLogicaContext()
}

type LogicaContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	val         interfaces.Expresion
	izq         ILogicaContext
	_relacional IRelacionalContext
	op          antlr.Token
	der         ILogicaContext
}

func NewEmptyLogicaContext() *LogicaContext {
	var p = new(LogicaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_logica
	return p
}

func (*LogicaContext) IsLogicaContext() {}

func NewLogicaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicaContext {
	var p = new(LogicaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_logica

	return p
}

func (s *LogicaContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicaContext) GetOp() antlr.Token { return s.op }

func (s *LogicaContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicaContext) GetIzq() ILogicaContext { return s.izq }

func (s *LogicaContext) Get_relacional() IRelacionalContext { return s._relacional }

func (s *LogicaContext) GetDer() ILogicaContext { return s.der }

func (s *LogicaContext) SetIzq(v ILogicaContext) { s.izq = v }

func (s *LogicaContext) Set_relacional(v IRelacionalContext) { s._relacional = v }

func (s *LogicaContext) SetDer(v ILogicaContext) { s.der = v }

func (s *LogicaContext) GetVal() interfaces.Expresion { return s.val }

func (s *LogicaContext) SetVal(v interfaces.Expresion) { s.val = v }

func (s *LogicaContext) Relacional() IRelacionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelacionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelacionalContext)
}

func (s *LogicaContext) AllLogica() []ILogicaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILogicaContext)(nil)).Elem())
	var tst = make([]ILogicaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILogicaContext)
		}
	}

	return tst
}

func (s *LogicaContext) Logica(i int) ILogicaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILogicaContext)
}

func (s *LogicaContext) S_AND() antlr.TerminalNode {
	return s.GetToken(sintacticoS_AND, 0)
}

func (s *LogicaContext) S_OR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_OR, 0)
}

func (s *LogicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterLogica(s)
	}
}

func (s *LogicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitLogica(s)
	}
}

func (p *sintactico) Logica() (localctx ILogicaContext) {
	return p.logica(0)
}

func (p *sintactico) logica(_p int) (localctx ILogicaContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLogicaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, sintacticoRULE_logica, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)

		var _x = p.relacional(0)

		localctx.(*LogicaContext)._relacional = _x
	}
	localctx.(*LogicaContext).val = localctx.(*LogicaContext).Get_relacional().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogicaContext(p, _parentctx, _parentState)
				localctx.(*LogicaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_logica)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(152)

					var _m = p.Match(sintacticoS_AND)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(153)

					var _x = p.logica(4)

					localctx.(*LogicaContext).der = _x
				}

				localctx.(*LogicaContext).val = expresion.NewOperacion(localctx.(*LogicaContext).GetIzq().GetVal(), localctx.(*LogicaContext).GetDer().GetVal(), (func() string {
					if localctx.(*LogicaContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*LogicaContext).GetOp().GetText()
					}
				}()), false)

			case 2:
				localctx = NewLogicaContext(p, _parentctx, _parentState)
				localctx.(*LogicaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_logica)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(157)

					var _m = p.Match(sintacticoS_OR)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(158)

					var _x = p.logica(3)

					localctx.(*LogicaContext).der = _x
				}

				localctx.(*LogicaContext).val = expresion.NewOperacion(localctx.(*LogicaContext).GetIzq().GetVal(), localctx.(*LogicaContext).GetDer().GetVal(), (func() string {
					if localctx.(*LogicaContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*LogicaContext).GetOp().GetText()
					}
				}()), false)

			}

		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IRelacionalContext is an interface to support dynamic dispatch.
type IRelacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetIzq returns the izq rule contexts.
	GetIzq() IRelacionalContext

	// Get_aritmetica returns the _aritmetica rule contexts.
	Get_aritmetica() IAritmeticaContext

	// GetDer returns the der rule contexts.
	GetDer() IRelacionalContext

	// SetIzq sets the izq rule contexts.
	SetIzq(IRelacionalContext)

	// Set_aritmetica sets the _aritmetica rule contexts.
	Set_aritmetica(IAritmeticaContext)

	// SetDer sets the der rule contexts.
	SetDer(IRelacionalContext)

	// GetVal returns the val attribute.
	GetVal() interfaces.Expresion

	// SetVal sets the val attribute.
	SetVal(interfaces.Expresion)

	// IsRelacionalContext differentiates from other interfaces.
	IsRelacionalContext()
}

type RelacionalContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	val         interfaces.Expresion
	izq         IRelacionalContext
	_aritmetica IAritmeticaContext
	op          antlr.Token
	der         IRelacionalContext
}

func NewEmptyRelacionalContext() *RelacionalContext {
	var p = new(RelacionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_relacional
	return p
}

func (*RelacionalContext) IsRelacionalContext() {}

func NewRelacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelacionalContext {
	var p = new(RelacionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_relacional

	return p
}

func (s *RelacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *RelacionalContext) GetOp() antlr.Token { return s.op }

func (s *RelacionalContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelacionalContext) GetIzq() IRelacionalContext { return s.izq }

func (s *RelacionalContext) Get_aritmetica() IAritmeticaContext { return s._aritmetica }

func (s *RelacionalContext) GetDer() IRelacionalContext { return s.der }

func (s *RelacionalContext) SetIzq(v IRelacionalContext) { s.izq = v }

func (s *RelacionalContext) Set_aritmetica(v IAritmeticaContext) { s._aritmetica = v }

func (s *RelacionalContext) SetDer(v IRelacionalContext) { s.der = v }

func (s *RelacionalContext) GetVal() interfaces.Expresion { return s.val }

func (s *RelacionalContext) SetVal(v interfaces.Expresion) { s.val = v }

func (s *RelacionalContext) Aritmetica() IAritmeticaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAritmeticaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAritmeticaContext)
}

func (s *RelacionalContext) AllRelacional() []IRelacionalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelacionalContext)(nil)).Elem())
	var tst = make([]IRelacionalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelacionalContext)
		}
	}

	return tst
}

func (s *RelacionalContext) Relacional(i int) IRelacionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelacionalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelacionalContext)
}

func (s *RelacionalContext) S_MENOR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MENOR, 0)
}

func (s *RelacionalContext) S_MAYOR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MAYOR, 0)
}

func (s *RelacionalContext) S_MENORI() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MENORI, 0)
}

func (s *RelacionalContext) S_MAYORI() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MAYORI, 0)
}

func (s *RelacionalContext) S_IGUAL() antlr.TerminalNode {
	return s.GetToken(sintacticoS_IGUAL, 0)
}

func (s *RelacionalContext) S_DIFERENTE() antlr.TerminalNode {
	return s.GetToken(sintacticoS_DIFERENTE, 0)
}

func (s *RelacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterRelacional(s)
	}
}

func (s *RelacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitRelacional(s)
	}
}

func (p *sintactico) Relacional() (localctx IRelacionalContext) {
	return p.relacional(0)
}

func (p *sintactico) relacional(_p int) (localctx IRelacionalContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelacionalContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelacionalContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, sintacticoRULE_relacional, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)

		var _x = p.aritmetica(0)

		localctx.(*RelacionalContext)._aritmetica = _x
	}
	localctx.(*RelacionalContext).val = localctx.(*RelacionalContext).Get_aritmetica().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelacionalContext(p, _parentctx, _parentState)
			localctx.(*RelacionalContext).izq = _prevctx
			p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_relacional)
			p.SetState(170)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(171)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*RelacionalContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<sintacticoS_MAYOR)|(1<<sintacticoS_MENOR)|(1<<sintacticoS_MAYORI)|(1<<sintacticoS_MENORI)|(1<<sintacticoS_IGUAL)|(1<<sintacticoS_DIFERENTE))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*RelacionalContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(172)

				var _x = p.relacional(3)

				localctx.(*RelacionalContext).der = _x
			}

			localctx.(*RelacionalContext).val = expresion.NewOperacion(localctx.(*RelacionalContext).GetIzq().GetVal(), localctx.(*RelacionalContext).GetDer().GetVal(), (func() string {
				if localctx.(*RelacionalContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*RelacionalContext).GetOp().GetText()
				}
			}()), false)

		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IAritmeticaContext is an interface to support dynamic dispatch.
type IAritmeticaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetIzq returns the izq rule contexts.
	GetIzq() IAritmeticaContext

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// Get_exp_valor returns the _exp_valor rule contexts.
	Get_exp_valor() IExp_valorContext

	// GetDer returns the der rule contexts.
	GetDer() IAritmeticaContext

	// SetIzq sets the izq rule contexts.
	SetIzq(IAritmeticaContext)

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// Set_exp_valor sets the _exp_valor rule contexts.
	Set_exp_valor(IExp_valorContext)

	// SetDer sets the der rule contexts.
	SetDer(IAritmeticaContext)

	// GetVal returns the val attribute.
	GetVal() interfaces.Expresion

	// SetVal sets the val attribute.
	SetVal(interfaces.Expresion)

	// IsAritmeticaContext differentiates from other interfaces.
	IsAritmeticaContext()
}

type AritmeticaContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	val        interfaces.Expresion
	izq        IAritmeticaContext
	_exp       IExpContext
	_exp_valor IExp_valorContext
	op         antlr.Token
	der        IAritmeticaContext
}

func NewEmptyAritmeticaContext() *AritmeticaContext {
	var p = new(AritmeticaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_aritmetica
	return p
}

func (*AritmeticaContext) IsAritmeticaContext() {}

func NewAritmeticaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AritmeticaContext {
	var p = new(AritmeticaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_aritmetica

	return p
}

func (s *AritmeticaContext) GetParser() antlr.Parser { return s.parser }

func (s *AritmeticaContext) GetOp() antlr.Token { return s.op }

func (s *AritmeticaContext) SetOp(v antlr.Token) { s.op = v }

func (s *AritmeticaContext) GetIzq() IAritmeticaContext { return s.izq }

func (s *AritmeticaContext) Get_exp() IExpContext { return s._exp }

func (s *AritmeticaContext) Get_exp_valor() IExp_valorContext { return s._exp_valor }

func (s *AritmeticaContext) GetDer() IAritmeticaContext { return s.der }

func (s *AritmeticaContext) SetIzq(v IAritmeticaContext) { s.izq = v }

func (s *AritmeticaContext) Set_exp(v IExpContext) { s._exp = v }

func (s *AritmeticaContext) Set_exp_valor(v IExp_valorContext) { s._exp_valor = v }

func (s *AritmeticaContext) SetDer(v IAritmeticaContext) { s.der = v }

func (s *AritmeticaContext) GetVal() interfaces.Expresion { return s.val }

func (s *AritmeticaContext) SetVal(v interfaces.Expresion) { s.val = v }

func (s *AritmeticaContext) S_RESTA() antlr.TerminalNode {
	return s.GetToken(sintacticoS_RESTA, 0)
}

func (s *AritmeticaContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AritmeticaContext) Exp_valor() IExp_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp_valorContext)
}

func (s *AritmeticaContext) S_APAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_APAR, 0)
}

func (s *AritmeticaContext) S_CPAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CPAR, 0)
}

func (s *AritmeticaContext) AllAritmetica() []IAritmeticaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAritmeticaContext)(nil)).Elem())
	var tst = make([]IAritmeticaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAritmeticaContext)
		}
	}

	return tst
}

func (s *AritmeticaContext) Aritmetica(i int) IAritmeticaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAritmeticaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAritmeticaContext)
}

func (s *AritmeticaContext) S_POR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_POR, 0)
}

func (s *AritmeticaContext) S_DIVISION() antlr.TerminalNode {
	return s.GetToken(sintacticoS_DIVISION, 0)
}

func (s *AritmeticaContext) S_MODULO() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MODULO, 0)
}

func (s *AritmeticaContext) S_SUMA() antlr.TerminalNode {
	return s.GetToken(sintacticoS_SUMA, 0)
}

func (s *AritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AritmeticaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterAritmetica(s)
	}
}

func (s *AritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitAritmetica(s)
	}
}

func (p *sintactico) Aritmetica() (localctx IAritmeticaContext) {
	return p.aritmetica(0)
}

func (p *sintactico) aritmetica(_p int) (localctx IAritmeticaContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAritmeticaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAritmeticaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, sintacticoRULE_aritmetica, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoS_RESTA:
		{
			p.SetState(181)
			p.Match(sintacticoS_RESTA)
		}
		{
			p.SetState(182)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		localctx.(*AritmeticaContext).val = expresion.NewOperacion(localctx.(*AritmeticaContext).Get_exp().GetVal(), nil, "-", true)

	case sintacticoR_TRUE, sintacticoR_FALSE, sintacticoNUMERO, sintacticoDECIMAL, sintacticoCADENA, sintacticoID:
		{
			p.SetState(185)

			var _x = p.Exp_valor()

			localctx.(*AritmeticaContext)._exp_valor = _x
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp_valor().GetVal()

	case sintacticoS_APAR:
		{
			p.SetState(188)
			p.Match(sintacticoS_APAR)
		}
		{
			p.SetState(189)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		{
			p.SetState(190)
			p.Match(sintacticoS_CPAR)
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp().GetVal()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAritmeticaContext(p, _parentctx, _parentState)
				localctx.(*AritmeticaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_aritmetica)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(196)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<sintacticoS_POR)|(1<<sintacticoS_DIVISION)|(1<<sintacticoS_MODULO))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(197)

					var _x = p.aritmetica(5)

					localctx.(*AritmeticaContext).der = _x
				}

				localctx.(*AritmeticaContext).val = expresion.NewOperacion(localctx.(*AritmeticaContext).GetIzq().GetVal(), localctx.(*AritmeticaContext).GetDer().GetVal(), (func() string {
					if localctx.(*AritmeticaContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*AritmeticaContext).GetOp().GetText()
					}
				}()), false)

			case 2:
				localctx = NewAritmeticaContext(p, _parentctx, _parentState)
				localctx.(*AritmeticaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_aritmetica)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(201)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AritmeticaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == sintacticoS_SUMA || _la == sintacticoS_RESTA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AritmeticaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(202)

					var _x = p.aritmetica(4)

					localctx.(*AritmeticaContext).der = _x
				}

				localctx.(*AritmeticaContext).val = expresion.NewOperacion(localctx.(*AritmeticaContext).GetIzq().GetVal(), localctx.(*AritmeticaContext).GetDer().GetVal(), (func() string {
					if localctx.(*AritmeticaContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*AritmeticaContext).GetOp().GetText()
					}
				}()), false)

			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IExp_valorContext is an interface to support dynamic dispatch.
type IExp_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// GetVal returns the val attribute.
	GetVal() interfaces.Expresion

	// SetVal sets the val attribute.
	SetVal(interfaces.Expresion)

	// IsExp_valorContext differentiates from other interfaces.
	IsExp_valorContext()
}

type Exp_valorContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	val        interfaces.Expresion
	_primitivo IPrimitivoContext
}

func NewEmptyExp_valorContext() *Exp_valorContext {
	var p = new(Exp_valorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_exp_valor
	return p
}

func (*Exp_valorContext) IsExp_valorContext() {}

func NewExp_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_valorContext {
	var p = new(Exp_valorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_exp_valor

	return p
}

func (s *Exp_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_valorContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Exp_valorContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Exp_valorContext) GetVal() interfaces.Expresion { return s.val }

func (s *Exp_valorContext) SetVal(v interfaces.Expresion) { s.val = v }

func (s *Exp_valorContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Exp_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_valorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterExp_valor(s)
	}
}

func (s *Exp_valorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitExp_valor(s)
	}
}

func (p *sintactico) Exp_valor() (localctx IExp_valorContext) {
	this := p
	_ = this

	localctx = NewExp_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, sintacticoRULE_exp_valor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)

		var _x = p.Primitivo()

		localctx.(*Exp_valorContext)._primitivo = _x
	}
	localctx.(*Exp_valorContext).val = localctx.(*Exp_valorContext).Get_primitivo().GetVal()

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetVal returns the val attribute.
	GetVal() interfaces.Expresion

	// SetVal sets the val attribute.
	SetVal(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	val      interfaces.Expresion
	_NUMERO  antlr.Token
	_DECIMAL antlr.Token
	_CADENA  antlr.Token
	_ID      antlr.Token
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *PrimitivoContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *PrimitivoContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrimitivoContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitivoContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *PrimitivoContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *PrimitivoContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrimitivoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitivoContext) GetVal() interfaces.Expresion { return s.val }

func (s *PrimitivoContext) SetVal(v interfaces.Expresion) { s.val = v }

func (s *PrimitivoContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(sintacticoNUMERO, 0)
}

func (s *PrimitivoContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(sintacticoDECIMAL, 0)
}

func (s *PrimitivoContext) CADENA() antlr.TerminalNode {
	return s.GetToken(sintacticoCADENA, 0)
}

func (s *PrimitivoContext) R_TRUE() antlr.TerminalNode {
	return s.GetToken(sintacticoR_TRUE, 0)
}

func (s *PrimitivoContext) R_FALSE() antlr.TerminalNode {
	return s.GetToken(sintacticoR_FALSE, 0)
}

func (s *PrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(sintacticoID, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *sintactico) Primitivo() (localctx IPrimitivoContext) {
	this := p
	_ = this

	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, sintacticoRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(225)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoNUMERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)

			var _m = p.Match(sintacticoNUMERO)

			localctx.(*PrimitivoContext)._NUMERO = _m
		}

		val, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMERO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(val, entorno.INTEGER)

	case sintacticoDECIMAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)

			var _m = p.Match(sintacticoDECIMAL)

			localctx.(*PrimitivoContext)._DECIMAL = _m
		}

		val, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DECIMAL().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(val, entorno.FLOAT)

	case sintacticoCADENA:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(217)

			var _m = p.Match(sintacticoCADENA)

			localctx.(*PrimitivoContext)._CADENA = _m
		}

		val := (func() string {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(val, entorno.STRING)

	case sintacticoR_TRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(219)
			p.Match(sintacticoR_TRUE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(true, entorno.BOOLEAN)

	case sintacticoR_FALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(221)
			p.Match(sintacticoR_FALSE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(false, entorno.BOOLEAN)

	case sintacticoID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(223)

			var _m = p.Match(sintacticoID)

			localctx.(*PrimitivoContext)._ID = _m
		}

		localctx.(*PrimitivoContext).val = expresion.NewIdentificador((func() string {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *sintactico) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *Lista_expContext = nil
		if localctx != nil {
			t = localctx.(*Lista_expContext)
		}
		return p.Lista_exp_Sempred(t, predIndex)

	case 11:
		var t *LogicaContext = nil
		if localctx != nil {
			t = localctx.(*LogicaContext)
		}
		return p.Logica_Sempred(t, predIndex)

	case 12:
		var t *RelacionalContext = nil
		if localctx != nil {
			t = localctx.(*RelacionalContext)
		}
		return p.Relacional_Sempred(t, predIndex)

	case 13:
		var t *AritmeticaContext = nil
		if localctx != nil {
			t = localctx.(*AritmeticaContext)
		}
		return p.Aritmetica_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *sintactico) Lista_exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *sintactico) Logica_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *sintactico) Relacional_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *sintactico) Aritmetica_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
