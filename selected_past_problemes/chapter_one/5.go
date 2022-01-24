// https://atcoder.jp/contests/abc095/tasks/arc096_a
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
	var a, b, c, x, y int
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &c)
	fmt.Fscan(r, &x)
	fmt.Fscan(r, &y)
	ans := math.MaxInt64
	for i := 0; i <= 200000; i++ {
		tmp := c * i
		t_x := x - i/2
		t_y := y - i/2
		if t_x > 0 {
			tmp += t_x * a
		}
		if t_y > 0 {
			tmp += t_y * b
		}
		if ans > tmp {
			ans = tmp
		}
	}
	fmt.Fprintln(w, ans)
}
