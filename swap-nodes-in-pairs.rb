#!/usr/bin/env ruby
#
# https://leetcode.com/problems/swap-nodes-in-pairs/#/description
#

require_relative 'list-helper'

def swap_pairs head
  return if head == nil
  return head if head.next == nil
  
  p1 = head
  p2 = head.next
  new_head = p2
  parent = nil

  while p2 != nil
    tmp = p2.next

    p2.next = p1
    p1.next = tmp
    parent.next = p2 if parent != nil

    break if tmp == nil
    parent = p1
    p1 = tmp
    p2 = p1.next
  end

  return new_head
end

a = create_lst [1, 2, 3,4, 5]
print_lst swap_pairs a
