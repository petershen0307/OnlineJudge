package basiccalculatorii

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	operators := "+-*/"
	// initial with '+' operation to add first digit number
	operation := "+"
	numStack := []int{}
	num := ""
	for i, char := range s {
		if char != rune(' ') && !strings.ContainsRune(operators, char) {
			num += string(char)
		}
		if strings.ContainsRune(operators, char) || i == len(s)-1 {
			v, _ := strconv.Atoi(num)
			num = ""
			switch operation {
			case "+":
				numStack = append(numStack, v)
			case "-":
				numStack = append(numStack, -1*v)
			case "*":
				// pop value from stack
				temp := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				numStack = append(numStack, temp*v)
			case "/":
				// pop value from stack
				temp := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				numStack = append(numStack, temp/v)
			}
			operation = string(char)
		}
	}
	sum := 0
	// 目前在 numStack 中只剩正數與負數, 所以最後再把 numStack 全部加起來就是結果
	for _, v := range numStack {
		sum += v
	}
	return sum
}
