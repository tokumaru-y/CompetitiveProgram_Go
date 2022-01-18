// https://atcoder.jp/contests/abs/tasks/abc087_b
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
	var a, b, c, x int
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &c)
	fmt.Fscan(r, &x)
	ans := 0
	for a_cnt := 0; a_cnt <= a; a_cnt++ {
		for b_cnt := 0; b_cnt <= b; b_cnt++ {
			for c_cnt := 0; c_cnt <= c; c_cnt++ {
				if a_cnt*500+b_cnt*100+c_cnt*50 == x {
					ans++
				}
			}
		}
	}
	fmt.Fprintln(w, ans)
}
