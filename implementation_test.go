package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPostfixToLisp(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"4 2 -", "(- 4 2)", false},
		{"3 2 ^", "(pow 3 2)", false},
		{"5 4 2 - 3 2 ^ * +", "(+ 5 (* (- 4 2) (pow 3 2)))", false},
		{"5", "5", false},
		{"", "", true},
		{"4 2", "", true},
		{"4 +", "", true},
		{"2 3 2 ^ *", "(* 2 (pow 3 2))", false},
		{"2 3 2 5 ^ *", "", true},
	}
	for _, tt := range tests {
		result, err := convertPostfixToLisp(tt.input)
		if tt.hasError {
			assert.Error(t, err, "Expected an error for input: %s", tt.input)
		} else {
			assert.NoError(t, err, "Did not expect an error for input: %s", tt.input)
			assert.Equal(t, tt.expected, result, "Mismatch for input: %s", tt.input)
		}
	}
}
