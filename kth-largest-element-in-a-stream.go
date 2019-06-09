package main

import "fmt"

func main() {
	input := []int{4, 5, 8, 2}
	instance := Constructor(3, input)
	fmt.Println(instance.Nums)
	/*
		instance.Add(3)
		fmt.Println(instance.Nums)
			instance.Add(5)
			fmt.Println(instance.Nums)
			instance.Add(10)
			fmt.Println(instance.Nums)
			instance.Add(9)
			fmt.Println(instance.Nums)
			instance.Add(4)
			fmt.Println(instance.Nums)
	*/
}

type KthLargest struct {
	K    int
	Nums []int
}

func Constructor(k int, nums []int) KthLargest {
	instance := &KthLargest{K: k, Nums: nums}
	instance.Init()
	size := len(nums)
	for size > k {
		instance.DeleteMin()
		size = len(instance.Nums)
	}
	return *instance
}

func (this *KthLargest) Add(val int) int {
	this.Insert(val)
	size := len(this.Nums)
	for size > this.K {
		this.DeleteMin()
		size = len(this.Nums)
		fmt.Println("size: ", size)
		fmt.Println("k: ", this.K)
	}
	return this.Nums[0]
}

func (this *KthLargest) Left(parent int) int {
	return 2*parent + 1
}

func (this *KthLargest) Right(parent int) int {
	return 2*parent + 2
}

func (this *KthLargest) Parent(left int) int {
	return (left - 1) / 2
}

func (this *KthLargest) Swap(a, b int) {
	this.Nums[a], this.Nums[b] = this.Nums[b], this.Nums[a]
}

func (this *KthLargest) Swim(key int) {
	parent := this.Parent(key)
	for this.Nums[parent] > this.Nums[key] && parent >= 0 {
		this.Swap(parent, key)
		key = parent
		parent = this.Parent(key)
	}
}

func (this *KthLargest) Less(a, b int) bool {
	return this.Nums[a] < this.Nums[b]
}

func (this *KthLargest) Sink(key int) {
	left := this.Left(key)
	right := this.Right(key)

	for (left < len(this.Nums) && this.Less(left, key)) ||
		(right < len(this.Nums) && this.Less(right, key)) {
		if left < len(this.Nums) && right < len(this.Nums) {
			if this.Less(left, right) {
				this.Swap(key, left)
				key = left
			} else {
				this.Swap(key, right)
				key = right
			}
		}
		if left < len(this.Nums) && !(right < len(this.Nums)) {
			this.Swap(key, left)
			key = left
		}
		if right < len(this.Nums) && !(right < len(this.Nums)) {
			this.Swap(key, right)
			key = right
		}
	}
}

func (this *KthLargest) Init() {
	for i := len(this.Nums) - 1; i >= 0; i-- {
		this.Sink(i)
	}
}

func (this *KthLargest) Insert(val int) {
	this.Nums = append(this.Nums, val)
	this.Swim(len(this.Nums) - 1)
}

func (this *KthLargest) DeleteMin() {
	if len(this.Nums) <= 0 {
		return
	}
	this.Nums[0] = this.Nums[len(this.Nums)-1]
	this.Nums = this.Nums[:len(this.Nums)-2]
	fmt.Println("this.Nums: ", this.Nums)
	this.Sink(0)
}

func (this *KthLargest) GetKth(K int) int {
	return this.Nums[K]
}
