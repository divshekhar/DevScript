package object

import "fmt"

// Define the Boolean struct
type Boolean struct {
	Value bool
}

// Implement the Inspect method for the Boolean struct
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Implement the Type method for the Boolean struct
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
