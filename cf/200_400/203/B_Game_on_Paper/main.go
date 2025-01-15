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

// link: https://codeforces.com/contest/203/problem/B
// time: 2025-01-14 22:00:43 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {

	n, m := input[int](), input[int]()
	a := make([][]int, n+10)
	for i := range a {
		a[i] = make([]int, n+10)
	}
	ans := -1
	for i := 0; i < m; i++ {
		x, y := input[int](), input[int]()
		x += 2
		y += 2
		a[x][y] = 1
		for u := x - 2; u <= x; u++ {
			for v := y - 2; v <= y; v++ { // 9 POINTS
				cnt := 0
				for s := 0; s < 3; s++ {
					for t := 0; t < 3; t++ { // 9 directions
						cnt += a[u+s][v+t]
					}
				}
				if cnt == 9 {
					ans = i + 1
					break
				}
			}
		}
	}
	println(ans)
	return
}

// ----------------------------- /* End of useful functions */ -------------------------------
