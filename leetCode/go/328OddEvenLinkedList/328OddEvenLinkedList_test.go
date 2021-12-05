package oddevenlinkedlist

import (
	"reflect"
	"testing"

	. "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
	"github.com/stretchr/testify/assert"
)

func Test_oddEvenList1(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{1, 2, 3, 4, 5})
	expect := []int{1, 3, 5, 2, 4}
	// act
	result := oddEvenList(head)
	// assert
	actual := LinkedListToSlice(result)
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_oddEvenList2(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{2, 1, 3, 5, 6, 4, 7})
	expect := []int{2, 3, 6, 7, 1, 5, 4}
	// act
	result := oddEvenList(head)
	// assert
	actual := LinkedListToSlice(result)
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_oddEvenList_evenLength(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{2, 1, 3, 5, 6, 4})
	expect := []int{2, 3, 6, 1, 5, 4}
	// act
	result := oddEvenList(head)
	// assert
	actual := LinkedListToSlice(result)
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}
