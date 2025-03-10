package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: AtCoder - AtCoder Beginner Contest 369 A - 369
// URL: https://atcoder.jp/contests/abc369/tasks/abc369_a
// Time: 2025-03-10 11:04:59

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "5 7\n",
			want:  "3\n",
		},
		{
			name:  "case2",
			input: "6 1\n",
			want:  "2\n",
		},
		{
			name:  "case3",
			input: "3 3\n",
			want:  "1\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
