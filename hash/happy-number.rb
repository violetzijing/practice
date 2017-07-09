#!/usr/bin/env ruby
#
# https://leetcode.com/problems/happy-number/#/description
#

def is_happy n
  return true if n == 1
  nums_hash = {}
  while n != 1
    sum = 0
    num = n
    while num > 0
      remain = num % 10
      sum += remain * remain
      num /= 10
    end
#    puts "key is #{n}, value is #{sum}"
    return true if sum == 1
    return false if nums_hash[n] != nil
    nums_hash[n] = sum
    n = sum
  end
end

puts is_happy 2
