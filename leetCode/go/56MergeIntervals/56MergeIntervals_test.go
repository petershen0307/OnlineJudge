package mergeintervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	expect := [][]int{{1, 6}, {8, 10}, {15, 18}}
	// act
	actual := merge(intervals)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	intervals := [][]int{{1, 4}, {4, 5}}
	expect := [][]int{{1, 5}}
	// act
	actual := merge(intervals)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	intervals := [][]int{{1, 4}}
	expect := [][]int{{1, 4}}
	// act
	actual := merge(intervals)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	intervals := [][]int{{1, 4}, {2, 4}}
	expect := [][]int{{1, 4}}
	// act
	actual := merge(intervals)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_5(t *testing.T) {
	// arrange
	intervals := [][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}
	expect := [][]int{{2, 4}, {5, 5}}
	// act
	actual := merge(intervals)
	// assert
	assert.Equal(t, expect, actual)
}
