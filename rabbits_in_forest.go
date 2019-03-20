package Leetcode

func numRabbits(answers []int) int {
	var result int
	mm := make(map[int]int)
	for _, rab := range answers {
		if num, ok := mm[rab+1]; !ok || num == 0 {
			result += rab + 1
			mm[rab+1] = rab
		} else {
			mm[rab+1] --
		}
	}

	return result
}
