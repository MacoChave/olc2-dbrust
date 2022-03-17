package analizador

type ErrorSemantico struct {
	Linea   int
	Columna int
	Mensaje string
}

func NewErrorSemantico(linea int, columna int, mensaje string) ErrorSemantico {
	return ErrorSemantico{
		Linea:   linea,
		Columna: columna,
		Mensaje: mensaje,
	}
}

var ListaErrores []ErrorSemantico
