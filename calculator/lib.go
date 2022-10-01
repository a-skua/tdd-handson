package calculator

import (
	"errors"
	"strconv"
	"strings"
)

// Stack
type stack []int

func (s stack) length() int {
	return len(s)
}

func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) push(v int) {
	*s = append(*s, v)
}

// 演算子の実行
func execOpecode(stack *stack, op func(int, int) int) error {
	if stack.length() < 2 {
		return errors.New("stack-length less than 2")
	}

	b, a := stack.pop(), stack.pop()
	stack.push(op(a, b))

	return nil
}

// 逆ポーランド記法電卓
// in : "1 2 + 3 4 + *"
// out: 21, nil
func Calc(exp string) (int, error) {
	stack := stack{}

	for _, op := range strings.Split(exp, " ") {
		err := func() error {
			switch op {
			case "+":
				return execOpecode(&stack, func(a, b int) int { return a + b })
			// case "-":
			// 	return execOpecode(&stack, func(a, b int) int { return a - b })
			case "/":
				return execOpecode(&stack, func(a, b int) int { return a / b })
			case "*":
				return execOpecode(&stack, func(a, b int) int { return a * b })
			default:
				num, err := strconv.Atoi(op)
				if err != nil {
					return err
				}
				stack.push(num)
				return nil
			}
		}()
		if err != nil {
			return 0, err
		}
	}

	if stack.length() != 1 {
		return 0, errors.New("failed result")
	}

	return stack.pop(), nil
}
