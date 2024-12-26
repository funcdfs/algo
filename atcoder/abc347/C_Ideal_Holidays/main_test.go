package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: AtCoder - AtCoder Beginner Contest 347 C - Ideal Holidays
// URL: https://atcoder.jp/contests/abc347/tasks/abc347_c
// Time: 2024-12-26 17:02:48

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "case1",
			input: "3 2 5\n1 2 9\n",
			want:  "Yes\n",
		},
		{
			name:  "case2",
			input: "2 5 10\n10 15\n",
			want:  "No\n",
		},
		{
			name:  "case3",
			input: "4 347 347\n347 700 705 710\n",
			want:  "Yes\n",
		},
		{
			name:  "case4",
			input: "2 5 10\n11 15\n",
			want:  "Yes\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
