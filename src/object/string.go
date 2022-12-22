package object

import "fmt"

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return fmt.Sprintf("%s", s.Value)
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}
