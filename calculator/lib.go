package calculator

import (
	"errors"
	"strconv"
	"strings"
)

// 演算子の実行
func execOpecode(stack *[]int, op func(int, int) int) error {
	if len(*stack) < 2 {
		return errors.New("stack-length less than 2")
	}

	a, b := (*stack)[len(*stack)-2], (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-2]
	*stack = append(*stack, op(a, b))

	return nil
}

// 逆ポーランド記法電卓
// in : "1 2 + 3 4 + *"
// out: 21, nil
func Calc(exp string) (int, error) {
	stack := []int{}

	for _, op := range strings.Split(exp, " ") {
		switch op {
		case "+":
			err := execOpecode(&stack, func(a, b int) int { return a + b })
			if err != nil {
				return 0, err
			}
		case "/":
			err := execOpecode(&stack, func(a, b int) int { return a / b })
			if err != nil {
				return 0, err
			}
		case "*":
			err := execOpecode(&stack, func(a, b int) int { return a * b })
			if err != nil {
				return 0, err
			}
		default:
			num, err := strconv.Atoi(op)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("failed result")
	}

	return stack[0], nil
}
