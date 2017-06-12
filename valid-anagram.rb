#!/usr/bin/env ruby
#
# https://leetcode.com/problems/valid-anagram/#/description
#

def is_anagram s, t
  return false if s.length != t.length
  return true if s == t

  i = 0
  s_hash = {}
  while i < s.length
    if s_hash[s[i]] == nil
      s_hash[s[i]] = 1
    else
      s_hash[s[i]] += 1
    end
    i += 1
  end

  i = 0
  while i < t.length
    if s_hash[t[i]] == nil
      return false
    else
      s_hash[t[i]] -= 1
      return false if s_hash[t[i]] == -1
    end
    i += 1
  end

  s_hash.each do |key, value|
    return false if value != 0
  end

  return true
end

puts is_anagram "car", "rat"
