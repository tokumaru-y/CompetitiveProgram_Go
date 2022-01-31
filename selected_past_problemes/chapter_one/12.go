// https://atcoder.jp/contests/joi2008yo/tasks/joi2008yo_e
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
	var R, C int
	fmt.Fscan(r, &R)
	fmt.Fscan(r, &C)
	grid := make([][]int, R)
	for i := 0; i < R; i++ {
		grid[i] = make([]int, C)
		for j := 0; j < C; j++ {
			var tmp int
			fmt.Fscan(r, &tmp)
			grid[i][j] = tmp
		}
	}
	var ans int
	for bit := 0; bit < (1 << uint(R)); bit++ {
		copy_grid := make([][]int, R)
		for i := 0; i < R; i++ {
			copy_grid[i] = make([]int, C)
			copy(copy_grid[i], grid[i])
		}
		for i := 0; i < R; i++ {
			if bit&(1<<uint(i)) > 0 {
				for j := 0; j < C; j++ {
					copy_grid[i][j]++
				}
			}
		}
		var tmp_ans int
		for i := 0; i < C; i++ {
			tmp_d_cnt := 0
			for j := 0; j < R; j++ {
				if copy_grid[j][i]%2 == 0 {
					tmp_d_cnt++
				}
			}
			if tmp_d_cnt >= R-tmp_d_cnt {
				tmp_ans += tmp_d_cnt
			} else {
				tmp_ans += R - tmp_d_cnt
			}
		}

		if tmp_ans > ans {
			ans = tmp_ans
		}
	}
	fmt.Fprintln(w, ans)
}
