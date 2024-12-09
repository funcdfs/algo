package testutil

import (
	"fmt"
	"log"
	"testing"
)

// Simulate a simple main function that reads two numbers from standard input and outputs their sum
func mockMain() {
	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		log.Printf("Failed to read input: %v", err)
		return
	}
	fmt.Printf("%d\n", a+b)
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "case1",
			input: "1 2\n",
			want:  "3",
		},
		{
			name:  "case2",
			input: "100 200\n",
			want:  "3010",
		},
		{
			name:  "case3",
			input: "-5 8\n",
			want:  "3",
		},
		{
			name:  "case4",
			input: "0 0\n",
			want:  "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunTest(t, tt.name, tt.input, tt.want, mockMain)
		})
	}
}

// TestPanic specifically tests panic scenarios
func TestPanic(t *testing.T) {
	panicMain := func() {
		var arr []int
		fmt.Printf("%d\n", arr[10])
	}

	t.Run("panic_case", func(t *testing.T) {
		RunTest(t, "panic_test", "", "", panicMain)
		// panic 会被 RunTest 捕获并通过 t.Error 报告
	})
}

// TestMultiline specifically tests multi-line output scenarios
func TestMultiline(t *testing.T) {
	multilineMain := func() {
		var n int
		fmt.Scanf("%d", &n)
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				fmt.Printf("Line %d\n", i)
			}
			fmt.Printf("Blcok %d\n", i)
		}
	}

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "multiline1",
			input: "3\n",
			want:  "Line 1\nLine 2\nLine 3",
		},
		{
			name:  "multiline2",
			input: "2\n",
			want:  "Line 1\nLine 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunTest(t, tt.name, tt.input, tt.want, multilineMain)
		})
	}
}
