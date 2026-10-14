package main

func findDifference(a, b []string) []string {
	var difference []string

	lookup := make(map[string]bool)

	for _, bValue := range b {
		lookup[bValue] = true
	}

	for _, aValue := range a {
		_, exist := lookup[aValue]
		if !exist {
			difference = append(difference, aValue)
		}
	}

	return difference
}
