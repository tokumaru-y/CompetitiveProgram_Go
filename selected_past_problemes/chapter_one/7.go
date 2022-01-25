// https://www.ioi-jp.org/joi/2006/2007-ho-prob_and_sol/2007-ho.pdf#page=5

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cordinate struct {
	xs int
	ys int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)
	X := make([]int, n)
	Y := make([]int, n)
	m := map[Cordinate]int{}
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(r, &x)
		fmt.Fscan(r, &y)
		X[i] = x
		Y[i] = y
		m[Cordinate{xs: x, ys: y}] = 1
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			lx := X[i] - X[j]
			ly := Y[i] - Y[j]
			rx := X[j] - X[i]
			ry := Y[j] - Y[i]
			_, ok1 := m[Cordinate{xs: -ly, ys: lx}]
			_, ok2 := m[Cordinate{xs: ry, ys: -rx}]
			_, ok3 := m[Cordinate{xs: ly, ys: -lx}]
			_, ok4 := m[Cordinate{xs: -ry, ys: rx}]
			if ok1 && ok2 || ok3 && ok4 {
				v := lx*lx + ly*ly
				if ans < v {
					ans = v
				}
			}
		}
	}
	fmt.Fprintln(w, ans)
}
