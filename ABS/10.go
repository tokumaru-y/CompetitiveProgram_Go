// https://atcoder.jp/contests/abs/tasks/arc065_a

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var s string
	li := []string{"dreamer", "dream", "erase", "eraser"}
	fmt.Fscan(r, &s)
	tmp := s
	for {
		for _, tm := range li {
			s = strings.TrimSuffix(s, tm)
		}
		if s == "" {
			fmt.Fprintln(w, "YES")
			break
		}
		if s == tmp {
			fmt.Fprintln(w, "NO")
			break
		}
		tmp = s
	}
}
