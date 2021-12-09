package rotatearray

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	actual := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	expect := []int{5, 6, 7, 1, 2, 3, 4}
	// act
	rotate(actual, k)
	// assert
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}
