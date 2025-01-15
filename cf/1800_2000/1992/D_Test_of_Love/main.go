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

// link: https://codeforces.com/contest/1992/problem/D
// time: 2025-01-12 13:33:41 https://github.com/funcdfs

// ----------------------------- /* Start of useful functions */ -----------------------------

func solve(_case int) {

	// swim a width 1m. lenght n. river
	// <= k
	// 0 -> n+1
	// logs, crocodile, water
	// L, C, W
	// on the bank or on the logs' jump m [expect to C]
	// on the water, swim to next water segment or move to the bank [nth]
	// cannot land in a segment with a C in any way
	n, m, k := input[int](), input[int](), input[int]()
	// m is the jump distance
	// k is the swim distance
	s := input[[]byte]()
	s = append([]byte{'L'}, s...)
	s = append(s, 'L')

	check := func() bool {
		f := make([]int, n+2)
		f[0] = 1 // represent can reach the first position
		// store the land positon first
		l := make([]int, n+2)
		for i, last := 1, 0; i < len(s); i++ {
			if s[i] != 'L' {
				l[i] = last
			} else {
				l[i] = last
				last = i
			}
		}
		for i := 1; i <= n+1; i++ {
			if s[i] == 'C' {
				continue
			}
			if i-l[i] <= m {
				// can jump to this position
				f[i] |= f[l[i]]
				if f[i] == 1 {
					continue
					// can jump to this position, don't consume swim point
				}
			}
			if k > 0 {
				k -= 1
				f[i] |= f[i-1]
			}
		}
		if f[n+1] == 1 {
			return true
		} else {
			return false
		}
	}
	if check() == true {
		println("YES")
	} else {
		println("NO")
	}

}

// ----------------------------- /* End of useful functions */ -------------------------------
