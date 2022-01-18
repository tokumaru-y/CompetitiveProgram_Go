// https://atcoder.jp/contests/abs/tasks/abc081_a
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

	var s string
	fmt.Fscan(r, &s)
	ans := 0
	for _, c := range s {
		if c == '1' {
			ans++
		}
	}
	fmt.Fprintln(w, ans)
}
