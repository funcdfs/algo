// <editor-fold desc="useless function"
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

// link: https://codeforces.com/contest/1722/problem/B
// time: 2024-12-11 12:08:12 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {

	n := input[int]()
	s1 := input[[]byte]()
	s2 := input[[]byte]()
	for i := 0; i < n; i++ {
		if s1[i] == 'B' {
			s1[i] = 'F'
		} else if s1[i] == 'G' {
			s1[i] = 'F'
		}
		if s2[i] == 'B' {
			s2[i] = 'F'
		} else if s2[i] == 'G' {
			s2[i] = 'F'
		}
	}

	check := func() bool {
		if slices.Equal(s1, s2) {
			return true
		}
		return false
	}
	if check() == true {
		println("YES")
	} else {
		println("NO")
	}

}

// ----------------------------- /* End of useful functions */ -------------------------------
