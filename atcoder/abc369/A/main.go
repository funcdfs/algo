// link: https://atcoder.jp/contests/abc369/tasks/abc369_a
// time: 2025-03-10 11:04:59 https://github.com/funcdfs

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

	a, b := input[int](), input[int]()

	if a > b {
		a, b = b, a
	} else if a == b {
		println(1)
		return
	}

	diff := b - a
	ans := 2

	calc := func(val int) int {
		x := 0
		for diff%2 == 0 {
			diff /= 2
			x += 1
		}
		return x
	}

	ans += calc(diff)

	println(ans)

}

// ----------------------------- /* End of useful functions */ -------------------------------
