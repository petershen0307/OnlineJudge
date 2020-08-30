package main

import "fmt"

func findComplement(num int) int {
	result := 0
	for iter := 1; iter <= num; iter <<= 1 {
		result |= (iter ^ (num & iter))
	}
	return result
}

func main() {
	fmt.Printf("%d\n", findComplement(5))
	fmt.Printf("%d", findComplement(2))
}
