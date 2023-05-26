package calculator

import (
	"strconv"
	"strings"
)

func BasicCalculator(operator string, a, b int) int {
	result := 0

	switch operator {
	case SUM_OPERATOR:
		result = a + b
	case SUBTRACT_OPERATOR:
		result = a - b
	case MULTIPLY_OPERATOR:
		result = a * b
	case DIVIDE_OPERATOR:
		result = a / b
	}

	return result
}

func GenerateInput(input string) (string, int, int, error) {
	array := strings.Split(input, " ")

	var (
		op         string
		num1, num2 int
		err        error
	)
	if len(array) > 2 {
		num1, err = strconv.Atoi(array[0])
		if err != nil {
			return "", 0, 0, err
		}

		num2, err = strconv.Atoi(array[2])
		if err != nil {
			return "", 0, 0, err
		}

		op = array[1]
	}

	return op, num1, num2, nil
}
