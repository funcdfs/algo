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
	//testCaseCnt := input[int]()
	testCaseCnt := 1
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

// link: https://codeforces.com/contest/958/problem/A1
// time: 2025-01-19 14:29:20 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {
	n := input[int]()
	a := make([][]byte, n)
	for i := range a {
		a[i] = input[[]byte]()
	}
	b := make([][]byte, n)
	for i := range b {
		b[i] = input[[]byte]()
	}

	ok := func() {
		println("Yes")
	}
	flip := func() {
		c := make([][]byte, n)
		for i := range c {
			c[i] = make([]byte, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				c[i][j] = a[n-i-1][j]
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a[i][j] = c[i][j]
			}
		}
	}
	rot := func() {
		c := make([][]byte, n)
		for i := range c {
			c[i] = make([]byte, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				c[i][j] = a[n-j-1][i]
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a[i][j] = c[i][j]
			}
		}
	}
	equal := func() bool {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if a[i][j] != b[i][j] {
					return false
				}
			}
		}
		return true
	}
	_ = func() {
		for i := 0; i < n; i++ {
			log.Println(string(a[i]))
		}
		log.Println("---")
		for i := 0; i < n; i++ {
			log.Println(string(b[i]))
		}
	}
	for i := 0; i < 4; i++ {
		rot()
		//log.Println("rot", i)
		//console()

		if equal() {
			ok()
			return
		}
		for j := 0; j < 4; j++ {
			flip()
			if equal() {
				ok()
				return
			}
		}
	}
	println("No")
	return
}

// ----------------------------- /* End of useful functions */ -------------------------------
