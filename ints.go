package cmp

func CrossInt64s(a, b []int64) bool {
	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				return true
			}
		}
	}

	return false
}

func InSInt64s(a []int64, b int64) bool {
	for _, av := range a {
		if av == b {
			return true
		}
	}

	return false
}
