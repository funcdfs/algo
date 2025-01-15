// <editor-fold desc="useless function"
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	_in = bufio.NewReader(os.Stdin)
	_out = bufio.NewWriter(os.Stdout)
	defer _out.Flush()
	log.SetPrefix("[dbg:] ")
	log.SetFlags(log.Lshortfile)
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

// link: https://codeforces.com/contest/1992/problem/C
// time: 2025-01-10 02:30:36 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {

	// g[i] is the sum of all the numbers in the permutation
	//      on a prefix of length [i] not greaterthan m .
	// f[i] is the sum of ...not less than k ;
	n, m, k := input[int](), input[int](), input[int]()
	for i := n; i >= k; i-- {
		print(i)
		print(" ")
	}
	for i := m + 1; i < k; i++ {
		print(i)
		print(" ")
	}
	for i := 1; i <= m; i++ {
		print(i)
		print(" ")
	}
	println("")
}

// ----------------------------- /* End of useful functions */ -------------------------------
