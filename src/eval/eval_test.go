package eval

import (
	"devscript/src/lexer"
	"devscript/src/object"
	"devscript/src/parser"
	"testing"
)

// testEval is a helper function to evaluate the input string and return the
// evaluated object.
func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"-(-10)", 10},
		{"-(10)", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		// Zero values
		// TODO: To be implemented
		// {"0 == true", false},
		// {"0 == false", true},
		// {"0 != false", false},
		// {"0 != true", true},
		// {"true == 0", false},
		// {"false == 0", true},
		// {"false != 0", false},
		// {"true != 0", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, evaluated object.Object, expected bool) bool {
	result, ok := evaluated.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", evaluated, evaluated)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

// Function to test evaluation of assignment expressions
//
//	a = 5;	// Assignment expression
func TestAssignmentEvaluation(t *testing.T) {
	input := `
	var a = 5;
	a = 10;
	`
	evaluated := testEval(input)

	testIntegerObject(t, evaluated, 10)
}

// Function to test string evaluation
func TestStringEvaluation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"\"hello\"", "hello"},
		{"\"hello world\"", "hello world"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected)
	}
}

func TestStringConcatenation(t *testing.T) {
	test := []struct {
		input    string
		expected string
	}{
		{"\"Hello\" + \" \" + \"World!\"", "Hello World!"},
	}

	for _, tt := range test {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, "Hello World!")
	}
}

func TestStringComparison(t *testing.T) {
	test := []struct {
		input    string
		expected bool
	}{
		{"\"Hello\" == \"Hello\"", true},
		{"\"Hello\" != \"Hello\"", false},
		{"\"Hello\" == \"World\"", false},
		{"\"Hello\" != \"World\"", true},
	}

	for _, tt := range test {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testStringObject(t *testing.T, evaluated object.Object, expected string) bool {
	result, ok := evaluated.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", evaluated, evaluated)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s", result.Value, expected)
		return false
	}

	return true
}
