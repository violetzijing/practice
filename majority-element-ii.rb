#!/usr/bin/env ruby
#
# https://leetcode.com/problems/majority-element-ii/#/description
#

def majority_element nums
  return [] if nums.length == 0
  return [nums.first] if nums.length == 1
  nums_hash = {}
  result = []

  i = 0
  while i < nums.length
    if nums_hash[nums[i]] != nil
      nums_hash[nums[i]] += 1
    else
      nums_hash[nums[i]] = 1
    end

    i += 1
  end

  nums_hash.each do |key, value|
    result << key if value > nums.length / 3
  end

  return result
end

a = [1, 2, 1, 2]
puts majority_element a
