#!/usr/bin/env ruby
#
# https://leetcode.com/problems/3sum-closest/description/
#

def three_sum_closest nums, target
  nums.sort!
  m = 0
  result = 1000000000
  minus = 1000000000

  while m < nums.length - 2
    i = m + 1
    j = nums.length - 1
    while i < j
      sum = nums[m] + nums[i] + nums[j]
      if sum > target
        j -= 1
      elsif sum < target
        i += 1
      else
        return sum
      end

      tmp = sum - target
      tmp = 0 - tmp if tmp < 0

      if tmp < minus
        result = sum
        minus = tmp
      end
    end

    m += 1
  end

  return result
end

puts three_sum_closest [1, 2, 5, 10, 11], 12
