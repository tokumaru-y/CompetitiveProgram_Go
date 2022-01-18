// https://atcoder.jp/contests/abs/tasks/abc081_b
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	ans := math.MaxInt64
	for i := 0; i < n; i++ {
		cnt := 0
		for a[i]%2 == 0 {
			cnt++
			a[i] = a[i] / 2
		}
		if ans > cnt {
			ans = cnt
		}
	}
	fmt.Fprintln(w, ans)
}
