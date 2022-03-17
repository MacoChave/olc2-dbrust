package entorno

type TipoDato int

const (
	INTEGER TipoDato = iota
	FLOAT
	CHAR
	BOOLEAN
	STRING
	NULL
	VOID
)

type Valor struct {
	Tipo  TipoDato
	Valor interface{}
}
