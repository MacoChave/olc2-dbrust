package expresion

import "db_rust/analizador/entorno"

type Identificador struct {
	Identificador string
}

func NewIdentificador(id string) Identificador {
	return Identificador{
		Identificador: id,
	}
}

func (id Identificador) GetValor(ent entorno.Entorno) entorno.Valor {
	var idEncontrado bool = ent.ExistSimbolo(id.Identificador)

	if !idEncontrado {
		return entorno.Valor{
			Valor: nil,
			Tipo:  entorno.NULL,
		}
	}

	simbolo := ent.GetSimbolo(id.Identificador)
	return entorno.Valor{
		Valor: simbolo.Valor,
		Tipo:  simbolo.Tipo,
	}
}
