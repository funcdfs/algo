package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 817 (Div. 4) A. Spell Check
// URL: https://codeforces.com/contest/1722/problem/A
// Time: 2024-12-11 12:03:59

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "10\n5\nTimur\n5\nmiurT\n5\nTrumi\n5\nmriTu\n5\ntimur\n4\nTimr\n6\nTimuur\n10\ncodeforces\n10\nTimurTimur\n5\nTIMUR\n",
			want:  "YES\nYES\nYES\nYES\nNO\nNO\nNO\nNO\nNO\nNO\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
