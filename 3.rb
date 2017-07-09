#!/usr/bin/env ruby
#
require 'pry'
def solution s, t
  return "NOTHING" if t.include? s

  # if s and t have common prefix
  if s[0] == t[0]
    compared_position = find_common_prefix s, t
    remain_t = t[compared_position, t.length - compared_position]
    remain_s = s[compared_position, s.length - compared_position]

    difference_prefix(remain_s, remain_t)
  else
    # s and t don't have common prefix
    difference_prefix(s, t)
  end
end

def find_common_prefix s, t
  i = 0
  while i < s.length
    if s[i] != t[i]
      return i
    else
      i += 1
    end
  end
  return s.length - 1
end

def difference_prefix remain_s, remain_t
  if remain_t.include? remain_s
    # insert
    if remain_t[1, remain_t.length - 1].include? remain_s
      return "INSERT #{remain_t[0]}"
    else
      return "IMPOSSIBLE"
    end
  elsif remain_t.include? remain_s[1, remain_s.length - 1]
    # delete
    if remain_t.include? remain_s[1, remain_s.length - 1]
      return "DELETE #{remain_s[0]}"
    else
      return "IMPOSSIBLE"
    end
  else
    # swap
    if remain_t[2, remain_t.length - 2].include? remain_s[2, remain_s.length - 2]
      if remain_t[0] == remain_s[1] and remain_s[0] == remain_t[1]
        return "SWAP #{remain_s[0]} #{remain_s[1]}"
      else
        return "IMPOSSIBLE"
      end
    else
      return "IMPOSSIBLE"
    end
  end
end

puts solution("test", "testt123")
