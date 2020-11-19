package q777

import "testing"

func Test_canTransform(t *testing.T) {
	caseTable := []struct {
		args struct {
			start string
			end   string
		}
		expect bool
	}{
		{
			args: struct {
				start string
				end   string
			}{
				start: "XLLR",
				end:   "LLXR",
			},
			expect: true,
		},
		{
			args: struct {
				start string
				end   string
			}{
				start: "RXXLRXRXL",
				end:   "XRLXXRRLX",
			},
			expect: true,
		},
		{
			args: struct {
				start string
				end   string
			}{
				start: "LLR",
				end:   "RRL",
			},
			expect: false,
		},
		{
			args: struct {
				start string
				end   string
			}{
				start: "LXXLXRLXXL",
				end:   "XLLXRXLXLX",
			},
			expect: false,
		},
		{
			args: struct {
				start string
				end   string
			}{
				start: "XRRXRX",
				end:   "RXLRRX",
			},
			expect: false,
		},
		{
			args: struct {
				start string
				end   string
			}{
				start: "RLX",
				end:   "XLR",
			},
			expect: false,
		},
	}
	for _, testCase := range caseTable {
		r := canTransform(testCase.args.start, testCase.args.end)
		if r != testCase.expect {
			t.Error("start:", testCase.args.start, " end:", testCase.args.end, " expect:", testCase.expect)
		}
	}
}
