//https://atcoder.jp/contests/abc126/tasks/abc126_a
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, k int
	var s string
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &k)
	fmt.Fscan(r, &s)
	ans := []rune(s)
	for i := 0; i < n; i++ {
		if i+1 == k {
			ans[i] = unicode.ToLower(ans[i])
		}
	}
	fmt.Fprintln(w, string(ans))
}
