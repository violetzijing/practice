#!/usr/bin/env ruby
#
# https://leetcode.com/problems/peak-index-in-a-mountain-array/description/
#

def peak_index_in_mountain_array a
  max = a[0]
  i = 0
  while i < a.length
    if a[i] < max
      return i
    else
      max = a[i]
    end

    i++
  end
end

puts peak_index_in_mountain_array([0,2,1,0])
