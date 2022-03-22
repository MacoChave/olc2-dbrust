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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 84, 350,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3,
	57, 10, 3, 12, 3, 14, 3, 60, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 6, 6, 77, 10, 6,
	13, 6, 14, 6, 78, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 5, 7, 104, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 121, 10, 9, 12, 9,
	14, 9, 124, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 156, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 168, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 205, 10, 13, 3, 14, 6, 14, 208, 10, 14, 13, 14, 14, 14, 209, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 231, 10,
	17, 13, 17, 14, 17, 232, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 268, 10, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 7, 22, 284, 10, 22, 12, 22, 14, 22, 287, 11, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 298, 10, 23, 12, 23,
	14, 23, 301, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 316, 10, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 328, 10,
	24, 12, 24, 14, 24, 331, 11, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26,
	348, 10, 26, 3, 26, 2, 6, 16, 42, 44, 46, 27, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 2,
	5, 3, 2, 16, 21, 3, 2, 13, 15, 3, 2, 11, 12, 2, 357, 2, 52, 3, 2, 2, 2,
	4, 58, 3, 2, 2, 2, 6, 63, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 76, 3, 2,
	2, 2, 12, 103, 3, 2, 2, 2, 14, 105, 3, 2, 2, 2, 16, 111, 3, 2, 2, 2, 18,
	155, 3, 2, 2, 2, 20, 167, 3, 2, 2, 2, 22, 169, 3, 2, 2, 2, 24, 204, 3,
	2, 2, 2, 26, 207, 3, 2, 2, 2, 28, 213, 3, 2, 2, 2, 30, 221, 3, 2, 2, 2,
	32, 230, 3, 2, 2, 2, 34, 236, 3, 2, 2, 2, 36, 244, 3, 2, 2, 2, 38, 251,
	3, 2, 2, 2, 40, 267, 3, 2, 2, 2, 42, 269, 3, 2, 2, 2, 44, 288, 3, 2, 2,
	2, 46, 315, 3, 2, 2, 2, 48, 332, 3, 2, 2, 2, 50, 347, 3, 2, 2, 2, 52, 53,
	5, 4, 3, 2, 53, 54, 8, 2, 1, 2, 54, 3, 3, 2, 2, 2, 55, 57, 5, 6, 4, 2,
	56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3,
	2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 62, 8, 3, 1, 2, 62,
	5, 3, 2, 2, 2, 63, 64, 5, 8, 5, 2, 64, 65, 8, 4, 1, 2, 65, 7, 3, 2, 2,
	2, 66, 67, 7, 44, 2, 2, 67, 68, 7, 45, 2, 2, 68, 69, 7, 25, 2, 2, 69, 70,
	7, 26, 2, 2, 70, 71, 7, 31, 2, 2, 71, 72, 5, 10, 6, 2, 72, 73, 7, 32, 2,
	2, 73, 74, 8, 5, 1, 2, 74, 9, 3, 2, 2, 2, 75, 77, 5, 12, 7, 2, 76, 75,
	3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2,
	79, 80, 3, 2, 2, 2, 80, 81, 8, 6, 1, 2, 81, 11, 3, 2, 2, 2, 82, 83, 5,
	14, 8, 2, 83, 84, 7, 5, 2, 2, 84, 85, 8, 7, 1, 2, 85, 104, 3, 2, 2, 2,
	86, 87, 5, 18, 10, 2, 87, 88, 7, 5, 2, 2, 88, 89, 8, 7, 1, 2, 89, 104,
	3, 2, 2, 2, 90, 91, 5, 22, 12, 2, 91, 92, 7, 5, 2, 2, 92, 93, 8, 7, 1,
	2, 93, 104, 3, 2, 2, 2, 94, 95, 5, 24, 13, 2, 95, 96, 8, 7, 1, 2, 96, 104,
	3, 2, 2, 2, 97, 98, 5, 30, 16, 2, 98, 99, 8, 7, 1, 2, 99, 104, 3, 2, 2,
	2, 100, 101, 5, 38, 20, 2, 101, 102, 8, 7, 1, 2, 102, 104, 3, 2, 2, 2,
	103, 82, 3, 2, 2, 2, 103, 86, 3, 2, 2, 2, 103, 90, 3, 2, 2, 2, 103, 94,
	3, 2, 2, 2, 103, 97, 3, 2, 2, 2, 103, 100, 3, 2, 2, 2, 104, 13, 3, 2, 2,
	2, 105, 106, 7, 51, 2, 2, 106, 107, 7, 25, 2, 2, 107, 108, 5, 16, 9, 2,
	108, 109, 7, 26, 2, 2, 109, 110, 8, 8, 1, 2, 110, 15, 3, 2, 2, 2, 111,
	112, 8, 9, 1, 2, 112, 113, 5, 40, 21, 2, 113, 114, 8, 9, 1, 2, 114, 122,
	3, 2, 2, 2, 115, 116, 12, 4, 2, 2, 116, 117, 7, 4, 2, 2, 117, 118, 5, 40,
	21, 2, 118, 119, 8, 9, 1, 2, 119, 121, 3, 2, 2, 2, 120, 115, 3, 2, 2, 2,
	121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	17, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 126, 7, 36, 2, 2, 126, 127,
	7, 37, 2, 2, 127, 128, 7, 82, 2, 2, 128, 129, 7, 27, 2, 2, 129, 130, 5,
	20, 11, 2, 130, 131, 7, 6, 2, 2, 131, 132, 5, 40, 21, 2, 132, 133, 8, 10,
	1, 2, 133, 156, 3, 2, 2, 2, 134, 135, 7, 36, 2, 2, 135, 136, 7, 37, 2,
	2, 136, 137, 7, 82, 2, 2, 137, 138, 7, 6, 2, 2, 138, 139, 5, 40, 21, 2,
	139, 140, 8, 10, 1, 2, 140, 156, 3, 2, 2, 2, 141, 142, 7, 36, 2, 2, 142,
	143, 7, 82, 2, 2, 143, 144, 7, 27, 2, 2, 144, 145, 5, 20, 11, 2, 145, 146,
	7, 6, 2, 2, 146, 147, 5, 40, 21, 2, 147, 148, 8, 10, 1, 2, 148, 156, 3,
	2, 2, 2, 149, 150, 7, 36, 2, 2, 150, 151, 7, 82, 2, 2, 151, 152, 7, 6,
	2, 2, 152, 153, 5, 40, 21, 2, 153, 154, 8, 10, 1, 2, 154, 156, 3, 2, 2,
	2, 155, 125, 3, 2, 2, 2, 155, 134, 3, 2, 2, 2, 155, 141, 3, 2, 2, 2, 155,
	149, 3, 2, 2, 2, 156, 19, 3, 2, 2, 2, 157, 158, 7, 38, 2, 2, 158, 168,
	8, 11, 1, 2, 159, 160, 7, 39, 2, 2, 160, 168, 8, 11, 1, 2, 161, 162, 7,
	41, 2, 2, 162, 168, 8, 11, 1, 2, 163, 164, 7, 42, 2, 2, 164, 168, 8, 11,
	1, 2, 165, 166, 7, 40, 2, 2, 166, 168, 8, 11, 1, 2, 167, 157, 3, 2, 2,
	2, 167, 159, 3, 2, 2, 2, 167, 161, 3, 2, 2, 2, 167, 163, 3, 2, 2, 2, 167,
	165, 3, 2, 2, 2, 168, 21, 3, 2, 2, 2, 169, 170, 7, 82, 2, 2, 170, 171,
	7, 6, 2, 2, 171, 172, 5, 40, 21, 2, 172, 173, 8, 12, 1, 2, 173, 23, 3,
	2, 2, 2, 174, 175, 7, 63, 2, 2, 175, 176, 5, 40, 21, 2, 176, 177, 7, 31,
	2, 2, 177, 178, 5, 10, 6, 2, 178, 179, 7, 32, 2, 2, 179, 180, 8, 13, 1,
	2, 180, 205, 3, 2, 2, 2, 181, 182, 7, 63, 2, 2, 182, 183, 5, 40, 21, 2,
	183, 184, 7, 31, 2, 2, 184, 185, 5, 10, 6, 2, 185, 186, 7, 32, 2, 2, 186,
	187, 7, 64, 2, 2, 187, 188, 7, 31, 2, 2, 188, 189, 5, 10, 6, 2, 189, 190,
	7, 32, 2, 2, 190, 191, 8, 13, 1, 2, 191, 205, 3, 2, 2, 2, 192, 193, 7,
	63, 2, 2, 193, 194, 5, 40, 21, 2, 194, 195, 7, 31, 2, 2, 195, 196, 5, 10,
	6, 2, 196, 197, 7, 32, 2, 2, 197, 198, 5, 26, 14, 2, 198, 199, 7, 64, 2,
	2, 199, 200, 7, 31, 2, 2, 200, 201, 5, 10, 6, 2, 201, 202, 7, 32, 2, 2,
	202, 203, 8, 13, 1, 2, 203, 205, 3, 2, 2, 2, 204, 174, 3, 2, 2, 2, 204,
	181, 3, 2, 2, 2, 204, 192, 3, 2, 2, 2, 205, 25, 3, 2, 2, 2, 206, 208, 5,
	28, 15, 2, 207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 207, 3, 2,
	2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 8, 14, 1, 2,
	212, 27, 3, 2, 2, 2, 213, 214, 7, 64, 2, 2, 214, 215, 7, 63, 2, 2, 215,
	216, 5, 40, 21, 2, 216, 217, 7, 31, 2, 2, 217, 218, 5, 10, 6, 2, 218, 219,
	7, 32, 2, 2, 219, 220, 8, 15, 1, 2, 220, 29, 3, 2, 2, 2, 221, 222, 7, 65,
	2, 2, 222, 223, 5, 40, 21, 2, 223, 224, 7, 31, 2, 2, 224, 225, 5, 32, 17,
	2, 225, 226, 5, 36, 19, 2, 226, 227, 7, 32, 2, 2, 227, 228, 8, 16, 1, 2,
	228, 31, 3, 2, 2, 2, 229, 231, 5, 34, 18, 2, 230, 229, 3, 2, 2, 2, 231,
	232, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 234,
	3, 2, 2, 2, 234, 235, 8, 17, 1, 2, 235, 33, 3, 2, 2, 2, 236, 237, 5, 40,
	21, 2, 237, 238, 7, 9, 2, 2, 238, 239, 7, 31, 2, 2, 239, 240, 5, 12, 7,
	2, 240, 241, 7, 32, 2, 2, 241, 242, 7, 4, 2, 2, 242, 243, 8, 18, 1, 2,
	243, 35, 3, 2, 2, 2, 244, 245, 7, 10, 2, 2, 245, 246, 7, 9, 2, 2, 246,
	247, 7, 31, 2, 2, 247, 248, 5, 12, 7, 2, 248, 249, 7, 32, 2, 2, 249, 250,
	8, 19, 1, 2, 250, 37, 3, 2, 2, 2, 251, 252, 7, 68, 2, 2, 252, 253, 5, 40,
	21, 2, 253, 254, 7, 31, 2, 2, 254, 255, 5, 10, 6, 2, 255, 256, 7, 32, 2,
	2, 256, 257, 8, 20, 1, 2, 257, 39, 3, 2, 2, 2, 258, 259, 5, 42, 22, 2,
	259, 260, 8, 21, 1, 2, 260, 268, 3, 2, 2, 2, 261, 262, 5, 44, 23, 2, 262,
	263, 8, 21, 1, 2, 263, 268, 3, 2, 2, 2, 264, 265, 5, 46, 24, 2, 265, 266,
	8, 21, 1, 2, 266, 268, 3, 2, 2, 2, 267, 258, 3, 2, 2, 2, 267, 261, 3, 2,
	2, 2, 267, 264, 3, 2, 2, 2, 268, 41, 3, 2, 2, 2, 269, 270, 8, 22, 1, 2,
	270, 271, 5, 44, 23, 2, 271, 272, 8, 22, 1, 2, 272, 285, 3, 2, 2, 2, 273,
	274, 12, 5, 2, 2, 274, 275, 7, 23, 2, 2, 275, 276, 5, 42, 22, 6, 276, 277,
	8, 22, 1, 2, 277, 284, 3, 2, 2, 2, 278, 279, 12, 4, 2, 2, 279, 280, 7,
	22, 2, 2, 280, 281, 5, 42, 22, 5, 281, 282, 8, 22, 1, 2, 282, 284, 3, 2,
	2, 2, 283, 273, 3, 2, 2, 2, 283, 278, 3, 2, 2, 2, 284, 287, 3, 2, 2, 2,
	285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 43, 3, 2, 2, 2, 287, 285,
	3, 2, 2, 2, 288, 289, 8, 23, 1, 2, 289, 290, 5, 46, 24, 2, 290, 291, 8,
	23, 1, 2, 291, 299, 3, 2, 2, 2, 292, 293, 12, 4, 2, 2, 293, 294, 9, 2,
	2, 2, 294, 295, 5, 44, 23, 5, 295, 296, 8, 23, 1, 2, 296, 298, 3, 2, 2,
	2, 297, 292, 3, 2, 2, 2, 298, 301, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 299,
	300, 3, 2, 2, 2, 300, 45, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 302, 303, 8,
	24, 1, 2, 303, 304, 7, 12, 2, 2, 304, 305, 5, 40, 21, 2, 305, 306, 8, 24,
	1, 2, 306, 316, 3, 2, 2, 2, 307, 308, 5, 48, 25, 2, 308, 309, 8, 24, 1,
	2, 309, 316, 3, 2, 2, 2, 310, 311, 7, 25, 2, 2, 311, 312, 5, 40, 21, 2,
	312, 313, 7, 26, 2, 2, 313, 314, 8, 24, 1, 2, 314, 316, 3, 2, 2, 2, 315,
	302, 3, 2, 2, 2, 315, 307, 3, 2, 2, 2, 315, 310, 3, 2, 2, 2, 316, 329,
	3, 2, 2, 2, 317, 318, 12, 6, 2, 2, 318, 319, 9, 3, 2, 2, 319, 320, 5, 46,
	24, 7, 320, 321, 8, 24, 1, 2, 321, 328, 3, 2, 2, 2, 322, 323, 12, 5, 2,
	2, 323, 324, 9, 4, 2, 2, 324, 325, 5, 46, 24, 6, 325, 326, 8, 24, 1, 2,
	326, 328, 3, 2, 2, 2, 327, 317, 3, 2, 2, 2, 327, 322, 3, 2, 2, 2, 328,
	331, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 47, 3,
	2, 2, 2, 331, 329, 3, 2, 2, 2, 332, 333, 5, 50, 26, 2, 333, 334, 8, 25,
	1, 2, 334, 49, 3, 2, 2, 2, 335, 336, 7, 78, 2, 2, 336, 348, 8, 26, 1, 2,
	337, 338, 7, 79, 2, 2, 338, 348, 8, 26, 1, 2, 339, 340, 7, 81, 2, 2, 340,
	348, 8, 26, 1, 2, 341, 342, 7, 34, 2, 2, 342, 348, 8, 26, 1, 2, 343, 344,
	7, 35, 2, 2, 344, 348, 8, 26, 1, 2, 345, 346, 7, 82, 2, 2, 346, 348, 8,
	26, 1, 2, 347, 335, 3, 2, 2, 2, 347, 337, 3, 2, 2, 2, 347, 339, 3, 2, 2,
	2, 347, 341, 3, 2, 2, 2, 347, 343, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 348,
	51, 3, 2, 2, 2, 19, 58, 78, 103, 122, 155, 167, 204, 209, 232, 267, 283,
	285, 299, 315, 327, 329, 347,
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
	"sentIf", "listaElseIf", "elseIf", "sentMatch", "casosMatch", "casoMatch",
	"matchDefecto", "sentWhile", "exp", "logica", "relacional", "aritmetica",
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
	sintacticoRULE_sentIf         = 11
	sintacticoRULE_listaElseIf    = 12
	sintacticoRULE_elseIf         = 13
	sintacticoRULE_sentMatch      = 14
	sintacticoRULE_casosMatch     = 15
	sintacticoRULE_casoMatch      = 16
	sintacticoRULE_matchDefecto   = 17
	sintacticoRULE_sentWhile      = 18
	sintacticoRULE_exp            = 19
	sintacticoRULE_logica         = 20
	sintacticoRULE_relacional     = 21
	sintacticoRULE_aritmetica     = 22
	sintacticoRULE_exp_valor      = 23
	sintacticoRULE_primitivo      = 24
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
		p.SetState(50)

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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sintacticoR_FN {
		{
			p.SetState(53)

			var _x = p.Procedimiento()

			localctx.(*ProcedimientosContext)._procedimiento = _x
		}
		localctx.(*ProcedimientosContext).proc = append(localctx.(*ProcedimientosContext).proc, localctx.(*ProcedimientosContext)._procedimiento)

		p.SetState(58)
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
		p.SetState(61)

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
		p.SetState(64)
		p.Match(sintacticoR_FN)
	}
	{
		p.SetState(65)
		p.Match(sintacticoR_MAIN)
	}
	{
		p.SetState(66)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(67)
		p.Match(sintacticoS_CPAR)
	}
	{
		p.SetState(68)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(69)

		var _x = p.Instrucciones()

		localctx.(*PrincipalContext)._instrucciones = _x
	}
	{
		p.SetState(70)
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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(sintacticoR_LET-34))|(1<<(sintacticoR_PRINTLN-34))|(1<<(sintacticoR_IF-34))|(1<<(sintacticoR_MATCH-34)))) != 0) || _la == sintacticoR_WHILE || _la == sintacticoID {
		{
			p.SetState(73)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).ins = append(localctx.(*InstruccionesContext).ins, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(76)
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

	// Get_sentIf returns the _sentIf rule contexts.
	Get_sentIf() ISentIfContext

	// Get_sentMatch returns the _sentMatch rule contexts.
	Get_sentMatch() ISentMatchContext

	// Get_sentWhile returns the _sentWhile rule contexts.
	Get_sentWhile() ISentWhileContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_sentIf sets the _sentIf rule contexts.
	Set_sentIf(ISentIfContext)

	// Set_sentMatch sets the _sentMatch rule contexts.
	Set_sentMatch(ISentMatchContext)

	// Set_sentWhile sets the _sentWhile rule contexts.
	Set_sentWhile(ISentWhileContext)

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
	_sentIf      ISentIfContext
	_sentMatch   ISentMatchContext
	_sentWhile   ISentWhileContext
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

func (s *InstruccionContext) Get_sentIf() ISentIfContext { return s._sentIf }

func (s *InstruccionContext) Get_sentMatch() ISentMatchContext { return s._sentMatch }

func (s *InstruccionContext) Get_sentWhile() ISentWhileContext { return s._sentWhile }

func (s *InstruccionContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_sentIf(v ISentIfContext) { s._sentIf = v }

func (s *InstruccionContext) Set_sentMatch(v ISentMatchContext) { s._sentMatch = v }

func (s *InstruccionContext) Set_sentWhile(v ISentWhileContext) { s._sentWhile = v }

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

func (s *InstruccionContext) SentIf() ISentIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentIfContext)
}

func (s *InstruccionContext) SentMatch() ISentMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentMatchContext)
}

func (s *InstruccionContext) SentWhile() ISentWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentWhileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentWhileContext)
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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)

			var _x = p.Imprimir()

			localctx.(*InstruccionContext)._imprimir = _x
		}
		{
			p.SetState(81)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_imprimir().GetInstr()

	case sintacticoR_LET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(85)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracion().GetInstr()

	case sintacticoID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(89)
			p.Match(sintacticoS_PTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_asignacion().GetInstr()

	case sintacticoR_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)

			var _x = p.SentIf()

			localctx.(*InstruccionContext)._sentIf = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_sentIf().GetInstr()

	case sintacticoR_MATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(95)

			var _x = p.SentMatch()

			localctx.(*InstruccionContext)._sentMatch = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_sentMatch().GetInstr()

	case sintacticoR_WHILE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)

			var _x = p.SentWhile()

			localctx.(*InstruccionContext)._sentWhile = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_sentWhile().GetInstr()

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
		p.SetState(103)
		p.Match(sintacticoR_PRINTLN)
	}
	{
		p.SetState(104)
		p.Match(sintacticoS_APAR)
	}
	{
		p.SetState(105)

		var _x = p.lista_exp(0)

		localctx.(*ImprimirContext)._lista_exp = _x
	}
	{
		p.SetState(106)
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
		p.SetState(110)

		var _x = p.Exp()

		localctx.(*Lista_expContext)._exp = _x
	}
	localctx.(*Lista_expContext).lista.Add(localctx.(*Lista_expContext).Get_exp().GetVal())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(120)
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
			p.SetState(113)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(114)
				p.Match(sintacticoS_COMA)
			}
			{
				p.SetState(115)

				var _x = p.Exp()

				localctx.(*Lista_expContext)._exp = _x
			}

			localctx.(*Lista_expContext).GetLISTA().GetLista().Add(localctx.(*Lista_expContext).Get_exp().GetVal())
			localctx.(*Lista_expContext).lista = localctx.(*Lista_expContext).GetLISTA().GetLista()

		}
		p.SetState(122)
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

	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(124)
			p.Match(sintacticoR_MUT)
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

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(133)
			p.Match(sintacticoR_MUT)
		}
		{
			p.SetState(134)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(135)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(136)

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

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(140)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(141)
			p.Match(sintacticoS_DOSPT)
		}
		{
			p.SetState(142)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(143)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(144)

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

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
			p.Match(sintacticoR_LET)
		}
		{
			p.SetState(148)

			var _m = p.Match(sintacticoID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(149)
			p.Match(sintacticoS_ASIGNAR)
		}
		{
			p.SetState(150)

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

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(sintacticoR_INT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.INTEGER

	case sintacticoR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.Match(sintacticoR_FLOAT)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.FLOAT

	case sintacticoR_CHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.Match(sintacticoR_CHAR)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.CHAR

	case sintacticoR_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
			p.Match(sintacticoR_STRING)
		}
		localctx.(*Tipo_datoContext).tipo = entorno.STRING

	case sintacticoR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(163)
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
		p.SetState(167)

		var _m = p.Match(sintacticoID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(168)
		p.Match(sintacticoS_ASIGNAR)
	}
	{
		p.SetState(169)

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

// ISentIfContext is an interface to support dynamic dispatch.
type ISentIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_exp returns the _exp rule contexts.
	Get_exp() IExpContext

	// GetS_then returns the s_then rule contexts.
	GetS_then() IInstruccionesContext

	// GetS_else returns the s_else rule contexts.
	GetS_else() IInstruccionesContext

	// Get_listaElseIf returns the _listaElseIf rule contexts.
	Get_listaElseIf() IListaElseIfContext

	// Set_exp sets the _exp rule contexts.
	Set_exp(IExpContext)

	// SetS_then sets the s_then rule contexts.
	SetS_then(IInstruccionesContext)

	// SetS_else sets the s_else rule contexts.
	SetS_else(IInstruccionesContext)

	// Set_listaElseIf sets the _listaElseIf rule contexts.
	Set_listaElseIf(IListaElseIfContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsSentIfContext differentiates from other interfaces.
	IsSentIfContext()
}

type SentIfContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruccion
	_exp         IExpContext
	s_then       IInstruccionesContext
	s_else       IInstruccionesContext
	_listaElseIf IListaElseIfContext
}

func NewEmptySentIfContext() *SentIfContext {
	var p = new(SentIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_sentIf
	return p
}

func (*SentIfContext) IsSentIfContext() {}

func NewSentIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentIfContext {
	var p = new(SentIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_sentIf

	return p
}

func (s *SentIfContext) GetParser() antlr.Parser { return s.parser }

func (s *SentIfContext) Get_exp() IExpContext { return s._exp }

func (s *SentIfContext) GetS_then() IInstruccionesContext { return s.s_then }

func (s *SentIfContext) GetS_else() IInstruccionesContext { return s.s_else }

func (s *SentIfContext) Get_listaElseIf() IListaElseIfContext { return s._listaElseIf }

func (s *SentIfContext) Set_exp(v IExpContext) { s._exp = v }

func (s *SentIfContext) SetS_then(v IInstruccionesContext) { s.s_then = v }

func (s *SentIfContext) SetS_else(v IInstruccionesContext) { s.s_else = v }

func (s *SentIfContext) Set_listaElseIf(v IListaElseIfContext) { s._listaElseIf = v }

func (s *SentIfContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *SentIfContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *SentIfContext) R_IF() antlr.TerminalNode {
	return s.GetToken(sintacticoR_IF, 0)
}

func (s *SentIfContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *SentIfContext) AllS_ALLAV() []antlr.TerminalNode {
	return s.GetTokens(sintacticoS_ALLAV)
}

func (s *SentIfContext) S_ALLAV(i int) antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, i)
}

func (s *SentIfContext) AllS_CLLAV() []antlr.TerminalNode {
	return s.GetTokens(sintacticoS_CLLAV)
}

func (s *SentIfContext) S_CLLAV(i int) antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, i)
}

func (s *SentIfContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *SentIfContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *SentIfContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(sintacticoR_ELSE, 0)
}

func (s *SentIfContext) ListaElseIf() IListaElseIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaElseIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaElseIfContext)
}

func (s *SentIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterSentIf(s)
	}
}

func (s *SentIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitSentIf(s)
	}
}

func (p *sintactico) SentIf() (localctx ISentIfContext) {
	this := p
	_ = this

	localctx = NewSentIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, sintacticoRULE_sentIf)

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

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(173)

			var _x = p.Exp()

			localctx.(*SentIfContext)._exp = _x
		}
		{
			p.SetState(174)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(175)

			var _x = p.Instrucciones()

			localctx.(*SentIfContext).s_then = _x
		}
		{
			p.SetState(176)
			p.Match(sintacticoS_CLLAV)
		}

		localctx.(*SentIfContext).instr = flujo.NewIf(localctx.(*SentIfContext).Get_exp().GetVal(), localctx.(*SentIfContext).GetS_then().GetLista(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(180)

			var _x = p.Exp()

			localctx.(*SentIfContext)._exp = _x
		}
		{
			p.SetState(181)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(182)

			var _x = p.Instrucciones()

			localctx.(*SentIfContext).s_then = _x
		}
		{
			p.SetState(183)
			p.Match(sintacticoS_CLLAV)
		}
		{
			p.SetState(184)
			p.Match(sintacticoR_ELSE)
		}
		{
			p.SetState(185)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(186)

			var _x = p.Instrucciones()

			localctx.(*SentIfContext).s_else = _x
		}
		{
			p.SetState(187)
			p.Match(sintacticoS_CLLAV)
		}

		localctx.(*SentIfContext).instr = flujo.NewIf(localctx.(*SentIfContext).Get_exp().GetVal(), localctx.(*SentIfContext).GetS_then().GetLista(), nil, localctx.(*SentIfContext).GetS_else().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(190)
			p.Match(sintacticoR_IF)
		}
		{
			p.SetState(191)

			var _x = p.Exp()

			localctx.(*SentIfContext)._exp = _x
		}
		{
			p.SetState(192)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(193)

			var _x = p.Instrucciones()

			localctx.(*SentIfContext).s_then = _x
		}
		{
			p.SetState(194)
			p.Match(sintacticoS_CLLAV)
		}
		{
			p.SetState(195)

			var _x = p.ListaElseIf()

			localctx.(*SentIfContext)._listaElseIf = _x
		}
		{
			p.SetState(196)
			p.Match(sintacticoR_ELSE)
		}
		{
			p.SetState(197)
			p.Match(sintacticoS_ALLAV)
		}
		{
			p.SetState(198)

			var _x = p.Instrucciones()

			localctx.(*SentIfContext).s_else = _x
		}
		{
			p.SetState(199)
			p.Match(sintacticoS_CLLAV)
		}

		localctx.(*SentIfContext).instr = flujo.NewIf(localctx.(*SentIfContext).Get_exp().GetVal(), localctx.(*SentIfContext).GetS_then().GetLista(), localctx.(*SentIfContext).Get_listaElseIf().GetLista(), localctx.(*SentIfContext).GetS_else().GetLista())

	}

	return localctx
}

// IListaElseIfContext is an interface to support dynamic dispatch.
type IListaElseIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_elseIf returns the _elseIf rule contexts.
	Get_elseIf() IElseIfContext

	// Set_elseIf sets the _elseIf rule contexts.
	Set_elseIf(IElseIfContext)

	// GetIns returns the ins rule context list.
	GetIns() []IElseIfContext

	// SetIns sets the ins rule context list.
	SetIns([]IElseIfContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListaElseIfContext differentiates from other interfaces.
	IsListaElseIfContext()
}

type ListaElseIfContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lista   *arrayList.List
	_elseIf IElseIfContext
	ins     []IElseIfContext
}

func NewEmptyListaElseIfContext() *ListaElseIfContext {
	var p = new(ListaElseIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_listaElseIf
	return p
}

func (*ListaElseIfContext) IsListaElseIfContext() {}

func NewListaElseIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaElseIfContext {
	var p = new(ListaElseIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_listaElseIf

	return p
}

func (s *ListaElseIfContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaElseIfContext) Get_elseIf() IElseIfContext { return s._elseIf }

func (s *ListaElseIfContext) Set_elseIf(v IElseIfContext) { s._elseIf = v }

func (s *ListaElseIfContext) GetIns() []IElseIfContext { return s.ins }

func (s *ListaElseIfContext) SetIns(v []IElseIfContext) { s.ins = v }

func (s *ListaElseIfContext) GetLista() *arrayList.List { return s.lista }

func (s *ListaElseIfContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListaElseIfContext) AllElseIf() []IElseIfContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfContext)(nil)).Elem())
	var tst = make([]IElseIfContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfContext)
		}
	}

	return tst
}

func (s *ListaElseIfContext) ElseIf(i int) IElseIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfContext)
}

func (s *ListaElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaElseIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterListaElseIf(s)
	}
}

func (s *ListaElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitListaElseIf(s)
	}
}

func (p *sintactico) ListaElseIf() (localctx IListaElseIfContext) {
	this := p
	_ = this

	localctx = NewListaElseIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, sintacticoRULE_listaElseIf)
	localctx.(*ListaElseIfContext).lista = arrayList.New()

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
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(204)

				var _x = p.ElseIf()

				localctx.(*ListaElseIfContext)._elseIf = _x
			}
			localctx.(*ListaElseIfContext).ins = append(localctx.(*ListaElseIfContext).ins, localctx.(*ListaElseIfContext)._elseIf)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	LISTA := localctx.(*ListaElseIfContext).GetIns()
	for _, elemento := range LISTA {
		localctx.(*ListaElseIfContext).lista.Add(elemento.GetInstr())
	}

	return localctx
}

// IElseIfContext is an interface to support dynamic dispatch.
type IElseIfContext interface {
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

	// IsElseIfContext differentiates from other interfaces.
	IsElseIfContext()
}

type ElseIfContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruccion
	_exp           IExpContext
	_instrucciones IInstruccionesContext
}

func NewEmptyElseIfContext() *ElseIfContext {
	var p = new(ElseIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_elseIf
	return p
}

func (*ElseIfContext) IsElseIfContext() {}

func NewElseIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfContext {
	var p = new(ElseIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_elseIf

	return p
}

func (s *ElseIfContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfContext) Get_exp() IExpContext { return s._exp }

func (s *ElseIfContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *ElseIfContext) Set_exp(v IExpContext) { s._exp = v }

func (s *ElseIfContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *ElseIfContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *ElseIfContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *ElseIfContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(sintacticoR_ELSE, 0)
}

func (s *ElseIfContext) R_IF() antlr.TerminalNode {
	return s.GetToken(sintacticoR_IF, 0)
}

func (s *ElseIfContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ElseIfContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *ElseIfContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *ElseIfContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *ElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterElseIf(s)
	}
}

func (s *ElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitElseIf(s)
	}
}

func (p *sintactico) ElseIf() (localctx IElseIfContext) {
	this := p
	_ = this

	localctx = NewElseIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, sintacticoRULE_elseIf)

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
		p.SetState(211)
		p.Match(sintacticoR_ELSE)
	}
	{
		p.SetState(212)
		p.Match(sintacticoR_IF)
	}
	{
		p.SetState(213)

		var _x = p.Exp()

		localctx.(*ElseIfContext)._exp = _x
	}
	{
		p.SetState(214)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(215)

		var _x = p.Instrucciones()

		localctx.(*ElseIfContext)._instrucciones = _x
	}
	{
		p.SetState(216)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*ElseIfContext).instr = flujo.NewIf(localctx.(*ElseIfContext).Get_exp().GetVal(), localctx.(*ElseIfContext).Get_instrucciones().GetLista(), nil, nil)

	return localctx
}

// ISentMatchContext is an interface to support dynamic dispatch.
type ISentMatchContext interface {
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

	// IsSentMatchContext differentiates from other interfaces.
	IsSentMatchContext()
}

type SentMatchContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	instr         interfaces.Instruccion
	_exp          IExpContext
	_casosMatch   ICasosMatchContext
	_matchDefecto IMatchDefectoContext
}

func NewEmptySentMatchContext() *SentMatchContext {
	var p = new(SentMatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_sentMatch
	return p
}

func (*SentMatchContext) IsSentMatchContext() {}

func NewSentMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentMatchContext {
	var p = new(SentMatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_sentMatch

	return p
}

func (s *SentMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *SentMatchContext) Get_exp() IExpContext { return s._exp }

func (s *SentMatchContext) Get_casosMatch() ICasosMatchContext { return s._casosMatch }

func (s *SentMatchContext) Get_matchDefecto() IMatchDefectoContext { return s._matchDefecto }

func (s *SentMatchContext) Set_exp(v IExpContext) { s._exp = v }

func (s *SentMatchContext) Set_casosMatch(v ICasosMatchContext) { s._casosMatch = v }

func (s *SentMatchContext) Set_matchDefecto(v IMatchDefectoContext) { s._matchDefecto = v }

func (s *SentMatchContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *SentMatchContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *SentMatchContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(sintacticoR_MATCH, 0)
}

func (s *SentMatchContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *SentMatchContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *SentMatchContext) CasosMatch() ICasosMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasosMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasosMatchContext)
}

func (s *SentMatchContext) MatchDefecto() IMatchDefectoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchDefectoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchDefectoContext)
}

func (s *SentMatchContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *SentMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterSentMatch(s)
	}
}

func (s *SentMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitSentMatch(s)
	}
}

func (p *sintactico) SentMatch() (localctx ISentMatchContext) {
	this := p
	_ = this

	localctx = NewSentMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, sintacticoRULE_sentMatch)

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
		p.SetState(219)
		p.Match(sintacticoR_MATCH)
	}
	{
		p.SetState(220)

		var _x = p.Exp()

		localctx.(*SentMatchContext)._exp = _x
	}
	{
		p.SetState(221)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(222)

		var _x = p.CasosMatch()

		localctx.(*SentMatchContext)._casosMatch = _x
	}
	{
		p.SetState(223)

		var _x = p.MatchDefecto()

		localctx.(*SentMatchContext)._matchDefecto = _x
	}
	{
		p.SetState(224)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*SentMatchContext).instr = flujo.NewSentMatch(localctx.(*SentMatchContext).Get_exp().GetVal(), localctx.(*SentMatchContext).Get_casosMatch().GetLista(), localctx.(*SentMatchContext).Get_matchDefecto().GetCaso())

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
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(sintacticoS_RESTA-10))|(1<<(sintacticoS_APAR-10))|(1<<(sintacticoR_TRUE-10))|(1<<(sintacticoR_FALSE-10)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(sintacticoNUMERO-76))|(1<<(sintacticoDECIMAL-76))|(1<<(sintacticoCADENA-76))|(1<<(sintacticoID-76)))) != 0) {
		{
			p.SetState(227)

			var _x = p.CasoMatch()

			localctx.(*CasosMatchContext)._casoMatch = _x
		}
		localctx.(*CasosMatchContext).caso = append(localctx.(*CasosMatchContext).caso, localctx.(*CasosMatchContext)._casoMatch)

		p.SetState(230)
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
		p.SetState(234)

		var _x = p.Exp()

		localctx.(*CasoMatchContext)._exp = _x
	}
	{
		p.SetState(235)
		p.Match(sintacticoS_MATCH_RET)
	}
	{
		p.SetState(236)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(237)

		var _x = p.Instruccion()

		localctx.(*CasoMatchContext)._instruccion = _x
	}
	{
		p.SetState(238)
		p.Match(sintacticoS_CLLAV)
	}
	{
		p.SetState(239)
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
		p.SetState(242)
		p.Match(sintacticoS_MATCH_DEFAULT)
	}
	{
		p.SetState(243)
		p.Match(sintacticoS_MATCH_RET)
	}
	{
		p.SetState(244)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(245)

		var _x = p.Instruccion()

		localctx.(*MatchDefectoContext)._instruccion = _x
	}
	{
		p.SetState(246)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*MatchDefectoContext).caso = flujo.NewCaseMatch(nil, localctx.(*MatchDefectoContext).Get_instruccion().GetInstr())

	return localctx
}

// ISentWhileContext is an interface to support dynamic dispatch.
type ISentWhileContext interface {
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

	// IsSentWhileContext differentiates from other interfaces.
	IsSentWhileContext()
}

type SentWhileContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruccion
	_exp           IExpContext
	_instrucciones IInstruccionesContext
}

func NewEmptySentWhileContext() *SentWhileContext {
	var p = new(SentWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sintacticoRULE_sentWhile
	return p
}

func (*SentWhileContext) IsSentWhileContext() {}

func NewSentWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentWhileContext {
	var p = new(SentWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sintacticoRULE_sentWhile

	return p
}

func (s *SentWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *SentWhileContext) Get_exp() IExpContext { return s._exp }

func (s *SentWhileContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *SentWhileContext) Set_exp(v IExpContext) { s._exp = v }

func (s *SentWhileContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *SentWhileContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *SentWhileContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *SentWhileContext) R_WHILE() antlr.TerminalNode {
	return s.GetToken(sintacticoR_WHILE, 0)
}

func (s *SentWhileContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *SentWhileContext) S_ALLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_ALLAV, 0)
}

func (s *SentWhileContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *SentWhileContext) S_CLLAV() antlr.TerminalNode {
	return s.GetToken(sintacticoS_CLLAV, 0)
}

func (s *SentWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.EnterSentWhile(s)
	}
}

func (s *SentWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sintacticoListener); ok {
		listenerT.ExitSentWhile(s)
	}
}

func (p *sintactico) SentWhile() (localctx ISentWhileContext) {
	this := p
	_ = this

	localctx = NewSentWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, sintacticoRULE_sentWhile)

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
		p.SetState(249)
		p.Match(sintacticoR_WHILE)
	}
	{
		p.SetState(250)

		var _x = p.Exp()

		localctx.(*SentWhileContext)._exp = _x
	}
	{
		p.SetState(251)
		p.Match(sintacticoS_ALLAV)
	}
	{
		p.SetState(252)

		var _x = p.Instrucciones()

		localctx.(*SentWhileContext)._instrucciones = _x
	}
	{
		p.SetState(253)
		p.Match(sintacticoS_CLLAV)
	}

	localctx.(*SentWhileContext).instr = flujo.NewSentWhile(localctx.(*SentWhileContext).Get_exp().GetVal(), localctx.(*SentWhileContext).Get_instrucciones().GetLista())

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
	p.EnterRule(localctx, 38, sintacticoRULE_exp)

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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)

			var _x = p.logica(0)

			localctx.(*ExpContext)._logica = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_logica().GetVal()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)

			var _x = p.relacional(0)

			localctx.(*ExpContext)._relacional = _x
		}
		localctx.(*ExpContext).val = localctx.(*ExpContext).Get_relacional().GetVal()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)

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
	_startState := 40
	p.EnterRecursionRule(localctx, 40, sintacticoRULE_logica, _p)

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
		p.SetState(268)

		var _x = p.relacional(0)

		localctx.(*LogicaContext)._relacional = _x
	}
	localctx.(*LogicaContext).val = localctx.(*LogicaContext).Get_relacional().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogicaContext(p, _parentctx, _parentState)
				localctx.(*LogicaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_logica)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(272)

					var _m = p.Match(sintacticoS_AND)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(273)

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
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(277)

					var _m = p.Match(sintacticoS_OR)

					localctx.(*LogicaContext).op = _m
				}
				{
					p.SetState(278)

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
		p.SetState(285)
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
	_startState := 42
	p.EnterRecursionRule(localctx, 42, sintacticoRULE_relacional, _p)
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
		p.SetState(287)

		var _x = p.aritmetica(0)

		localctx.(*RelacionalContext)._aritmetica = _x
	}
	localctx.(*RelacionalContext).val = localctx.(*RelacionalContext).Get_aritmetica().GetVal()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(297)
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
			p.SetState(290)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(291)

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
				p.SetState(292)

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
		p.SetState(299)
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, sintacticoRULE_aritmetica, _p)
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
	p.SetState(313)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoS_RESTA:
		{
			p.SetState(301)
			p.Match(sintacticoS_RESTA)
		}
		{
			p.SetState(302)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		localctx.(*AritmeticaContext).val = expresion.NewOperacion(localctx.(*AritmeticaContext).Get_exp().GetVal(), nil, "-", true)

	case sintacticoR_TRUE, sintacticoR_FALSE, sintacticoNUMERO, sintacticoDECIMAL, sintacticoCADENA, sintacticoID:
		{
			p.SetState(305)

			var _x = p.Exp_valor()

			localctx.(*AritmeticaContext)._exp_valor = _x
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp_valor().GetVal()

	case sintacticoS_APAR:
		{
			p.SetState(308)
			p.Match(sintacticoS_APAR)
		}
		{
			p.SetState(309)

			var _x = p.Exp()

			localctx.(*AritmeticaContext)._exp = _x
		}
		{
			p.SetState(310)
			p.Match(sintacticoS_CPAR)
		}
		localctx.(*AritmeticaContext).val = localctx.(*AritmeticaContext).Get_exp().GetVal()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAritmeticaContext(p, _parentctx, _parentState)
				localctx.(*AritmeticaContext).izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, sintacticoRULE_aritmetica)
				p.SetState(315)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(316)

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
					p.SetState(317)

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
				p.SetState(320)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(321)

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
					p.SetState(322)

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
		p.SetState(329)
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
	p.EnterRule(localctx, 46, sintacticoRULE_exp_valor)

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
		p.SetState(330)

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
	p.EnterRule(localctx, 48, sintacticoRULE_primitivo)

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

	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sintacticoNUMERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)

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
			p.SetState(335)

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
			p.SetState(337)

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
			p.SetState(339)
			p.Match(sintacticoR_TRUE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(true, entorno.BOOLEAN)

	case sintacticoR_FALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(341)
			p.Match(sintacticoR_FALSE)
		}

		localctx.(*PrimitivoContext).val = expresion.NewPrimitivo(false, entorno.BOOLEAN)

	case sintacticoID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(343)

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

	case 20:
		var t *LogicaContext = nil
		if localctx != nil {
			t = localctx.(*LogicaContext)
		}
		return p.Logica_Sempred(t, predIndex)

	case 21:
		var t *RelacionalContext = nil
		if localctx != nil {
			t = localctx.(*RelacionalContext)
		}
		return p.Relacional_Sempred(t, predIndex)

	case 22:
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
