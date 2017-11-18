#!/usr/bin/env ruby
#
# https://leetcode.com/problems/longest-palindromic-substring/description/
#

def longest_palindrome s
  stack = []
  hash = {}

  i = 0
  while i < s.length
    if hash[s[i]] == nil
      hash[s[i]] = i
      # push into stack
      stack << s[i]
      i += 1
    else
      # pop
      
    end
  end
end
