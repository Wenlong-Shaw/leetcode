package leetcode

func firstUniqChar(s string) int {
	ht := make(map[byte][]int, len(s))
	results := []int{}
	for i := 0; i < len(s); i++ {
		ht[s[i]] = append(ht[s[i]], i)
	}
	for i := range ht {
		v := ht[i]
		if len(v) < 2 {
			results = append(results, v[0])
		}
	}
	for i := range results {
		if results[0] > results[i] {
			results[0] = results[i]
		}
	}
	if len(results) == 0 {
		return -1
	}
	return results[0]
}
