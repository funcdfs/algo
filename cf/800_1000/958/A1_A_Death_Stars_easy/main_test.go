package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Helvetic Coding Contest 2018 online mirror (teams allowed, unrated) A1. Death Stars (easy)
// URL: https://codeforces.com/contest/958/problem/A1
// Time: 2025-01-19 14:29:20

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "4\nXOOO\nXXOO\nOOOO\nXXXX\nXOOO\nXOOO\nXOXO\nXOXX\n",
			want:  "Yes\n",
		},
		{
			name:  "case2",
			input: "2\nXX\nOO\nXO\nOX\n",
			want:  "No\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
