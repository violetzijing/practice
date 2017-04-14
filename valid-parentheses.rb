#!/usr/bin/env ruby
#
# https://leetcode.com/problems/valid-parentheses/#/description
#

def is_pair s1, s2
  if s1 == "{" and s2 == "}"
    return true
  elsif s1 == "(" and s2 == ")"
    return true
  elsif s1 == "[" and s2 == "]"
    return true
  else
    return false
  end
end

def is_valid s
  return false if s.length == 1 or s.length == 0
  i = 0
  stack = []

  while i < s.length do
    if s[i] == "(" or s[i] == "{" or s[i] == "["
      stack << s[i]
    elsif s[i] == ")" or s[i] == "}" or s[i] == "]"
      pop = stack.pop
      puts "#{pop} and #{s[i]}"
      return false unless is_pair pop, s[i]
    end
    i += 1
  end
  return false unless stack.empty?

  return true
end

puts is_valid "(("
