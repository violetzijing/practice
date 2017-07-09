#!/usr/bin/env ruby
#
# https://leetcode.com/problems/reverse-linked-list/#/description
#

class ListNode
    attr_accessor :val, :next
    def initialize(val)
        @val = val
        @next = nil
    end
end

def print_lst head
  node = head
  while node != nil
    print "#{node.val} -> "
    node = node.next
  end
end

def reverse_list head
  return if head == nil
  return head if head.next == nil

  node = head
  parent = nil

  while node != nil
    tmp = node.next
    node.next = parent
    parent = node
    node = tmp
  end

  return parent
end

n1 = ListNode.new 1
n2 = ListNode.new 2
n3 = ListNode.new 3
n4 = ListNode.new 4
n5 = ListNode.new 5

n1.next = n2
n2.next = n3
n3.next = n4
n4.next = n5

print_lst reverse_list n4

