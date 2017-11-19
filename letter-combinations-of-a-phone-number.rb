#!/usr/bin/env ruby
#
# https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
#

def letter_combinations digits
  return [] if digits.length == 0
  hash = {
    "2" => %w(a b c),
    "3" => %w(d e f),
    "4" => %w(g h i),
    "5" => %w(j k l),
    "6" => %w(m n o),
    "7" => %w(p q r s),
    "8" => %w(t u v),
    "9" => %w(w x y z)
  }

  result = []

  stack = [-1]
  while !stack.empty?
    stack[-1] += 1

    while stack.length != digits.length
      stack << 0
    end

    str = ""
    stack.each_index do |i|
      str += hash[digits[i]][stack[i]]
    end

    result << str

    while hash[digits[stack.length - 1]].length - 1 == stack[-1]
      stack.pop
    end
  end

  return result
end

result = letter_combinations ""

result.each {|i| print "#{i}, " }
