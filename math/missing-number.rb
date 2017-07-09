#!/usr/bin/env ruby
#
# https://leetcode.com/problems/missing-number/#/description
#

def missing_number nums
  (n = nums.size) * (n + 1) / 2 - nums.reduce(:+)
end
