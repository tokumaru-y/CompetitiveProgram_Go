// https://atcoder.jp/contests/abc106/tasks/abc106_b
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func make_divisor_list(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			res = append(res, i)
			res = append(res, n/i)
		}
	}
	return res
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, ans int
	fmt.Fscan(r, &n)
	for i := 1; i <= n; i += 2 {
		divisor_list := make_divisor_list(i)
		if len(divisor_list) == 8 {
			ans++
		}
	}
	fmt.Fprintln(w, ans)
}
