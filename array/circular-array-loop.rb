#!/usr/bin/env ruby
#
# https://leetcode.com/problems/circular-array-loop/#/description
#

def circular_array_loop(nums)
  nums_hash = {}
  step = nums[0]
  while step < nums.length
    return true if nums[step] == "a"
    next_step = (step + nums[step]) % nums.length
    return false if next_step == step
    nums[step] = "a"
    step = next_step
  end

  return true
end

a =  [-1, 2]
puts circular_array_loop a
