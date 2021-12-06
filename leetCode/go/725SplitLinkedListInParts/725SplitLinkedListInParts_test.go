package splitlinkedlistinparts

import (
	"reflect"
	"testing"

	. "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{1, 2, 3})
	expect := [][]int{{1}, {2}, {3}, {}, {}}
	k := len(expect)
	// act
	result := splitListToParts(head, k)
	// assert
	assert.Equal(t, k, len(result))
	actual := [][]int{}
	for _, v := range result {
		t := LinkedListToSlice(v)
		actual = append(actual, t)
	}
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	head := SliceToLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expect := [][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}}
	k := len(expect)
	// act
	result := splitListToParts(head, k)
	// assert
	assert.Equal(t, k, len(result))
	actual := [][]int{}
	for _, v := range result {
		t := LinkedListToSlice(v)
		actual = append(actual, t)
	}
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}
