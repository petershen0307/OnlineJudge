package validmountainarray

func validMountainArray(arr []int) bool {
	type stateType int
	const (
		none stateType = iota
		up
		down
	)
	state := none
	prev := arr[0]
	for _, v := range arr[1:] {
		if state == up {
			if prev < v {
				prev = v
				continue
			}
			if prev > v {
				state = down
				prev = v
				continue
			} else {
				state = none
				break
			}
		}
		if state == down {
			if prev > v {
				prev = v
				continue
			}
			if prev <= v {
				state = none
				break
			}
		}
		if state == none {
			if prev < v {
				prev = v
				state = up
			} else {
				break
			}
		}
	}
	return state == down
}
