package findtheduplicatenumber

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	nums := []int{1, 3, 4, 2, 2}
	expect := 2
	// act
	actual := findDuplicate(nums)
	// assert
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	nums := []int{3, 1, 3, 4, 2}
	expect := 3
	// act
	actual := findDuplicate(nums)
	// assert
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	nums := []int{2, 2, 2, 2, 2}
	expect := 2
	// act
	actual := findDuplicate(nums)
	// assert
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}
