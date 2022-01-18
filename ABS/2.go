// https://atcoder.jp/contests/abs/tasks/abc086_a
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

	var a, b int
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	var ans string
	if res := a * b; res%2 == 0 {
		ans = "Even"
	} else {
		ans = "Odd"
	}
	fmt.Fprintln(w, ans)
}
