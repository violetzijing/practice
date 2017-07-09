#!/usr/bin/env ruby
#
# https://leetcode.com/problems/remove-element/#/description
#

def remove_element nums, val
  head = 0
  tail = nums.length - 1

  while head <= tail do
    if nums[head] == val
      if nums[tail] != val
        nums[head] = nums[tail]
        nums[tail] = nil
        head += 1
      end
      tail -= 1
    else
      head += 1
    end
  end

  return head
end

puts remove_element [2], 2
