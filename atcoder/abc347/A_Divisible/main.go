// link: https://atcoder.jp/contests/abc347/tasks/abc347_a
// time: 2024-12-26 16:44:00 https://github.com/funcdfs

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
	n, k := input[int](), input[int]()
	a := inputSlice[int](n)

	log.Println(n, k, a)
	for i := range a {
		if a[i]%k == 0 {
			print(a[i] / k)
			print(" ")
		}
	}
	println("")
}

// ----------------------------- /* End of useful functions */ -------------------------------
