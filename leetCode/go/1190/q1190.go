package q1190

func reverseParentheses(s string) string {
	stack := make([]rune, 0)
	for _, char := range s {
		if char == ')' {
			// pop to '('
			// push back to stack
			pop(&stack)
			continue
		}
		stack = append(stack, char)
	}
	return string(stack)
}

func pop(stack *[]rune) {
	a := make([]rune, 0)
	for {
		r := (*stack)[len(*stack)-1]
		(*stack) = (*stack)[:len(*stack)-1]
		if r == '(' {
			*stack = append(*stack, a...)
			break
		}
		a = append(a, r)
	}
}
