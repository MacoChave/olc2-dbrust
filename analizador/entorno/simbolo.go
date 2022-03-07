package entorno

import "github.com/colegno/arraylist"

type Simbolo struct {
	Linea      int
	Columna    int
	Id         string
	Valor      interface{}
	Tipo       TipoDato
	IsConst    bool
	IsFunc     bool
	Parametros *arraylist.List
}

func NewSimboloId(linea int, columna int, id string, isConst bool, valor interface{}, tipo TipoDato) Simbolo {
	simbolo := Simbolo{
		Linea:   linea,
		Columna: columna,
		Id:      id,
		IsConst: isConst,
		Valor:   valor,
		Tipo:    tipo,
		IsFunc:  false,
	}
	return simbolo
}

func NewSimboloFunc(linea int, columna int, id string, tipo TipoDato, parametros *arraylist.List) Simbolo {
	simbolo := Simbolo{
		Linea:      linea,
		Columna:    columna,
		Id:         id,
		Tipo:       tipo,
		Parametros: parametros,
		IsFunc:     true,
		IsConst:    false,
		Valor:      nil,
	}
	return simbolo
}
