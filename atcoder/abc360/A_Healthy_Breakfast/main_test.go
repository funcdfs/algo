package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: AtCoder - AtCoder Beginner Contest 360 A - A Healthy Breakfast
// URL: https://atcoder.jp/contests/abc360/tasks/abc360_a
// Time: 2025-02-12 15:44:47

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "RSM\n",
			want:  "Yes\n",
		},
		{
			name:  "case2",
			input: "SMR\n",
			want:  "No\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
