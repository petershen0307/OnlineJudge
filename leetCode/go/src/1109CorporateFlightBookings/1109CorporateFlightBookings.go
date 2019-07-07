func corpFlightBookings(bookings [][]int, n int) []int {
    r := make([]int, n)
    for _, v := range bookings {
        for j := v[0] - 1; j < v[1]; j++ {
            r[j] += v[2]
        }
    }
    return r
}
