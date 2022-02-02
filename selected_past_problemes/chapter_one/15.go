//
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x int
	y int
}

// 参考: https://raahii.github.io/posts/permutations-in-go/
func permute(ori []Point) [][]Point {
	ret := make([][]Point, 0)
	ret = append(ret, makeCopy(ori))

	n := len(ori)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		ori[i], ori[j] = ori[j], ori[i]
		ret = append(ret, makeCopy(ori))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
	return ret
}

func makeCopy(ori []Point) []Point {
	return append([]Point{}, ori...)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)
	n_list := make([]Point, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(r, &x)
		fmt.Fscan(r, &y)
		n_list[i] = Point{x, y}
	}
	var sum float64
	perm_list := permute(n_list)
	for i := 0; i < len(perm_list); i++ {
		for j := 0; j < n-1; j++ {
			a := perm_list[i][j]
			b := perm_list[i][j+1]
			x := a.x - b.x
			y := a.y - b.y
			sum += math.Sqrt(float64(x*x + y*y))
		}
	}
	fmt.Fprintln(w, sum/float64(len(perm_list)))
}
