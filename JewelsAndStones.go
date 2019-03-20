package Leetcode

func numJewelsInStones(J string, S string) int {
	m := make(map[string]struct{})
	for _, v := range J {
		m[string(v)] = struct{}{}
	}

	var sum int
	for _, v := range S {
		if _, ok := m[string(v)]; ok {
			sum += 1
		}
	}
	return sum
}
