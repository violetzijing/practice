#!/usr/bin/env ruby
#
# https://leetcode.com/problems/two-sum/#/description
#

def two_sum nums, target
  nums_set = {}
  minus_result = []

  i = 0
  while i < nums.length do
    minus = target - nums[i]
    if nums_set[minus] != nil
      return [nums_set[minus], i]
    end
    
    nums_set[nums[i]] = i
    i += 1
  end
end

puts two_sum [2, 2, 3, 4], 7
