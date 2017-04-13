#!/usr/bin/env ruby
#
# https://leetcode.com/problems/string-to-integer-atoi/#/description
#

def my_atoi str
  i = 0
  negative = false
  result = 0
  int_max = 2147483647
  int_min = -2147483648

  while i < str.length do
    break unless str[i] == " " or str[i] == "\t"
    i += 1
  end

  if str[i] == "+"
    i += 1
  elsif str[i] == "-"
    negative = true
    i += 1
  end

  str[str.length] = "x"

  while str[i] >= "0" and str[i] <= "9" do
    if !negative and result > int_max / 10
      return int_max
    elsif negative and -result < int_min / 10
      return int_min
    end

    digit = str[i].ord - 48
    
    if !negative and result * 10 > int_max - digit
      return int_max
    elsif negative and -result * 10 < int_min + digit
      return int_min
    end

    result = result * 10 + digit
    i += 1
  end

  result = 0 - result if negative

  return result
end


puts my_atoi "-2147483648"  
