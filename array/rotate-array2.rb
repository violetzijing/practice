#!/usr/bin/env ruby
#
# https://leetcode.com/problems/rotate-array/description/
#

def rotate nums, k
  return nums if nums.length < 2

  k %= nums.length
  reverse nums, 0, nums.length - 1
  reverse nums, 0, k - 1
  reverse nums, k, nums.length - 1
  return nums
end

def reverse nums, head, tail
  while head < tail
    nums[head], nums[tail] = nums[tail], nums[head]
    head += 1
    tail -= 1
  end

  return nums
end

array = [1, 2, 3]

rotate array, 1

array.each {|i| print "#{i}, " }
puts 
