package binarynumberwithalternatingbits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	assert.True(t, hasAlternatingBits(5))
}

func Test_2(t *testing.T) {
	assert.False(t, hasAlternatingBits(7))
}

func Test_3(t *testing.T) {
	assert.True(t, hasAlternatingBits(1))
}

func Test_4(t *testing.T) {
	assert.False(t, hasAlternatingBits(0b100))
}
