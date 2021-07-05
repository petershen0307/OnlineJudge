package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"

	"github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
)

type args struct {
	digits string
}
type testCase struct {
	name string
	args args
	want []string
}

func run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
		}
	})
}

func Test_empty(t *testing.T) {
	run(t, testCase{
		name: general.GetCallerName(),
		args: args{digits: ""},
		want: []string{},
	})
}

func Test_2(t *testing.T) {
	run(t, testCase{
		name: general.GetCallerName(),
		args: args{digits: "2"},
		want: []string{"a", "b", "c"},
	})
}

func Test_23(t *testing.T) {
	run(t, testCase{
		name: general.GetCallerName(),
		args: args{digits: "23"},
		want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	})
}
