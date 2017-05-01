/*
Package HammingDistance :
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, calculate the Hamming distance.
Note: 0 ≤ x, y < 2^31.
*/
package HammingDistance

func hammingDistance(x int, y int) int {
	if x < 0 || y < 0 {
		return -1
	}
	if int64(x) >= int64(1<<31) || int64(y) >= int64(1<<31) {
		return -1
	}
	bitsWithXOR := x ^ y
	result := 0
	for iter := 1; iter <= x+y; iter <<= 1 {
		if bitsWithXOR&iter != 0 {
			result++
		}
	}
	return result
}
