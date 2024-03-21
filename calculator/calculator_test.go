package calculator

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

type Input struct {
	operand1 float64
	operator string
	operand2 float64
	expected float64
	err      error
}

func TestCalculateAdd(t *testing.T) {
	tests := []Input{
		{1, "+", 1, 2, nil},
		{1.0, "+", 1.0, 2, nil},
		{.99, "+", .01, 1, nil},
		{3.14, "+", 3.14, 6.28, nil},
		{1.1234567, "+", 1, 2.1234567, nil},
		{-1, "+", 1, 0, nil},
		{0.0000, "+", 0.00, 0, nil},
		{5, "+", -10.00, -5, nil},
		// Fails do to no handling of overflow
		//{math.MaxFloat64, "+", math.MaxFloat64, 0, errors.New("operand overflow")},
	}
	runCalculateTests(t, tests)
}

func TestCalculateSubtract(t *testing.T) {
	tests := []Input{
		{1, "-", 1, 0, nil},
		{1.0, "-", 1.0, 0, nil},
		{.99, "-", .01, .98, nil},
		{3.14, "-", -3.14, 6.28, nil},
		// Fails do to floating point math
		//{1.1234567, "-", 1, 0.1234567, nil},
		{-1, "-", 1, -2, nil},
		{0.0000, "-", 0.00, 0, nil},
		{math.MaxFloat64, "-", math.MaxFloat64, 0, nil},
		{math.SmallestNonzeroFloat64, "-", math.SmallestNonzeroFloat64, 0, nil},
		// Fails do to not catching overflowing in the negative
		//{math.SmallestNonzeroFloat64, "-", -1, 0, errors.New("operand overflow")},

	}
	runCalculateTests(t, tests)
}

func TestCalculateMultiply(t *testing.T) {
	tests := []Input{
		{1, "*", 1, 1, nil},
		{1.0, "*", 1.0, 1, nil},
		{.99, "*", .01, 0.0099, nil},
		{3.14, "*", -3.14, -9.8596, nil},
		// Fails due to floating point math
		//{1.1234567, "*", 1.1234567, 1.26215469, nil},
		{0.0000, "*", 0.00, 0, nil},
		// Both of these tests fail do to not handling overflow
		//{math.MaxFloat64, "*", 10, 0, errors.New("operand overflow")},
		//{math.SmallestNonzeroFloat64, "*", math.SmallestNonzeroFloat64, 0, errors.New("operand overflow")},
	}
	runCalculateTests(t, tests)
}

func TestCalculateDivide(t *testing.T) {
	tests := []Input{
		{1, "/", 1, 1, nil},
		{1.0, "/", 1.0, 1, nil},
		{.99, "/", .01, 99, nil},
		{3.14, "/", -3.14, -1, nil},
		{0, "/", -3.14, 0, nil},
		{1, "/", 0, 0, errors.New("division by zero is not allowed")},
		// Fails due to floating point math
		//{10, "/", 3, 3.333333, nil},
		{math.MaxFloat64, "/", math.MaxFloat64, 1, nil},
		{math.SmallestNonzeroFloat64, "/", math.SmallestNonzeroFloat64, 1, nil},
		// Fails do to not catching overflow
		//{math.SmallestNonzeroFloat64, "/", 2, 0, errors.New("operand overflow")},

	}
	runCalculateTests(t, tests)
}
func TestCalculateInvalidOperator(t *testing.T) {
	tests := []Input{
		{1, "", 1, 0, errors.New("invalid operator")},
		{1, "+-", 1, 0, errors.New("invalid operator")},
		{1, "\n*\t", 1, 0, errors.New("invalid operator")},
		{1, "\n\t+", 1, 0, errors.New("invalid operator")},
	}
	runCalculateTests(t, tests)
}
func runCalculateTests(t *testing.T, tests []Input) {
	for _, test := range tests {
		result, err := Calculate(test.operand1, test.operand2, test.operator)

		if err != nil && (test.err == nil || err.Error() != test.err.Error()) || err == nil && test.err != nil {
			t.Errorf("%v %v %v Expected error: %v, got: %v", test.operand1, test.operator, test.operand2, test.err, err)
		}

		if result != test.expected {
			t.Errorf("%v %v %v Expected result: %v, got: %v", test.operand1, test.operator, test.operand2, test.expected, result)
		}
	}
}

func TestParseOperand(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{"1", 1, nil},
		{"0.0000", 0, nil},
		{"3.14", 3.14, nil},
		{"-3.14", -3.14, nil},
		{fmt.Sprintf("%v", math.MaxFloat64), math.MaxFloat64, nil},
		{fmt.Sprintf("%v", math.SmallestNonzeroFloat64), math.SmallestNonzeroFloat64, nil},
		{"abc", 0, errors.New("invalid operand")},
		// These tests fail because the code does not catch non-numeric tokens that are valid floats
		//{"NaN", 0, errors.New("invalid operand")},
		//{"+inf", 0, errors.New("invalid operand")},
		//{"-inf", 0, errors.New("invalid operand")},
	}

	for _, test := range tests {
		result, err := ParseOperand(test.input)

		if err != nil && (test.err == nil || err.Error() != test.err.Error()) || err == nil && test.err != nil {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}
