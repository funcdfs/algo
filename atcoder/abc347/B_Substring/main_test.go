package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: AtCoder - AtCoder Beginner Contest 347 B - Substring
// URL: https://atcoder.jp/contests/abc347/tasks/abc347_b
// Time: 2024-12-26 16:49:03

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "yay\n",
			want:  "5\n",
		},
		{
			name:  "case2",
			input: "aababc\n",
			want:  "17\n",
		},
		{
			name:  "case3",
			input: "abracadabra\n",
			want:  "54\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
