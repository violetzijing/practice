#!/usr/bin/env ruby
#
# https://leetcode.com/problems/roman-to-integer/#/description
#

def roman_to_int s
  roman_num = {
                "I" => 1,
                "V" => 5,
                "X" => 10,
                "L" => 50,
                "C" => 100,
                "D" => 500,
                "M" => 1000
              }

  result = roman_num[s[0]]
  return result if s.length == 1

  i = 1
  while i < s.length do
    if roman_num[s[i - 1]] >= roman_num[s[i]]
      result += roman_num[s[i]]
    else
      result -= roman_num[s[i - 1]]
      result += roman_num[s[i]] - roman_num[s[i - 1]]
    end
    i += 1
  end
  
  return result
end

puts roman_to_int "MMMCCCXXXIII"
