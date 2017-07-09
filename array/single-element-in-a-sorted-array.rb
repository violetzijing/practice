#!/usr/bin/env ruby
#
# https://leetcode.com/problems/single-element-in-a-sorted-array/#/description
#

def single_non_duplicate nums
  start = 0
  _end = nums.length - 1
  while start < _end
    tmp_mid = (start + _end) / 2
    mid = tmp_mid + tmp_mid % 2
    if nums[mid] == nums[mid - 1]
      # left shifta
      _end = mid - 2
    elsif nums[mid] == nums[mid + 1]
      # right shift
      start = mid + 2
    else
      return nums[mid]
    end
  end

  return nums[start]
end

a = [0,1,1]
puts single_non_duplicate a
