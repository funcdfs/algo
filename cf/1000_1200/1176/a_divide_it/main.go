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

// link: https://codeforces.com/contest/1176/problem/A
// time: 2024-12-10 11:30:47 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {
	n := input[int]()

	cnt := 0
	for {
		if n == 1 {
			println(cnt)
			return
		} else if n%2 == 0 {
			n = (n * 1) / 2
			cnt++
		} else if n%3 == 0 {
			n = (n * 2) / 3
			cnt++
		} else if n%5 == 0 {
			n = (n * 4) / 5
			cnt++
		} else if n%2 != 0 && n%3 != 0 && n%5 != 0 {
			println(-1)
			return
		}
	}
}

// ----------------------------- /* End of useful functions */ -------------------------------
