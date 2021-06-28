package numberofprovinces

func findCircleNum(isConnected [][]int) int {
	tracedMap := make([]int, len(isConnected))

	provinces := 0
	for i := 0; i < len(isConnected); i++ {
		if tracedMap[i] == 0 {
			provinces++
			trace(isConnected, i, tracedMap)
		}
	}
	return provinces

}

func trace(isConnected [][]int, i int, tracedMap []int) {
	if i >= len(isConnected) {
		return
	}
	for j := 0; j < len(isConnected); j++ {
		if tracedMap[j] != 0 {
			continue
		}
		if isConnected[i][j] == 1 {
			// connected
			tracedMap[j] = 1
			trace(isConnected, j, tracedMap)
		}
	}
}
