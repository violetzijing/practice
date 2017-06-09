#!/usr/bin/env ruby
#
# https://leetcode.com/problems/gray-code/#/description
#

def gray_code n
  array = []

  array << 0

  i = 0
  while i < n
    inc = 1 << i
    j = array.length - 1
    while j >= 0
      array << array[j] + inc
      j -= 1
    end

    i += 1
  end

  return array
end

puts gray_code 3
