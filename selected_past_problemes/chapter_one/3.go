// https://atcoder.jp/contests/abc122/tasks/abc122_b

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
	var S string
	fmt.Fscan(r, &S)
	var ans, cnt int
	for i := 0; i < len(S); i++ {
		if S[i] == 'A' || S[i] == 'C' || S[i] == 'G' || S[i] == 'T' {
			cnt++
		} else {
			cnt = 0
		}
		if ans < cnt {
			ans = cnt
		}
	}
	fmt.Fprintln(w, ans)
}
