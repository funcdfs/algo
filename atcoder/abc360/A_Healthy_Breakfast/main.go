// link: https://atcoder.jp/contests/abc360/tasks/abc360_a
// time: 2025-02-12 15:44:47 https://github.com/funcdfs

// <editor-fold desc="useless function"
package main

import (
	"bufio"
	"fmt"
	"log"
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
	log.SetPrefix("[dbg:] ")
	log.SetFlags(log.Lshortfile)
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
	// check if R is left at M
	rIdx, mIdx := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			rIdx = i
		} else if s[i] == 'M' {
			mIdx = i
		}
	}
	check := func() bool {
		if rIdx < mIdx {
			return true
		}
		return false
	}
	if check() == true {
		println("Yes")
	} else {
		println("No")
	}

}

// ----------------------------- /* End of useful functions */ -------------------------------
