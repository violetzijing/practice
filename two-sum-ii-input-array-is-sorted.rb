#!/usr/bin/env ruby
#
# https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/#/description
#

def two_sum numbers, target
  num_set = {}
  
  i = 0
  if numbers.first > 0
    while i < numbers.length
      break if numbers[i] > target
      i += 1
    end
    sub_nums = numbers[0...i]
  else
    sub_nums = numbers
  end

  i = 0
  minus_result = []
  while i < sub_nums.length
    minus = target - sub_nums[i]
    if num_set[minus] != nil
      if num_set[minus] > i
        return [i + 1, num_set[minus] + 1]
      else
        return [num_set[minus] + 1, i + 1]
      end
    else
      num_set[sub_nums[i]] = i
    end
    i += 1
  end
end

puts two_sum [-3,3,4,90], 0
