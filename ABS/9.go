// https://atcoder.jp/contests/abs/tasks/abc085_c
package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n, y int) (int, int, int) {
	for a := 0; a <= n; a++ {
		for b := 0; b <= n-a; b++ {
			c := n - (a + b)
			if a*10000+b*5000+c*1000 == y {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, y int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &y)
	a, b, c := solve(n, y)
	fmt.Fprintln(w, a, b, c)
}
