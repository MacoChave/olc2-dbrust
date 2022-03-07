package entorno

type Entorno struct {
	Nombre    string
	Anterior  *Entorno
	Tabla     map[string]interface{}
	Funciones map[string]interface{}
}

func NewEntorno(nombre string, anterior *Entorno) Entorno {
	tabla := make(map[string]interface{})
	funciones := make(map[string]interface{})

	ent := Entorno{
		Nombre:    nombre,
		Anterior:  anterior,
		Tabla:     tabla,
		Funciones: funciones,
	}

	return ent
}

func (ent *Entorno) ExistSimbolo(id string) bool {
	for entActual := ent; entActual != nil; entActual = entActual.Anterior {
		for key := range entActual.Tabla {
			if key == id {
				return true
			}
		}
	}
	return false
}

func (ent *Entorno) AddSimbolo(id string, simbolo Simbolo) {
	ent.Tabla[id] = simbolo
}

func (ent *Entorno) GetSimbolo(id string) Simbolo {
	for entActual := ent; entActual != nil; entActual = entActual.Anterior {
		for key, simbolo := range entActual.Tabla {
			if key == id {
				return simbolo.(Simbolo)
			}
		}
	}

	var simbolo Simbolo
	return simbolo
}

func (ent *Entorno) UpdateSimbolo(id string, simbolo Simbolo) {
	// TODO: Get symbol and update
	for entActual := ent; entActual != nil; entActual = entActual.Anterior {
		for key := range entActual.Tabla {
			if key == id {
				entActual.Tabla[id] = simbolo
				return
			}
		}
	}
}
