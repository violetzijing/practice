#!/usr/bin/env ruby
#
# https://leetcode.com/problems/valid-anagram/#/description
#

def print_array array
  array.each do |i|
    if i == nil
      i = "nil"
    end
    print "#{i} -> "
  end
  puts
end

def is_anagram s, t
  return false if s.length != t.length
  return true if s == t

  array = Array.new(26)

  i = 0
  while i < s.length
    index = s[i].ord - "a".ord
    if array[index] == nil
      array[index] = 1
    else
      array[index] += 1
    end
    i += 1
  end

  i = 0
  while i < t.length
    index = t[i].ord - "a".ord
    if array[index] == nil
      return false
    else
      array[index] -= 1
      return false if array[index] == -1
    end
    i += 1
  end

  array.each do |i|
    next if i == nil
    return false if i > 0
  end

  return true
end

puts is_anagram "anagram", "nagaram"
