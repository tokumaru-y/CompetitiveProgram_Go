// https://atcoder.jp/contests/s8pc-6/tasks/s8pc_6_b

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
	A := make([]int, n)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		A[i] = a
		B[i] = b
	}
	var ans = math.MaxInt64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var tmp_sum int
			dd := A[i] - B[j]
			for k := 0; k < n; k++ {
				da := A[k] - A[i]
				db := B[k] - B[j]
				if da < 0 {
					da = -da * 2
				}
				if db < 0 {
					db = -db
				}
				tmp_sum += db + da
			}
			if ans > tmp_sum {
				ans = tmp_sum
			}
		}
	}
	fmt.Fprintln(w, ans)
}
