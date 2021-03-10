package cmp

func CrossStrings(a, b []string) bool {
	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				return true
			}
		}
	}

	return false
}
