// https://atcoder.jp/contests/abc002/tasks/abc002_4
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
	fmt.Fscan(r, &m)
	connection := make([][]bool, n)
	for i := 0; i < n; i++ {
		connection[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(r, &x)
		fmt.Fscan(r, &y)
		x--
		y--
		connection[x][y] = true
		connection[y][x] = true
	}
	var ans int
	for bit := 0; bit < (1 << uint(n)); bit++ {
		member := make([]int, 0, n)
		for j := 0; j < n; j++ {
			if bit&(1<<j) > 0 {
				member = append(member, j)
			}
		}
		is_ok := true
		for i := 0; i < len(member); i++ {
			for j := 0; j < len(member); j++ {
				if i == j {
					continue
				}
				if !connection[member[i]][member[j]] {
					is_ok = false
				}
			}
		}
		if is_ok && len(member) > ans {
			ans = len(member)
		}
	}
	fmt.Fprintln(w, ans)
}
