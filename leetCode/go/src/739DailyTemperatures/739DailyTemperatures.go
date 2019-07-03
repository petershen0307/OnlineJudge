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
