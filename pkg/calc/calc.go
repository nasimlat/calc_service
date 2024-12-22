package calc

import (
    "fmt"
)

func IsDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	r := rune(s[0])
	return '0' <= r && r <= '9'
}

func Pop(stack []string) ([]string, string) {
	last_index := len(stack) - 1
	last := stack[last_index]
	new_stack := stack[:last_index]
	return new_stack, last
}

func PopNum(stack []float64) ([]float64, float64) {
	last_index := len(stack) - 1
	last := stack[last_index]
	new_stack := stack[:last_index]
	return new_stack, last
}

func Top(stack []string) (string) {
	last_index := len(stack) - 1
    last := stack[last_index]
    return last
}

func isEmpty(stack []string) bool {
	return len(stack) == 0
}

func Calc(expression string) (float64, error) {
	numbers := []float64{}
	operators := []string{}

	for _, c := range expression {
		char := string(c)
		if IsDigit(char) {
			num, _ := Atoica(char)
			numbers = append(numbers, num)
		} else if char == "(" {
			operators = append(operators, char)
		} else if char == ")" {
			for !isEmpty(operators) && Top(operators) != "(" {
                var op string
                operators, op = Pop(operators)
                if len(numbers) < 2 {
                    return 0, fmt.Errorf("not enough operands for operator %s", op)
                }

				var a, b float64
                numbers, b = PopNum(numbers)
                numbers, a = PopNum(numbers)

				result, err := Operation(a, b, op)
                if err!= nil {
                    return 0, err
                }
                numbers = append(numbers, result)
            }
			if isEmpty(operators) {
				return 0, fmt.Errorf("expected")
			}
			operators, _ = Pop(operators) // удаляем ()
		} else {
			
			for !isEmpty(operators) && (Top(operators) == "*" || Top(operators) == "/") {
				var op string
				operators, op = Pop(operators)
				if len(numbers) < 2 {
					return 0, fmt.Errorf("not enough operands for operator %s", op)
				}

				var a, b float64
				numbers, b = PopNum(numbers)
				numbers, a = PopNum(numbers)

				result, err := Operation(a, b, op)
				if err != nil {
					return 0, err
				}
				numbers = append(numbers, result)
			}
			operators = append(operators, char)
		}
	}

	for !isEmpty(operators) {
		var op string
		operators, op = Pop(operators)
		if len(numbers) < 2 {
			return 0, fmt.Errorf("not enough operands for operator %s", op)
		}

		var a, b float64
        numbers, b = PopNum(numbers)
        numbers, a = PopNum(numbers)

		result, err := Operation(a, b, op)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, result)
	}

	if len(numbers) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return numbers[0], nil
}


func Atoica(s string) (float64, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("пустая строка")
	}

	c := rune(s[0])
	if c < '0' || c > '9' {
		return 0, fmt.Errorf("неверный формат числа")
	}
	return float64(c - '0'), nil
}

func Operation(a, b float64, operator string) (float64, error) {
	switch operator {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
			return 0, fmt.Errorf("на ноль делить нельзя")
		}
		return a / b, nil
    default:
        return 0, fmt.Errorf("неверный оператор")
    }
}
