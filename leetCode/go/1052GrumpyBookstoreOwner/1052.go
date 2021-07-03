package grumpybookstoreowner

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	satisfyCustomer := 0
	lastConsecutiveGrumpyTime := 0
	// sum the first minutes consecutive grumpy time
	for i := 0; i < len(customers) && i < minutes; i++ {
		if grumpy[i] == 1 {
			lastConsecutiveGrumpyTime += customers[i]
		} else {
			satisfyCustomer += customers[i]
		}
	}
	maxConsecutiveGrumpyTime := lastConsecutiveGrumpyTime
	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 0 {
			satisfyCustomer += customers[i]
		}
		// slide window
		if grumpy[i-minutes] == 1 {
			lastConsecutiveGrumpyTime -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			lastConsecutiveGrumpyTime += customers[i]
		}
		if maxConsecutiveGrumpyTime < lastConsecutiveGrumpyTime {
			maxConsecutiveGrumpyTime = lastConsecutiveGrumpyTime
		}
	}
	return satisfyCustomer + maxConsecutiveGrumpyTime
}
