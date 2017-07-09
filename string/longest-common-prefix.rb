#!/usr/bin/env ruby
#
# https://leetcode.com/problems/longest-common-prefix/#/description
#

def longest_common_prefix strs
  i = 0
  j = 0

  return "" if strs.length == 0
  length = strs[0].length
  
  while j < length do
    i = 1
    while i < strs.length do
      if strs[i][j] != strs[0][j]
        common_prefix = (j == 1) ? strs[0][0] : strs[0][0, j]
        return "" if common_prefix.nil?
        return common_prefix
      end
      i += 1
    end
    j += 1
  end

  return strs[0][0, j]
end

puts (longest_common_prefix ["123", "fsdhk123w", "1234"])
