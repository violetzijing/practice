#!/usr/bin/env ruby
#
# https://leetcode.com/problems/poor-pigs/#/description
#

def poor_pigs(buckets, minutes_to_die, minutes_to_test)
  pigs = 0
  while (minutes_to_test / minutes_to_die + 1) ** pigs < buckets
    pigs += 1
  end
  
  pigs
end

puts poor_pigs 1000, 15, 60
