//https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_5_A&lang=ja

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
	var n, m int
	fmt.Fscan(r, &n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		A[i] = x
	}
	fmt.Fscan(r, &m)
	sum_map := map[int]int{}
	for bit := 0; bit < (1 << uint(n)); bit++ {
		tmp_sum := 0
		for j := 0; j < n; j++ {
			if (bit & (1 << uint(j))) > 0 {
				tmp_sum += A[j]
			}
		}
		sum_map[tmp_sum] = 1
	}
	for i := 0; i < m; i++ {
		var target int
		fmt.Fscan(r, &target)
		_, ok := sum_map[target]
		ans := "no"
		if ok {
			ans = "yes"
		}
		fmt.Fprintln(w, ans)
	}

}
