package minimumincrementtomakearrayunique

import "testing"

type args struct {
	nums []int
}
type testCase struct {
	name string
	args args
	want int
}

func run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		if got := minIncrementForUnique(tt.args.nums); got != tt.want {
			t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
		}
	})
}

func Test_1(t *testing.T) {
	run(t, testCase{
		name: "Test_1",
		args: args{nums: []int{1, 2, 2}},
		want: 1,
	})
}

func Test_2(t *testing.T) {
	run(t, testCase{
		name: "Test_2",
		args: args{nums: []int{3, 2, 1, 2, 1, 7}},
		want: 6,
	})
}

func Test_3(t *testing.T) {
	run(t, testCase{
		name: "Test_3",
		args: args{nums: []int{2, 2, 2, 1}},
		want: 3,
	})
}
