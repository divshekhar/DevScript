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
	// points to the outer environment
	outer *Environment
}

// NewEnvironment returns a new Environment
func NewEnvironment() *Environment {
	store := make(map[string]Object)
	return &Environment{store: store, outer: nil}
}

// returns a new Environment with the given store
func NewEnclosedEnvironment(parentEnv *Environment) *Environment {
	env := NewEnvironment()
	env.outer = parentEnv
	return env
}

// Get returns the Object associated with the given name
func (env *Environment) Get(name string) (Object, bool) {
	obj, ok := env.store[name]
	if !ok && env.outer != nil {
		obj, ok = env.outer.Get(name)
	}
	return obj, ok
}

// Set sets the Object associated with the given name
func (env *Environment) Set(name string, val Object) Object {
	env.store[name] = val
	return val
}
