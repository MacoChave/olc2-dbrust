package utils

import "github.com/antlr/antlr4/runtime/Go/antlr"

type CustomSyntaxError struct {
	Linea, Columna int
	Mensaje        string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []CustomSyntaxError
}

func (cerror *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	cerror.Errors = append(cerror.Errors, CustomSyntaxError{
		Linea:   line,
		Columna: column,
		Mensaje: msg,
	})
}
