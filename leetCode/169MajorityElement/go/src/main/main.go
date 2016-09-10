package main

import "fmt"

func majorityElement(nums []int) int {
	countMap := make(map[int]int)

	for _, key := range nums {
		_, exist := countMap[key]
		if !exist {
			countMap[key] = 1
		} else {
			countMap[key]++
		}
	}
	majoritySIze := (len(nums) / 2)
	fmt.Println(majoritySIze)
	for k, v := range countMap {
		fmt.Printf("key:%d value:%d\n", k, v)
		if v > majoritySIze {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{1, 4, 3, 4, 5, 5, 3, 3, 4, 4, 4, 4, 4}))
}
