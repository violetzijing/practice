#!/usr/bin/env ruby
#
# https://leetcode.com/problems/add-digits/#/description
#

def add_digits num
  return num if num < 10

  (num - 1) % 9 + 1
end

puts add_digits 38
