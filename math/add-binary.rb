#!/usr/bin/env ruby
#
# https://leetcode.com/problems/add-binary/#/description
#

require 'pry'

def add_binary a, b
  i = -1
  carry = 0
  result = ""

  while i > 0 - a.length - 1 or i > 0 - b.length - 1
    tmp_a = (a[i].nil? ? 0 : a[i].to_i)
    tmp_b = (b[i].nil? ? 0 : b[i].to_i)
    binding.pry

    tmp = tmp_a + tmp_b + carry

    carry = 0 if carry != 0

    if tmp > 2
      carry = 1
      tmp = 1
    elsif tmp > 1
      carry = 1
      tmp = 0
    end

    result.prepend("#{tmp}")

    i -= 1
  end

  remain = (a.length > b.length ? a[0..i] : b[0..i])
  result.prepend("1") if carry == 1
  result.prepend(remain)

  return result
end

puts add_binary "0", "1"
