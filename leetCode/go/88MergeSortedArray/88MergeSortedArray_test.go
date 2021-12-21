package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	expect := []int{1, 2, 2, 3, 5, 6}
	// act
	merge(nums1, m, nums2, n)
	// assert
	assert.Equal(t, expect, nums1)
}

func Test_2(t *testing.T) {
	// arrange
	nums1 := []int{1}
	m := 1
	nums2 := []int{}
	n := 0
	expect := []int{1}
	// act
	merge(nums1, m, nums2, n)
	// assert
	assert.Equal(t, expect, nums1)
}

func Test_3(t *testing.T) {
	// arrange
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	expect := []int{1}
	// act
	merge(nums1, m, nums2, n)
	// assert
	assert.Equal(t, expect, nums1)
}
