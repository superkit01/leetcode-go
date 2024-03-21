package middle

// nums [0,0,3,0,0,0]
//       0 1 2 3 4 5
// freq [0,0,0,1,0,0]
//       0 1 2 3 4 5
type FrequencyTracker struct {
	nums []int
	freq []int
}

func Constructor5() FrequencyTracker {
	return FrequencyTracker{
		nums: make([]int, 100001),
		freq: make([]int, 100001),
	}
}

func (this *FrequencyTracker) Add(number int) {
	prev := this.nums[number]
	this.nums[number]++
	if prev > 0 {
		this.freq[prev]--
	}
	this.freq[prev+1]++

}

func (this *FrequencyTracker) DeleteOne(number int) {

	prev := this.nums[number]
	if prev > 0 {
		this.nums[number]--
		if this.freq[prev] > 0 {
			this.freq[prev]--
			this.freq[prev-1]++
		}
	}

}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {

	return this.freq[frequency] > 0

}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
