#!/usr/bin/env ruby
#
# https://leetcode.com/problems/length-of-last-word/description/
#

def length_of_last_word(s)
  array = s.split(" ")
  return 0 if array.length == 0
  return array.last.length    
end
