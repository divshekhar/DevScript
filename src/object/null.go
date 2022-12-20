package object

// Define the Null struct
type Null struct{}

// Implement the Inspect method for the Null struct
func (n *Null) Inspect() string {
	return "null"
}

// Implement the Type method for the Null struct
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}
