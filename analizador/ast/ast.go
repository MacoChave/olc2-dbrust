package ast

import "github.com/colegno/arraylist"

type Ast struct {
	Instrucciones *arraylist.List
}

func NewAst(lista *arraylist.List) Ast {
	return Ast{
		Instrucciones: lista,
	}
}
