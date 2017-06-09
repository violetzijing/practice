#!/usr/bin/env ruby
#
# https://leetcode.com/problems/factorial-trailing-zeroes/#/description
#

def trailing_zeroes(n)
  i = 5
  result = 0

  while n / i > 0
    result += n / i
    i *= 5
  end

  return result
end

puts trailing_zeroes 100
