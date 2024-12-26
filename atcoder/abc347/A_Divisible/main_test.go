package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: AtCoder - AtCoder Beginner Contest 347 A - Divisible
// URL: https://atcoder.jp/contests/abc347/tasks/abc347_a
// Time: 2024-12-26 16:44:00

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "5 2\n2 5 6 7 10\n",
			want:  "1 3 5\n",
		},
		{
			name:  "case2",
			input: "3 1\n3 4 7\n",
			want:  "3 4 7\n",
		},
		{
			name:  "case3",
			input: "5 10\n50 51 54 60 65\n",
			want:  "5 6\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
