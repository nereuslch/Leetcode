package Leetcode

func sortArrayByParity(A []int) []int {
	var odd int = 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[odd], A[i] = A[i], A[odd]
			odd ++
		}
	}
	return A
}
