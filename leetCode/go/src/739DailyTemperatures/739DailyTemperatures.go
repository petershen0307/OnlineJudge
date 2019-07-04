import "math"
func dailyTemperatures2(T []int) []int {
    tIndex := make([]int, 101) //[30, 100]
    for i := 0; i < len(tIndex); i++ {
        tIndex[i] = -1
    }
    r := []int {}
    for i := len(T)-1; i >= 0; i-- {
        minIndex := math.MaxInt64
        for j := T[i] + 1; j < len(tIndex); j++ {
            if tIndex[j] != -1 && minIndex > tIndex[j] && minIndex > i {
                minIndex = tIndex[j]
            }
        }
        if math.MaxInt64 == minIndex {
            r = append([]int{0}, r...)
        } else {
            r = append([]int{minIndex - i}, r...)
        }
        tIndex[T[i]] = i
    }
    return r
}

func dailyTemperatures(T []int) []int {
    r := []int {}
    for i, v := range T {
        found := false
        for j := i + 1; j < len(T); j++ {
            if T[j] > v {
                found = true
                r = append(r, j - i)
                break
            }
        }
        if !found {
            r = append(r, 0)
        }
    }
    return r
}
