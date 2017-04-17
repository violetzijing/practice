#!/usr/bin/env ruby
#
# https://leetcode.com/problems/two-sum/#/description
#

def two_sum nums, target
  nums_set = {}
  minus_result = []
  results = []

  i = 0
  while i < nums.length do
    nums_set[i] = nums[i]
    i += 1
  end

  i = 0
  while i < nums.length do
    minus_result << target - nums[i]
    i += 1
  end

  i = 0
  while i < minus_result.length do
    tmp = minus_result[i]
    if nums_set.has_value? tmp
      pair1 = i
      pair2 = nums_set.key(tmp)
      results << [pair1, pair2]
    end
    i += 1
  end

  results.each do |result|
    return result if result[0] != result[1]
  end
end

puts two_sum [3,2,4], 6
