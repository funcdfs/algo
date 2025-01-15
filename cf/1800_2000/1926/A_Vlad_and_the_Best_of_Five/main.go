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

// link: https://codeforces.com/contest/1926/problem/A
// time: 2024-12-26 18:24:45 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {
	s := input[[]byte]()
	if len(s) != 5 {
		panic("assert failed, fix me")
	}
	cnt1, cnt2 := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			cnt1++
		} else if s[i] == 'B' {
			cnt2++
		}
	}
	if cnt1 > cnt2 {
		println("A")
	} else {
		println("B")
	}
}

// ----------------------------- /* End of useful functions */ -------------------------------
