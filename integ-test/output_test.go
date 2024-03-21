package integtest

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	operand1 string
	operator string
	operand2 string
	result   string
	err      bool
}

func TestSunnyDayOutput(t *testing.T) {
	tests := []testData{
		{"1", "+", "1", "1 + 1 = 2.000000\n", false},
		{"2.5", "/", "2", "2.5 / 2 = 1.250000\n", false},
		{"2.5", "-", "2", "2.5 - 2 = 0.500000\n", false},
		{"2.5", "*", "2", "2.5 * 2 = 5.000000\n", false},
	}
	runTestCalculator(t, tests)
}

func TestErrorOutput(t *testing.T) {
	tests := []testData{
		{"1", "ab", "3", "invalid operator\n", true},
		{"1", "+", "x", "invalid operand\n", true},
		{"1", "/", "0", "division by zero is not allowed\n", true},
	}
	runTestCalculator(t, tests)

}

func runTestCalculator(t *testing.T, tests []testData) {
	for _, test := range tests {
		var actual string
		cmd := exec.Command("../test-calculator", test.operand1, test.operator, test.operand2)
		var output bytes.Buffer
		if test.err {
			cmd.Stderr = &output
		} else {
			cmd.Stdout = &output
		}
		err := cmd.Run()
		if test.err {
			assert.NotNil(t, err)
			actual = stripTimestamp(output.String())
		} else {
			assert.Nil(t, err)
			actual = output.String()
		}
		assert.Equal(t, test.result, actual)

	}
}

func stripTimestamp(in string) string {
	parsed := strings.Split(in, " ")
	return strings.Join(parsed[2:], " ")
}
