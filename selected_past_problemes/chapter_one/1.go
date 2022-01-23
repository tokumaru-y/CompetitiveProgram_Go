// https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ITP1_7_B&lang=ja

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
	var n, x int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &x)
	for n > 0 || x > 0 {
		var ans int
		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				for k := j + 1; k <= n; k++ {
					if i+j+k == x {
						ans++
					}
				}
			}
		}
		fmt.Fprintln(w, ans)
		fmt.Fscan(r, &n)
		fmt.Fscan(r, &x)
	}
}
