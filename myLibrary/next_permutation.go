// 引用元:https://qiita.com/ktateish/items/ab2df3e0864d2e931bf2#next_permutation
// 解説:https://qiita.com/siser/items/a91022071b24952d27d9#3next_permutation%E3%81%AE%E4%BB%95%E7%B5%84%E3%81%BF%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6
func next_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return x < y })
}

func prev_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return y < x })
}

func _permutation_pandita(a []int, less func(x, y int) bool) bool {
	i := len(a) - 2
	// Find the right most incresing elements a[i] and a[i+1]
	for 0 <= i && !less(a[i], a[i+1]) {
		i--
	}
	if i == -1 {
		return false
	}
	j := i + 1
	// Find the smallest element that is greater than a[i] in a[i+1:]
	for k := j + 1; k < len(a); k++ {
		// a[i] < a[k] && a[k] <= a[j]
		if less(a[i], a[k]) && !less(a[j], a[k]) {
			j = k
		}
	}
	a[i], a[j] = a[j], a[i]
	for p, q := i+1, len(a)-1; p < q; p, q = p+1, q-1 {
		a[p], a[q] = a[q], a[p]
	}
	return true
}