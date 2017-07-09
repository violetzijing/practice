#!/usr/bin/env ruby
#
# https://leetcode.com/problems/sum-of-two-integers/#/description
#

def get_sum(a, b)
  return b if a == 0
  return a if b == 0

  max_int = 0x7FFFFFFF
  min_int = 0x80000000
  mask = 0x100000000
  while b != 0
    a, b = (a ^ b) % mask, ((a & b) << 1) % mask
  end
  if a <= max_int
    return a
  else
    return ~((a % min_int) ^ max_int)
  end
end

puts get_sum 2, 3
