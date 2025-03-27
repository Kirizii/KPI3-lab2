package lab2

import (
	"bufio"
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	if !scanner.Scan() {
		return fmt.Errorf("empty input")
	}
	expression := scanner.Text()

	result, err := ConvertPostfixToLisp(expression)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(ch.Output, result)
	return err
}
