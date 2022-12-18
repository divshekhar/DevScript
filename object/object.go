package object

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
