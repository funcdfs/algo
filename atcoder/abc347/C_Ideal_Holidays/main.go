// link: https://atcoder.jp/contests/abc347/tasks/abc347_c
// time: 2024-12-26 17:02:48 https://github.com/funcdfs

// <editor-fold desc="useless function"
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	log.SetPrefix("[dbg:] ")
	log.SetFlags(log.Lshortfile)
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
	n, a, b := input[int](), input[int](), input[int]()
	d := inputSlice[int](n)
	totalWeek := a + b

	log.Println(a, b)
	log.Println(d)
	reminder := make([]int, 0)
	sort.Ints(d)
	for i := 0; i < n; i++ {
		reminder = append(reminder, (d[i]-d[0])%totalWeek)
	}
	sort.Ints(reminder)

	log.Println(reminder)

	check := func() bool {
		return false
	}
	if check() == true {
		println("Yes")
	} else {
		println("No")
	}
}

// ----------------------------- /* End of useful functions */ -------------------------------
