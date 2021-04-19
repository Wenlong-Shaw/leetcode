package leetcode

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	var ns Minnum = nums
	numstr := ""
	sort.Sort(ns)
	for i := range ns {
		numstr += strconv.Itoa(ns[i])
	}
	return numstr
}

/*
type Minnum []int

func (n Minnum) Len() int { return len(n) }
func (n Minnum) Less(i, j int) bool {
	ico, jco := 1, 1

	if n[i] == 0 {
		jco = 10
	} else {
		for ni := n[i]; ni > 0; ni /= 10 {
			jco *= 10
		}
	}

	if n[j] == 0 {
		ico = 10
	} else {
		for nj := n[j]; nj > 0; nj /= 10 {
			ico *= 10
		}
	}

	if n[i]*ico+n[j] < n[j]*jco+n[i] {
		return true
	}
	return false
}
func (n Minnum) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
*/
