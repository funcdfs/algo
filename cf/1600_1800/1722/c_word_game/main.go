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

// link: https://codeforces.com/contest/1722/problem/C
// time: 2024-12-11 12:14:14 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {

	n := input[int]()
	a := make([][]string, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]string, n)
		for idx := range a[i] {
			a[i][idx] = string(input[[]byte]())
		}
	}
	info := make(map[string][]int)
	for i := 0; i < 3; i++ {
		for j := 0; j < n; j++ {
			info[a[i][j]] = append(info[a[i][j]], i+1)
		}
	}
	// output
	for i := 0; i < 3; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			x := len(info[a[i][j]])
			if x == 3 {
				cnt += 0
			} else if x == 2 {
				cnt += 1
			} else if x == 1 {
				cnt += 3
			}
		}
		print(cnt)
		if i != 3-1 {
			print(" ")
		}
	}
	println("")
}

// ----------------------------- /* End of useful functions */ -------------------------------
