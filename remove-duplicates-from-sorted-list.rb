#!/usr/bin/env ruby
#
# https://leetcode.com/problems/remove-duplicates-from-sorted-list/#/description
#

require_relative "list-helper"

def delete_duplicates head
  return if head == nil
  return head if head.next == nil

  value = head.val
  node = head.next
  parent = head
  while node != nil
    if value != node.val
      value = node.val
      parent = node
      node = node.next
    else
      # delete the node
      tmp = node.next
      parent.next = node.next
      node = tmp
    end
  end

  return head
end

a = [1]
print_lst delete_duplicates(create_lst a)
