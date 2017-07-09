#!/usr/bin/env ruby
#
# https://leetcode.com/problems/house-robber/#/description
#

require 'pry'

def rob(nums)
  return 0 if nums.length == 0

  d = Array.new(0)
  d[0] = 0
  d[1] = nums[0]
  i = 2
  while i <= nums.length
    d[i] = [d[i - 1], d[i - 2] + nums[i - 1]].max
    i += 1
  end

  return d[nums.length]
end

puts rob([1, 3, 4, 5])
