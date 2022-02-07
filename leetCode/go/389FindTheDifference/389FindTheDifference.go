package findthedifference

func findTheDifference(s string, t string) byte {
	var sum int
	for i := 0; i != len(t); i++ {
		sum += int(t[i])
	}
	for i := 0; i != len(s); i++ {
		sum -= int(s[i])
	}
	return byte(sum)
}
