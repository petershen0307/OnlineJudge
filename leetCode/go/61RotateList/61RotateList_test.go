package rotatelist

import (
	"reflect"
	"testing"

	. "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{1, 2, 3, 4, 5})
	k := 2
	expect := []int{4, 5, 1, 2, 3}
	// act
	result := rotateRight(head, k)
	// assert
	actual := LinkedListToSlice(result)
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{1, 2, 3})
	k := 4
	expect := []int{3, 1, 2}
	// act
	result := rotateRight(head, k)
	// assert
	actual := LinkedListToSlice(result)
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{})
	k := 4
	expect := []int{}
	// act
	result := rotateRight(head, k)
	// assert
	actual := LinkedListToSlice(result)
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{5})
	k := 4
	expect := []int{5}
	// act
	result := rotateRight(head, k)
	// assert
	actual := LinkedListToSlice(result)
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}
