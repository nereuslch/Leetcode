package Leetcode

func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	sum = sum / 2
	dp := make([]bool,sum+1)
	dp[0] = true
	for _,v :=range nums {
		for i := sum;i>=v;i-- {
			dp[i] = dp[i] || dp[i- v]
		}
	}

	return dp[sum]


}