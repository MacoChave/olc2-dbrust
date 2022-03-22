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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 84, 340,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 55, 10, 3, 12,
	3, 14, 3, 58, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 6, 6, 75, 10, 6, 13, 6, 14, 6, 76,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 101,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 118, 10, 9, 12, 9, 14, 9, 121, 11, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 153,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 5, 11, 165, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 202, 10, 13,
	3, 14, 6, 14, 205, 10, 14, 13, 14, 14, 14, 206, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 228, 10, 17, 13, 17, 14, 17,
	229, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 258, 10, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 7, 21, 274, 10, 21, 12, 21, 14, 21, 277, 11, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 288, 10, 22,
	12, 22, 14, 22, 291, 11, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 306, 10, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23,
	318, 10, 23, 12, 23, 14, 23, 321, 11, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 338, 10, 25, 3, 25, 2, 6, 16, 40, 42, 44, 26, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	2, 5, 3, 2, 16, 21, 3, 2, 13, 15, 3, 2, 11, 12, 2, 347, 2, 50, 3, 2, 2,
	2, 4, 56, 3, 2, 2, 2, 6, 61, 3, 2, 2, 2, 8, 64, 3, 2, 2, 2, 10, 74, 3,
	2, 2, 2, 12, 100, 3, 2, 2, 2, 14, 102, 3, 2, 2, 2, 16, 108, 3, 2, 2, 2,
	18, 152, 3, 2, 2, 2, 20, 164, 3, 2, 2, 2, 22, 166, 3, 2, 2, 2, 24, 201,
	3, 2, 2, 2, 26, 204, 3, 2, 2, 2, 28, 210, 3, 2, 2, 2, 30, 218, 3, 2, 2,
	2, 32, 227, 3, 2, 2, 2, 34, 233, 3, 2, 2, 2, 36, 241, 3, 2, 2, 2, 38, 257,
	3, 2, 2, 2, 40, 259, 3, 2, 2, 2, 42, 278, 3, 2, 2, 2, 44, 305, 3, 2, 2,
	2, 46, 322, 3, 2, 2, 2, 48, 337, 3, 2, 2, 2, 50, 51, 5, 4, 3, 2, 51, 52,
	8, 2, 1, 2, 52, 3, 3, 2, 2, 2, 53, 55, 5, 6, 4, 2, 54, 53, 3, 2, 2, 2,
	55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 59, 3,
	2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 8, 3, 1, 2, 60, 5, 3, 2, 2, 2, 61,
	62, 5, 8, 5, 2, 62, 63, 8, 4, 1, 2, 63, 7, 3, 2, 2, 2, 64, 65, 7, 44, 2,
	2, 65, 66, 7, 45, 2, 2, 66, 67, 7, 25, 2, 2, 67, 68, 7, 26, 2, 2, 68, 69,
	7, 31, 2, 2, 69, 70, 5, 10, 6, 2, 70, 71, 7, 32, 2, 2, 71, 72, 8, 5, 1,
	2, 72, 9, 3, 2, 2, 2, 73, 75, 5, 12, 7, 2, 74, 73, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2,
	78, 79, 8, 6, 1, 2, 79, 11, 3, 2, 2, 2, 80, 81, 5, 14, 8, 2, 81, 82, 7,
	5, 2, 2, 82, 83, 8, 7, 1, 2, 83, 101, 3, 2, 2, 2, 84, 85, 5, 18, 10, 2,
	85, 86, 7, 5, 2, 2, 86, 87, 8, 7, 1, 2, 87, 101, 3, 2, 2, 2, 88, 89, 5,
	22, 12, 2, 89, 90, 7, 5, 2, 2, 90, 91, 8, 7, 1, 2, 91, 101, 3, 2, 2, 2,
	92, 93, 5, 24, 13, 2, 93, 94, 7, 5, 2, 2, 94, 95, 8, 7, 1, 2, 95, 101,
	3, 2, 2, 2, 96, 97, 5, 30, 16, 2, 97, 98, 7, 5, 2, 2, 98, 99, 8, 7, 1,
	2, 99, 101, 3, 2, 2, 2, 100, 80, 3, 2, 2, 2, 100, 84, 3, 2, 2, 2, 100,
	88, 3, 2, 2, 2, 100, 92, 3, 2, 2, 2, 100, 96, 3, 2, 2, 2, 101, 13, 3, 2,
	2, 2, 102, 103, 7, 51, 2, 2, 103, 104, 7, 25, 2, 2, 104, 105, 5, 16, 9,
	2, 105, 106, 7, 26, 2, 2, 106, 107, 8, 8, 1, 2, 107, 15, 3, 2, 2, 2, 108,
	109, 8, 9, 1, 2, 109, 110, 5, 38, 20, 2, 110, 111, 8, 9, 1, 2, 111, 119,
	3, 2, 2, 2, 112, 113, 12, 4, 2, 2, 113, 114, 7, 4, 2, 2, 114, 115, 5, 38,
	20, 2, 115, 116, 8, 9, 1, 2, 116, 118, 3, 2, 2, 2, 117, 112, 3, 2, 2, 2,
	118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120,
	17, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 7, 36, 2, 2, 123, 124,
	7, 37, 2, 2, 124, 125, 7, 82, 2, 2, 125, 126, 7, 27, 2, 2, 126, 127, 5,
	20, 11, 2, 127, 128, 7, 6, 2, 2, 128, 129, 5, 38, 20, 2, 129, 130, 8, 10,
	1, 2, 130, 153, 3, 2, 2, 2, 131, 132, 7, 36, 2, 2, 132, 133, 7, 37, 2,
	2, 133, 134, 7, 82, 2, 2, 134, 135, 7, 6, 2, 2, 135, 136, 5, 38, 20, 2,
	136, 137, 8, 10, 1, 2, 137, 153, 3, 2, 2, 2, 138, 139, 7, 36, 2, 2, 139,
	140, 7, 82, 2, 2, 140, 141, 7, 27, 2, 2, 141, 142, 5, 20, 11, 2, 142, 143,
	7, 6, 2, 2, 143, 144, 5, 38, 20, 2, 144, 145, 8, 10, 1, 2, 145, 153, 3,
	2, 2, 2, 146, 147, 7, 36, 2, 2, 147, 148, 7, 82, 2, 2, 148, 149, 7, 6,
	2, 2, 149, 150, 5, 38, 20, 2, 150, 151, 8, 10, 1, 2, 151, 153, 3, 2, 2,
	2, 152, 122, 3, 2, 2, 2, 152, 131, 3, 2, 2, 2, 152, 138, 3, 2, 2, 2, 152,
	146, 3, 2, 2, 2, 153, 19, 3, 2, 2, 2, 154, 155, 7, 38, 2, 2, 155, 165,
	8, 11, 1, 2, 156, 157, 7, 39, 2, 2, 157, 165, 8, 11, 1, 2, 158, 159, 7,
	41, 2, 2, 159, 165, 8, 11, 1, 2, 160, 161, 7, 42, 2, 2, 161, 165, 8, 11,
	1, 2, 162, 163, 7, 40, 2, 2, 163, 165, 8, 11, 1, 2, 164, 154, 3, 2, 2,
	2, 164, 156, 3, 2, 2, 2, 164, 158, 3, 2, 2, 2, 164, 160, 3, 2, 2, 2, 164,
	162, 3, 2, 2, 2, 165, 21, 3, 2, 2, 2, 166, 167, 7, 82, 2, 2, 167, 168,
	7, 6, 2, 2, 168, 169, 5, 38, 20, 2, 169, 170, 8, 12, 1, 2, 170, 23, 3,
	2, 2, 2, 171, 172, 7, 63, 2, 2, 172, 173, 5, 38, 20, 2, 173, 174, 7, 31,
	2, 2, 174, 175, 5, 10, 6, 2, 175, 176, 7, 32, 2, 2, 176, 177, 8, 13, 1,
	2, 177, 202, 3, 2, 2, 2, 178, 179, 7, 63, 2, 2, 179, 180, 5, 38, 20, 2,
	180, 181, 7, 31, 2, 2, 181, 182, 5, 10, 6, 2, 182, 183, 7, 32, 2, 2, 183,
	184, 7, 64, 2, 2, 184, 185, 7, 31, 2, 2, 185, 186, 5, 10, 6, 2, 186, 187,
	7, 32, 2, 2, 187, 188, 8, 13, 1, 2, 188, 202, 3, 2, 2, 2, 189, 190, 7,
	63, 2, 2, 190, 191, 5, 38, 20, 2, 191, 192, 7, 31, 2, 2, 192, 193, 5, 10,
	6, 2, 193, 194, 7, 32, 2, 2, 194, 195, 5, 26, 14, 2, 195, 196, 7, 64, 2,
	2, 196, 197, 7, 31, 2, 2, 197, 198, 5, 10, 6, 2, 198, 199, 7, 32, 2, 2,
	199, 200, 8, 13, 1, 2, 200, 202, 3, 2, 2, 2, 201, 171, 3, 2, 2, 2, 201,
	178, 3, 2, 2, 2, 201, 189, 3, 2, 2, 2, 202, 25, 3, 2, 2, 2, 203, 205, 5,
	28, 15, 2, 204, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 204, 3, 2,
	2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 8, 14, 1, 2,
	209, 27, 3, 2, 2, 2, 210, 211, 7, 64, 2, 2, 211, 212, 7, 63, 2, 2, 212,
	213, 5, 38, 20, 2, 213, 214, 7, 31, 2, 2, 214, 215, 5, 10, 6, 2, 215, 216,
	7, 32, 2, 2, 216, 217, 8, 15, 1, 2, 217, 29, 3, 2, 2, 2, 218, 219, 7, 65,
	2, 2, 219, 220, 5, 38, 20, 2, 220, 221, 7, 31, 2, 2, 221, 222, 5, 32, 17,
	2, 222, 223, 5, 36, 19, 2, 223, 224, 7, 32, 2, 2, 224, 225, 8, 16, 1, 2,
	225, 31, 3, 2, 2, 2, 226, 228, 5, 34, 18, 2, 227, 226, 3, 2, 2, 2, 228,
	229, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231,
	3, 2, 2, 2, 231, 232, 8, 17, 1, 2, 232, 33, 3, 2, 2, 2, 233, 234, 5, 38,
	20, 2, 234, 235, 7, 9, 2, 2, 235, 236, 7, 31, 2, 2, 236, 237, 5, 12, 7,
	2, 237, 238, 7, 32, 2, 2, 238, 239, 7, 4, 2, 2, 239, 240, 8, 18, 1, 2,
	240, 35, 3, 2, 2, 2, 241, 242, 7, 10, 2, 2, 242, 243, 7, 9, 2, 2, 243,
	244, 7, 31, 2, 2, 244, 245, 5, 12, 7, 2, 245, 246, 7, 32, 2, 2, 246, 247,
	8, 19, 1, 2, 247, 37, 3, 2, 2, 2, 248, 249, 5, 40, 21, 2, 249, 250, 8,
	20, 1, 2, 250, 258, 3, 2, 2, 2, 251, 252, 5, 42, 22, 2, 252, 253, 8, 20,
	1, 2, 253, 258, 3, 2, 2, 2, 254, 255, 5, 44, 23, 2, 255, 256, 8, 20, 1,
	2, 256, 258, 3, 2, 2, 2, 257, 248, 3, 2, 2, 2, 257, 251, 3, 2, 2, 2, 257,
	254, 3, 2, 2, 2, 258, 39, 3, 2, 2, 2, 259, 260, 8, 21, 1, 2, 260, 261,
	5, 42, 22, 2, 261, 262, 8, 21, 1, 2, 262, 275, 3, 2, 2, 2, 263, 264, 12,
	5, 2, 2, 264, 265, 7, 23, 2, 2, 265, 266, 5, 40, 21, 6, 266, 267, 8, 21,
	1, 2, 267, 274, 3, 2, 2, 2, 268, 269, 12, 4, 2, 2, 269, 270, 7, 22, 2,
	2, 270, 271, 5, 40, 21, 5, 271, 272, 8, 21, 1, 2, 272, 274, 3, 2, 2, 2,
	273, 263, 3, 2, 2, 2, 273, 268, 3, 2, 2, 2, 274, 277, 3, 2, 2, 2, 275,
	273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 41, 3, 2, 2, 2, 277, 275, 3,
	2, 2, 2, 278, 279, 8, 22, 1, 2, 279, 280, 5, 44, 23, 2, 280, 281, 8, 22,
	1, 2, 281, 289, 3, 2, 2, 2, 282, 283, 12, 4, 2, 2, 283, 284, 9, 2, 2, 2,
	284, 285, 5, 42, 22, 5, 285, 286, 8, 22, 1, 2, 286, 288, 3, 2, 2, 2, 287,
	282, 3, 2, 2, 2, 288, 291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290,
	3, 2, 2, 2, 290, 43, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 293, 8, 23,
	1, 2, 293, 294, 7, 12, 2, 2, 294, 295, 5, 38, 20, 2, 295, 296, 8, 23, 1,
	2, 296, 306, 3, 2, 2, 2, 297, 298, 5, 46, 24, 2, 298, 299, 8, 23, 1, 2,
	299, 306, 3, 2, 2, 2, 300, 301, 7, 25, 2, 2, 301, 302, 5, 38, 20, 2, 302,
	303, 7, 26, 2, 2, 303, 304, 8, 23, 1, 2, 304, 306, 3, 2, 2, 2, 305, 292,
	3, 2, 2, 2, 305, 297, 3, 2, 2, 2, 305, 300, 3, 2, 2, 2, 306, 319, 3, 2,
	2, 2, 307, 308, 12, 6, 2, 2, 308, 309, 9, 3, 2, 2, 309, 310, 5, 44, 23,
	7, 310, 311, 8, 23, 1, 2, 311, 318, 3, 2, 2, 2, 312, 313, 12, 5, 2, 2,
	313, 314, 9, 4, 2, 2, 314, 315, 5, 44, 23, 6, 315, 316, 8, 23, 1, 2, 316,
	318, 3, 2, 2, 2, 317, 307, 3, 2, 2, 2, 317, 312, 3, 2, 2, 2, 318, 321,
	3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 45, 3, 2,
	2, 2, 321, 319, 3, 2, 2, 2, 322, 323, 5, 48, 25, 2, 323, 324, 8, 24, 1,
	2, 324, 47, 3, 2, 2, 2, 325, 326, 7, 78, 2, 2, 326, 338, 8, 25, 1, 2, 327,
	328, 7, 79, 2, 2, 328, 338, 8, 25, 1, 2, 329, 330, 7, 81, 2, 2, 330, 338,
	8, 25, 1, 2, 331, 332, 7, 34, 2, 2, 332, 338, 8, 25, 1, 2, 333, 334, 7,
	35, 2, 2, 334, 338, 8, 25, 1, 2, 335, 336, 7, 82, 2, 2, 336, 338, 8, 25,
	1, 2, 337, 325, 3, 2, 2, 2, 337, 327, 3, 2, 2, 2, 337, 329, 3, 2, 2, 2,
	337, 331, 3, 2, 2, 2, 337, 333, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 338,
	49, 3, 2, 2, 2, 19, 56, 76, 100, 119, 152, 164, 201, 206, 229, 257, 273,
	275, 289, 305, 317, 319, 337,
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
	"sent_if", "lista_elseif", "elseif", "sent_match", "casosMatch", "casoMatch",
	"matchDefecto", "exp", "logica", "relacional", "aritmetica", "exp_valor",
	"primitivo",
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
	sintacticoRULE_sent_match     = 14
	sintacticoRULE_casosMatch     = 15
	sintacticoRULE_casoMatch      = 16
	sintacticoRULE_matchDefecto   = 17
	sintacticoRULE_exp            = 18
	sintacticoRULE_logica         = 19
	sintacticoRULE_relacional     = 20
	sintacticoRULE_aritmetica     = 21
	sintacticoRULE_exp_valor      = 22
	sintacticoRULE_primitivo      = 23
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
		p.SetState(48)

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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sintacticoR_FN {
		{
			p.SetState(51)

			var _x = p.Procedimiento()

			localctx.(*ProcedimientosContext)._procedimiento = _x
		}
		localctx.(*ProcedimientosContext).proc = append(localctx.(*ProcedimientosContext).proc, localctx.(*ProcedimientosContext)._procedimiento)

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	LISTA := localctx.(*ProcedimientosContext).GetProc()
	for _, elemento := range LISTA {
		localctx.(*ProcedimientosContext).lista.Add(elemento.GetInstr())
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
		p.SetState(59)

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
		p.SetState(62)
		p.Match(sintacticoR_FN)
	}
	{
		p.SetState(63)
		p.Match(sintacticoR_MAIN)
	}
	{
		p.SetState(64)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(65)
		p.Match(sintacticoS_CPAR)
	}
	{
		p.SetState(66)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(67)

		var _x = p.Instrucciones()

		localctx.(*PrincipalContext)._instrucciones = _x
	}
	{
		p.SetState(68)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(sintacticoR_LET-34))|(1<<(sintacticoR_PRINTLN-34))|(1<<(sintacticoR_IF-34))|(1<<(sintacticoR_MATCH-34)))) != 0) || _la == sintacticoID {
		{
			p.SetState(71)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).ins = append(localctx.(*InstruccionesContext).ins, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	LISTA := localctx.(*InstruccionesContext).GetIns()
	for _, elemento := range LISTA {
		localctx.(*InstruccionesContext).lista.Add(elemento.GetInstr())
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

	// Get_sent_match returns the _sent_match rule contexts.
	Get_sent_match() ISent_matchContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_sent_if sets the _sent_if rule contexts.
	Set_sent_if(ISent_ifContext)

	// Set_sent_match sets the _sent_match rule contexts.
	Set_sent_match(ISent_matchContext)

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
	_sent_match  ISent_matchContext
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

func (s *InstruccionContext) Get_sent_match() ISent_matchContext { return s._sent_match }

func (s *InstruccionContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_sent_if(v ISent_ifContext) { s._sent_if = v }

func (s *InstruccionContext) Set_sent_match(v ISent_matchContext) { s._sent_match = v }

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

func (s *InstruccionContext) Sent_match() ISent_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISent_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISent_matchContext)
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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)

			var _x = p.Imprimir()

			localctx.(*InstruccionContext)._imprimir = _x
		}
		{
			p.SetState(79)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_imprimir().GetInstr()

	case sintacticoR_LET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(83)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracion().GetInstr()

	case sintacticoID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(87)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_asignacion().GetInstr()

	case sintacticoR_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)

			var _x = p.Sent_if()

			localctx.(*InstruccionContext)._sent_if = _x
		}
		{
			p.SetState(91)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_sent_if().GetInstr()

	case sintacticoR_MATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)

			var _x = p.Sent_match()

			localctx.(*InstruccionContext)._sent_match = _x
		}
		{
			p.SetState(95)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_sent_match().GetInstr()

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
		p.SetState(100)
		p.Match(sintacticoR_PRINTLN)
	}
	{
		p.SetState(101)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(102)

		var _x = p.lista_exp(0)

		localctx.(*ImprimirContext)._lista_exp = _x
	}
	{
		p.SetState(103)
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
		p.SetState(107)

		var _x = p.Exp()

		localctx.(*Lista_expContext)._exp = _x
	}
	localctx.(*Lista_expContext).lista.Add(localctx.(*Lista_expContext).Get_exp().GetVal())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
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
			p.SetState(110)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(111)
				p.Match(sintacticoS_COMA)
			}
			{
				p.SetState(112)

				var _x = p.Exp()

				localctx.(*Lista_expContext)._exp = _x
			}

			localctx.(*Lista_expContext).GetLISTA().GetLista().Add(localctx.(*Lista_expContext).Get_exp().GetVal())
			localctx.(*Lista_expContext).lista = localctx.(*Lista_expContext).GetLISTA().GetLista()

		}
		p.SetState(119)
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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(121)
			p.Match(sintacticoR_MUT)
		}
		{
			p.SetState(122)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(123)
			p.Match(sintacticoS_DOSPT)
		}
		{
			p.SetState(124)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(125)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(126)

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
			p.SetState(129)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(130)
			p.Match(sintacticoR_MUT)
		}
		{
			p.SetState(131)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(132)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(133)

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
			p.SetState(136)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(137)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(138)
			p.Match(sintacticoS_DOSPT)
		}
		{
			p.SetState(139)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(140)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(141)

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
			p.SetState(144)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(145)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(146)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(147)

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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(sintacticoR_INT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.INTEGER

	case sintacticoR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(sintacticoR_FLOAT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.FLOAT

	case sintacticoR_CHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)
			p.Match(sintacticoR_CHAR)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.CHAR

	case sintacticoR_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.Match(sintacticoR_STRING)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.STRING

	case sintacticoR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(160)
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
		p.SetState(164)

		var _m = p.Match(sintacticoID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(165)
		p.Match(sintacticoS_ASIGNAR)
	}
	{
		p.SetState(166)

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

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(170)

			var _x = p.Exp()

			localctx.(*Sent_ifContext)._exp = _x
		}
		{
			p.SetState(171)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(172)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_then = _x
		}
		{
			p.SetState(173)
			p.Match(sintacticoS_CLLAV)
		}

		localctx.(*Sent_ifContext).instr = flujo.NewIf(localctx.(*Sent_ifContext).Get_exp().GetVal(), localctx.(*Sent_ifContext).GetS_then().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(177)

			var _x = p.Exp()

			localctx.(*Sent_ifContext)._exp = _x
		}
		{
			p.SetState(178)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(179)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_then = _x
		}
		{
			p.SetState(180)
			p.Match(sintacticoS_CLLAV)
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

		localctx.(*Sent_ifContext).instr = flujo.NewIf(localctx.(*Sent_ifContext).Get_exp().GetVal(), localctx.(*Sent_ifContext).GetS_then().GetLista(), nil, localctx.(*Sent_ifContext).GetS_else().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(188)

			var _x = p.Exp()

			localctx.(*Sent_ifContext)._exp = _x
		}
		{
			p.SetState(189)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(190)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_then = _x
		}
		{
			p.SetState(191)
			p.Match(sintacticoS_CLLAV)
		}
		{
			p.SetState(192)

			var _x = p.Lista_elseif()

			localctx.(*Sent_ifContext)._lista_elseif = _x
		}
		{
			p.SetState(193)
			p.Match(sintacticoR_ELSE)
		}
		{
			p.SetState(194)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(195)

			var _x = p.Instrucciones()

			localctx.(*Sent_ifContext).s_else = _x
		}
		{
			p.SetState(196)
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
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(201)

				var _x = p.Elseif()

				localctx.(*Lista_elseifContext)._elseif = _x
			}
			localctx.(*Lista_elseifContext).ins = append(localctx.(*Lista_elseifContext).ins, localctx.(*Lista_elseifContext)._elseif)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	LISTA := localctx.(*Lista_elseifContext).GetIns()
	for _, elemento := range LISTA {
		localctx.(*Lista_elseifContext).lista.Add(elemento.GetInstr())
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
		p.SetState(208)
		p.Match(sintacticoR_ELSE)
	}
	{
		p.SetState(209)
		p.Match(sintacticoR_IF)
	}
	{
		p.SetState(210)

		var _x = p.Exp()

		localctx.(*ElseifContext)._exp = _x
	}
	{
		p.SetState(211)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(212)

		var _x = p.Instrucciones()

		localctx.(*ElseifContext)._instrucciones = _x
	}
	{
		p.SetState(213)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*ElseifContext).instr = flujo.NewIf(localctx.(*ElseifContext).Get_exp().GetVal(), localctx.(*ElseifContext).Get_instrucciones().GetLista(), nil, nil)

	return localctx
}

// ISent_matchContext is an interface to support dynamic dispatch.
type ISent_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// Get_casosMatch returns the _casosMatch rule contexts.
	Get_casosMatch() ICasosMatchContext

	// Get_matchDefecto returns the _matchDefecto rule contexts.
	Get_matchDefecto() IMatchDefectoContext

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// Set_casosMatch sets the _casosMatch rule contexts.
	Set_casosMatch(ICasosMatchContext)

	// Set_matchDefecto sets the _matchDefecto rule contexts.
	Set_matchDefecto(IMatchDefectoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsSent_matchContext differentiates from other interfaces.
	IsSent_matchContext()
}

type Sent_matchContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	instr         interfaces.Instruccion
	_exp          IExpContext
	_casosMatch   ICasosMatchContext
	_matchDefecto IMatchDefectoContext
}

func NewEmptySent_matchContext() *Sent_matchContext {
	var p = new(Sent_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_sent_match
	return p
}

func (*Sent_matchContext) IsSent_matchContext() {}

func NewSent_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sent_matchContext {
	var p = new(Sent_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_sent_match

	return p
}

func (s *Sent_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Sent_matchContext) Get_exp() IExpContext { return s._exp }

func (s *Sent_matchContext) Get_casosMatch() ICasosMatchContext { return s._casosMatch }

func (s *Sent_matchContext) Get_matchDefecto() IMatchDefectoContext { return s._matchDefecto }

func (s *Sent_matchContext) Set_exp(v IExpContext) { s._exp = v }

func (s *Sent_matchContext) Set_casosMatch(v ICasosMatchContext) { s._casosMatch = v }

func (s *Sent_matchContext) Set_matchDefecto(v IMatchDefectoContext) { s._matchDefecto = v }

func (s *Sent_matchContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *Sent_matchContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *Sent_matchContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(sintacticoR_MATCH, 0)
}

func (s *Sent_matchContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Sent_matchContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *Sent_matchContext) CasosMatch() ICasosMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasosMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasosMatchContext)
}

func (s *Sent_matchContext) MatchDefecto() IMatchDefectoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchDefectoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchDefectoContext)
}

func (s *Sent_matchContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *Sent_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sent_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sent_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterSent_match(s)
	}
}

func (s *Sent_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitSent_match(s)
	}
}

func (p *sintactico) Sent_match() (localctx ISent_matchContext) {
	this := p
	_ = this

	localctx = NewSent_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, sintacticoRULE_sent_match)

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
		p.SetState(216)
		p.Match(sintacticoR_MATCH)
	}
	{
		p.SetState(217)

		var _x = p.Exp()

		localctx.(*Sent_matchContext)._exp = _x
	}
	{
		p.SetState(218)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(219)

		var _x = p.CasosMatch()

		localctx.(*Sent_matchContext)._casosMatch = _x
	}
	{
		p.SetState(220)

		var _x = p.MatchDefecto()

		localctx.(*Sent_matchContext)._matchDefecto = _x
	}
	{
		p.SetState(221)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*Sent_matchContext).instr = flujo.NewSentMatch(localctx.(*Sent_matchContext).Get_exp().GetVal(), localctx.(*Sent_matchContext).Get_casosMatch().GetLista(), localctx.(*Sent_matchContext).Get_matchDefecto().GetCaso())

	return localctx
}

// ICasosMatchContext is an interface to support dynamic dispatch.
type ICasosMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_casoMatch returns the _casoMatch rule contexts.
	Get_casoMatch() ICasoMatchContext

	// Set_casoMatch sets the _casoMatch rule contexts.
	Set_casoMatch(ICasoMatchContext)

	// GetCaso returns the caso rule context list.
	GetCaso() []ICasoMatchContext

	// SetCaso sets the caso rule context list.
	SetCaso([]ICasoMatchContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsCasosMatchContext differentiates from other interfaces.
	IsCasosMatchContext()
}

type CasosMatchContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	lista      *arrayList.List
	_casoMatch ICasoMatchContext
	caso       []ICasoMatchContext
}

func NewEmptyCasosMatchContext() *CasosMatchContext {
	var p = new(CasosMatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_casosMatch
	return p
}

func (*CasosMatchContext) IsCasosMatchContext() {}

func NewCasosMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasosMatchContext {
	var p = new(CasosMatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_casosMatch

	return p
}

func (s *CasosMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *CasosMatchContext) Get_casoMatch() ICasoMatchContext { return s._casoMatch }

func (s *CasosMatchContext) Set_casoMatch(v ICasoMatchContext) { s._casoMatch = v }

func (s *CasosMatchContext) GetCaso() []ICasoMatchContext { return s.caso }

func (s *CasosMatchContext) SetCaso(v []ICasoMatchContext) { s.caso = v }

func (s *CasosMatchContext) GetLista() *arrayList.List { return s.lista }

func (s *CasosMatchContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *CasosMatchContext) AllCasoMatch() []ICasoMatchContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICasoMatchContext)(nil)).Elem())
	var tst = make([]ICasoMatchContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICasoMatchContext)
		}
	}

	return tst
}

func (s *CasosMatchContext) CasoMatch(i int) ICasoMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasoMatchContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICasoMatchContext)
}

func (s *CasosMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasosMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasosMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterCasosMatch(s)
	}
}

func (s *CasosMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitCasosMatch(s)
	}
}

func (p *sintactico) CasosMatch() (localctx ICasosMatchContext) {
	this := p
	_ = this

	localctx = NewCasosMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, sintacticoRULE_casosMatch)
	localctx.(*CasosMatchContext).lista = arrayList.New()
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
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(sintacticoS_RESTA-10))|(1<<(sintacticoS_APAR-10))|(1<<(sintacticoR_TRUE-10))|(1<<(sintacticoR_FALSE-10)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(sintacticoNUMERO-76))|(1<<(sintacticoDECIMAL-76))|(1<<(sintacticoCADENA-76))|(1<<(sintacticoID-76)))) != 0) {
		{
			p.SetState(224)

			var _x = p.CasoMatch()

			localctx.(*CasosMatchContext)._casoMatch = _x
		}
		localctx.(*CasosMatchContext).caso = append(localctx.(*CasosMatchContext).caso, localctx.(*CasosMatchContext)._casoMatch)

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	LISTA := localctx.(*CasosMatchContext).GetCaso()
	for _, elemento := range LISTA {
		localctx.(*CasosMatchContext).lista.Add(elemento.GetCaso())
	}

	return localctx
}

// ICasoMatchContext is an interface to support dynamic dispatch.
type ICasoMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetCaso returns the caso attribute.
	GetCaso() flujo.CaseMatch

	// SetCaso sets the caso attribute.
	SetCaso(flujo.CaseMatch)

	// IsCasoMatchContext differentiates from other interfaces.
	IsCasoMatchContext()
}

type CasoMatchContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	caso         flujo.CaseMatch
	_exp         IExpContext
	_instruccion IInstruccionContext
}

func NewEmptyCasoMatchContext() *CasoMatchContext {
	var p = new(CasoMatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_casoMatch
	return p
}

func (*CasoMatchContext) IsCasoMatchContext() {}

func NewCasoMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasoMatchContext {
	var p = new(CasoMatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_casoMatch

	return p
}

func (s *CasoMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *CasoMatchContext) Get_exp() IExpContext { return s._exp }

func (s *CasoMatchContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *CasoMatchContext) Set_exp(v IExpContext) { s._exp = v }

func (s *CasoMatchContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *CasoMatchContext) GetCaso() flujo.CaseMatch { return s.caso }

func (s *CasoMatchContext) SetCaso(v flujo.CaseMatch) { s.caso = v }

func (s *CasoMatchContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *CasoMatchContext) S_MATCH_RET() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MATCH_RET, 0)
}

func (s *CasoMatchContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *CasoMatchContext) Instruccion() IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *CasoMatchContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *CasoMatchContext) S_COMA() antlr.TerminalNode {
	return s.GetToken(sintacticoS_COMA, 0)
}

func (s *CasoMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasoMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasoMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterCasoMatch(s)
	}
}

func (s *CasoMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitCasoMatch(s)
	}
}

func (p *sintactico) CasoMatch() (localctx ICasoMatchContext) {
	this := p
	_ = this

	localctx = NewCasoMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, sintacticoRULE_casoMatch)

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
		p.SetState(231)

		var _x = p.Exp()

		localctx.(*CasoMatchContext)._exp = _x
	}
	{
		p.SetState(232)
		p.Match(sintacticoS_MATCH_RET)
	}
	{
		p.SetState(233)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(234)

		var _x = p.Instruccion()

		localctx.(*CasoMatchContext)._instruccion = _x
	}
	{
		p.SetState(235)
		p.Match(sintacticoS_CLLAV)
	}
	{
		p.SetState(236)
		p.Match(sintacticoS_COMA)
	}

	localctx.(*CasoMatchContext).caso = flujo.NewCaseMatch(localctx.(*CasoMatchContext).Get_exp().GetVal(), localctx.(*CasoMatchContext).Get_instruccion().GetInstr())

	return localctx
}

// IMatchDefectoContext is an interface to support dynamic dispatch.
type IMatchDefectoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetCaso returns the caso attribute.
	GetCaso() flujo.CaseMatch

	// SetCaso sets the caso attribute.
	SetCaso(flujo.CaseMatch)

	// IsMatchDefectoContext differentiates from other interfaces.
	IsMatchDefectoContext()
}

type MatchDefectoContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	caso         flujo.CaseMatch
	_instruccion IInstruccionContext
}

func NewEmptyMatchDefectoContext() *MatchDefectoContext {
	var p = new(MatchDefectoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_matchDefecto
	return p
}

func (*MatchDefectoContext) IsMatchDefectoContext() {}

func NewMatchDefectoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchDefectoContext {
	var p = new(MatchDefectoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_matchDefecto

	return p
}

func (s *MatchDefectoContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchDefectoContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *MatchDefectoContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *MatchDefectoContext) GetCaso() flujo.CaseMatch { return s.caso }

func (s *MatchDefectoContext) SetCaso(v flujo.CaseMatch) { s.caso = v }

func (s *MatchDefectoContext) S_MATCH_DEFAULT() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MATCH_DEFAULT, 0)
}

func (s *MatchDefectoContext) S_MATCH_RET() antlr.TerminalNode {
	return s.GetToken(sintacticoS_MATCH_RET, 0)
}

func (s *MatchDefectoContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *MatchDefectoContext) Instruccion() IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *MatchDefectoContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *MatchDefectoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchDefectoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchDefectoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterMatchDefecto(s)
	}
}

func (s *MatchDefectoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitMatchDefecto(s)
	}
}

func (p *sintactico) MatchDefecto() (localctx IMatchDefectoContext) {
	this := p
	_ = this

	localctx = NewMatchDefectoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, sintacticoRULE_matchDefecto)

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
		p.SetState(239)
		p.Match(sintacticoS_MATCH_DEFAULT)
	}
	{
		p.SetState(240)
		p.Match(sintacticoS_MATCH_RET)
	}
	{
		p.SetState(241)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(242)

		var _x = p.Instruccion()

		localctx.(*MatchDefectoContext)._instruccion = _x
	}
	{
		p.SetState(243)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*MatchDefectoContext).caso = flujo.NewCaseMatch(nil, localctx.(*MatchDefectoContext).Get_instruccion().GetInstr())

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
	p.EnterRule(localctx, 36, sintacticoRULE_exp)

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

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)

			var _x = p.logica(0)

			localctx.(*ExpContext)._logica = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_logica().GetVal()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)

			var _x = p.relacional(0)

			localctx.(*ExpContext)._relacional = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_relacional().GetVal()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(252)

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
	_startState := 38
	p.EnterRecursionRule(localctx, 38, sintacticoRULE_logica, _p)

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
		p.SetState(258)

		var _x = p.relacional(0)

		localctx.(*LogicaContext)._relacional = _x
	}
	localctx.(*LogicaContext).val = localctx.(*LogicaContext).Get_relacional().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogicaContext(p, _parentctx, _parentState)
				localctx.(*LogicaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_logica)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(262)

					var _m = p.Match(sintacticoS_AND)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(263)

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
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(267)

					var _m = p.Match(sintacticoS_OR)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(268)

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
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	_startState := 40
	p.EnterRecursionRule(localctx, 40, sintacticoRULE_relacional, _p)
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
		p.SetState(277)

		var _x = p.aritmetica(0)

		localctx.(*RelacionalContext)._aritmetica = _x
	}
	localctx.(*RelacionalContext).val = localctx.(*RelacionalContext).Get_aritmetica().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelacionalContext(p, _parentctx, _parentState)
			localctx.(*RelacionalContext).izq = _prevctx
			p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_relacional)
			p.SetState(280)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(281)

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
				p.SetState(282)

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
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	_startState := 42
	p.EnterRecursionRule(localctx, 42, sintacticoRULE_aritmetica, _p)
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
	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoS_RESTA:
		{
			p.SetState(291)
			p.Match(sintacticoS_RESTA)
		}
		{
			p.SetState(292)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		localctx.(*AritmeticaContext).val = expresion.NewOperacion(localctx.(*AritmeticaContext).Get_exp().GetVal(), nil, "-", true)

	case sintacticoR_TRUE, sintacticoR_FALSE, sintacticoNUMERO, sintacticoDECIMAL, sintacticoCADENA, sintacticoID:
		{
			p.SetState(295)

			var _x = p.Exp_valor()

			localctx.(*AritmeticaContext)._exp_valor = _x
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp_valor().GetVal()

	case sintacticoS_APAR:
		{
			p.SetState(298)
			p.Match(sintacticoS_APAR)
		}
		{
			p.SetState(299)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		{
			p.SetState(300)
			p.Match(sintacticoS_CPAR)
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp().GetVal()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAritmeticaContext(p, _parentctx, _parentState)
				localctx.(*AritmeticaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_aritmetica)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(306)

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
					p.SetState(307)

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
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(311)

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
					p.SetState(312)

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
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 44, sintacticoRULE_exp_valor)

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
		p.SetState(320)

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
	p.EnterRule(localctx, 46, sintacticoRULE_primitivo)

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

	p.SetState(335)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoNUMERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)

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
			p.SetState(325)

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
			p.SetState(327)

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
			p.SetState(329)
			p.Match(sintacticoR_TRUE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(true, entorno.BOOLEAN)

	case sintacticoR_FALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(331)
			p.Match(sintacticoR_FALSE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(false, entorno.BOOLEAN)

	case sintacticoID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(333)

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

	case 19:
		var t *LogicaContext = nil
		if localctx != nil {
			t = localctx.(*LogicaContext)
		}
		return p.Logica_Sempred(t, predIndex)

	case 20:
		var t *RelacionalContext = nil
		if localctx != nil {
			t = localctx.(*RelacionalContext)
		}
		return p.Relacional_Sempred(t, predIndex)

	case 21:
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
