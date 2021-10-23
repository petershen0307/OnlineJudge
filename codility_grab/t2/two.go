package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

const insertedVal = 5

func Solution(N int) int {
	n := []int{}
	r := N
	if N < 0 {
		r = N * -1
	}
	for r > 0 {
		n = append(n, r%10)
		r = r / 10
	}
	if N == 0 {
		n = append(n, 0)
	}

	result := 0
	if N >= 0 {
		// positive
		result = check(true, n)
	} else {
		// negative
		result = check(false, n)
	}
	return result
}

func check(isPositive bool, n []int) int {
	index := 0
	for i := len(n) - 1; i >= 0; i-- {
		if isPositive {
			if n[i] < insertedVal {
				index = i
				break
			}
		} else {
			if n[i] > insertedVal {
				index = i
				break
			}
		}
	}
	t := append([]int{}, n[:index+1]...)
	t = append(t, insertedVal)
	t = append(t, n[index+1:]...)
	r := 0
	exp := 1
	for _, v := range t {
		r += v * exp
		exp *= 10
	}
	if !isPositive {
		r *= -1
	}
	return r
}
