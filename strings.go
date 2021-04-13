package cmp

import "strings"

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

func InStrings(a []string, b string) bool {
	for _, av := range a {
		if av == b {
			return true
		}
	}

	return false
}

func InStringsNonStrict(a []string, b string) bool {
	b = strings.ToLower(b)
	for _, av := range a {
		av = strings.ToLower(av)
		if av == b {
			return true
		}
	}

	return false
}

func ContainStrings(a []string, b string) bool {
	for _, av := range a {
		if strings.Contains(av, b) {
			return true
		}
	}

	return false
}

func ContainStringsNonStrict(a []string, b string) bool {
	b = strings.ToLower(b)
	for _, av := range a {
		av = strings.ToLower(av)
		if strings.Contains(av, b) {
			return true
		}
	}

	return false
}
