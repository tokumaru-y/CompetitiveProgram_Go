// https://atcoder.jp/contests/pakencamp-2019-day3/tasks/pakencamp_2019_day3_c
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
	var N, M int
	fmt.Fscan(r, &N)
	fmt.Fscan(r, &M)
	var A = make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(r, &A[i][j])
		}
	}
	var ans int
	for i := 0; i < M; i++ {
		for j := i + 1; j < M; j++ {
			var tmp_sum int
			for k := 0; k < N; k++ {
				if A[k][i] > A[k][j] {
					tmp_sum += A[k][i]
				} else {
					tmp_sum += A[k][j]
				}
			}
			if ans < tmp_sum {
				ans = tmp_sum
			}
		}
	}
	fmt.Fprintln(w, ans)
}
