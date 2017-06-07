#!/usr/bin/env ruby
#
# https://leetcode.com/problems/plus-one/#/description
#

require 'pry'

def plus_one digits
  i = -1
  carry = 0

  digits[-1] = digits[-1] + 1

  while i > 0 - digits.length - 1
    if carry == 1
      digits[i] = digits[i] + 1
      carry = 0
    end

    if digits[i] > 9
      digits[i] = digits[i] - 10
      carry = 1
    end

    i -= 1
  end
  
  digits.unshift(1) if carry == 1

  return digits
end

a = [9, 9, 9]

puts plus_one(a)
