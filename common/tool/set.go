package tool

func DiffSet(a []int64, b []int64) []int64 {
	var diff []int64
	mp := make(map[int64]int)
	for _, n := range b {
		if _, ok := mp[n]; !ok {
			mp[n] = 1
		}
	}

	for _, n := range a {
		if _, ok := mp[n]; !ok {
			diff = append(diff, n)
		}
	}

	return diff
}
