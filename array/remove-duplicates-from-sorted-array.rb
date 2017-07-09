#!/usr/bin/env ruby
#
# https://leetcode.com/problems/remove-duplicates-from-sorted-array/#/description
#

def remove_duplicates nums
  return 0 if nums.length == 0

  i = 1
  pointer = 0
  length = nums.length

  while i < nums.length do
    if nums[i - 1] != nums[i]
      pointer += 1
      nums[pointer] = nums[i]
    end
    i += 1
  end
  return pointer + 1
end

puts remove_duplicates [1]
