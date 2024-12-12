// <editor-fold desc="useless function"
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_in = bufio.NewReader(os.Stdin)
	_out = bufio.NewWriter(os.Stdout)
	defer _out.Flush()
	testCaseCnt := input[int]()
	//testCaseCnt := 1
	for i := 0; i < testCaseCnt; i++ {
		solve(i + 1)
	}
}
var _in, _out = new(bufio.Reader), new(bufio.Writer)

func _github_funcdfs[T any](sep, end string, arr ...T) {
	for idx := range arr {
		fmt.Fprint(_out, arr[idx])
		if idx == len(arr)-1 {
			fmt.Fprint(_out, end)
		} else {
			fmt.Fprint(_out, sep)
		}
	}
}
func input[T any]() T { var value T; fmt.Fscan(_in, &value); return value }
func inputSlice[T any](size int) []T {
	data := make([]T, size)
	for idx := 0; idx < size; idx++ {
		data[idx] = input[T]()
	}
	return data
}
func print[T any](arr ...T)   { _github_funcdfs("", "", arr...) }
func println[T any](arr ...T) { _github_funcdfs(" ", "\n", arr...) }
//</editor-fold>

// link: https://codeforces.com/contest/1850/problem/B
// time: 2024-12-12 01:19:04 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {

	n := input[int]()
	maxVal := 0
	maxIdx :=0 
	for i := 0; i < n; i++ {
		x, y := input[int](), input[int]()
		if x <= 10 {
			if y > maxVal {
				maxVal = y
				maxIdx = i
			}
		}
	}
	println(maxIdx+1)
	
}

// ----------------------------- /* End of useful functions */ -------------------------------
