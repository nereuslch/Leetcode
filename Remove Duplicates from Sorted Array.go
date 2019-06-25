package Leetcode


func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	for i := 0 ;i < len(nums); {
		flag := false
		j := i+1
		for ;j < len(nums);j ++ {
			if nums[i] != nums[j] {
				break
			}
			flag = true
		}

		if j == i+1 {
			i= i+1
			continue
		}

		if flag {
			nums = append(nums[:i+1],nums[j:]...)
			i = i+1
		}else {
			break
		}
	}
	return len(nums)
}

func removeDuplicates(a []int) int {
	left, right, size := 0, 1, len(a)
	for ; right < size; right++ {
		if a[left] == a[right] {
			continue
		}
		left++
		a[left], a[right] = a[right], a[left]
	}
	return left + 1
}