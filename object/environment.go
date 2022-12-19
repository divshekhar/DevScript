package object

// Environment is a map of string to Object
//
//	{
//		"foo": Integer{Value: 1},
//		"bar": Integer{Value: 2},
//	}
type Environment struct {
	// store is a map of string to Object
	store map[string]Object
}

// NewEnvironment returns a new Environment
func NewEnvironment() *Environment {
	store := make(map[string]Object)
	return &Environment{store: store}
}

// Get returns the Object associated with the given name
func (env *Environment) Get(name string) (Object, bool) {
	obj, ok := env.store[name]
	return obj, ok
}

// Set sets the Object associated with the given name
func (env *Environment) Set(name string, val Object) Object {
	env.store[name] = val
	return val
}
