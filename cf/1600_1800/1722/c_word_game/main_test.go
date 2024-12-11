package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 817 (Div. 4) C. Word Game
// URL: https://codeforces.com/contest/1722/problem/C
// Time: 2024-12-11 12:14:14

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "3\n1\nabc\ndef\nabc\n3\norz for qaq\nqaq orz for\ncod for ces\n5\niat roc hem ica lly\nbac ter iol ogi sts\nbac roc lly iol iat\n",
			want:  "1 3 1\n2 2 6\n9 11 5\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
