#!/usr/bin/env ruby
#
# https://leetcode.com/problems/climbing-stairs/#/description
#

def climb_stairs n
  return 0 if n == 0
  return 1 if n == 1
  return 2 if n == 2

  move = [1, 2]

  all_path = 0
  one_step = 2
  two_steps = 1

  i = 2
  while i < n
    all_path = one_step + two_steps
    two_steps = one_step
    one_step = all_path
    i += 1
  end

  return all_path
end

puts climb_stairs(4)
