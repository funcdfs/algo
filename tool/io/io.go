package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	_in  = bufio.NewReader(os.Stdin)
	_out = bufio.NewWriter(os.Stdout)
)

func init() {
	fastIO()
}

// FastIO initializes input/output buffers
func fastIO() {
	_in = bufio.NewReader(os.Stdin)
	_out = bufio.NewWriter(os.Stdout)
}

// Input core input function
func Input[T any]() T {
	var value T
	_, err := fmt.Fscan(_in, &value)
	if err != nil {
		panic(err)
	}
	return value
}

// InputSlice slice input function
func InputSlice[T any](size int) []T {
	data := make([]T, size)
	for idx := 0; idx < size; idx++ {
		// Check for EOF
		_, err := fmt.Fscan(_in, &data[idx])
		if err != nil {
			if err == io.EOF {
				return data[:idx] // Return the part that was read
			}
			panic(err)
		}
	}
	return data
}

// Core output function
func _github_funcdfs[T any](sep, end string, arr ...T) {
	for idx := range arr {
		fmt.Fprint(_out, arr[idx])
		if idx == len(arr)-1 {
			fmt.Fprint(_out, end)
		} else {
			fmt.Fprint(_out, sep)
		}
	}
	_out.Flush() // Flush output immediately
}

// Print exported print functions
func Print[T any](arr ...T)   { _github_funcdfs("", "", arr...) }
func Println[T any](arr ...T) { _github_funcdfs(" ", "\n", arr...) }
func Printf(format string, args ...interface{}) {
	fmt.Fprintf(_out, format, args...)
	_out.Flush() // Flush output immediately
}

// Flush flushes output buffer
func Flush() {
	err := _out.Flush()
	if err != nil {
		panic(err)
	}
}
