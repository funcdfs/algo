package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 966 (Div. 3) A. Primary Task
// URL: https://codeforces.com/contest/2000/problem/A
// Time: 2025-01-13 17:19:58

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "7\n100\n1010\n101\n105\n2033\n1019\n1002\n",
			want:  "NO\nYES\nNO\nYES\nNO\nYES\nNO\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
