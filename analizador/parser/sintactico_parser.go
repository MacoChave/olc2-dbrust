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
import "db_rust/analizador/ast/flujo"
import "db_rust/analizador/ast/interfaces"
import "db_rust/analizador/entorno"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 84, 298,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 3, 7,
	3, 47, 10, 3, 12, 3, 14, 3, 50, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 6, 6, 67, 10,
	6, 13, 6, 14, 6, 68, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 89, 10,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 7, 9, 106, 10, 9, 12, 9, 14, 9, 109, 11, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 141, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 153, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 190, 10, 13, 3, 14,
	6, 14, 193, 10, 14, 13, 14, 14, 14, 194, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 216, 10, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	7, 17, 232, 10, 17, 12, 17, 14, 17, 235, 11, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 246, 10, 18, 12, 18, 14,
	18, 249, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 264, 10, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 276, 10, 19,
	12, 19, 14, 19, 279, 11, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 296,
	10, 21, 3, 21, 2, 6, 16, 32, 34, 36, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 2, 5, 3, 2, 16, 21, 3, 2, 13,
	15, 3, 2, 11, 12, 2, 307, 2, 42, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 53,
	3, 2, 2, 2, 8, 56, 3, 2, 2, 2, 10, 66, 3, 2, 2, 2, 12, 88, 3, 2, 2, 2,
	14, 90, 3, 2, 2, 2, 16, 96, 3, 2, 2, 2, 18, 140, 3, 2, 2, 2, 20, 152, 3,
	2, 2, 2, 22, 154, 3, 2, 2, 2, 24, 189, 3, 2, 2, 2, 26, 192, 3, 2, 2, 2,
	28, 198, 3, 2, 2, 2, 30, 215, 3, 2, 2, 2, 32, 217, 3, 2, 2, 2, 34, 236,
	3, 2, 2, 2, 36, 263, 3, 2, 2, 2, 38, 280, 3, 2, 2, 2, 40, 295, 3, 2, 2,
	2, 42, 43, 5, 4, 3, 2, 43, 44, 8, 2, 1, 2, 44, 3, 3, 2, 2, 2, 45, 47, 5,
	6, 4, 2, 46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48,
	49, 3, 2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 52, 8, 3, 1,
	2, 52, 5, 3, 2, 2, 2, 53, 54, 5, 8, 5, 2, 54, 55, 8, 4, 1, 2, 55, 7, 3,
	2, 2, 2, 56, 57, 7, 44, 2, 2, 57, 58, 7, 45, 2, 2, 58, 59, 7, 25, 2, 2,
	59, 60, 7, 26, 2, 2, 60, 61, 7, 31, 2, 2, 61, 62, 5, 10, 6, 2, 62, 63,
	7, 32, 2, 2, 63, 64, 8, 5, 1, 2, 64, 9, 3, 2, 2, 2, 65, 67, 5, 12, 7, 2,
	66, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 71, 8, 6, 1, 2, 71, 11, 3, 2, 2, 2, 72,
	73, 5, 14, 8, 2, 73, 74, 7, 5, 2, 2, 74, 75, 8, 7, 1, 2, 75, 89, 3, 2,
	2, 2, 76, 77, 5, 18, 10, 2, 77, 78, 7, 5, 2, 2, 78, 79, 8, 7, 1, 2, 79,
	89, 3, 2, 2, 2, 80, 81, 5, 22, 12, 2, 81, 82, 7, 5, 2, 2, 82, 83, 8, 7,
	1, 2, 83, 89, 3, 2, 2, 2, 84, 85, 5, 24, 13, 2, 85, 86, 7, 5, 2, 2, 86,
	87, 8, 7, 1, 2, 87, 89, 3, 2, 2, 2, 88, 72, 3, 2, 2, 2, 88, 76, 3, 2, 2,
	2, 88, 80, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 89, 13, 3, 2, 2, 2, 90, 91,
	7, 51, 2, 2, 91, 92, 7, 25, 2, 2, 92, 93, 5, 16, 9, 2, 93, 94, 7, 26, 2,
	2, 94, 95, 8, 8, 1, 2, 95, 15, 3, 2, 2, 2, 96, 97, 8, 9, 1, 2, 97, 98,
	5, 30, 16, 2, 98, 99, 8, 9, 1, 2, 99, 107, 3, 2, 2, 2, 100, 101, 12, 4,
	2, 2, 101, 102, 7, 4, 2, 2, 102, 103, 5, 30, 16, 2, 103, 104, 8, 9, 1,
	2, 104, 106, 3, 2, 2, 2, 105, 100, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107,
	105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 17, 3, 2, 2, 2, 109, 107, 3,
	2, 2, 2, 110, 111, 7, 36, 2, 2, 111, 112, 7, 37, 2, 2, 112, 113, 7, 82,
	2, 2, 113, 114, 7, 27, 2, 2, 114, 115, 5, 20, 11, 2, 115, 116, 7, 6, 2,
	2, 116, 117, 5, 30, 16, 2, 117, 118, 8, 10, 1, 2, 118, 141, 3, 2, 2, 2,
	119, 120, 7, 36, 2, 2, 120, 121, 7, 37, 2, 2, 121, 122, 7, 82, 2, 2, 122,
	123, 7, 6, 2, 2, 123, 124, 5, 30, 16, 2, 124, 125, 8, 10, 1, 2, 125, 141,
	3, 2, 2, 2, 126, 127, 7, 36, 2, 2, 127, 128, 7, 82, 2, 2, 128, 129, 7,
	27, 2, 2, 129, 130, 5, 20, 11, 2, 130, 131, 7, 6, 2, 2, 131, 132, 5, 30,
	16, 2, 132, 133, 8, 10, 1, 2, 133, 141, 3, 2, 2, 2, 134, 135, 7, 36, 2,
	2, 135, 136, 7, 82, 2, 2, 136, 137, 7, 6, 2, 2, 137, 138, 5, 30, 16, 2,
	138, 139, 8, 10, 1, 2, 139, 141, 3, 2, 2, 2, 140, 110, 3, 2, 2, 2, 140,
	119, 3, 2, 2, 2, 140, 126, 3, 2, 2, 2, 140, 134, 3, 2, 2, 2, 141, 19, 3,
	2, 2, 2, 142, 143, 7, 38, 2, 2, 143, 153, 8, 11, 1, 2, 144, 145, 7, 39,
	2, 2, 145, 153, 8, 11, 1, 2, 146, 147, 7, 41, 2, 2, 147, 153, 8, 11, 1,
	2, 148, 149, 7, 42, 2, 2, 149, 153, 8, 11, 1, 2, 150, 151, 7, 40, 2, 2,
	151, 153, 8, 11, 1, 2, 152, 142, 3, 2, 2, 2, 152, 144, 3, 2, 2, 2, 152,
	146, 3, 2, 2, 2, 152, 148, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 153, 21, 3,
	2, 2, 2, 154, 155, 7, 82, 2, 2, 155, 156, 7, 6, 2, 2, 156, 157, 5, 30,
	16, 2, 157, 158, 8, 12, 1, 2, 158, 23, 3, 2, 2, 2, 159, 160, 7, 63, 2,
	2, 160, 161, 5, 30, 16, 2, 161, 162, 7, 31, 2, 2, 162, 163, 5, 10, 6, 2,
	163, 164, 7, 32, 2, 2, 164, 165, 8, 13, 1, 2, 165, 190, 3, 2, 2, 2, 166,
	167, 7, 63, 2, 2, 167, 168, 5, 30, 16, 2, 168, 169, 7, 31, 2, 2, 169, 170,
	5, 10, 6, 2, 170, 171, 7, 32, 2, 2, 171, 172, 7, 64, 2, 2, 172, 173, 7,
	31, 2, 2, 173, 174, 5, 10, 6, 2, 174, 175, 7, 32, 2, 2, 175, 176, 8, 13,
	1, 2, 176, 190, 3, 2, 2, 2, 177, 178, 7, 63, 2, 2, 178, 179, 5, 30, 16,
	2, 179, 180, 7, 31, 2, 2, 180, 181, 5, 10, 6, 2, 181, 182, 7, 32, 2, 2,
	182, 183, 5, 26, 14, 2, 183, 184, 7, 64, 2, 2, 184, 185, 7, 31, 2, 2, 185,
	186, 5, 10, 6, 2, 186, 187, 7, 32, 2, 2, 187, 188, 8, 13, 1, 2, 188, 190,
	3, 2, 2, 2, 189, 159, 3, 2, 2, 2, 189, 166, 3, 2, 2, 2, 189, 177, 3, 2,
	2, 2, 190, 25, 3, 2, 2, 2, 191, 193, 5, 28, 15, 2, 192, 191, 3, 2, 2, 2,
	193, 194, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	196, 3, 2, 2, 2, 196, 197, 8, 14, 1, 2, 197, 27, 3, 2, 2, 2, 198, 199,
	7, 64, 2, 2, 199, 200, 7, 63, 2, 2, 200, 201, 5, 30, 16, 2, 201, 202, 7,
	31, 2, 2, 202, 203, 5, 10, 6, 2, 203, 204, 7, 32, 2, 2, 204, 205, 8, 15,
	1, 2, 205, 29, 3, 2, 2, 2, 206, 207, 5, 32, 17, 2, 207, 208, 8, 16, 1,
	2, 208, 216, 3, 2, 2, 2, 209, 210, 5, 34, 18, 2, 210, 211, 8, 16, 1, 2,
	211, 216, 3, 2, 2, 2, 212, 213, 5, 36, 19, 2, 213, 214, 8, 16, 1, 2, 214,
	216, 3, 2, 2, 2, 215, 206, 3, 2, 2, 2, 215, 209, 3, 2, 2, 2, 215, 212,
	3, 2, 2, 2, 216, 31, 3, 2, 2, 2, 217, 218, 8, 17, 1, 2, 218, 219, 5, 34,
	18, 2, 219, 220, 8, 17, 1, 2, 220, 233, 3, 2, 2, 2, 221, 222, 12, 5, 2,
	2, 222, 223, 7, 23, 2, 2, 223, 224, 5, 32, 17, 6, 224, 225, 8, 17, 1, 2,
	225, 232, 3, 2, 2, 2, 226, 227, 12, 4, 2, 2, 227, 228, 7, 22, 2, 2, 228,
	229, 5, 32, 17, 5, 229, 230, 8, 17, 1, 2, 230, 232, 3, 2, 2, 2, 231, 221,
	3, 2, 2, 2, 231, 226, 3, 2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 231, 3, 2,
	2, 2, 233, 234, 3, 2, 2, 2, 234, 33, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2,
	236, 237, 8, 18, 1, 2, 237, 238, 5, 36, 19, 2, 238, 239, 8, 18, 1, 2, 239,
	247, 3, 2, 2, 2, 240, 241, 12, 4, 2, 2, 241, 242, 9, 2, 2, 2, 242, 243,
	5, 34, 18, 5, 243, 244, 8, 18, 1, 2, 244, 246, 3, 2, 2, 2, 245, 240, 3,
	2, 2, 2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2,
	2, 248, 35, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 251, 8, 19, 1, 2, 251,
	252, 7, 12, 2, 2, 252, 253, 5, 30, 16, 2, 253, 254, 8, 19, 1, 2, 254, 264,
	3, 2, 2, 2, 255, 256, 5, 38, 20, 2, 256, 257, 8, 19, 1, 2, 257, 264, 3,
	2, 2, 2, 258, 259, 7, 25, 2, 2, 259, 260, 5, 30, 16, 2, 260, 261, 7, 26,
	2, 2, 261, 262, 8, 19, 1, 2, 262, 264, 3, 2, 2, 2, 263, 250, 3, 2, 2, 2,
	263, 255, 3, 2, 2, 2, 263, 258, 3, 2, 2, 2, 264, 277, 3, 2, 2, 2, 265,
	266, 12, 6, 2, 2, 266, 267, 9, 3, 2, 2, 267, 268, 5, 36, 19, 7, 268, 269,
	8, 19, 1, 2, 269, 276, 3, 2, 2, 2, 270, 271, 12, 5, 2, 2, 271, 272, 9,
	4, 2, 2, 272, 273, 5, 36, 19, 6, 273, 274, 8, 19, 1, 2, 274, 276, 3, 2,
	2, 2, 275, 265, 3, 2, 2, 2, 275, 270, 3, 2, 2, 2, 276, 279, 3, 2, 2, 2,
	277, 275, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 37, 3, 2, 2, 2, 279, 277,
	3, 2, 2, 2, 280, 281, 5, 40, 21, 2, 281, 282, 8, 20, 1, 2, 282, 39, 3,
	2, 2, 2, 283, 284, 7, 78, 2, 2, 284, 296, 8, 21, 1, 2, 285, 286, 7, 79,
	2, 2, 286, 296, 8, 21, 1, 2, 287, 288, 7, 81, 2, 2, 288, 296, 8, 21, 1,
	2, 289, 290, 7, 34, 2, 2, 290, 296, 8, 21, 1, 2, 291, 292, 7, 35, 2, 2,
	292, 296, 8, 21, 1, 2, 293, 294, 7, 82, 2, 2, 294, 296, 8, 21, 1, 2, 295,
	283, 3, 2, 2, 2, 295, 285, 3, 2, 2, 2, 295, 287, 3, 2, 2, 2, 295, 289,
	3, 2, 2, 2, 295, 291, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 296, 41, 3, 2,
	2, 2, 18, 48, 68, 88, 107, 140, 152, 189, 194, 215, 231, 233, 247, 263,
	275, 277, 295,
}
var literalNames = []string{
	"", "'.'", "','", "';'", "'='", "'&'", "'|'", "'=>'", "'_'", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'||'",
	"'&&'", "'->'", "'('", "')'", "':'", "'::'", "'['", "']'", "'{'", "'}'",
	"'..'", "'true'", "'false'", "'let'", "'mut'", "'i64'", "'f64'", "'bool'",
	"'char'", "'String'", "'&str'", "'fn'", "'main'", "'as'", "'to_owned'",
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
	"start", "procedimientos", "procedimiento", "principal", "instrucciones",
	"instruccion", "imprimir", "lista_exp", "declaracion", "tipo_dato", "asignacion",
	"sent_if", "lista_elseif", "elseif", "exp", "logica", "relacional", "aritmetica",
	"exp_valor", "primitivo",
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
	sintacticoRULE_start          = 0
	sintacticoRULE_procedimientos = 1
	sintacticoRULE_procedimiento  = 2
	sintacticoRULE_principal      = 3
	sintacticoRULE_instrucciones  = 4
	sintacticoRULE_instruccion    = 5
	sintacticoRULE_imprimir       = 6
	sintacticoRULE_lista_exp      = 7
	sintacticoRULE_declaracion    = 8
	sintacticoRULE_tipo_dato      = 9
	sintacticoRULE_asignacion     = 10
	sintacticoRULE_sent_if        = 11
	sintacticoRULE_lista_elseif   = 12
	sintacticoRULE_elseif         = 13
	sintacticoRULE_exp            = 14
	sintacticoRULE_logica         = 15
	sintacticoRULE_relacional     = 16
	sintacticoRULE_aritmetica     = 17
	sintacticoRULE_exp_valor      = 18
	sintacticoRULE_primitivo      = 19
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_procedimientos returns the _procedimientos rule contexts.
	Get_procedimientos() IProcedimientosContext

	// Set_procedimientos sets the _procedimientos rule contexts.
	Set_procedimientos(IProcedimientosContext)

	// GetRoot returns the root attribute.
	GetRoot() ast.Ast

	// SetRoot sets the root attribute.
	SetRoot(ast.Ast)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	root            ast.Ast
	_procedimientos IProcedimientosContext
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

func (s *StartContext) Get_procedimientos() IProcedimientosContext { return s._procedimientos }

func (s *StartContext) Set_procedimientos(v IProcedimientosContext) { s._procedimientos = v }

func (s *StartContext) GetRoot() ast.Ast { return s.root }

func (s *StartContext) SetRoot(v ast.Ast) { s.root = v }

func (s *StartContext) Procedimientos() IProcedimientosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedimientosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedimientosContext)
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
		p.SetState(40)

		var _x = p.Procedimientos()

		localctx.(*StartContext)._procedimientos = _x
	}

	localctx.(*StartContext).root = ast.NewAst(localctx.(*StartContext).Get_procedimientos().GetLista())

	return localctx
}

// IProcedimientosContext is an interface to support dynamic dispatch.
type IProcedimientosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_procedimiento returns the _procedimiento rule contexts.
	Get_procedimiento() IProcedimientoContext

	// Set_procedimiento sets the _procedimiento rule contexts.
	Set_procedimiento(IProcedimientoContext)

	// GetProc returns the proc rule context list.
	GetProc() []IProcedimientoContext

	// SetProc sets the proc rule context list.
	SetProc([]IProcedimientoContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsProcedimientosContext differentiates from other interfaces.
	IsProcedimientosContext()
}

type ProcedimientosContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_procedimiento IProcedimientoContext
	proc           []IProcedimientoContext
}

func NewEmptyProcedimientosContext() *ProcedimientosContext {
	var p = new(ProcedimientosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_procedimientos
	return p
}

func (*ProcedimientosContext) IsProcedimientosContext() {}

func NewProcedimientosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedimientosContext {
	var p = new(ProcedimientosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_procedimientos

	return p
}

func (s *ProcedimientosContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedimientosContext) Get_procedimiento() IProcedimientoContext { return s._procedimiento }

func (s *ProcedimientosContext) Set_procedimiento(v IProcedimientoContext) { s._procedimiento = v }

func (s *ProcedimientosContext) GetProc() []IProcedimientoContext { return s.proc }

func (s *ProcedimientosContext) SetProc(v []IProcedimientoContext) { s.proc = v }

func (s *ProcedimientosContext) GetLista() *arrayList.List { return s.lista }

func (s *ProcedimientosContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ProcedimientosContext) AllProcedimiento() []IProcedimientoContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcedimientoContext)(nil)).Elem())
	var tst = make([]IProcedimientoContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcedimientoContext)
		}
	}

	return tst
}

func (s *ProcedimientosContext) Procedimiento(i int) IProcedimientoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedimientoContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcedimientoContext)
}

func (s *ProcedimientosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedimientosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedimientosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterProcedimientos(s)
	}
}

func (s *ProcedimientosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitProcedimientos(s)
	}
}

func (p *sintactico) Procedimientos() (localctx IProcedimientosContext) {
	this := p
	_ = this

	localctx = NewProcedimientosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sintacticoRULE_procedimientos)
	localctx.(*ProcedimientosContext).lista = arrayList.New()
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sintacticoR_FN {
		{
			p.SetState(43)

			var _x = p.Procedimiento()

			localctx.(*ProcedimientosContext)._procedimiento = _x
		}
		localctx.(*ProcedimientosContext).proc = append(localctx.(*ProcedimientosContext).proc, localctx.(*ProcedimientosContext)._procedimiento)

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	LISTA := localctx.(*ProcedimientosContext).GetProc()
	for _, i := range LISTA {
		localctx.(*ProcedimientosContext).lista.Add(i.GetInstr())
	}

	return localctx
}

// IProcedimientoContext is an interface to support dynamic dispatch.
type IProcedimientoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_principal returns the _principal rule contexts.
	Get_principal() IPrincipalContext

	// Set_principal sets the _principal rule contexts.
	Set_principal(IPrincipalContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsProcedimientoContext differentiates from other interfaces.
	IsProcedimientoContext()
}

type ProcedimientoContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	instr      interfaces.Instruccion
	_principal IPrincipalContext
}

func NewEmptyProcedimientoContext() *ProcedimientoContext {
	var p = new(ProcedimientoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_procedimiento
	return p
}

func (*ProcedimientoContext) IsProcedimientoContext() {}

func NewProcedimientoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedimientoContext {
	var p = new(ProcedimientoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_procedimiento

	return p
}

func (s *ProcedimientoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedimientoContext) Get_principal() IPrincipalContext { return s._principal }

func (s *ProcedimientoContext) Set_principal(v IPrincipalContext) { s._principal = v }

func (s *ProcedimientoContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ProcedimientoContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ProcedimientoContext) Principal() IPrincipalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrincipalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrincipalContext)
}

func (s *ProcedimientoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedimientoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedimientoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterProcedimiento(s)
	}
}

func (s *ProcedimientoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitProcedimiento(s)
	}
}

func (p *sintactico) Procedimiento() (localctx IProcedimientoContext) {
	this := p
	_ = this

	localctx = NewProcedimientoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sintacticoRULE_procedimiento)

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
		p.SetState(51)

		var _x = p.Principal()

		localctx.(*ProcedimientoContext)._principal = _x
	}
	localctx.(*ProcedimientoContext).instr = localctx.(*ProcedimientoContext).Get_principal().GetInstr()

	return localctx
}

// IPrincipalContext is an interface to support dynamic dispatch.
type IPrincipalContext interface {
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

	// IsPrincipalContext differentiates from other interfaces.
	IsPrincipalContext()
}

type PrincipalContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruccion
	_instrucciones IInstruccionesContext
}

func NewEmptyPrincipalContext() *PrincipalContext {
	var p = new(PrincipalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_principal
	return p
}

func (*PrincipalContext) IsPrincipalContext() {}

func NewPrincipalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrincipalContext {
	var p = new(PrincipalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_principal

	return p
}

func (s *PrincipalContext) GetParser() antlr.Parser { return s.parser }

func (s *PrincipalContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *PrincipalContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *PrincipalContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *PrincipalContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *PrincipalContext) R_FN() antlr.TerminalNode {
	return s.GetToken(sintacticoR_FN, 0)
}

func (s *PrincipalContext) R_MAIN() antlr.TerminalNode {
	return s.GetToken(sintacticoR_MAIN, 0)
}

func (s *PrincipalContext) S_APAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_APAR, 0)
}

func (s *PrincipalContext) S_CPAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CPAR, 0)
}

func (s *PrincipalContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *PrincipalContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *PrincipalContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *PrincipalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrincipalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrincipalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterPrincipal(s)
	}
}

func (s *PrincipalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitPrincipal(s)
	}
}

func (p *sintactico) Principal() (localctx IPrincipalContext) {
	this := p
	_ = this

	localctx = NewPrincipalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sintacticoRULE_principal)
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
		p.SetState(54)
		p.Match(sintacticoR_FN)
	}
	{
		p.SetState(55)
		p.Match(sintacticoR_MAIN)
	}
	{
		p.SetState(56)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(57)
		p.Match(sintacticoS_CPAR)
	}
	{
		p.SetState(58)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(59)

		var _x = p.Instrucciones()

		localctx.(*PrincipalContext)._instrucciones = _x
	}
	{
		p.SetState(60)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*PrincipalContext).instr = funcion.NewFuncion(entorno.VOID, "main", params, localctx.(*PrincipalContext).Get_instrucciones().GetLista())

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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(sintacticoR_LET-34))|(1<<(sintacticoR_PRINTLN-34))|(1<<(sintacticoR_IF-34)))) != 0) || _la == sintacticoID {
		{
			p.SetState(63)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).ins = append(localctx.(*InstruccionesContext).ins, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(66)
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

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_sent_if returns the _sent_if rule contexts.
	Get_sent_if() ISent_ifContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_sent_if sets the _sent_if rule contexts.
	Set_sent_if(ISent_ifContext)

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
	_asignacion  IAsignacionContext
	_sent_if     ISent_ifContext
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

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Get_sent_if() ISent_ifContext { return s._sent_if }

func (s *InstruccionContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_sent_if(v ISent_ifContext) { s._sent_if = v }

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

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) Sent_if() ISent_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISent_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISent_ifContext)
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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)

			var _x = p.Imprimir()

			localctx.(*InstruccionContext)._imprimir = _x
		}
		{
			p.SetState(71)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_imprimir().GetInstr()

	case sintacticoR_LET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(75)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracion().GetInstr()

	case sintacticoID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(79)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_asignacion().GetInstr()

	case sintacticoR_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)

			var _x = p.Sent_if()

			localctx.(*InstruccionContext)._sent_if = _x
		}
		{
			p.SetState(83)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_sent_if().GetInstr()

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
		p.SetState(88)
		p.Match(sintacticoR_PRINTLN)
	}
	{
		p.SetState(89)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(90)

		var _x = p.lista_exp(0)

		localctx.(*ImprimirContext)._lista_exp = _x
	}
	{
		p.SetState(91)
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
		p.SetState(95)

		var _x = p.Exp()

		localctx.(*Lista_expContext)._exp = _x
	}
	localctx.(*Lista_expContext).lista.Add(localctx.(*Lista_expContext).Get_exp().GetVal())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(105)
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
			p.SetState(98)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(99)
				p.Match(sintacticoS_COMA)
			}
			{
				p.SetState(100)

				var _x = p.Exp()

				localctx.(*Lista_expContext)._exp = _x
			}

			localctx.(*Lista_expContext).GetLISTA().GetLista().Add(localctx.(*Lista_expContext).Get_exp().GetVal())
			localctx.(*Lista_expContext).lista = localctx.(*Lista_expContext).GetLISTA().GetLista()

		}
		p.SetState(107)
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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(109)
			p.Match(sintacticoR_MUT)
		}
		{
			p.SetState(110)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(111)
			p.Match(sintacticoS_DOSPT)
		}
		{
			p.SetState(112)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(113)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(114)

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
			p.SetState(117)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(118)
			p.Match(sintacticoR_MUT)
		}
		{
			p.SetState(119)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(120)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(121)

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
			p.SetState(124)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(125)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(126)
			p.Match(sintacticoS_DOSPT)
		}
		{
			p.SetState(127)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(128)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(129)

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
			p.SetState(132)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(133)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(134)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(135)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(sintacticoR_INT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.INTEGER

	case sintacticoR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Match(sintacticoR_FLOAT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.FLOAT

	case sintacticoR_CHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Match(sintacticoR_CHAR)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.CHAR

	case sintacticoR_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(146)
			p.Match(sintacticoR_STRING)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.STRING

	case sintacticoR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(148)
			p.Match(sintacticoR_BOOL)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.BOOLEAN

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  interfaces.Instruccion
	_ID    antlr.Token
	_exp   IExpContext
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignacionContext) Get_exp() IExpContext { return s._exp }

func (s *AsignacionContext) Set_exp(v IExpContext) { s._exp = v }

func (s *AsignacionContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *AsignacionContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(sintacticoID, 0)
}

func (s *AsignacionContext) S_ASIGNAR() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ASIGNAR, 0)
}

func (s *AsignacionContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *sintactico) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, sintacticoRULE_asignacion)

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
		p.SetState(152)

		var _m = p.Match(sintacticoID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(153)
		p.Match(sintacticoS_ASIGNAR)
	}
	{
		p.SetState(154)

		var _x = p.Exp()

		localctx.(*AsignacionContext)._exp = _x
	}

	id := expresion.NewIdentificador((func() string {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetText()
		}
	}()))
	localctx.(*AsignacionContext).instr = variables.NewAsignacion(id, localctx.(*AsignacionContext).Get_exp().GetVal())

	return localctx
}

// ISent_ifContext is an interface to support dynamic dispatch.
type ISent_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// GetS_then returns the s_then rule contexts.
	GetS_then() IInstruccionesContext

	// GetS_else returns the s_else rule contexts.
	GetS_else() IInstruccionesContext

	// Get_lista_elseif returns the _lista_elseif rule contexts.
	Get_lista_elseif() ILista_elseifContext

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// SetS_then sets the s_then rule contexts.
	SetS_then(IInstruccionesContext)

	// SetS_else sets the s_else rule contexts.
	SetS_else(IInstruccionesContext)

	// Set_lista_elseif sets the _lista_elseif rule contexts.
	Set_lista_elseif(ILista_elseifContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsSent_ifContext differentiates from other interfaces.
	IsSent_ifContext()
}

type Sent_ifContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	instr         interfaces.Instruccion
	_exp          IExpContext
	s_then        IInstruccionesContext
	s_else        IInstruccionesContext
	_lista_elseif ILista_elseifContext
}

func NewEmptySent_ifContext() *Sent_ifContext {
	var p = new(Sent_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_sent_if
	return p
}

func (*Sent_ifContext) IsSent_ifContext() {}

func NewSent_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sent_ifContext {
	var p = new(Sent_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_sent_if

	return p
}

func (s *Sent_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Sent_ifContext) Get_exp() IExpContext { return s._exp }

func (s *Sent_ifContext) GetS_then() IInstruccionesContext { return s.s_then }

func (s *Sent_ifContext) GetS_else() IInstruccionesContext { return s.s_else }

func (s *Sent_ifContext) Get_lista_elseif() ILista_elseifContext { return s._lista_elseif }

func (s *Sent_ifContext) Set_exp(v IExpContext) { s._exp = v }

func (s *Sent_ifContext) SetS_then(v IInstruccionesContext) { s.s_then = v }

func (s *Sent_ifContext) SetS_else(v IInstruccionesContext) { s.s_else = v }

func (s *Sent_ifContext) Set_lista_elseif(v ILista_elseifContext) { s._lista_elseif = v }

func (s *Sent_ifContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *Sent_ifContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *Sent_ifContext) R_IF() antlr.TerminalNode {
	return s.GetToken(sintacticoR_IF, 0)
}

func (s *Sent_ifContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Sent_ifContext) AllS_ALLAV() []antlr.TerminalNode {
	return s.GetTokens(sintacticoS_ALLAV)
}

func (s *Sent_ifContext) S_ALLAV(i int) antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, i)
}

func (s *Sent_ifContext) AllS_CLLAV() []antlr.TerminalNode {
	return s.GetTokens(sintacticoS_CLLAV)
}

func (s *Sent_ifContext) S_CLLAV(i int) antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, i)
}

func (s *Sent_ifContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *Sent_ifContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Sent_ifContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(sintacticoR_ELSE, 0)
}

func (s *Sent_ifContext) Lista_elseif() ILista_elseifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILista_elseifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILista_elseifContext)
}

func (s *Sent_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sent_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sent_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterSent_if(s)
	}
}

func (s *Sent_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitSent_if(s)
	}
}

func (p *sintactico) Sent_if() (localctx ISent_ifContext) {
	this := p
	_ = this

	localctx = NewSent_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, sintacticoRULE_sent_if)

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

	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(158)

			var _x = p.Exp()

			localctx.(*Sent_ifContext)._exp = _x
		}
		{
			p.SetState(159)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(160)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_then = _x
		}
		{
			p.SetState(161)
			p.Match(sintacticoS_CLLAV)
		}

		localctx.(*Sent_ifContext).instr = flujo.NewIf(localctx.(*Sent_ifContext).Get_exp().GetVal(), localctx.(*Sent_ifContext).GetS_then().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(165)

			var _x = p.Exp()

			localctx.(*Sent_ifContext)._exp = _x
		}
		{
			p.SetState(166)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(167)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_then = _x
		}
		{
			p.SetState(168)
			p.Match(sintacticoS_CLLAV)
		}
		{
			p.SetState(169)
			p.Match(sintacticoR_ELSE)
		}
		{
			p.SetState(170)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(171)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_else = _x
		}
		{
			p.SetState(172)
			p.Match(sintacticoS_CLLAV)
		}

		localctx.(*Sent_ifContext).instr = flujo.NewIf(localctx.(*Sent_ifContext).Get_exp().GetVal(), localctx.(*Sent_ifContext).GetS_then().GetLista(), nil, localctx.(*Sent_ifContext).GetS_else().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(176)

			var _x = p.Exp()

			localctx.(*Sent_ifContext)._exp = _x
		}
		{
			p.SetState(177)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(178)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_then = _x
		}
		{
			p.SetState(179)
			p.Match(sintacticoS_CLLAV)
		}
		{
			p.SetState(180)

			var _x = p.Lista_elseif()

			localctx.(*Sent_ifContext)._lista_elseif = _x
		}
		{
			p.SetState(181)
			p.Match(sintacticoR_ELSE)
		}
		{
			p.SetState(182)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(183)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_else = _x
		}
		{
			p.SetState(184)
			p.Match(sintacticoS_CLLAV)
		}

		localctx.(*Sent_ifContext).instr = flujo.NewIf(localctx.(*Sent_ifContext).Get_exp().GetVal(), localctx.(*Sent_ifContext).GetS_then().GetLista(), localctx.(*Sent_ifContext).Get_lista_elseif().GetLista(), localctx.(*Sent_ifContext).GetS_else().GetLista())

	}

	return localctx
}

// ILista_elseifContext is an interface to support dynamic dispatch.
type ILista_elseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_elseif returns the _elseif rule contexts.
	Get_elseif() IElseifContext

	// Set_elseif sets the _elseif rule contexts.
	Set_elseif(IElseifContext)

	// GetIns returns the ins rule context list.
	GetIns() []IElseifContext

	// SetIns sets the ins rule context list.
	SetIns([]IElseifContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsLista_elseifContext differentiates from other interfaces.
	IsLista_elseifContext()
}

type Lista_elseifContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lista   *arrayList.List
	_elseif IElseifContext
	ins     []IElseifContext
}

func NewEmptyLista_elseifContext() *Lista_elseifContext {
	var p = new(Lista_elseifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_lista_elseif
	return p
}

func (*Lista_elseifContext) IsLista_elseifContext() {}

func NewLista_elseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_elseifContext {
	var p = new(Lista_elseifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_lista_elseif

	return p
}

func (s *Lista_elseifContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_elseifContext) Get_elseif() IElseifContext { return s._elseif }

func (s *Lista_elseifContext) Set_elseif(v IElseifContext) { s._elseif = v }

func (s *Lista_elseifContext) GetIns() []IElseifContext { return s.ins }

func (s *Lista_elseifContext) SetIns(v []IElseifContext) { s.ins = v }

func (s *Lista_elseifContext) GetLista() *arrayList.List { return s.lista }

func (s *Lista_elseifContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *Lista_elseifContext) AllElseif() []IElseifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseifContext)(nil)).Elem())
	var tst = make([]IElseifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseifContext)
		}
	}

	return tst
}

func (s *Lista_elseifContext) Elseif(i int) IElseifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseifContext)
}

func (s *Lista_elseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_elseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_elseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterLista_elseif(s)
	}
}

func (s *Lista_elseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitLista_elseif(s)
	}
}

func (p *sintactico) Lista_elseif() (localctx ILista_elseifContext) {
	this := p
	_ = this

	localctx = NewLista_elseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, sintacticoRULE_lista_elseif)
	localctx.(*Lista_elseifContext).lista = arrayList.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(189)

				var _x = p.Elseif()

				localctx.(*Lista_elseifContext)._elseif = _x
			}
			localctx.(*Lista_elseifContext).ins = append(localctx.(*Lista_elseifContext).ins, localctx.(*Lista_elseifContext)._elseif)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	LISTA := localctx.(*Lista_elseifContext).GetIns()
	for _, i := range LISTA {
		localctx.(*Lista_elseifContext).lista.Add(i.GetInstr())
	}

	return localctx
}

// IElseifContext is an interface to support dynamic dispatch.
type IElseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsElseifContext differentiates from other interfaces.
	IsElseifContext()
}

type ElseifContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruccion
	_exp           IExpContext
	_instrucciones IInstruccionesContext
}

func NewEmptyElseifContext() *ElseifContext {
	var p = new(ElseifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_elseif
	return p
}

func (*ElseifContext) IsElseifContext() {}

func NewElseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifContext {
	var p = new(ElseifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_elseif

	return p
}

func (s *ElseifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifContext) Get_exp() IExpContext { return s._exp }

func (s *ElseifContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *ElseifContext) Set_exp(v IExpContext) { s._exp = v }

func (s *ElseifContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *ElseifContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ElseifContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ElseifContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(sintacticoR_ELSE, 0)
}

func (s *ElseifContext) R_IF() antlr.TerminalNode {
	return s.GetToken(sintacticoR_IF, 0)
}

func (s *ElseifContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ElseifContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *ElseifContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *ElseifContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *ElseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterElseif(s)
	}
}

func (s *ElseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitElseif(s)
	}
}

func (p *sintactico) Elseif() (localctx IElseifContext) {
	this := p
	_ = this

	localctx = NewElseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, sintacticoRULE_elseif)

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
		p.SetState(196)
		p.Match(sintacticoR_ELSE)
	}
	{
		p.SetState(197)
		p.Match(sintacticoR_IF)
	}
	{
		p.SetState(198)

		var _x = p.Exp()

		localctx.(*ElseifContext)._exp = _x
	}
	{
		p.SetState(199)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(200)

		var _x = p.Instrucciones()

		localctx.(*ElseifContext)._instrucciones = _x
	}
	{
		p.SetState(201)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*ElseifContext).instr = flujo.NewIf(localctx.(*ElseifContext).Get_exp().GetVal(), localctx.(*ElseifContext).Get_instrucciones().GetLista(), nil, nil)

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
	p.EnterRule(localctx, 28, sintacticoRULE_exp)

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

	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)

			var _x = p.logica(0)

			localctx.(*ExpContext)._logica = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_logica().GetVal()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)

			var _x = p.relacional(0)

			localctx.(*ExpContext)._relacional = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_relacional().GetVal()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(210)

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
	_startState := 30
	p.EnterRecursionRule(localctx, 30, sintacticoRULE_logica, _p)

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
		p.SetState(216)

		var _x = p.relacional(0)

		localctx.(*LogicaContext)._relacional = _x
	}
	localctx.(*LogicaContext).val = localctx.(*LogicaContext).Get_relacional().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogicaContext(p, _parentctx, _parentState)
				localctx.(*LogicaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_logica)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(220)

					var _m = p.Match(sintacticoS_AND)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(221)

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
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(225)

					var _m = p.Match(sintacticoS_OR)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(226)

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
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, sintacticoRULE_relacional, _p)
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
		p.SetState(235)

		var _x = p.aritmetica(0)

		localctx.(*RelacionalContext)._aritmetica = _x
	}
	localctx.(*RelacionalContext).val = localctx.(*RelacionalContext).Get_aritmetica().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelacionalContext(p, _parentctx, _parentState)
			localctx.(*RelacionalContext).izq = _prevctx
			p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_relacional)
			p.SetState(238)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(239)

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
				p.SetState(240)

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
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	_startState := 34
	p.EnterRecursionRule(localctx, 34, sintacticoRULE_aritmetica, _p)
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
	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoS_RESTA:
		{
			p.SetState(249)
			p.Match(sintacticoS_RESTA)
		}
		{
			p.SetState(250)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		localctx.(*AritmeticaContext).val = expresion.NewOperacion(localctx.(*AritmeticaContext).Get_exp().GetVal(), nil, "-", true)

	case sintacticoR_TRUE, sintacticoR_FALSE, sintacticoNUMERO, sintacticoDECIMAL, sintacticoCADENA, sintacticoID:
		{
			p.SetState(253)

			var _x = p.Exp_valor()

			localctx.(*AritmeticaContext)._exp_valor = _x
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp_valor().GetVal()

	case sintacticoS_APAR:
		{
			p.SetState(256)
			p.Match(sintacticoS_APAR)
		}
		{
			p.SetState(257)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		{
			p.SetState(258)
			p.Match(sintacticoS_CPAR)
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp().GetVal()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(273)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAritmeticaContext(p, _parentctx, _parentState)
				localctx.(*AritmeticaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_aritmetica)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(264)

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
					p.SetState(265)

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
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(269)

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
					p.SetState(270)

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
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, sintacticoRULE_exp_valor)

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
		p.SetState(278)

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
	p.EnterRule(localctx, 38, sintacticoRULE_primitivo)

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

	p.SetState(293)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoNUMERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)

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
			p.SetState(283)

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
			p.SetState(285)

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
			p.SetState(287)
			p.Match(sintacticoR_TRUE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(true, entorno.BOOLEAN)

	case sintacticoR_FALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(289)
			p.Match(sintacticoR_FALSE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(false, entorno.BOOLEAN)

	case sintacticoID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(291)

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

	case 15:
		var t *LogicaContext = nil
		if localctx != nil {
			t = localctx.(*LogicaContext)
		}
		return p.Logica_Sempred(t, predIndex)

	case 16:
		var t *RelacionalContext = nil
		if localctx != nil {
			t = localctx.(*RelacionalContext)
		}
		return p.Relacional_Sempred(t, predIndex)

	case 17:
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
