// link: https://atcoder.jp/contests/abc369/tasks/abc369_b
// time: 2025-03-10 11:14:15 https://github.com/funcdfs

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
	type motion struct {
		position int
		hand     string
	}
	n := input[int]()
	in := make([]motion, n)
	for i := range in {
		in[i].position = input[int]()
		in[i].hand = string(input[[]byte]())
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	ans := 0
	posl, posr := -1, -1
	for i := range in {
		if in[i].hand == "L" {
			if posl == -1 {
				posl = in[i].position
			} else {
				ans += abs(in[i].position - posl)
				posl = in[i].position
			}
		} else if in[i].hand == "R" {
			if posr == -1 {
				posr = in[i].position
			} else {
				ans += abs(in[i].position - posr)
				posr = in[i].position
			}
		}
	}
	println(ans)
}

// ----------------------------- /* End of useful functions */ -------------------------------
