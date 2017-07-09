#!/usr/bin/env ruby
#
# https://leetcode.com/problems/power-of-two/#/description
#

def is_power_of_two n
  return (n > 0) && ((n & (n - 1)) == 0)
end

puts is_power_of_two 8
