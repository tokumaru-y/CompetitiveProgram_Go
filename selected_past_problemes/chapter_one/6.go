// https://atcoder.jp/contests/sumitrust2019/tasks/sumitb2019_d
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
	var n int
	var s string
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &s)
	var ans int
	num_list := "0123456789"
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				cnt := 0
				cnt_list := []byte{num_list[i], num_list[j], num_list[k]}
				for s_i := 0; s_i < n; s_i++ {
					if cnt < 3 && s[s_i] == cnt_list[cnt] {
						cnt++
					}
				}
				if cnt == 3 {
					ans++
				}
			}
		}
	}
	fmt.Fprintln(w, ans)
}
