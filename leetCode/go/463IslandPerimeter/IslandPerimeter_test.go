package IslandPerimeter

import "testing"

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Example", args{[][]int{}}, 16},
	}
	tests[0].args.grid = append(tests[0].args.grid, []int{0, 1, 0, 0})
	tests[0].args.grid = append(tests[0].args.grid, []int{1, 1, 1, 0})
	tests[0].args.grid = append(tests[0].args.grid, []int{0, 1, 0, 0})
	tests[0].args.grid = append(tests[0].args.grid, []int{1, 1, 0, 0})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
