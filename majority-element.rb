#!/usr/bin/env ruby
#
# https://leetcode.com/problems/majority-element/#/description 
#

def majority_element nums
  return if nums.length == 0
  return nums.first if nums.length == 1
  nums_hash = {}

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
    return key if value > nums.length / 2
  end

  return
end

a = [ 1, 1, 2, 3, 4, 5]
puts majority_element a
