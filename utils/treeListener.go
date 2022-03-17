package utils

import (
	"db_rust/analizador/ast"
	"db_rust/analizador/parser"
)

type TreeListener struct {
	*parser.BasesintacticoListener
	Ast ast.Ast
}

func NewTreeListener() *TreeListener {
	return new(TreeListener)
}

func (listener *TreeListener) ExitStart(context *parser.StartContext) {
	listener.Ast = context.GetRoot()
}
