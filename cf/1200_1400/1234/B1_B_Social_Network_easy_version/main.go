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

// link: https://codeforces.com/contest/1234/problem/B1
// time: 2025-01-15 22:39:43 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {

	n, k := input[int](), input[int]()
	a := inputSlice[int](n)

	hs := make(map[int]int, 0)
	ans := make([]int, 0)

	for i := 0; i < n; i++ {
		if _, ok := hs[a[i]]; !ok {
			ans = append([]int{a[i]}, ans...)
			hs[a[i]] += 1
		}
		if len(ans) > k {
			top := ans[len(ans)-1]
			hs[top] -= 1
			if hs[top] == 0 {
				delete(hs, top)
			}
			ans = ans[:len(ans)-1]
		}
	}
	println(len(ans))
	println(ans...)

}

// ----------------------------- /* End of useful functions */ -------------------------------
