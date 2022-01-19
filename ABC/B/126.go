// https://atcoder.jp/contests/abc126/tasks/abc126_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func is_mm(s string) bool {
	yy, _ := strconv.Atoi(s)
	return 1 <= yy && yy <= 12
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var s string
	fmt.Fscan(r, &s)
	x, y := is_mm(s[0:2]), is_mm(s[2:4])
	var ans string
	ans = "NA"
	if x && y {
		ans = "AMBIGUOUS"
	} else if y {
		ans = "YYMM"
	} else if x {
		ans = "MMYY"
	}
	fmt.Fprintln(w, ans)
}
