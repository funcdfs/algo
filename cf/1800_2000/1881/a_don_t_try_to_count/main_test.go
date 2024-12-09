package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 903 (Div. 3) A. Don't Try to Count
// URL: https://codeforces.com/contest/1881/problem/A
// Time: 2024-12-09 17:30:31

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "case1",
			input: "12\n1 5\na\naaaaa\n5 5\neforc\nforce\n2 5\nab\nababa\n3 5\naba\nababa\n4 3\nbabb\nbbb\n5 1\naaaaa\na\n4 2\naabb\nba\n2 8\nbk\nkbkbkbkb\n12 2\nfjdgmujlcont\ntf\n2 2\naa\naa\n3 5\nabb\nbabba\n1 19\nm\nmmmmmmmmmmmmmmmmmmm\n",
			want:  "3\n1\n2\n-1\n1\n0\n1\n3\n1\n0\n2\n5\n",
		},
		{
			name:  "error case:",
			input: "1\n2 3\nbc\ncbc\n",
			want:  "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
