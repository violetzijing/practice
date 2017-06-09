#!/usr/bin/env ruby
#
# https://leetcode.com/problems/length-of-last-word/#/description
#

def length_of_last_word s
  return 0 if s.length == 0

  i = s.length - 1
  while i >= 0 and s[i] == " "
    i -= 1
  end

  return 0 if i == -1

  start = i
  while i >= 0 and s[i] != " "
    i -= 1
  end

  return start - i
end

puts length_of_last_word "     "
