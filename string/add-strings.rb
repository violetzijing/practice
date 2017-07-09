#!/usr/bin/env ruby
#
# https://leetcode.com/problems/add-strings/#/description
#

def add_strings num1, num2
  return "" if num1 == "" and num2 == ""

  i = num1.length - 1
  j = num2.length - 1
  sum = ""
  carry = 0
  while i >= 0 or j >= 0
    tmp1 = (i < 0 ? 0 : num1[i].to_i)
    tmp2 = (j < 0 ? 0 : num2[j].to_i)

    tmp = tmp1 + tmp2 + carry
    carry = 0 if carry == 1
  
    if tmp > 9
      carry = 1
      tmp = tmp % 10
    end
    
    sum = "#{tmp}" + sum
    i -= 1
    j -= 1
  end

  sum = "1" + sum if carry == 1
  return sum
end

puts add_strings "12345", "665"
