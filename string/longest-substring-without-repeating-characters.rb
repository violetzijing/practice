#!/usr/bin/env ruby
#
# https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
#

def length_of_longest_substring(s)
  max = 0
  hash = {}
  left = 0
  right = 0

  while right != s.length
    if hash[s[right]] == nil or hash[s[right]] < left
      hash[s[right]] = right
      right += 1
      next
    end

    max = (max > right - left) ? max : right - left
    left = hash[s[right]] + 1
  end

  max = (max > right - left) ? max : right - left
end

puts length_of_longest_substring "zyhormwjcxfbobwrcvehbitnsdgacjpeiysbmrnoqssfvoyxke"
