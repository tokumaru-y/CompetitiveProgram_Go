package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		var inp int
		fmt.Fscan(r, &inp)
		A[i] = inp
	}
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	var alice_p, bob_p int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice_p += A[i]
		} else {
			bob_p += A[i]
		}
	}
	fmt.Fprintln(w, alice_p-bob_p)
}
