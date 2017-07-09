#!/usr/bin/env ruby
#
# https://leetcode.com/problems/excel-sheet-column-title/#/description
#

def convert_to_title(n)
  alphabet = []
  ("A".."Z").each {|i| alphabet << i }

  i = 0
  result = ""
  while n > 0
    n -= 1
    mod_num = n % 26
    result = alphabet[mod_num] + result
    n = n / 26
  end

  return result
end

puts convert_to_title 1
