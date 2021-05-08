package leetcode

func decode(encoded []int, first int) []int {
	results := []int{}
	results = append(results, first)
	for i := 0; i < len(encoded); i++ {
		second := first ^ encoded[i]
		first = second
		results = append(results, first)
	}
	return results
}
