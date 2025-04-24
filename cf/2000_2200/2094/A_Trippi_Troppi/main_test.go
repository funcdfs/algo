package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: Codeforces - Codeforces Round 1017 (Div. 4) A. Trippi Troppi
// URL: https://codeforces.com/contest/2094/problem/A
// Time: 2025-04-24 20:52:35

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		
		{
			name:  "case1",
			input: "7\nunited states america\noh my god\ni cant lie\nbinary indexed tree\nbelieve in yourself\nskibidi slay sigma\ngod bless america\n",
			want:  "usa\nomg\nicl\nbit\nbiy\nsss\ngba\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
