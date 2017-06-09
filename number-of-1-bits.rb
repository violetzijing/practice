#!/usr/bin/env ruby
#
# https://leetcode.com/problems/number-of-1-bits/#/description
#

def hamming_weight(n)
  i = 0
  while n != 0
    n &= n - 1
    i += 1
  end

  return i
end

puts hamming_weight 0
