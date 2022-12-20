package eval

import (
	"devscript/object"
	"testing"
)

func TestFunctionExpression(testing *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`
		func add(x, y) {
			return x + y;
		}
	
		add(2, 3);
		`, 5},

		{`
		func add(x, y, z) {
			return x + y + z;
		}
	
		add(2, 3, 5);
		`, 10},

		{`
		func doNothing(x, y, z) {
		}
	
		doNothing(2, 3, 5);
		`, nil},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		integer, ok := test.expected.(int)
		if ok {
			testIntegerObject(testing, evaluated, int64(integer))
		} else {
			testNullObject(testing, evaluated)
		}
	}
}

func TestFunctionObject(t *testing.T) {
	input := `func(x) { x + 2; };`

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var identity = func(x) { x; }; identity(5);", 5},
		{"var identity = func(x) { return x; }; identity(5);", 5},
		{"var double = func(x) { x * 2; }; double(5);", 10},
		{"var add = func(x, y) { x + y; }; add(5, 5);", 10},
		{"var add = func(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"func(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
