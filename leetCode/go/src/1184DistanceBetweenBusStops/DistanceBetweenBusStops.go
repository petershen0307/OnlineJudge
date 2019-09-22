package q1184

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	betweenStartDest := false
	distanceBetween := 0
	distanceOutside := 0
	if start > destination {
		destination, start = start, destination
	}
	for i, d := range distance {
		if i == start {
			betweenStartDest = true
		}
		if i == destination {
			betweenStartDest = false
		}
		if betweenStartDest {
			distanceBetween += d
		} else {
			distanceOutside += d
		}
	}
	if distanceBetween > distanceOutside {
		return distanceOutside
	}
	return distanceBetween
}
