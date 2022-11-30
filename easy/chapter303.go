package easy

type NumArray struct {
	nums    []int
	sumNums []int
}

func Constructor3(nums []int) NumArray {
	sumNums := make([]int, len(nums))
	sumNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sumNums[i] = sumNums[i-1] + nums[i]
	}

	return NumArray{
		nums:    nums,
		sumNums: sumNums,
	}

}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sumNums[right] - this.sumNums[left] + this.nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
