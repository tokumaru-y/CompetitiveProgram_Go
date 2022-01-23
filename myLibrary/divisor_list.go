func make_divisor_list(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			res = append(res, i)
			res = append(res, n/i)
		}
	}
	return res
}
