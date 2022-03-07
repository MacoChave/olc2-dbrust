package expresion

import "db_rust/analizador/entorno"

type Primitivo struct {
	Valor interface{}
	Tipo  entorno.TipoDato
}

func NewPrimitivo(valor interface{}, tipo entorno.TipoDato) Primitivo {
	return Primitivo{
		Valor: valor,
		Tipo:  tipo,
	}
}

func (prim Primitivo) GetValor(ent entorno.Entorno) entorno.Valor {
	return entorno.Valor{
		Tipo:  prim.Tipo,
		Valor: prim.Valor,
	}
}
