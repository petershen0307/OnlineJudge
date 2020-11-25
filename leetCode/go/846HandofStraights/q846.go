package q846

import "sort"

func isNStraightHand(hand []int, W int) bool {
	if W == 1 {
		return true
	}
	sort.Ints(hand)
	w := 0
	c := 0
	l := 0
	for l < len(hand) && c < len(hand) {
		if hand[l] == hand[c] {
			c++
			continue
		}
		if hand[l]+1 == hand[c] {
			// pop hand[l]
			if l == 0 {
				hand = hand[1:]
			} else {
				hand = append(hand[0:l], hand[l+1:]...)
			}
			l = c - 1
			w++
		} else {
			return false
		}
		if w == W-1 {
			// pop hand[l]
			if l == 0 {
				hand = hand[1:]
			} else {
				hand = append(hand[0:l], hand[l+1:]...)
			}
			l = 0
			c = 0
			w = 0
		}
	}
	return len(hand) == 0 && w == 0
}
