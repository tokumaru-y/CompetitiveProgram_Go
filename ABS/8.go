//https://atcoder.jp/contests/abs/tasks/abc085_b
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
	fmt.Fscan(r, &n)
	D := make([]int, n)
	for i := 0; i < n; i++ {
		var inp int
		fmt.Fscan(r, &inp)
		D[i] = inp
	}
	num_map := make(map[int]bool)
	uniq_slice := []int{}
	for _, d := range D {
		if _, ok := num_map[d]; !ok {
			num_map[d] = true
			uniq_slice = append(uniq_slice, d)
		}
	}

	fmt.Fprintln(w, len(uniq_slice))
}
