// https://atcoder.jp/contests/joi2008yo/tasks/joi2008yo_d
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
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
		n_list[i] = Point{x: x, y: y}
	}
	var m int
	fmt.Fscan(r, &m)
	grid := make([]Point, m)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(r, &x)
		fmt.Fscan(r, &y)
		grid[i] = Point{x: x, y: y}
	}
	for i := 0; i < m; i++ {
		map_list := map[Point]int{}
		dif_x := n_list[0].x - grid[i].x
		dif_y := n_list[0].y - grid[i].y
		for j := 0; j < m; j++ {
			map_list[Point{x: grid[j].x + dif_x, y: grid[j].y + dif_y}] = 1
		}

		is_correct := true
		for j := 0; j < n; j++ {
			_, ok := map_list[Point{x: n_list[j].x, y: n_list[j].y}]
			if !ok {
				is_correct = false
			}
		}
		if is_correct {
			fmt.Fprintln(w, -dif_x, -dif_y)
		}
	}
}
