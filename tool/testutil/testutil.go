package testutil

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// Color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"

	// Status icons
	iconPanic  = "🔥 PANIC"
	iconWrong  = "❌ WRONG"
	iconAccept = "✅ ACCEPT"

	// Output formats
	formatPanic = `%s:
%s%s%s, %s%s%s

Input:
%s%s%s`

	formatWrongMultiline = `%s (first mismatch at line %d):
Got:
%s%s%s
Expect:
%s%s%s
Input:
%s%s%s`

	formatWrongSingleline = `%s (first mismatch at index %d):
Got:    %s
Expect: %s%s%s

Input:
%s%s%s`

	formatAccept = "%s: %s%s !%s"
)

// RunTest runs test cases by simulating standard input/output through pipes
// t: test object
// name: test name
// input: simulated standard input
// want: expected output
// fn: function to be tested
func RunTest(t *testing.T, name string, input, want string, fn func()) {
	// Save original stdin/stdout
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		// Restore original stdin/stdout after test
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// Create input pipe
	inReader, inWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("%s Failed to create input pipe: %v", iconWrong, err)
	}

	// Create output pipe
	outReader, outWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("%s Failed to create output pipe: %v", iconWrong, err)
	}

	// Replace standard input/output
	os.Stdin = inReader
	os.Stdout = outWriter

	// Create channel for synchronization and panic info
	done := make(chan error)

	// Start goroutine to write input data
	go func() {
		defer inWriter.Close()
		io.Copy(inWriter, strings.NewReader(input))
	}()

	// Start goroutine to read output data
	var output bytes.Buffer
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(&output, outReader)
	}()

	// Start goroutine to run test function (with panic capture)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// Get stack trace
				buf := make([]byte, 4096)
				n := runtime.Stack(buf, false)
				stackTrace := string(buf[:n])

				// Extract line number and error message
				lineNum := "unknown"
				lines := strings.Split(stackTrace, "\n")
				for _, line := range lines {
					// Look for main.go call line
					if strings.Contains(line, "/main.go:") {
						if start := strings.LastIndex(line, "/main.go:"); start != -1 {
							start += len("/main.go:")
							if end := strings.Index(line[start:], " "); end != -1 {
								lineNum = line[start : start+end]
							} else {
								// Line number might be at the end if no space
								lineNum = line[start:]
							}
							break
						}
					}
				}

				// Format error message
				errMsg := fmt.Sprintf("Error: %v", r)
				formattedInput := formatInput(input)

				t.Errorf(formatPanic,
					iconPanic,
					colorBlue, lineNum, colorReset,
					colorRed, errMsg, colorReset,
					colorBlue, formattedInput, colorReset)
				done <- fmt.Errorf("panic occurred")
				return
			}
			outWriter.Close()
			done <- nil
		}()
		fn()
	}()

	// Wait for test function to complete and check for panic
	if err := <-done; err != nil {
		t.Error(err)
		return
	}
	wg.Wait()

	// Compare output results
	compareOutput(t, name, output.String(), want, input)
}

// formatInput formats input data with line numbers
func formatInput(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var formatted strings.Builder
	maxLineNumWidth := len(fmt.Sprint(len(lines)))

	for i, line := range lines {
		// Right-align line numbers
		lineNum := fmt.Sprintf("%*d", maxLineNumWidth, i+1)
		formatted.WriteString(fmt.Sprintf("│ %s) %s\n", lineNum, line))
	}
	return formatted.String()
}

// compareOutput compares output results and provides detailed difference display
func compareOutput(t *testing.T, name, got, want, input string) {
	got = strings.TrimSpace(got)
	want = strings.TrimSpace(want)

	if got == want {
		t.Logf(formatAccept, name, colorGreen, iconAccept, colorReset)
		return
	}

	formattedInput := formatInput(input)

	gotLines := strings.Split(got, "\n")
	wantLines := strings.Split(want, "\n")

	if len(gotLines) > 1 || len(wantLines) > 1 {
		var formattedGot strings.Builder
		maxLineNumWidth := len(fmt.Sprint(len(gotLines)))
		firstMismatchLine := -1

		for i, line := range gotLines {
			lineNum := fmt.Sprintf("%*d", maxLineNumWidth, i+1)
			color := colorBlue
			if i < len(wantLines) && line == wantLines[i] {
				color = colorGreen
			} else if firstMismatchLine == -1 {
				firstMismatchLine = i + 1 // Use 1-based index
			}
			formattedGot.WriteString(fmt.Sprintf("│ %s) %s%s%s\n", lineNum, color, line, colorReset))
		}

		if firstMismatchLine == -1 && len(gotLines) != len(wantLines) {
			firstMismatchLine = min(len(gotLines), len(wantLines)) + 1
		}

		formattedWant := formatWithLineNumbers(want)

		t.Errorf(formatWrongMultiline,
			iconWrong,
			firstMismatchLine,
			"", formattedGot.String(), "",
			colorRed, formattedWant, colorReset,
			colorBlue, formattedInput, colorReset)
		return
	}

	// Single line comparison with mismatch index
	var gotStr strings.Builder
	maxLen := max(len(got), len(want))
	firstMismatch := -1

	for i := 0; i < maxLen; i++ {
		if i < len(got) && i < len(want) {
			if got[i] == want[i] {
				gotStr.WriteString(fmt.Sprintf("%s%c%s", colorGreen, got[i], colorReset))
			} else {
				if firstMismatch == -1 {
					firstMismatch = i
				}
				gotStr.WriteString(fmt.Sprintf("%s%c%s", colorBlue, got[i], colorReset))
			}
		} else {
			if firstMismatch == -1 {
				firstMismatch = i
			}
			if i < len(got) {
				gotStr.WriteString(fmt.Sprintf("%s%c%s", colorBlue, got[i], colorReset))
			}
		}
	}

	if firstMismatch == -1 {
		firstMismatch = maxLen
	}

	t.Errorf(formatWrongSingleline,
		iconWrong,
		firstMismatch,
		gotStr.String(),
		colorRed, want, colorReset,
		colorBlue, formattedInput, colorReset)
}

// formatWithLineNumbers formats output with line numbers
func formatWithLineNumbers(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	var formatted strings.Builder
	maxLineNumWidth := len(fmt.Sprint(len(lines)))

	for i, line := range lines {
		lineNum := fmt.Sprintf("%*d", maxLineNumWidth, i+1)
		formatted.WriteString(fmt.Sprintf("│ %s) %s\n", lineNum, line))
	}
	return formatted.String()
}
