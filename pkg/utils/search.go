package utils

func Search(data []string, arg string) []string {
	var result []string
	for _, line := range data {
		if v, found := isFound(line, arg); found {
			result = append(result, v)
		}
	}
	return result
}
