package general

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListGenerator(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	head := SliceToLinkedList(source)
	actual := LinkedListToSlice(head)
	assert.True(t, reflect.DeepEqual(source, actual))
	assert.Equal(t, source, actual)
}
