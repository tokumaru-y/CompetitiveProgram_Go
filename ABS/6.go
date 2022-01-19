// https://atcoder.jp/contests/abs/tasks/abc083_b
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, a, b int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	var ans int
	for i := 1; i <= n; i++ {
		x := i
		tmp_sum := 0
		for x > 0 {
			tmp_sum += x % 10
			x /= 10
		}
		if a <= tmp_sum && tmp_sum <= b {
			ans += i
		}
	}
	fmt.Fprintln(w, ans)
}
