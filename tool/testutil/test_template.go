package testutil

const TestTemplate = `package main

import (
	"github.com/funcdfs/algo/tool/testutil"
	"testing"
)

// Contest: {{.Problem.Group}} {{.Problem.Name}}
// URL: {{.Problem.URL}}
// Time: {{.CurrentTime}}

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{{range $index, $test := .Tests}}
		{
			name:  "case{{add $index 1}}",
			input: {{printf "%q" $test.Input}},
			want:  {{printf "%q" $test.Output}},
		},{{end}}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.RunTest(t, tt.name, tt.input, tt.want, main)
		})
	}
}
`
