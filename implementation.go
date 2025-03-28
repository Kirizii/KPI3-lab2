package lab2

import (
	"fmt"
	"strings"
)

func ConvertPostfixToLisp(expression string) (string, error) {
	tokens := strings.Fields(expression)
	stack := []string{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid postfix expression")
			}
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if token == "^" {
				token = "pow"
			}
			stack = append(stack, fmt.Sprintf("(%s %s %s)", token, a, b))
		default:
			stack = append(stack, token)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("invalid postfix expression")
	}
	return stack[0], nil
}
