// <editor-fold desc="useless function"
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

// link: https://codeforces.com/contest/1810/problem/B
// time: 2025-01-19 16:05:16 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {
	// 1: x -> 2x-1
	// 2: x -> 2x+1
	// max is 40
	// construct a sequence

	n := input[int]()
	// the final distination
	// output:
	// 1: the total count
	// 2: the specific sequence
	if n%2 == 0 { // 2x-1 && 2x+1 is odd
		println(-1)
		return
	}
	ans := make([]int, 0)
	for n > 1 {
		if ((n+1)/2)%2 == 1 {
			// if from path 1, the ans is odd
			ans = append(ans, 1)
			// push method one to the answer
			n++
			n /= 2
		} else {
			// else, from path 2
			ans = append(ans, 2)
			n--
			n /= 2
		}
	}
	slices.Reverse(ans)
	if len(ans) > 40 {
		println(-1)
		return
	}
	println(len(ans))
	println(ans...)

}

// ----------------------------- /* End of useful functions */ -------------------------------
