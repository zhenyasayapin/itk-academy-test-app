package main

func intersection(a, b []int) (bool, []int) {
	result := []int{}

	lookup := make(map[int]bool)

	for _, v := range b {
		lookup[v] = true
	}

	for _, v := range a {
		_, exist := lookup[v]
		if exist {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}
