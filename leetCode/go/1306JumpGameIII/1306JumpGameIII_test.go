package jumpgameiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 5
	expect := true
	// act
	actual := canReach(arr, start)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 0
	expect := true
	// act
	actual := canReach(arr, start)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	arr := []int{3, 0, 2, 1, 2}
	start := 2
	expect := false
	// act
	actual := canReach(arr, start)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	arr := []int{}
	start := 2
	expect := false
	// act
	actual := canReach(arr, start)
	// assert
	assert.Equal(t, expect, actual)
}
