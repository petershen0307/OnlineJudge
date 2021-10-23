package solution

func Solution(A []int) int {
	a := [100000]int{}
	a[0] = 1
	count := 0
	for _, v := range A {
		a[v] = 1
		consequence := true
		for i, vv := range a {
			if vv == 0 && i < v {
				consequence = false
				break
			}
		}
		if consequence {
			count++
		}
	}
	return count
}
