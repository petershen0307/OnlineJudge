package binarywatch

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	/*
		從 0 數到 1023 2^10-1
		判斷 bit 數是否符合要求
		判斷前 4 bit 與後 6 bit 是否符合要求 (0-11), (0-59)
		最後轉換成字串
	*/
	results := []string{}
	for i := range 1024 {
		// count bits
		c := 0
		for b := 1; b <= (1 << 9); b = b << 1 {
			if (b & i) != 0 {
				c++
			}
			if c > turnedOn {
				break
			}
		}
		if c != turnedOn {
			continue
		}
		if r := hourMinute(i); r != "" {
			results = append(results, r)
		}
	}
	return results
}

func hourMinute(i int) string {
	// bit 9 8 7 6 mask 0b1111000000
	hourBitsMask := 0b1111000000
	hour := i & hourBitsMask
	// should count from 0
	hour = hour >> 6
	// bit 5 4 3 2 1 0 mask 0b111111
	minuteBitsMask := 0b111111
	minute := i & minuteBitsMask
	if hour >= 0 && hour <= 11 && minute >= 0 && minute <= 59 {
		// the result
		return fmt.Sprintf("%d:%02d", hour, minute)
	}
	return ""
}
