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

// link: https://codeforces.com/contest/1881/problem/A
// time: 2024-12-09 17:30:31 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {
	n, m := input[int](), input[int]()
	short := input[[]byte]()
	if len(short) != n {
		panic("assert failed, fix me")
	}
	long := input[[]byte]()
	if len(long) != m {
		panic("assert failed, fix me")
	}
	// when short contains long
	cnt := 0
	//if _case == 46 {
	//	println(string(short))
	//	println(string(long))
	//	return
	//}
	partSuccess := func(x []byte, y []byte) bool {
		for i := 0; i <= len(x)-len(y); i++ {
			//log.Println(i, string(x), string(x[i:i+len(y)]))
			//log.Println(string(y[0 : 0+len(y)]))
			if slices.Equal(x[i:i+len(y)], y) {
				return true
			}
		}
		return false
	}
	for i := 0; i < 10; i++ {
		if slices.Equal(short, long) || partSuccess(short, long) {
			println(cnt)
			return
		} else {
			short = slices.Concat(short, short)
			cnt += 1
			//log.Println(string(short))
		}
	}
	println(-1)
}

// ----------------------------- /* End of useful functions */ -------------------------------
