// https://atcoder.jp/contests/abc150/tasks/abc150_c
package main

import (
	"bufio"
	"fmt"
	"os"
)

func permute(ori []int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, makeCopy(ori))

	n := len(ori)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		ori[i], ori[j] = ori[j], ori[i]
		ret = append(ret, makeCopy(ori))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
	return ret
}

func makeCopy(ori []int) []int {
	return append([]int{}, ori...)
}
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)
	P := make([]int, n)
	Q := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscan(r, &tmp)
		P[i] = tmp
	}
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscan(r, &tmp)
		Q[i] = tmp
	}
	origin_list := make([]int, n)
	for i := 1; i <= n; i++ {
		origin_list[i-1] = i
	}
	perm_list := permute(origin_list)
	var a, b int
	for i := 0; i < len(perm_list); i++ {
		target := perm_list[i]
		is_ok_a := true
		is_ok_b := true
		for j := 0; j < n; j++ {
			if target[j] != P[j] {
				is_ok_a = false
				break
			}
		}
		for j := 0; j < n; j++ {
			if target[j] != Q[j] {
				is_ok_b = false
				break
			}
		}
		if is_ok_a {
			a = i
		}
		if is_ok_b {
			b = i
		}
	}
	ans := a - b
	if ans < 0 {
		ans *= -1
	}
	fmt.Fprintln(w, ans)
}
