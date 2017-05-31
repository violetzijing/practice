#!/usr/bin/env ruby
#
# https://leetcode.com/problems/house-robber/#/description
#

def rob nums
  a, b = 0, 0
 
  i = 0
  while i < nums.length
    if i % 2 == 0
      a = [a + nums[i], b].max
    else
      b = [a, b + nums[i]].max
    end
    i += 1
  end

  return (a > b ? a : b)
end

puts rob([1, 3, 4, 5])
