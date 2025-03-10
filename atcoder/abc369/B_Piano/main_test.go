package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: AtCoder - AtCoder Beginner Contest 369 B - Piano 3
// URL: https://atcoder.jp/contests/abc369/tasks/abc369_b
// Time: 2025-03-10 11:14:15

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "4\n3 L\n6 R\n9 L\n1 R\n",
			want:  "11\n",
		},
		{
			name:  "case2",
			input: "3\n2 L\n2 L\n100 L\n",
			want:  "98\n",
		},
		{
			name:  "case3",
			input: "8\n22 L\n75 L\n26 R\n45 R\n72 R\n81 R\n47 L\n29 R\n",
			want:  "188\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
