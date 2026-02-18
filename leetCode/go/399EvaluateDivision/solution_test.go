package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	expected := []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}
	assert.Equal(t, expected, calcEquation(equations, values, queries))
}

func Test_2(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	expected := []float64{3.75000, 0.40000, 5.00000, 0.20000}
	assert.Equal(t, expected, calcEquation(equations, values, queries))
}

func Test_3(t *testing.T) {
	equations := [][]string{{"a", "b"}}
	values := []float64{0.5}
	queries := [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}
	expected := []float64{0.50000, 2.00000, -1.00000, -1.00000}
	assert.Equal(t, expected, calcEquation(equations, values, queries))
}
