package grumpybookstoreowner

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	satisfyCustomer := 0
	maxConsecutiveGrumpyTime := 0
	for i, v := range customers {
		if grumpy[i] == 0 {
			satisfyCustomer += v
		} else {
			temp := 0
			for j := i; j < len(customers) && j < i+minutes; j++ {
				if grumpy[j] == 1 {
					temp += customers[j]
				}
			}
			if temp > maxConsecutiveGrumpyTime {
				maxConsecutiveGrumpyTime = temp
			}
		}
	}
	return satisfyCustomer + maxConsecutiveGrumpyTime
}
