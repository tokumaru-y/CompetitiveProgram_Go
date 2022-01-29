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
			var st, en int
			if A[i] <= B[j] {
				st = A[i]
				en = B[j]
			}
			for k := 0; k < n; k++ {
				var nex int
				if st > A[k] {
					tmp_sum += (st - A[k]) * 2
					nex = st
				} else {
					tmp_sum += A[k] - st
					nex = A[k]
				}
				if nex < B[k] {
					tmp_sum += B[k] - nex
					nex = B[k]
				}
				if nex > en {
					tmp_sum += nex - en
				} else {
					tmp_sum += en - nex
				}
			}
			if ans > tmp_sum {
				ans = tmp_sum
			}
		}
	}
	fmt.Fprintln(w, ans)
}
