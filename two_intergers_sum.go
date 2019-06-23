package Leetcode

func getSum(a int, b int) int {
	var c int

	for b != 0 {
		c = a & b
		a = a ^ b
		b = c << 1
	}

	return a
}


// 3 ^2        =>  001
// (3&2)<<1    =>  100

