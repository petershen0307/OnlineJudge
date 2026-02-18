package binarywatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	readBinaryWatch(1)
}

func Test_hour_minute(t *testing.T) {
	i := 0b1000000
	assert.Equal(t, "1:00", hourMinute(i))
}
