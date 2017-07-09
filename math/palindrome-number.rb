#!/usr/bin/env ruby
#
# https://leetcode.com/problems/palindrome-number/#/description
#

def is_palindrome x
  return false if x < 0

  i = 0
  power = 0
  tmp = x
  while tmp > 0 do
    power += 1
    tmp /= 10
  end

  first = 0
  last = 0
  power_num = 10 ** (power - 1)

  while x > 0 do
    last = x % 10
    first = x / power_num
    return false if last != first
    x = x - first * power_num
    x = x / 10
    power_num /= 100
  end

  return true
end

puts is_palindrome 100031
