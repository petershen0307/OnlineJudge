package grumpybookstoreowner

import (
	"testing"

	"github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
)

type args struct {
	customers []int
	grumpy    []int
	minutes   int
}
type testCase struct {
	name string
	args args
	want int
}

func run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes); got != tt.want {
			t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
		}
	})
}

func Test_1(t *testing.T) {
	run(t, testCase{
		name: general.GetCallerName(),
		args: args{
			customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
			grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
			minutes:   3,
		},
		want: 16,
	})
}

func Test_2(t *testing.T) {
	run(t, testCase{
		name: general.GetCallerName(),
		args: args{
			customers: []int{3},
			grumpy:    []int{1},
			minutes:   1,
		},
		want: 3,
	})
}
