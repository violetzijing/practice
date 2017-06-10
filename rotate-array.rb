#!/usr/bin/env ruby
#
# https://leetcode.com/problems/rotate-array/#/description
#

def rotate nums, k
  return if k == 0

  tmp = nums[0...nums.length - k]
  first = nums[nums.length - k..nums.length - 1]
  nums = first + tmp
end

a = [1, 2]
rotate a, 1
