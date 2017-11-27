#!/usr/bin/env ruby
#
# https://leetcode.com/problems/implement-strstr/description/
#

def str_str haystack, needle
  return -1 if needle.length > haystack.length

  i = 0
  loop do
    j = 0
    loop do
      return i if j == needle.length
      return -1 if i + j == haystack.length
      break if needle[j] != haystack[i + j]
      j += 1
    end

    i += 1
  end
end

puts str_str "aaaa", "ll"
