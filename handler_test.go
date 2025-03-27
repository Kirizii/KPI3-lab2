package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_ValidExpression(t *testing.T) {
	input := strings.NewReader("3 2 ^")
	output := bytes.Buffer{}

	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}

	err := handler.Compute()
	assert.NoError(t, err)
	assert.Equal(t, "(pow 3 2)\n", output.String())
}

func TestComputeHandler_InvalidExpression(t *testing.T) {
	input := strings.NewReader("3 +")
	output := bytes.Buffer{}

	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}

	err := handler.Compute()
	assert.Error(t, err)
}
