// https://atcoder.jp/contests/abc150/tasks/abc150_c
package main

import (
	"bufio"
	"fmt"
	"os"
)

func next_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return x < y })
}

func prev_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return y < x })
}

func _permutation_pandita(a []int, less func(x, y int) bool) bool {
	i := len(a) - 2
	// Find the right most incresing elements a[i] and a[i+1]
	for 0 <= i && !less(a[i], a[i+1]) {
		i--
	}
	if i == -1 {
		return false
	}
	j := i + 1
	// Find the smallest element that is greater than a[i] in a[i+1:]
	for k := j + 1; k < len(a); k++ {
		// a[i] < a[k] && a[k] <= a[j]
		if less(a[i], a[k]) && !less(a[j], a[k]) {
			j = k
		}
	}
	a[i], a[j] = a[j], a[i]
	for p, q := i+1, len(a)-1; p < q; p, q = p+1, q-1 {
		a[p], a[q] = a[q], a[p]
	}
	return true
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
		fmt.Fscan(r, &P[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &Q[i])
	}
	origin_list := make([]int, n)
	for i := 1; i <= n; i++ {
		origin_list[i-1] = i
	}
	var a, b int
	for i := 1; next_permutation(origin_list); i++ {
		is_ok_a := true
		is_ok_b := true
		for j := 0; j < n; j++ {
			if origin_list[j] != P[j] {
				is_ok_a = false
				break
			}
		}
		for j := 0; j < n; j++ {
			if origin_list[j] != Q[j] {
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
