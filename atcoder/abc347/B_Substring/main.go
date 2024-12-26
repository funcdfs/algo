// link: https://atcoder.jp/contests/abc347/tasks/abc347_b
// time: 2024-12-26 16:49:03 https://github.com/funcdfs

// <editor-fold desc="useless function"
package main

import (
	"bufio"
	"fmt"
	"os"
)

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
func main() {
	_in = bufio.NewReader(os.Stdin)
	_out = bufio.NewWriter(os.Stdout)
	defer _out.Flush()
	solve()
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

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve() {

	s := input[[]byte]()
	n := len(s)

	// different continue' string count
	// len=1:100
	// len=100:
	hs := make(map[string]int)
	for size := 1; size <= n; size++ {
		for i := 0; i+size <= n; i++ {
			// start position is i
			// end position is i+size
			hs[string(s[i:i+size])]++
		}
	}
	//log.Println(hs)

	println(len(hs))
}

// ----------------------------- /* End of useful functions */ -------------------------------
