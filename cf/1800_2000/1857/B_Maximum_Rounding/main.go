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

// link: https://codeforces.com/contest/1857/problem/B
// time: 2025-02-12 15:26:05 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {
	// output the max value can be output, the final number contains no front zero
	s := input[[]byte]()
	carry := 0
	k := len(s)
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i]+byte(carry) >= '5' {
			k = i
			// record the last idx
			carry = 1
			// from behind to front, if >= '5'. success
		} else {
			s[i] += byte(carry)
			carry = 0
		}
	}
	for i := k; i < n; i++ {
		s[i] = '0'
	}
	if carry > 0 {
		s = append([]byte{'1'}, s...)
	}
	println(string(s))
}

// ----------------------------- /* End of useful functions */ -------------------------------
