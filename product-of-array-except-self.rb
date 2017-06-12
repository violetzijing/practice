#!/usr/bin/env ruby
#
# https://leetcode.com/problems/product-of-array-except-self/#/description
#

def product_except_self nums
  output = []
  i = 0
  tmp = 1
  while i < nums.length
    output[i] = tmp
    tmp *= nums[i]
    i += 1
  end

  i = nums.length - 1
  tmp = 1
  while i >= 0
    output[i] *= tmp
    tmp *= nums[i]
    i -= 1
  end

  return output
end

a = [1]
puts product_except_self a
